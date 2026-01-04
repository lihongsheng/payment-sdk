package payment_sdk

import (
	_ "github.com/lihongsheng/payment-sdk/adapter/alipay/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/fuiou/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/lakala/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/wxpay/driver"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
)

func Payment(channelName string, options ...config.Option) (iface.Pay, error) {
	cf := config.NewPayment()
	for _, option := range options {
		option(cf)
	}
	if channelName == "" {
		return nil, errors2.ErrorNoSupport("payment: unknown channel")
	}
	if cf.Payment == payment.Payment_Payment_UNKNOWN {
		return nil, errors2.ErrorNoSupport("payment: unknown payment method")
	}
	return driver.Payment(channelName, *cf)
}

func Refund(channelName string, options ...config.Option) (iface.Refund, error) {
	cf := config.NewRefund()
	for _, option := range options {
		option(cf)
	}
	if channelName == "" {
		return nil, errors2.ErrorNoSupport("refund: unknown channel")
	}
	if cf.Payment == payment.Payment_Payment_UNKNOWN {
		return nil, errors2.ErrorNoSupport("refund: unknown payment method")
	}
	return driver.Refund(channelName, *cf)
}
