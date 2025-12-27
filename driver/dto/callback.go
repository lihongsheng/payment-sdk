package dto

import (
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	"time"
)

type CallbackPayDetail struct {
	// 实付金额
	PayAmount Amount
	// 订单号
	OrderNo string
	// 交易单号
	TradeNo     string
	Status      enum.Status
	SuccessTime int64
	// H5 | JSAPI | NATIVE | APP | 扫码
	PaymentProduct string
	// 支付发卡行，不一定都有
	BankType       string
	OriginResponse string
	EventAction    enum1.Event
	EventRefund    *EventRefundActionParams
}

type EventRefundActionParams struct {
	// 退款单号
	RefundNo string `json:"refund_no"`
	// 订单号
	OrderNo string `json:"order_no"`
}

type CallbackRefundDetail struct {
	// 微信返回的退款交易号
	TradeRefundNo string `json:"refund_id"`
	// 商家自己的退款单号
	RefundNo string `json:"out_refund_no"`
	// 微信支付交易单号
	TradeNo string `json:"transaction_id"`
	// 商家自己的订单号
	OrderNo string `json:"order_no"`
	//ORIGINAL: 原路退款
	//
	//BALANCE: 退回到余额
	//
	//OTHER_BALANCE: 原账户异常退到其他余额账户
	//
	//OTHER_BANKCARD: 原银行卡异常退到其他银行卡(发起异常退款成功后返回)
	Channel             refund.RefundChannel `json:"channel"`
	UserReceivedAccount string               `json:"user_received_account"`
	SuccessTime         time.Time            `json:"success_time"`
	CreateTime          time.Time            `json:"create_time"`
	Status              refund.Status        `json:"status"`
	FundsAccount        string               `json:"funds_account"`
	Amount              Amount               `json:"amount"`
	OriginResponse      string
}
