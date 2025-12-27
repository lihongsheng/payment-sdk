package dto

import (
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"time"
)

type BatchTransferRequest struct {
	TransferNO string `json:"transfer_no"`
	Subject    string `json:"subject"`
	Remark     string `json:"remark"`
	Amount     Amount `json:"amount"`
	// 场景id
	SceneID   string                       `json:"scene_id"`
	NotifyUrl string                       `json:"notify_url"`
	Extend    map[string]string            `json:"extend"`
	Details   []BatchTransferDetailRequest `json:"details"`
	//WechatTransferSceneReportInfos *WechatTransferSceneReportInfos `json:"wechat_transfer_scene_report_infos"`
	// 子商户账号，富有可以选传
	SubAccountIn string `json:"sub_account_in"`
}

//type WechatTransferSceneReportInfos struct {
//	InfoType    string `json:"info_type"`
//	InfoContent string `json:"info_content"`
//}

type BatchTransferDetailRequest struct {
	// 微信或者支付宝传 openid,富有传 Account no
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	TradeNO  string `json:"trade_no"`
	Amount   Amount `json:"amount"`
	Remark   string `json:"remark"`
	// 银行卡号，富有自动结算必传
	BankNO string `json:"bank_no"`
	// 富有结算必传 01 转账金额不参与费率，2分账金额参与费率
	AllocateType transfer.AllocateType `json:"allocate_type"`
}

type BatchTransferResponse struct {
	TransferNO      string                  `json:"transfer_no"`
	ThirdTransferNO string                  `json:"third_transfer_no"`
	CreateTime      time.Time               `json:"create_time"`
	TransferStatus  transfer.TransferStatus `json:"transfer_status"`
	PackageInfo     string                  `json:"package_info"`
	OriginResponse  string
}

type BatchTransferQueryRequest struct {
	TransferNO      string `json:"transfer_no"`
	ThirdTransferNO string `json:"third_transfer_no"`
	PageNum         int64  `json:"page_num"`
	PageSize        int64  `json:"page_size"`
	// 富有必传
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type BatchTransferQueryResponse struct {
	TransferNO      string                     `json:"transfer_no"`
	ThirdTransferNO string                     `json:"third_transfer_no"`
	Total           int64                      `json:"total"`
	Details         []BatchTransferQueryDetail `json:"details"`
	TransferStatus  transfer.TransferStatus    `json:"transfer_status"`
	OriginResponse  string
}

type BatchTransferQueryDetail struct {
	// 支付宝传 openid,富有传 Account no.
	UserID          string                `json:"user_id"`
	UserName        string                `json:"user_name"`
	TradeNO         string                `json:"trade_no"`
	ThirdTransferNO string                `json:"third_transfer_no"`
	Amount          Amount                `json:"amount"`
	Remark          string                `json:"remark"`
	TransferStatus  transfer.DetailStatus `json:"transfer_status"`
	ErrorRemark     string                `json:"error_remark"`
}

type CallbackBatchTransferRequest struct {
	TransferNO      string `json:"transfer_no"`
	ThirdTransferNO string `json:"third_transfer_no"`
	// 支付宝没有 details，需要主动查询。富有单条结果，微信单条结果（微信批量转账接口不在支持）
	Details        []BatchTransferQueryDetail `json:"details"`
	TransferStatus transfer.TransferStatus    `json:"transfer_status"`
	OriginResponse string
}
