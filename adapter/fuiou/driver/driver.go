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

func init() {
	driver.PaymentRegister(channel.Channel_Fuiou.String(), Payment{})
	driver.RefundRegister(channel.Channel_Fuiou.String(), Refund{})
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
	if cf.OrderPrefix == "" {
		return nil, errors2.ErrorParamError("order_prefix is empty", nil)
	}
	return payment2.NewJsApi(*cf, c.PaymentProduct, c.Payment)
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
	cf, err := initConfig(c)
	if err != nil {
		return nil, err
	}
	return refund.NewRefund(*cf, c.Payment)
}

func initConfig(c config.Config) (*conf.Config, error) {
	var cf conf.Config
	if c.LakalaConfig != nil {
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
	if c.Proxy != nil {
		cf.Proxy = *c.Proxy
	}
	return &cf, nil
}
