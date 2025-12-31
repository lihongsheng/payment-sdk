package refund

import "time"

type CallbackRefund struct {
	Mchid               string       `json:"mchid"`
	OutTradeNo          string       `json:"out_trade_no"`
	TransactionId       string       `json:"transaction_id"`
	OutRefundNo         string       `json:"out_refund_no"`
	RefundId            string       `json:"refund_id"`
	RefundStatus        string       `json:"refund_status"`
	SuccessTime         time.Time    `json:"success_time"`
	Amount              RefundAmount `json:"amount"`
	UserReceivedAccount string       `json:"user_received_account"`
}

type RefundAmount struct {
	Total       int `json:"total"`
	Refund      int `json:"refund"`
	PayerTotal  int `json:"payer_total"`
	PayerRefund int `json:"payer_refund"`
}
