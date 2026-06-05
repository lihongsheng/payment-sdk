package iface

import (
	"context"
	"github.com/lihongsheng/payment-sdk/config/params"
	"github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"net/http"

	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
)

type PaymentMethod struct {
	Method  string           `json:"method"`
	Label   string           `json:"label"`
	Product []PaymentProduct `json:"product"`
}

type PaymentProduct struct {
	Product string `json:"product"`
	Label   string `json:"label"`
}

type ChannelOption struct {
	Channel string `json:"channel"`
	Label   string `json:"label"`
	Options []params.Option
}
type PaymentDriver interface {
	// Open 实例化支付接口
	Open(c config.Config) (Pay, error)
	// GetConfigOptions 获取支付配置项
	GetConfigOptions() *ChannelOption
	GetSupportProduct() []PaymentMethod
	IsSupportPayment(payment.PaymentProduct, enum.Device) bool
	CallbackResponse() string
}

type Pay interface {
	Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
	Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error)
	Close(ctx context.Context, req dto.CloseQuery) error
	Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
	Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error)
}
