package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type Logger interface {
	Log(level Level, keyvals ...any) error
}

type Format int

const (
	FormatText Format = iota
	FormatJSON
)

type Option func(*handlerConfig)

type handlerConfig struct {
	writer      io.Writer
	format      Format
	level       Level
	addSource   bool
	replaceAttr func(groups []string, a slog.Attr) slog.Attr
}

type Valuer func(context.Context) any

type callerKey struct{}

type Field struct {
	Key   string
	Value any
}

const (
	MessageKey = slog.MessageKey

	FieldKeyChannel  = "channel"
	FieldKeyMethod   = "method"
	FieldKeyOrderNo  = "order_no"
	FieldKeyTradeNo  = "trade_no"
	FieldKeyRequest  = "request"
	FieldKeyResponse = "response"
	FieldKeyError    = "error"
)

func F(key string, value any) Field {
	return Field{Key: key, Value: value}
}

func WithWriter(w io.Writer) Option {
	return func(c *handlerConfig) {
		if w != nil {
			c.writer = w
		}
	}
}

func WithFormat(format Format) Option {
	return func(c *handlerConfig) {
		c.format = format
	}
}

func WithAddSource(addSource bool) Option {
	return func(c *handlerConfig) {
		c.addSource = addSource
	}
}

func WithReplaceAttr(replaceAttr func(groups []string, a slog.Attr) slog.Attr) Option {
	return func(c *handlerConfig) {
		c.replaceAttr = replaceAttr
	}
}

var globalLogger Logger = NewSlogLogger(LevelInfo)

func SetLogger(l Logger) {
	if l != nil {
		globalLogger = l
	}
}

func GetLogger() Logger {
	return globalLogger
}

func Log(level Level, keyvals ...any) error {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	vals := make([]any, 0, len(keyvals)+2)
	vals = append(vals, callerKey{}, pcs[0])
	vals = append(vals, keyvals...)
	return globalLogger.Log(level, vals...)
}

func With(l Logger, keyvals ...any) Logger {
	if l == nil {
		l = globalLogger
	}
	return &withLogger{logger: l, keyvals: keyvals}
}

func WithContext(ctx context.Context, l Logger) Logger {
	if ctx == nil {
		ctx = context.Background()
	}
	return &contextLogger{logger: l, ctx: ctx}
}

func Value(ctx context.Context, keyvals ...any) []any {
	if ctx == nil {
		ctx = context.Background()
	}
	vals := make([]any, 0, len(keyvals))
	for i := 0; i < len(keyvals); i += 2 {
		vals = append(vals, keyvals[i])
		if i+1 >= len(keyvals) {
			vals = append(vals, "")
			continue
		}
		if v, ok := keyvals[i+1].(Valuer); ok {
			vals = append(vals, v(ctx))
			continue
		}
		vals = append(vals, keyvals[i+1])
	}
	return vals
}

func Debug(ctx context.Context, msg string, fields ...Field) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = logWithContext(ctx, LevelDebug, pcs[0], msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...Field) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = logWithContext(ctx, LevelInfo, pcs[0], msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...Field) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = logWithContext(ctx, LevelWarn, pcs[0], msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...Field) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	_ = logWithContext(ctx, LevelError, pcs[0], msg, fields...)
}

func logWithContext(ctx context.Context, level Level, pc uintptr, msg string, fields ...Field) error {
	logger := WithContext(ctx, globalLogger)
	keyvals := make([]any, 0, 4+len(fields)*2)
	keyvals = append(keyvals, callerKey{}, pc, MessageKey, msg)
	for _, field := range fields {
		keyvals = append(keyvals, field.Key, field.Value)
	}
	return logger.Log(level, keyvals...)
}

type nopLogger struct{}

func (l *nopLogger) Log(level Level, keyvals ...any) error { return nil }

type withLogger struct {
	logger  Logger
	keyvals []any
}

func (l *withLogger) Log(level Level, keyvals ...any) error {
	vals := make([]any, 0, len(l.keyvals)+len(keyvals))
	vals = append(vals, l.keyvals...)
	vals = append(vals, keyvals...)
	return l.logger.Log(level, vals...)
}

type contextLogger struct {
	logger Logger
	ctx    context.Context
}

func (l *contextLogger) Log(level Level, keyvals ...any) error {
	if l.logger == nil {
		l.logger = globalLogger
	}
	return l.logger.Log(level, Value(l.ctx, keyvals...)...)
}

type SlogLogger struct {
	Level   Level
	Handler slog.Handler
}

type StdLogger = SlogLogger

func NewStdLogger(level Level, opts ...Option) *StdLogger {
	return NewSlogLogger(level, opts...)
}

func NewSlogLogger(level Level, opts ...Option) *SlogLogger {
	cfg := &handlerConfig{
		writer:    os.Stderr,
		format:    FormatText,
		level:     level,
		addSource: true,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &SlogLogger{
		Level:   cfg.level,
		Handler: newBaseHandler(cfg),
	}
}

func newBaseHandler(cfg *handlerConfig) slog.Handler {
	hopts := &slog.HandlerOptions{
		Level:       slogLevel(cfg.level),
		AddSource:   cfg.addSource,
		ReplaceAttr: cfg.replaceAttr,
	}
	switch cfg.format {
	case FormatJSON:
		return slog.NewJSONHandler(cfg.writer, hopts)
	default:
		return slog.NewTextHandler(cfg.writer, hopts)
	}
}

func (l *SlogLogger) Log(level Level, keyvals ...any) error {
	if level < l.Level {
		return nil
	}
	handler := l.Handler
	if handler == nil {
		handler = slog.Default().Handler()
	}
	msg, attrs, pc := slogAttrs(keyvals...)
	if pc == 0 {
		var pcs [1]uintptr
		runtime.Callers(2, pcs[:])
		pc = pcs[0]
	}
	record := slog.NewRecord(time.Now(), slogLevel(level), msg, pc)
	record.AddAttrs(attrs...)
	return handler.Handle(context.Background(), record)
}

func slogAttrs(keyvals ...any) (string, []slog.Attr, uintptr) {
	attrs := make([]slog.Attr, 0, len(keyvals)/2)
	msg := ""
	var pc uintptr
	for i := 0; i < len(keyvals); i += 2 {
		var value any = ""
		if i+1 < len(keyvals) {
			value = keyvals[i+1]
		}
		if _, ok := keyvals[i].(callerKey); ok {
			if v, ok := value.(uintptr); ok {
				pc = v
			}
			continue
		}
		key := fmt.Sprint(keyvals[i])
		if key == MessageKey {
			msg = fmt.Sprint(value)
			continue
		}
		attrs = append(attrs, slog.Any(key, value))
	}
	return msg, attrs, pc
}

func slogLevel(level Level) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelWarn:
		return slog.LevelWarn
	case LevelError, LevelFatal:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
