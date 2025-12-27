package transfer

type AllocateType string

const (
	// 自动结算手续费，不充转账金额中扣除
	AllocateType_Auto AllocateType = "01"
	// 手续费由商户承担,在此次转账金额中扣除
	AllocateType_Merchant AllocateType = "02"
)
