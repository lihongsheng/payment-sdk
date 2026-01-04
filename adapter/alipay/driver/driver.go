package driver

import (
	"encoding/json"
	"errors"
	conf "github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	payment2 "github.com/lihongsheng/payment-sdk/adapter/alipay/payment"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/refund"
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
	if c.WxConfig != nil {
		cf = *c.AlipayConfig
	} else {
		if c.Config == "" {
			return nil, errors.New("payment: config is empty")
		}
		err := json.Unmarshal([]byte(c.Config), &cf)
		if err != nil {
			return nil, err
		}
	}
	return payment2.NewJsApi(cf)
}

func init() {
	driver.PaymentRegister(channel.Channel_Alipay.String(), Payment{})
	driver.RefundRegister(channel.Channel_Alipay.String(), Refund{})
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
	var cf conf.Config
	if c.WxConfig != nil {
		cf = *c.AlipayConfig
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
