package driver

import (
	"encoding/json"
	"errors"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/client"
	conf "github.com/lihongsheng/payment-sdk/adapter/lakala/config"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/prepay"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/refund"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/config/params"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
)

func init() {
	driver.PaymentRegister(channel.Channel_Lakala, Payment{})
	driver.RefundRegister(channel.Channel_Lakala, Refund{})
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
	return prepay.NewPay(cf, c.PaymentProduct, c.Payment)
}
func (p Payment) GetConfigOptions() *params.Option {
	return options
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
	cf, err := initConfig(c)
	if err != nil {
		return nil, err
	}
	return refund.NewRefund(cf)
}

func initConfig(c config.Config) (*client.Client, error) {
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

	cl, err := client.NewClient(cf, c.Proxy)
	if err != nil {
		return nil, err
	}
	return cl, nil
}
