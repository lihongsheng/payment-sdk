package dto

import (
	refund "github.com/lihongsheng/payment-sdk/enum/refund"
	"time"
)

type RefundRequest struct {
	// 退款单号，每笔订单唯一
	RefundNo string
	// 商户返回的支付交易号
	TradeNo string
	// 订单号
	OrderNo string
	// 退款原因
	Reason string
	// 退款回调地址
	NotifyUrl string
	Amount    Amount
	Goods     []Goods
	// 订单支付金额，微信要求必传
	OrderAmount Amount
}

type RefundQuery struct {
	// 商家自己的退款单号
	RefundNo string
	OrderNo  string
	TradeNo  string
}

type RefundDetail struct {
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
