package driver

import (
	"encoding/json"
	"errors"
	conf "github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	payment2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/prepay"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/refund"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
)

type Payment struct{}

func (p Payment) Open(c config.Config) (iface.Pay, error) {
	if c.PaymentProduct == payment.PaymentProduct_PaymentMethod_UNKNOWN {
		return nil, errors.New("payment: unknown payment product")
	}
	var cf conf.Config
	if c.FuiouConfig != nil {
		cf = *c.FuiouConfig
	} else {
		if c.Config == "" {
			return nil, errors.New("payment: config is empty")
		}
		err := json.Unmarshal([]byte(c.Config), &cf)
		if err != nil {
			return nil, err
		}
	}
	if cf.OrderPrefix == "" {
		return nil, errors2.ErrorParamError("order_prefix is empty", nil)
	}
	return payment2.NewJsApi(cf, c.PaymentProduct, c.Payment)
}

func init() {
	driver.PaymentRegister(channel.Channel_Fuiou.String(), Payment{})
	driver.RefundRegister(channel.Channel_Fuiou.String(), Refund{})
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
	var cf conf.Config
	if c.FuiouConfig != nil {
		cf = *c.FuiouConfig
	} else {
		if c.Config == "" {
			return nil, errors.New("payment: config is empty")
		}
		err := json.Unmarshal([]byte(c.Config), &cf)
		if err != nil {
			return nil, err
		}
	}
	return refund.NewRefund(cf, c.Payment)
}
