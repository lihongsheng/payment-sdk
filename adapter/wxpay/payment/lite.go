package payment

import (
	"context"

	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
)

type Lite struct {
	*Jsapi
}

func NewLite(conf config.Config) (driver.Pay, error) {
	api, err := wxpay.InitClient(conf)
	if err != nil {
		return nil, err
	}
	return &Lite{
		Jsapi: &Jsapi{
			Api: api,
		},
	}, nil
}

func (l *Lite) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	result, err := l.Jsapi.Pay(ctx, req)
	if err != nil {
		return nil, err
	}
	result.PaymentProduct = enum.PaymentProduct_LITE.String()
	return result, nil
}
func (l *Lite) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	result, err := l.Jsapi.Query(ctx, req)
	if err != nil {
		return nil, err
	}
	result.PaymentProduct = enum.PaymentProduct_LITE.String()
	return result, nil
}
