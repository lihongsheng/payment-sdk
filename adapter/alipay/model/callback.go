package model

import (
	"encoding/json"
	"github.com/singer-stack-lab/payment-sdk/adapter/alipay/enum"
	"github.com/singer-stack-lab/payment-sdk/adapter/alipay/util"
	"time"
)

// AlipayNotify 支付宝通知参数结构体
type AlipayNotifyBody struct {
	NotifyTime        string `json:"notify_time"`                   // 通知时间，格式为 yyyy-MM-dd HH:mm:ss
	NotifyType        string `json:"notify_type"`                   // 通知类型，枚举值：trade_status_sync
	NotifyId          string `json:"notify_id"`                     // 通知校验 ID
	SignType          string `json:"sign_type"`                     // 签名类型，支持 RSA2、RSA
	Sign              string `json:"sign"`                          // 签名
	TradeNo           string `json:"trade_no"`                      // 支付宝交易号
	AppId             string `json:"app_id"`                        // 开发者的 app_id
	AuthAppId         string `json:"auth_app_id"`                   // 开发者的 app_id（服务商调用场景）
	OutTradeNo        string `json:"out_trade_no"`                  // 商户订单号
	OutBizNo          string `json:"out_biz_no"`                    // 商家业务号（可选）
	BuyerId           string `json:"buyer_id"`                      // 买家支付宝用户号（可选）
	BuyerLogonId      string `json:"buyer_logon_id"`                // 买家支付宝账号（可选）
	SellerId          string `json:"seller_id"`                     // 卖家支付宝用户号（可选）
	SellerEmail       string `json:"seller_email"`                  // 卖家支付宝账号（可选）
	TradeStatus       string `json:"trade_status"`                  // 交易状态
	TotalAmount       string `json:"total_amount"`                  // 订单金额（元）
	ReceiptAmount     string `json:"receipt_amount,omitempty"`      // 实收金额，单位：元，支持小数点后两位
	InvoiceAmount     string `json:"invoice_amount,omitempty"`      // 开票金额，单位：元，支持小数点后两位
	BuyerPayAmount    string `json:"buyer_pay_amount,omitempty"`    // 付款金额，单位：元，支持小数点后两位
	PointAmount       string `json:"point_amount,omitempty"`        // 集分宝金额，单位：元，支持小数点后两位
	RefundFee         string `json:"refund_fee,omitempty"`          // 总退款金额，单位：元，支持小数点后两位
	SendBackFee       string `json:"send_back_fee,omitempty"`       // 实际退款金额，单位：元，支持小数点后两位
	Subject           string `json:"subject,omitempty"`             // 订单标题
	Body              string `json:"body,omitempty"`                // 商品描述
	GmtCreate         string `json:"gmt_create,omitempty"`          // 交易创建时间，格式：yyyy-MM-dd HH:mm:ss
	GmtPayment        string `json:"gmt_payment,omitempty"`         // 交易付款时间，格式：yyyy-MM-dd HH:mm:ss
	GmtRefund         string `json:"gmt_refund,omitempty"`          // 交易退款时间，格式：yyyy-MM-dd HH:mm:ss.SS
	GmtClose          string `json:"gmt_close,omitempty"`           // 交易结束时间，格式：yyyy-MM-dd HH:mm:ss
	FundBillList      string `json:"fund_bill_list,omitempty"`      // 支付金额信息
	VoucherDetailList string `json:"voucher_detail_list,omitempty"` // 优惠券信息
	// biz_settle_mode
	BizSettleMode string `json:"biz_settle_mode,omitempty"`
}

func (a AlipayNotifyBody) IsRefund() bool {
	if (a.TradeStatus == enum.PAYMENT_STATUS_TRADE_SUCCESS || a.TradeStatus == enum.PAYMENT_STATUS_TRADE_FINISHED) && a.OutBizNo != "" {
		return true
	}
	return false
}

func (a AlipayNotifyBody) GetTotalAmount() int64 {
	amount, _ := util.AmountToCents(a.TotalAmount)
	return int64(amount)
}

func (a AlipayNotifyBody) GetReceiptAmount() int64 {
	amount, _ := util.AmountToCents(a.ReceiptAmount)
	return int64(amount)
}

