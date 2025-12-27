package dto

import (
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"time"
)

type UintTransferRequest struct {
	// 转账金额 分单位
	TransferAmount Amount `json:"transfer_amount"`
	// 用户相关信息
	User User `json:"user"`
	// 转账流水号
	TransferNo string `json:"transfer_no"`
	// 给用户的提示
	UserTips string `json:"user_tips"`
	// 回调地址
	NotifyUrl string `json:"notify_url"`
	// 转账场景
	SceneId string `json:"scene_id"`
	// 转账备注
	Remark string `json:"transfer_remark"`
	// 转账场景(微信必填)
	SceneReport []SceneReport `json:"scene_report"`
	// Subject
	Subject string `json:"subject"`
}

type User struct {
	// 用户在商户的标识
	UnionID string
	// 用户在商户下某个应用的标识；微信，支付宝，富有都需要
	OpenID string
	// user name
	UserName string
	Phone    string
}

type SceneReport struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type UintTransferResponse struct {
	// 转账流水号
	TransferNo string `json:"transfer_no"`
	// 微信等转账流水号
	ThirdTransferNo string `json:"transfer_bill_no"`
	// 转账受理时间，有可能没有。微信会返回
	CreateTime *time.Time `json:"create_time"`
	// 转账状态
	Status transfer.UnitTransferStatus `json:"status"`
	// 转账后续动作
	Action *Action
}

type UintTransferQueryRequest struct {
	TransferNo      string `json:"transfer_no"`
	ThirdTransferNo string `json:"third_transfer_no"`
}

type UintTransferDetailResponse struct {
	TransferNo      string `json:"transfer_no"`
	ThirdTransferNo string `json:"third_transfer_no"`
	// 转账状态
	Status         transfer.UnitTransferStatus `json:"status"`
	TransferAmount Amount                      `json:"transfer_amount"`
	Remark         string                      `json:"transfer_remark"`
	FailReason     string                      `json:"fail_reason"`
	User           User                        `json:"user"`
	CreateTime     *time.Time                  `json:"create_time"`
	ArrivalTime    *time.Time                  `json:"arrival_time"`
	OriginResponse string                      `json:"origin_response"`
}

type UnitTransferCancelRequest struct {
	TransferNo string `json:"transfer_no"`
}

type UintTransferCancelResponse struct {
	// 转账流水号
	TransferNo string `json:"transfer_no"`
	// 微信等转账流水号
	ThirdTransferNo string `json:"transfer_bill_no"`
	// 转账状态
	Status transfer.UnitTransferStatus `json:"status"`
}

type CallbackUnitTransferDetail struct {
	TransferNo      string                      `json:"transfer_no"`
	ThirdTransferNo string                      `json:"transfer_bill_no"`
	State           transfer.UnitTransferStatus `json:"state"`
	TransferAmount  int                         `json:"transfer_amount"`
	UserID          string                      `json:"openid"`
	FailReason      string                      `json:"fail_reason"`
	OriginResponse  string
}
