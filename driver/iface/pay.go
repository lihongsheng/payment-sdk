package iface

import (
	"context"
	"net/http"

	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
)

type PaymentDriver interface {
	Open(c config.Config) (Pay, error)
}

type Pay interface {
	Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
	Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error)
	Close(ctx context.Context, req dto.CloseQuery) error
	Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
	Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error)
}
