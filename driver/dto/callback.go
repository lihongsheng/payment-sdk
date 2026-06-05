package dto

import (
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	"time"
)

type CallbackPayDetail struct {
	// 实付金额
	PayAmount Amount `json:"pay_amount"`
	// 订单号
	OrderNo string `json:"order_no"`
	// 交易单号
	TradeNo     string      `json:"trade_no"`
	Status      enum.Status `json:"status"`
	SuccessTime int64       `json:"success_time"`
	// H5 | JSAPI | NATIVE | APP | 扫码
	PaymentProduct string `json:"payment_product"`
	// 支付发卡行，不一定都有
	BankType       string                   `json:"bank_type"`
	OriginResponse string                   `json:"origin_response"`
	EventAction    enum1.Event              `json:"event_action"`
	EventRefund    *EventRefundActionParams `json:"event_refund"`
	// 正确返回的body内容
	Response string `json:"response"`
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
	OriginResponse      string               `json:"origin_response"`
	Response            string               `json:"response"`
}
