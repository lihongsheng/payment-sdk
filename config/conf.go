package config

import (
	ali "github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	fuiou "github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	lakala "github.com/lihongsheng/payment-sdk/adapter/lakala/config"
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
	FuiouConfig    *fuiou.Config          `json:"fuiou_config"`
	AlipayConfig   *ali.Config            `json:"alipay_config"`
	LakalaConfig   *lakala.Config         `json:"lakala_config"`
}

// Option 选项函数类型
type Option func(*Config)

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

func WithFuiouConfig(cf *fuiou.Config) Option {
	return func(c *Config) {
		c.FuiouConfig = cf
	}
}

func WithLakalaConfig(cf *fuiou.Config) Option {
	return func(c *Config) {
		c.FuiouConfig = cf
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
		Payment:        payment.Payment_Wxpay,
		PaymentProduct: payment.PaymentProduct_JSAPI,
		Config:         "",
	}
}

func NewRefund() *Config {
	return &Config{
		Payment:        payment.Payment_Wxpay,
		PaymentProduct: payment.PaymentProduct_JSAPI,
		Config:         "",
	}
}
