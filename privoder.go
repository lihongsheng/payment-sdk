package payment_sdk

import (
	_ "github.com/lihongsheng/payment-sdk/adapter/alipay/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/fuiou/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/lakala/driver"
	_ "github.com/lihongsheng/payment-sdk/adapter/wxpay/driver"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
)

func Payment(chl channel.Channel, options ...config.Option) (iface.Pay, error) {
	cf := config.NewPayment()
	for _, option := range options {
		option(cf)
	}
	if chl == channel.Channel_Channel_UNKNOWN {
		return nil, errors2.ErrorNoSupport("payment: unknown channel")
	}
	if cf.Payment == payment.Payment_Payment_UNKNOWN {
		return nil, errors2.ErrorNoSupport("payment: unknown payment method")
	}
	return driver.Payment(chl, *cf)
}

func GetPaymentDriver(channelName channel.Channel) (iface.PaymentDriver, error) {
	return driver.GetPaymentDriver(channelName)
}

func GetRefundDriver(channelName channel.Channel) (iface.RefundDriver, error) {
	return driver.GetRefundDriver(channelName)
}

func Refund(chl channel.Channel, options ...config.Option) (iface.Refund, error) {
	cf := config.NewRefund()
	for _, option := range options {
		option(cf)
	}
	if chl == channel.Channel_Channel_UNKNOWN {
		return nil, errors2.ErrorNoSupport("payment: unknown channel")
	}
	if cf.Payment == payment.Payment_Payment_UNKNOWN {
		return nil, errors2.ErrorNoSupport("refund: unknown payment method")
	}
	return driver.Refund(chl, *cf)
}
