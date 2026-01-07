package driver

import (
	"encoding/json"
	"errors"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/client"
	conf "github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	payment2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/prepay"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/refund"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/config/params"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"strings"
)

func init() {
	driver.PaymentRegister(channel.Channel_Fuiou, Payment{})
	driver.RefundRegister(channel.Channel_Fuiou, Refund{})
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
	return payment2.NewJsApi(cf, c.PaymentProduct, c.Payment)
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
	return refund.NewRefund(cf, c.Payment)
}

func initConfig(c config.Config) (*client.Client, error) {
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
	if cf.OrderPrefix == "" {
		return nil, errors2.ErrorParamError("order_prefix is empty", nil)
	}
	if cf.ApiHost == "" {
		cf.ApiHost = enum2.ApiHost
	} else {
		cf.ApiHost = strings.TrimRight(cf.ApiHost, "/")
	}
	if cf.Version == "" {
		cf.Version = enum2.Version
	}
	if c.Proxy != nil {
		cf.Proxy = *c.Proxy
	}
	cl, err := client.NewClient(cf, c.Proxy)
	if err != nil {
		return nil, err
	}
	return cl, nil
}
