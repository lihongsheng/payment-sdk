package iface

import (
	"context"
	"github.com/lihongsheng/payment-sdk/config/params"
	"net/http"

	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
)

type PaymentDriver interface {
	// Open 实例化支付接口
	Open(c config.Config) (Pay, error)
	// GetConfigOptions 获取支付配置项
	GetConfigOptions() *params.Option
	//SupportProduct(payment.PaymentProduct, enum.Device) bool
}

type Pay interface {
	Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
	Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error)
	Close(ctx context.Context, req dto.CloseQuery) error
	Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
	Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error)
}
