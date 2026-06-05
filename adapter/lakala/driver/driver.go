package driver

import (
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/client"
	conf "github.com/lihongsheng/payment-sdk/adapter/lakala/config"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/prepay"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/refund"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/errors"
)

func init() {
	driver.PaymentRegister(channel.Channel_Lakala, Payment{})
	driver.RefundRegister(channel.Channel_Lakala, Refund{})
}

type Payment struct{}

func (p Payment) Open(c config.Config) (iface.Pay, error) {
	if c.PaymentProduct == payment.PaymentProduct_PaymentMethod_UNKNOWN {
		return nil, errors.ErrorParamError("payment: unknown payment product")
	}
	cf, err := initConfig(c)
	if err != nil {
		return nil, err
	}
	return prepay.NewPay(cf, c.PaymentProduct, c.Payment)
}
func (p Payment) GetConfigOptions() *iface.ChannelOption {
	return &iface.ChannelOption{
		Channel: channel.Channel_Lakala.String(),
		Label:   "拉卡拉支付",
		Options: options,
	}
}

func (p Payment) IsSupportPayment(product payment.PaymentProduct, device enum.Device) bool {
	switch product {
	case payment.PaymentProduct_JSAPI:
		switch device {
		case enum.Device_Wechat, enum.Device_Wechat_Lite, enum.Device_Alipay, enum.Device_Alipay_Lite:
			return true
		}
	case payment.PaymentProduct_LITE:
		switch device {
		case enum.Device_Wechat, enum.Device_Wechat_Lite, enum.Device_Alipay, enum.Device_Alipay_Lite:
			return true
		}
	}
	return false
}

func (p Payment) GetSupportProduct() []iface.PaymentMethod {
	return []iface.PaymentMethod{
		{
			Method: payment.Payment_Wechat.String(),
			Label:  "微信支付",
			Product: []iface.PaymentProduct{
				{
					Product: payment.PaymentProduct_JSAPI.String(),
					Label:   "公众号支付(JSAPI)",
				},
				{
					Product: payment.PaymentProduct_LITE.String(),
					Label:   "小程序支付",
				},
			},
		},
		{
			Method: payment.Payment_Alipay.String(),
			Label:  "支付宝支付",
			Product: []iface.PaymentProduct{
				{
					Product: payment.PaymentProduct_JSAPI.String(),
					Label:   "公众号支付(JSAPI)",
				},
				{
					Product: payment.PaymentProduct_LITE.String(),
					Label:   "小程序支付",
				},
			},
		},
	}
}

func (p Payment) CallbackResponse() string {
	return ""
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
			return nil, errors.ErrorParamError("payment: config is empty")
		}
		err := json.Unmarshal([]byte(c.Config), &cf)
		if err != nil {
			return nil, errors.ErrorParamError("parse config err: %v", err)
		}
	}

	cl, err := client.NewClient(cf, c.Proxy)
	if err != nil {
		return nil, err
	}
	return cl, nil
}

func (p Refund) CallbackResponse() string {
	return ""
}
