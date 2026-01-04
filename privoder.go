package payment_sdk

import (
	"errors"
	_ "github.com/lihongsheng/payment-sdk/adapter/alipay/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/fuiou/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/lakala/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/wxpay/driver"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/payment"
)

func Payment(channelName string, options ...config.Option) (iface.Pay, error) {
	cf := config.NewPayment()
	for _, option := range options {
		option(cf)
	}
	if channelName == "" {
		return nil, errors.New("payment: unknown channel")
	}
	if cf.Payment == payment.Payment_Payment_UNKNOWN {
		return nil, errors.New("payment: unknown payment")
	}
	return driver.Payment(channelName, *cf)
}

func Refund(channelName string, options ...config.Option) (iface.Refund, error) {
	cf := config.NewRefund()
	for _, option := range options {
		option(cf)
	}
	if channelName == "" {
		return nil, errors.New("refund: unknown channel")
	}
	if cf.Payment == payment.Payment_Payment_UNKNOWN {
		return nil, errors.New("payment: unknown payment")
	}
	return driver.Refund(channelName, *cf)
}
