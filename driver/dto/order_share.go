package dto

import (
	shareEnum "github.com/lihongsheng/payment-sdk/enum/order_share"
)

type CreateShareOrderRequest struct {
	// 分账流水号，支付宝需要。
	ShareOrderNo string `json:"share_order_no"`
	// 微信等订单号
	TradeNo string `json:"trade_no"`
	// 分账接收方，用户
	Users []ShareUserReceiveRequest `json:"receivers"`
	// 分账接收方，商户
	Mch []ShareMchReceiveRequest `json:"mch_receivers"`
	// 是否解冻剩余,微信才有
	UnfreezeUnsplit bool `json:"unfreeze_unsplit"`
}

type ShareUserReceiveRequest struct {
	User   User   `json:"user"`
	Amount Amount `json:"amount"`
	Remark string `json:"remark"`
}

type ShareMchReceiveRequest struct {
	Account     string `json:"account"`
	AccountName string `json:"name"`
	Amount      Amount `json:"amount"`
	Remark      string `json:"remark"`
}

type ShareOrderResponse struct {
	// 分账流水号，支付宝需要。
	ShareOrderNo string `json:"share_order_no"`
	// 微信等订单号
	TradeNo string `json:"trade_no"`
	// 微信等分账单号
	ThirdShareOrderNo string `json:"third_share_order_no"`
	Status            shareEnum.Status
}

type ShareOrderQueryRequest struct {
	// 分账流水号，支付宝需要。
	ShareOrderNo string `json:"share_order_no"`
	// 微信等分账单号
	ThirdShareOrderNo string `json:"third_share_order_no"`
}

type ShareOrderDetailResponse struct {
	// 分账流水号，支付宝需要。
	ShareOrderNo string `json:"share_order_no"`
	// 微信等订单号
	TradeNo string `json:"trade_no"`
	// 微信等分账单号
	ThirdShareOrderNo string `json:"third_share_order_no"`
	Status            shareEnum.Status
	Receivers         []ShareUserReceive `json:"receivers"`
	MchReceivers      []ShareMchReceive  `json:"mch_receivers"`
}

type ShareUserReceive struct {
	User         User   `json:"user"`
	Amount       Amount `json:"amount"`
	Remark       string `json:"remark"`
	Status       shareEnum.ShareOrderStaus
	FailReason   string `json:"fail_reason"`
	ShareOrderNO string `json:"share_order_no"`
}

type ShareMchReceive struct {
	Account      string `json:"account"`
	Amount       Amount `json:"amount"`
	Remark       string `json:"remark"`
	Status       shareEnum.ShareOrderStaus
	FailReason   string `json:"fail_reason"`
	ShareOrderNO string `json:"share_order_no"`
}

//type T struct {
//	TransactionId string `json:"transaction_id"`
//	OutOrderNo    string `json:"out_order_no"`
//	OrderId       string `json:"order_id"`
//	State         string `json:"state"`
//	Receivers     []struct {
//		Amount      int       `json:"amount"`
//		Description string    `json:"description"`
//		Type        string    `json:"type"`
//		Account     string    `json:"account"`
//		Result      string    `json:"result"`
//		FailReason  string    `json:"fail_reason"`
//		CreateTime  time.Time `json:"create_time"`
//		FinishTime  time.Time `json:"finish_time"`
//		DetailId    string    `json:"detail_id"`
//	} `json:"receivers"`
//}
