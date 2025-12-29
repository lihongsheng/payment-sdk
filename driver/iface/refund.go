package iface

import (
	"context"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
)

type RefundDriver interface {
	Open(c config.Config) (Refund, error)
}

type Refund interface {
	Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error)
	Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error)
}
