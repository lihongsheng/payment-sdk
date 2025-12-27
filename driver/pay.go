package driver

import (
	"context"

	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
)

type PayDriver interface {
	Open(c config.Config) (Pay, error)
}

type Pay interface {
	Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
	Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error)
	Close(ctx context.Context, req dto.CloseQuery) error
	Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)
}
