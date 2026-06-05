package iface

import (
	"context"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"net/http"
)

type RefundDriver interface {
	Open(c config.Config) (Refund, error)
	CallbackResponse() string
}

type Refund interface {
	Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error)
	Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error)
	Callback(ctx context.Context, req *http.Request) (*dto.CallbackRefundDetail, error)
	IsSupportCallback() bool
}
