package log

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
)

func TestSlogLogger_TextFormat(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(LevelInfo, WithWriter(&buf), WithFormat(FormatText))

	err := logger.Log(LevelInfo, MessageKey, "hello", FieldKeyChannel, "wxpay")
	if err != nil {
		t.Fatal(err)
	}

	out := buf.String()
	for _, want := range []string{"time=", "level=INFO", "source=", "msg=hello", "channel=wxpay"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in output %q", want, out)
		}
	}
}

func TestSlogLogger_JSONFormat(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(LevelInfo, WithWriter(&buf), WithFormat(FormatJSON))

	err := logger.Log(LevelError, MessageKey, "failed", FieldKeyError, "boom")
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]any
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got[slog.TimeKey] == "" {
		t.Fatalf("expected time field in output %v", got)
	}
	if got[slog.LevelKey] != "ERROR" {
		t.Fatalf("expected ERROR level, got %v", got[slog.LevelKey])
	}
	if got[slog.SourceKey] == nil {
		t.Fatalf("expected source field in output %v", got)
	}
	if got[MessageKey] != "failed" {
		t.Fatalf("expected failed msg, got %v", got[MessageKey])
	}
	if got[FieldKeyError] != "boom" {
		t.Fatalf("expected boom error, got %v", got[FieldKeyError])
	}
	t.Log(got)
}

func TestLogger_WithAndWithContext(t *testing.T) {
	type traceKey struct{}

	var buf bytes.Buffer
	base := NewSlogLogger(LevelDebug, WithWriter(&buf), WithAddSource(false))
	logger := WithContext(context.WithValue(context.Background(), traceKey{}, "trace-1"), With(base, FieldKeyChannel, "alipay"))

	err := logger.Log(LevelInfo, MessageKey, "with context", "trace_id", Valuer(func(ctx context.Context) any {
		return ctx.Value(traceKey{})
	}))
	if err != nil {
		t.Fatal(err)
	}

	out := buf.String()
	for _, want := range []string{"level=INFO", "msg=\"with context\"", "channel=alipay", "trace_id=trace-1"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in output %q", want, out)
		}
	}
	if strings.Contains(out, "source=") {
		t.Fatalf("did not expect source field in output %q", out)
	}
}

func TestInfoSourceUsesCallSite(t *testing.T) {
	var buf bytes.Buffer
	old := GetLogger()
	SetLogger(NewSlogLogger(LevelInfo, WithWriter(&buf)))
	defer SetLogger(old)

	Info(context.Background(), "call site", F(FieldKeyChannel, "wxpay"))

	out := buf.String()
	if !strings.Contains(out, "logger_test.go") {
		t.Fatalf("expected source to use log.Info call site, got %q", out)
	}
	if strings.Contains(out, "logger.go") {
		t.Fatalf("expected source not to use logger internals, got %q", out)
	}
}

func TestSlogLogger_LevelFilter(t *testing.T) {
	var buf bytes.Buffer
	logger := NewSlogLogger(LevelWarn, WithWriter(&buf))

	if err := logger.Log(LevelInfo, MessageKey, "ignore"); err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 0 {
		t.Fatalf("expected empty output, got %q", buf.String())
	}

	if err := logger.Log(LevelWarn, MessageKey, "warn"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "level=WARN") {
		t.Fatalf("expected warn output, got %q", buf.String())
	}
}
