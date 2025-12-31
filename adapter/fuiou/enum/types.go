package enum

import (
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"strings"
)

const (
	Version      = "1.0"
	ApiHost      = "https://aipay-cloud.fuioupay.com"
	OrderPrefix  = "F"
	TransferHost = "https://richfront.fuioupay.com"
)

var (
	WxPaymentProductMap = map[enum.PaymentProduct]string{
		enum.PaymentProduct_JSAPI: JSAPI,
		enum.PaymentProduct_LITE:  LETPAY,
	}
	AliPaymentProductMap = map[enum.PaymentProduct]string{
		enum.PaymentProduct_JSAPI: FWC,
		enum.PaymentProduct_LITE:  FWC,
	}
)

func GenOrder(orderPrefix string, orderNo string) string {
	return orderPrefix + OrderPrefix + orderNo
}

func ParseOrder(orderPrefix string, orderNo string) string {
	return strings.Replace(orderNo, orderPrefix+OrderPrefix, "", 1)
}
