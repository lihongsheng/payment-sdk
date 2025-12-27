package model

import "time"

type UnitTransferCallback struct {
	OutBillNo      string    `json:"out_bill_no"`
	TransferBillNo string    `json:"transfer_bill_no"`
	State          string    `json:"state"`
	MchId          string    `json:"mch_id"`
	TransferAmount int       `json:"transfer_amount"`
	Openid         string    `json:"openid"`
	FailReason     string    `json:"fail_reason"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

type ComplaintCallbackRequest struct {
	ComplaintId string `json:"complaint_id"`
	ActionType  string `json:"action_type"`
}