func (a AlipayNotifyBody) GetBuyerPayAmount() int64 {
	amount, _ := util.AmountToCents(a.BuyerPayAmount)
	return int64(amount)
}
func (a AlipayNotifyBody) GetInvoiceAmount() int64 {
	amount, _ := util.AmountToCents(a.InvoiceAmount)
	return int64(amount)
}
func (a AlipayNotifyBody) GetPointAmount() int64 {
	amount, _ := util.AmountToCents(a.PointAmount)
	return int64(amount)
}
func (a AlipayNotifyBody) GetRefundFee() int64 {
	amount, _ := util.AmountToCents(a.RefundFee)
	return int64(amount)
}

func (a AlipayNotifyBody) GetNotifyTime() time.Time {
	if a.NotifyTime != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", a.NotifyTime)
		return t
	}
	return time.Time{}
}

func (a AlipayNotifyBody) GetGmtCreate() time.Time {
	if a.GmtCreate != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", a.GmtCreate)
		return t
	}
	return time.Time{}
}

func (a AlipayNotifyBody) GetGmtPayment() time.Time {
	if a.GmtPayment != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", a.GmtPayment)
		return t
	}
	return time.Time{}
}
func (a AlipayNotifyBody) GetGmtRefund() time.Time {
	if a.GmtRefund != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", a.GmtRefund)
		return t
	}
	return time.Time{}
}
func (a AlipayNotifyBody) GetGmtClose() time.Time {
	if a.GmtClose != "" {
		t, _ := time.Parse("2006-01-02 15:04:05", a.GmtClose)
		return t
	}
	return time.Time{}
}

func (a AlipayNotifyBody) GetFundBillList() []NotifyFundBill {
	var list []NotifyFundBill
	if a.FundBillList != "" {
		_ = json.Unmarshal([]byte(a.FundBillList), &list)
	}
	return list
}
func (a AlipayNotifyBody) GetVoucherDetailList() []NotifyVoucherDetail {
	var list []NotifyVoucherDetail
	if a.VoucherDetailList != "" {
		_ = json.Unmarshal([]byte(a.VoucherDetailList), &list)
	}
	return list
}

type NotifyFundBill struct {
	Amount      string `json:"amount"`
	FundChannel string `json:"fundChannel"`
}

type NotifyVoucherDetail struct {
	Amount             string `json:"amount"`
	MerchantContribute string `json:"merchantContribute"`
	Name               string `json:"name"`
	OtherContribute    string `json:"otherContribute"`
	Type               string `json:"type"`
	Memo               string `json:"memo"`
}

type AlipayCallbackEventBody struct {
	// utc_timestamp
	UtcTimestamp string `json:"utc_timestamp"`
	// biz_content
	BizContent string `json:"biz_content"`
	// sign
	Sign string `json:"sign"`
	// app_id
	AppId string `json:"app_id"`
	// version
	Version string `json:"version"`
	// sign_type
	SignType string `json:"sign_type"`
	// notify_id
	NotifyId string `json:"notify_id"`
	// msg_method
	MsgMethod string `json:"msg_method"`
	// merchant_app_id
	MerchantAppId string `json:"merchant_app_id"`
}

type AlipayTransferOrderAmountChangeEvent struct {
	// out_biz_no
	OutBizNo string `json:"out_biz_no"`
	// product_code
	ProductCode string `json:"product_code"`
	// biz_scene
	BizScene string `json:"biz_scene"`
	// origin_interface
	OriginInterface string `json:"origin_interface"`
	// pay_fund_order_id
	PayFundOrderId string `json:"pay_fund_order_id"`
	// order_id
	OrderId string `json:"order_id"`
	// status
	Status string `json:"status"`
	// sub_status
	SubStatus string `json:"sub_status"`
	// receiver_open_id
	ReceiverOpenId string `json:"receiver_open_id"`
	// receiver_user_id
	ReceiverUserId string `json:"receiver_user_id"`
	// action_type
	ActionType string `json:"action_type"`
	// trans_amount
	TransAmount string `json:"trans_amount"`
	// pay_date
	PayDate string `json:"pay_date"`
	// refund_date
	RefundDate string `json:"refund_date"`
	// entrust_order_id
	EntrustOrderId string `json:"entrust_order_id"`
	// settle_serial_no
	SettleSerialNo string `json:"settle_serial_no"`
	// sub_order_status
	SubOrderStatus string `json:"sub_order_status"`
	// sub_order_error_code
	SubOrderErrorCode string `json:"sub_order_error_code"`
	// sub_order_fail_reason
	SubOrderFailReason string `json:"sub_order_fail_reason"`
}
