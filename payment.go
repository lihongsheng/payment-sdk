package payment_sdk

import (
	"errors"
	_ "github.com/lihongsheng/payment-sdk/adapter/wxpay"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
)

func Payment(options ...config.Option) (iface.Pay, error) {
	cf := config.NewPayment()
	for _, option := range options {
		option(cf)
	}
	if cf.Channel == channel.Channel_Channel_UNKNOWN {
		return nil, errors.New("payment: unknown channel")
	}
	if cf.Payment == payment.Payment_Payment_UNKNOWN {
		return nil, errors.New("payment: unknown payment")
	}
	return driver.Payment(cf.Channel, *cf)
}
