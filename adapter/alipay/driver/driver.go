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

func init() {
	driver.PaymentRegister(channel.Channel_Alipay, Payment{})
	driver.RefundRegister(channel.Channel_Alipay, Refund{})
}

type Payment struct{}

func (p Payment) Open(c config.Config) (iface.Pay, error) {
	if c.PaymentProduct == payment.PaymentProduct_PaymentMethod_UNKNOWN {
		return nil, errors.New("payment: unknown payment product")
	}
	cf, err := initConfig(c)
	if err != nil {
		return nil, err
	}
	switch c.PaymentProduct {
	case payment.PaymentProduct_JSAPI:
		return payment2.NewJsApi(*cf)
	case payment.PaymentProduct_LITE:
		return payment2.NewJsApi(*cf)
	case payment.PaymentProduct_H5:
		return payment2.NewH5(*cf)
	case payment.PaymentProduct_Qrcode:
		return payment2.NewQrcode(*cf)
	case payment.PaymentProduct_APP:
		return payment2.NewApp(*cf)
	}
	return nil, errors.New("alipay-payment: unknown payment product")
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
	cf, err := initConfig(c)
	if err != nil {
		return nil, err
	}
	return refund.NewRefund(*cf)
}

func initConfig(c config.Config) (*conf.Config, error) {
	var cf conf.Config
	if c.LakalaConfig != nil {
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
	if c.Proxy != nil {
		cf.Proxy = *c.Proxy
	}
	return &cf, nil
}
