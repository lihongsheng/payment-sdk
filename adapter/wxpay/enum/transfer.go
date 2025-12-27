package enum

// ACCEPTED:  转账已受理，可原单重试（非终态）。
//
// PROCESSING:  转账锁定资金中。如果一直停留在该状态，建议检查账户余额是否足够，如余额不足，可充值后再原单重试（非终态）。
//
// WAIT_USER_CONFIRM:  待收款用户确认，当前转账单据资金已锁定，可拉起微信收款确认页面进行收款确认（非终态）。
//
// TRANSFERING:  转账中，可拉起微信收款确认页面再次重试确认收款（非终态）。
//
// SUCCESS:  转账成功，表示转账单据已成功（终态）。
//
// FAIL:  转账失败，表示该笔转账单据已失败。若需重新向用户转账，请重新生成单据并再次发起（终态）。
//
// CANCELING:  转账撤销中，商户撤销请求受理成功，该笔转账正在撤销中，需查单确认撤销的转账单据状态（非终态）。
//
// CANCELLED:  转账撤销完成，代表转账单据已撤销成功（终态）。

const (
	TransferStatus_ACCEPTED   = "ACCEPTED"
	TransferStatus_PROCESSING = "PROCESSING"
	// FINISHED
	TransferStatus_FINISHED = "FINISHED"
	// CLOSED
	TransferStatus_CLOSED = "CLOSED"
	// WAIT_PAY
	TransferStatus_WAIT_PAY = "WAIT_PAY"

	// WAIT_USER_CONFIRM
	TransferStatus_WAIT_USER_CONFIRM = "WAIT_USER_CONFIRM"
	// SUCCESS
	TransferStatus_SUCCESS = "SUCCESS"
	// FAIL
	TransferStatus_FAIL = "FAIL"
	// CANCELING
	TransferStatus_CANCELING = "CANCELING"
	// CANCELLED
	TransferStatus_CANCELLED = "CANCELLED"
	// TRANSFERING
	TransferStatus_TRANSFERING = "TRANSFERING"
)

//
//【批次状态】
//WAIT_PAY: 待付款确认。需要付款出资商户在商家助手小程序或服务商助手小程序进行付款确认
//ACCEPTED:已受理。批次已受理成功，若发起批量转账的30分钟后，转账批次单仍处于该状态，可能原因是商户账户余额不足等。商户可查询账户资金流水，若该笔转账批次单的扣款已经发生，则表示批次已经进入转账中，请再次查单确认
//PROCESSING:转账中。已开始处理批次内的转账明细单
//FINISHED:已完成。批次内的所有转账明细单都已处理完成
//CLOSED:已关闭。可查询具体的批次关闭原因确认

const (
	BatchTransferStatus_WAIT_PAY   = "WAIT_PAY"
	BatchTransferStatus_ACCEPTED   = "ACCEPTED"
	BatchTransferStatus_PROCESSING = "PROCESSING"
	BatchTransferStatus_FINISHED   = "FINISHED"
	BatchTransferStatus_CLOSED     = "CLOSED"
)
