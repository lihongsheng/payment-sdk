package payment

import (
	"context"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client"
	"github.com/lihongsheng/payment-sdk/driver/iface"

	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
)

type Lite struct {
	*Jsapi
}

func NewLite(api *client.Api) (iface.Pay, error) {
	api2, err := newJsApi(api)
	if err != nil {
		return nil, err
	}
	return &Lite{
		Jsapi: api2,
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
