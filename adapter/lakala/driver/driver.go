package driver

import (
	"encoding/json"
	"errors"
	conf "github.com/lihongsheng/payment-sdk/adapter/lakala/config"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/prepay"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/refund"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
)

type Payment struct{}

func (p Payment) Open(c config.Config) (iface.Pay, error) {
	if c.PaymentProduct == payment.PaymentProduct_PaymentMethod_UNKNOWN {
		return nil, errors.New("payment: unknown payment product")
	}
	var cf conf.Config
	if c.LakalaConfig != nil {
		cf = *c.LakalaConfig
	} else {
		if c.Config == "" {
			return nil, errors.New("payment: config is empty")
		}
		err := json.Unmarshal([]byte(c.Config), &cf)
		if err != nil {
			return nil, err
		}
	}
	return prepay.NewPay(cf, c.PaymentProduct, c.Payment)
}

func init() {
	driver.PaymentRegister(channel.Channel_Wxpay, Payment{})
	driver.RefundRegister(channel.Channel_Wxpay, Refund{})
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
	var cf conf.Config
	if c.LakalaConfig != nil {
		cf = *c.LakalaConfig
	} else {
		if c.Config == "" {
			return nil, errors.New("payment: config is empty")
		}
		err := json.Unmarshal([]byte(c.Config), &cf)
		if err != nil {
			return nil, err
		}
	}
	return refund.NewRefund(cf)
}
