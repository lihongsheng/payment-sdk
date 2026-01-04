package driver

import (
	"encoding/json"
	"errors"
	conf "github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	payment2 "github.com/lihongsheng/payment-sdk/adapter/wxpay/payment"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/refund"
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
		cf = *c.WxConfig
	} else {
		if c.Config == "" {
			return nil, errors.New("payment: config is empty")
		}
		err := json.Unmarshal([]byte(c.Config), &cf)
		if err != nil {
			return nil, err
		}
	}
	switch c.PaymentProduct {
	case payment.PaymentProduct_JSAPI:
		return payment2.NewJsApi(cf)
	case payment.PaymentProduct_LITE:
		return payment2.NewLite(cf)
	case payment.PaymentProduct_H5:
		return payment2.NewH5(cf)
	case payment.PaymentProduct_Qrcode:
		return payment2.NewNative(cf)
	case payment.PaymentProduct_APP:
		return payment2.NewApp(cf)
	}
	return nil, errors.New("payment: unknown payment product")
}

func init() {
	driver.PaymentRegister(channel.Channel_Wxpay.String(), Payment{})
	driver.RefundRegister(channel.Channel_Wxpay.String(), Refund{})
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
	var cf conf.Config
	if c.WxConfig != nil {
		cf = *c.WxConfig
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
