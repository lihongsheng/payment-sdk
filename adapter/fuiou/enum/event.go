package enum

type EventType string

const (
	// alipay.fund.trans.app.pay
	EventTypeUnKnown EventType = "UNKNOWN"
	// alipay.fund.trans.order.changed
	EventTypeUserChanged     EventType = "accountInResult"
	EventTypeTransferChanged EventType = "accountInSettleResult"
)
