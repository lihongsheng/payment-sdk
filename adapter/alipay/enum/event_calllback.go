package enum

type EventType string

const (
	// alipay.fund.trans.app.pay
	EventTypeUnKnown EventType = "UNKNOWN"
	// alipay.fund.trans.order.changed
	EventTypeOrderChanged EventType = "alipay.fund.trans.order.changed"
)
