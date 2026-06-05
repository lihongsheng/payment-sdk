package driver

import (
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/client"
	conf "github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	payment2 "github.com/lihongsheng/payment-sdk/adapter/alipay/payment"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/refund"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/errors"
)

func init() {
	driver.PaymentRegister(channel.Channel_Alipay, Payment{})
	driver.RefundRegister(channel.Channel_Alipay, Refund{})
}

type Payment struct{}

func (p Payment) Open(c config.Config) (iface.Pay, error) {
	if c.PaymentProduct == payment.PaymentProduct_PaymentMethod_UNKNOWN {
		return nil, errors.ErrorParamError("payment: unknown payment product")
	}
	cl, err := initConfig(c)
	if err != nil {
		return nil, err
	}
	switch c.PaymentProduct {
	case payment.PaymentProduct_JSAPI, payment.PaymentProduct_LITE:
		return payment2.NewJsApi(cl)
	case payment.PaymentProduct_H5:
		return payment2.NewH5(cl)
	case payment.PaymentProduct_Qrcode:
		return payment2.NewQrcode(cl)
	case payment.PaymentProduct_APP:
		return payment2.NewApp(cl)
	}
	return nil, errors.ErrorParamError("alipay-payment: unknown payment product")
}

func (p Payment) GetConfigOptions() *iface.ChannelOption {
	return &iface.ChannelOption{
		Channel: channel.Channel_Alipay.String(),
		Label:   "支付宝支付",
		Options: options,
	}
}

func (p Payment) IsSupportPayment(product payment.PaymentProduct, device enum.Device) bool {
	switch product {
	case payment.PaymentProduct_JSAPI:
		switch device {
		case enum.Device_Alipay, enum.Device_Alipay_Lite:
			return true
		}
	case payment.PaymentProduct_LITE:
		switch device {
		case enum.Device_Alipay, enum.Device_Alipay_Lite:
			return true
		}
	case payment.PaymentProduct_Qrcode:
		switch device {
		case enum.Device_PC:
			return true
		}
	case payment.PaymentProduct_H5:
		switch device {
		case enum.Device_H5:
			return true
		}
	case payment.PaymentProduct_APP:
		switch device {
		case enum.Device_APP:
			return true
		}
	}
	return false
}

func (p Payment) GetSupportProduct() []iface.PaymentMethod {
	return []iface.PaymentMethod{
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
				{
					Product: payment.PaymentProduct_H5.String(),
					Label:   "H5（手机H5支付）",
				},
				{
					Product: payment.PaymentProduct_Qrcode.String(),
					Label:   "二维码支付",
				},
			},
		},
	}
}

func (p Payment) CallbackResponse() string {
	return "success"
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
		cf = *c.AlipayConfig
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
	return "success"
}
