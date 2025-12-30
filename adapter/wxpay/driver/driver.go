package driver

import (
  "encoding/json"
  "errors"
  conf "github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
  payment2 "github.com/lihongsheng/payment-sdk/adapter/wxpay/payment"
  "github.com/lihongsheng/payment-sdk/config"
  "github.com/lihongsheng/payment-sdk/driver/iface"
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
  case payment.PaymentProduct_JSAPI, payment.PaymentProduct_LITE:
    return payment2.NewJsApi(cf)
  case payment.PaymentProduct_H5:
    return payment2.NewH5(cf)
  }
  return nil, errors.New("payment: unknown payment product")
}
