package iface

import (
	"context"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
)

type UnitTransferDriver interface {
	Open(c config.Config) (UnitTransfer, error)
}
type UnitTransfer interface {
	Transfer(ctx context.Context, req *dto.UintTransferRequest) (*dto.UintTransferResponse, error)
	Query(ctx context.Context, req dto.UintTransferQueryRequest) (*dto.UintTransferDetailResponse, error)
	Cancel(ctx context.Context, req dto.UnitTransferCancelRequest) (*dto.UintTransferCancelResponse, error)
}
