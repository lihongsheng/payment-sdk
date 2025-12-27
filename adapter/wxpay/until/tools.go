package until

import (
	enum2 "github.com/singer-stack-lab/payment-sdk/adapter/wxpay/enum"
	shareEnum "github.com/singer-stack-lab/payment-sdk/enum/order_share"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
	"github.com/singer-stack-lab/payment-sdk/enum/refund"
	"github.com/singer-stack-lab/payment-sdk/enum/transfer"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

var PaymentStatus = map[string]enum.Status{
	"SUCCESS":    enum.Status_Success,
	"REFUND":     enum.Status_Refund,
	"NOTPAY":     enum.Status_Pending,
	"CLOSED":     enum.Status_Close,
	"REVOKED":    enum.Status_Cancel,
	"USERPAYING": enum.Status_Pending,
	"PAYERROR":   enum.Status_Failed,
}

func Int64Point(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func StringPoint(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func GetRefundStatus(status refunddomestic.Status) refund.Status {
	switch status {
	case refunddomestic.STATUS_SUCCESS:
		return refund.Status_Success
	case refunddomestic.STATUS_CLOSED:
		return refund.Status_Closed
	case refunddomestic.STATUS_PROCESSING:
		return refund.Status_Pending
	case refunddomestic.STATUS_ABNORMAL:
		return refund.Status_Abnormal
	default:
		return refund.Status_Status_UNKNOWN
	}
}

func GetRefundStatus2(status string) refund.Status {
	switch status {
	case string(refunddomestic.STATUS_SUCCESS):
		return refund.Status_Success
	case string(refunddomestic.STATUS_CLOSED):
		return refund.Status_Closed
	case string(refunddomestic.STATUS_PROCESSING):
		return refund.Status_Pending
	case string(refunddomestic.STATUS_ABNORMAL):
		return refund.Status_Abnormal
	default:
		return refund.Status_Status_UNKNOWN
	}
}

func GetBatchStatus(status string) transfer.TransferStatus {
	switch status {
	case enum2.TransferStatus_PROCESSING:
		return transfer.TransferStatus_TransferStatus_PROCESSING
	case enum2.TransferStatus_ACCEPTED:
		return transfer.TransferStatus_TransferStatus_ACCEPTED
	case enum2.TransferStatus_WAIT_PAY:
		return transfer.TransferStatus_TransferStatus_WAIT_PAY
	case enum2.TransferStatus_FINISHED:
		return transfer.TransferStatus_TransferStatus_FINISHED
	case enum2.TransferStatus_CLOSED:
		return transfer.TransferStatus_TransferStatus_CLOSE
	case enum2.TransferStatus_CANCELING:
		return transfer.TransferStatus_TransferStatus_CANCELING
	case enum2.TransferStatus_CANCELLED:
		return transfer.TransferStatus_TransferStatus_CANCELED
	case enum2.TransferStatus_WAIT_USER_CONFIRM:
		return transfer.TransferStatus_TransferStatus_WAIT_USER_CONFIRM
	case enum2.TransferStatus_TRANSFERING:
		return transfer.TransferStatus_TransferStatus_TRANSFERING
	case enum2.TransferStatus_FAIL:
		return transfer.TransferStatus_TransferStatus_FAILED

	}
	return transfer.TransferStatus_TransferStatus_UNKNOWN
}

func GetTransferStatus(status string) transfer.UnitTransferStatus {
	switch status {
	case enum2.TransferStatus_PROCESSING:
		return transfer.UnitTransferStatus_UnitTransferStatus_PROCESSING
	case enum2.TransferStatus_ACCEPTED:
		return transfer.UnitTransferStatus_UnitTransferStatus_ACCEPTED
	case enum2.TransferStatus_WAIT_PAY:
		return transfer.UnitTransferStatus_UnitTransferStatus_WAIT_PAY
	case enum2.TransferStatus_FINISHED:
		return transfer.UnitTransferStatus_UnitTransferStatus_FINISHED
	case enum2.TransferStatus_CLOSED:
		return transfer.UnitTransferStatus_UnitTransferStatus_CLOSE
	case enum2.TransferStatus_CANCELING:
		return transfer.UnitTransferStatus_UnitTransferStatus_CANCELING
	case enum2.TransferStatus_CANCELLED:
		return transfer.UnitTransferStatus_UnitTransferStatus_CANCELED
	case enum2.TransferStatus_WAIT_USER_CONFIRM:
		return transfer.UnitTransferStatus_UnitTransferStatus_WAIT_USER_CONFIRM
	case enum2.TransferStatus_TRANSFERING:
		return transfer.UnitTransferStatus_UnitTransferStatus_TRANSFERING
	case enum2.TransferStatus_SUCCESS:
		return transfer.UnitTransferStatus_UnitTransferStatus_SUCCESS
	case enum2.TransferStatus_FAIL:
		return transfer.UnitTransferStatus_UnitTransferStatus_FAILED

	}
	return transfer.UnitTransferStatus_UnitTransferStatus_UNKNOWN
}

func GetScorePaymentStatus(orderStatus string, completeStatus string, collectionStatus string) enum.Status {
	status := enum.Status_Status_UNKNOWN
	switch orderStatus {
	case "CREATED":
		status = enum.Status_Pending
	case "DOING":
		status = enum.Status_WaitConfirm
		// 【订单状态说明】此参数用于对服务订单处于DOING状态时的附加说明，非DOIING状态将不会返回该参数。具体状态如下：
		//USER_CONFIRM：用户已确认状态，表示用户成功确认订单后所处状态。
		//MCH_COMPLETE：商户已完结状态，指商户调用完结接口成功后至扣款成功前的状态。
		if completeStatus == "MCH_COMPLETE" {
			status = enum.Status_WaitPay
		}
		// USER_PAYING：待支付
		//USER_PAID：已支付
		if collectionStatus == "USER_PAID" {
			status = enum.Status_Success
		}
	case "DONE":
		status = enum.Status_WaitPay
		// USER_PAYING：待支付
		//USER_PAID：已支付
		if collectionStatus == "USER_PAID" {
			status = enum.Status_Success
		}
	case "REVOKED":
		status = enum.Status_Cancel
	case "EXPIRED":
		status = enum.Status_Close
	}

	return status
}

func GetShareTransferStatus(status *profitsharing.OrderStatus) shareEnum.Status {
	if status == nil {
		return shareEnum.Status_UNKNOWN
	}
	s := *status
	switch s {
	case profitsharing.ORDERSTATUS_PROCESSING:
		return shareEnum.Status_PROCESSING
	case profitsharing.ORDERSTATUS_FINISHED:
		return shareEnum.Status_FINISHED
	}
	return shareEnum.Status_UNKNOWN
}

func GetShareTransferOrderStatus(status *profitsharing.DetailStatus, fail *profitsharing.DetailFailReason) shareEnum.ShareOrderStaus {
	if status == nil {
		return shareEnum.ShareOrderStaus_SHARE_ORDER_UNKNOWN
	}
	s := *status
	switch s {
	case profitsharing.DETAILSTATUS_PENDING:
		return shareEnum.ShareOrderStaus_SHARE_ORDER_PROCESSING
	case profitsharing.DETAILSTATUS_SUCCESS:
		return shareEnum.ShareOrderStaus_SHARE_ORDER_SUCCESS
	case profitsharing.DETAILSTATUS_CLOSED:
		if fail != nil {
			return shareEnum.ShareOrderStaus_SHARE_ORDER_FAILED
		}
	}
	return shareEnum.ShareOrderStaus_SHARE_ORDER_UNKNOWN
}

func GetTransferDetailStatus(status string) transfer.DetailStatus {
	// INIT: 初始态。 系统转账校验中
	//WAIT_PAY: 待确认。待商户确认, 符合免密条件时, 系统会自动扭转为转账中
	//PROCESSING:转账中。正在处理中，转账结果尚未明确
	//SUCCESS:转账成功
	//FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单）
	switch status {
	case "INIT", "WAIT_PAY":
		return transfer.DetailStatus_DetailStatus_WAITING
	case "PROCESSING":
		return transfer.DetailStatus_DetailStatus_PROCESSING
	case "SUCCESS":
		return transfer.DetailStatus_DetailStatus_SUCCESS
	case "FAIL":
		return transfer.DetailStatus_DetailStatus_FAILED
	}
	return transfer.DetailStatus_DetailStatus_UNKNOWN
}
