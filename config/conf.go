package config

import (
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/lihongsheng/payment-sdk/config/proxy"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
)

type Config struct {
	Proxy          proxy.Proxy
	Channel        channel.Channel        `json:"channel"`
	Payment        payment.Payment        `json:"payment"`
	PaymentProduct payment.PaymentProduct `json:"payment_product"`
	Config         string                 `json:"config"`
	WxConfig       *config.Config         `json:"wx_config"`
	//
}

// Option 选项函数类型
type Option func(*Config)

func WithChannel(ch channel.Channel) Option {
	return func(c *Config) {
		c.Channel = ch
	}
}
func WithPayment(pay payment.Payment) Option {
	return func(c *Config) {
		c.Payment = pay
	}
}

func WithPaymentProduct(pay payment.PaymentProduct) Option {
	return func(c *Config) {
		c.PaymentProduct = pay
	}
}

func WithWxConfig(cf *config.Config) Option {
	return func(c *Config) {
		c.WxConfig = cf
	}
}

func WithConfig(cf string) Option {
	return func(c *Config) {
		c.Config = cf
	}
}

func WithProxy(cf proxy.Proxy) Option {
	return func(c *Config) {
		c.Proxy = cf
	}
}

func NewPayment() *Config {
	return &Config{
		Channel:        channel.Channel_Wxpay,
		Payment:        payment.Payment_Wxpay,
		PaymentProduct: payment.PaymentProduct_JSAPI,
		Config:         "",
	}
}

func NewRefund() *Config {
	return &Config{
		Channel:        channel.Channel_Wxpay,
		Payment:        payment.Payment_Wxpay,
		PaymentProduct: payment.PaymentProduct_JSAPI,
		Config:         "",
	}
}
