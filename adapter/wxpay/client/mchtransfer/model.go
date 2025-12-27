package mchtransfer

import (
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay/enum"
	"github.com/singer-stack-lab/payment-sdk/enum/transfer"
	"time"
)

type TransferBillRequest struct {
	TransferSceneReportInfos []TransferSceneReportInfo `json:"transfer_scene_report_infos,omitempty"`
	TransferAmount           int64                     `json:"transfer_amount,omitempty"`
	Openid                   string                    `json:"openid,omitempty"`
	UserName                 string                    `json:"user_name,omitempty"`
	Appid                    string                    `json:"appid,omitempty"`
	OutBillNo                string                    `json:"out_bill_no,omitempty"`
	UserRecvPerception       string                    `json:"user_recv_perception,omitempty"`
	NotifyUrl                string                    `json:"notify_url,omitempty"`
	TransferSceneId          string                    `json:"transfer_scene_id,omitempty"`
	TransferRemark           string                    `json:"transfer_remark,omitempty"`
}

type TransferSceneReportInfo struct {
	InfoType    string `json:"info_type"`
	InfoContent string `json:"info_content"`
}

type TransferBillResponse struct {
	OutBillNo      string     `json:"out_bill_no"`
	TransferBillNo string     `json:"transfer_bill_no"`
	CreateTime     *time.Time `json:"create_time"`
	State          string     `json:"state"`
	PackageInfo    string     `json:"package_info"`
}

type TransferBillCancelRequest struct {
	OutBillNo string `json:"out_bill_no"`
}

type TransferBillCancelResponse struct {
	OutBillNo      string    `json:"out_bill_no"`
	TransferBillNo string    `json:"transfer_bill_no"`
	State          string    `json:"state"`
	UpdateTime     time.Time `json:"update_time"`
}

type GetTransferBilByOutBillNORequest struct {
	OutBillNo string `json:"out_bill_no"`
}

type GetTransferBilByNORequest struct {
	TransferBillNo string `json:"out_bill_no"`
}

type TransferBillDetailResponse struct {
	MchId          string     `json:"mch_id"`
	OutBillNo      string     `json:"out_bill_no"`
	TransferBillNo string     `json:"transfer_bill_no"`
	Appid          string     `json:"appid"`
	State          string     `json:"state"`
	TransferAmount int64      `json:"transfer_amount"`
	TransferRemark string     `json:"transfer_remark"`
	FailReason     string     `json:"fail_reason"`
	Openid         string     `json:"openid"`
	UserName       string     `json:"user_name"`
	CreateTime     *time.Time `json:"create_time"`
	UpdateTime     *time.Time `json:"update_time"`
}

type CallBackRequest struct {
	OutBatchNo    string    `json:"out_batch_no"`
	BatchId       string    `json:"batch_id"`
	BatchStatus   string    `json:"batch_status"`
	TotalNum      int       `json:"total_num"`
	TotalAmount   int       `json:"total_amount"`
	SuccessAmount int       `json:"success_amount"`
	SuccessNum    int       `json:"success_num"`
	FailAmount    int       `json:"fail_amount"`
	FailNum       int       `json:"fail_num"`
	Mchid         string    `json:"mchid"`
	UpdateTime    time.Time `json:"update_time"`
}

func (c CallBackRequest) GetTransferStatus() transfer.TransferStatus {
	switch c.BatchStatus {
	case enum.BatchTransferStatus_ACCEPTED:
		return transfer.TransferStatus_TransferStatus_ACCEPTED
	case enum.BatchTransferStatus_WAIT_PAY:
		return transfer.TransferStatus_TransferStatus_WAIT_PAY
	case enum.BatchTransferStatus_FINISHED:
		return transfer.TransferStatus_TransferStatus_FINISHED
	case enum.BatchTransferStatus_PROCESSING:
		return transfer.TransferStatus_TransferStatus_PROCESSING
	case enum.BatchTransferStatus_CLOSED:
		return transfer.TransferStatus_TransferStatus_CLOSE
	}
	return transfer.TransferStatus_TransferStatus_UNKNOWN
}
