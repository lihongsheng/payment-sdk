package enum

import "github.com/lihongsheng/payment-sdk/enum/payment"

const (
	Version = "3.0"
	ApiHost = "https://s2.lakala.com"
)

const (
	// INIT-初始化 CREATE-下单成功 SUCCESS-交易成功 FAIL-交易失败 DEAL-交易处理中 UNKNOWN-未知状态 CLOSE-订单关闭 PART_REFUND-部分退款 REFUND-全部退款(或订单被撤销）
	PaymentTrade_INIT        = "INIT"
	PaymentTrade_CREATE      = "CREATE"
	PaymentTrade_SUCCESS     = "SUCCESS"
	PaymentTrade_FAIL        = "FAIL"
	PaymentTrade_DEAL        = "DEAL"
	PaymentTrade_UNKNOWN     = "UNKNOWN"
	PaymentTrade_CLOSE       = "CLOSE"
	PaymentTrade_PART_REFUND = "PART_REFUND"
	PaymentTrade_REFUND      = "REFUND"
)

func GetPaymentStatus(state string) payment.Status {
	switch state {
	case PaymentTrade_INIT, PaymentTrade_CREATE, PaymentTrade_DEAL:
		return payment.Status_Pending
	case PaymentTrade_SUCCESS:
		return payment.Status_Success
	case PaymentTrade_FAIL:
		return payment.Status_Failed
	case PaymentTrade_CLOSE:
		return payment.Status_Close
	case PaymentTrade_PART_REFUND, PaymentTrade_REFUND:
		return payment.Status_Refund
	}
	return payment.Status_Status_UNKNOWN
}

var PaymentMap = map[payment.Payment]string{
	payment.Payment_Wechat: "WECHAT",
	payment.Payment_Alipay: "ALIPAY",
}

var ProductMap = map[payment.PaymentProduct]string{
	payment.PaymentProduct_JSAPI: "51",
	payment.PaymentProduct_LITE:  "71",
	payment.PaymentProduct_APP:   "61",
}
