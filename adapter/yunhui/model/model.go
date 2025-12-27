package model

import (
	"errors"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/enum/payment"
	"github.com/singer-stack-lab/payment-sdk/enum/refund"
)

type CommonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Sign string `json:"sign"`
}

func (c CommonResp) Error() error {
	if c.Code != 0 {
		return errors.New(fmt.Sprintf("code:%d;msg:%s", c.Code, c.Msg))
	}
	return nil
}

type Error struct {
	ErrCode string `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

func (e Error) Error() error {
	if e.ErrCode != "" {
		return errors.New(fmt.Sprintf("code:%s;msg:%s", e.ErrCode, e.ErrMsg))
	}
	return nil
}

// {
// "amount": 100,
// "mchNo": "Y1715088123",
// "ifCode": "alipay",
// "apiInfo": "123456"
// "mchOrderNo": "123456789"
// "notifyUrl": "https://business.yiyunhuipay.com/notify.html"
// }
type PaymentRequest struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	MchNo       string `json:"mchNo"`
	WayCode     string `json:"wayCode"`
	ApiInfo     string `json:"apiInfo"`
	MchOrderNo  string `json:"mchOrderNo"`
	NotifyUrl   string `json:"notifyUrl"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	AppId       string `json:"appId"`
	ExpiredTime int64  `json:"expiredTime,omitempty"`
	// {"openid": "o6BcIwvSiRpfS8e_UyfQNrYuk2LI"}
	ChannelExtra string `json:"channelExtra"`
	ReqTime      int64  `json:"reqTime"`
	ClientIp     string `json:"clientIp,omitempty"`
	ReturnUrl    string `json:"returnUrl,omitempty"`
}

type PaymentResponse struct {
	Error
	MchOrderNo  string `json:"mchOrderNo"`
	OrderState  int    `json:"orderState"`
	PayOrderId  string `json:"payOrderId"`
	PayDataType string `json:"payDataType"`
	PayData     string `json:"payData"`
}

type PaymentResp struct {
	CommonResp
	Data *PaymentResponse `json:"data"`
}

type PaymentQuery struct {
	MchNo      string `json:"mchNo"`
	ApiInfo    string `json:"apiInfo"`
	MchOrderNo string `json:"mchOrderNo"`
	AppId      string `json:"appId"`
	ReqTime    int64  `json:"reqTime"`
}
type PaymentQueryResp struct {
	CommonResp
	Data *PaymentQueryResponse `json:"data"`
}
type PaymentQueryResponse struct {
	Amount         int    `json:"amount"`
	AppId          string `json:"appId"`
	Body           string `json:"body"`
	ChannelOrderNo string `json:"channelOrderNo"`
	ClientIp       string `json:"clientIp"`
	CreatedAt      int64  `json:"createdAt"`
	Currency       string `json:"currency"`
	ExtParam       string `json:"extParam"`
	IfCode         string `json:"ifCode"`
	MchNo          string `json:"mchNo"`
	MchOrderNo     string `json:"mchOrderNo"`
	PayOrderId     string `json:"payOrderId"`
	State          int    `json:"state"`
	Subject        string `json:"subject"`
	SuccessTime    int64  `json:"successTime"`
	WayCode        string `json:"wayCode"`
	Error          Error
}

func GetPaymentStatus(state int) payment.Status {
	switch state {
	case 0, 1:
		return payment.Status_Pending
	case 2:
		return payment.Status_Success
	case 3:
		return payment.Status_Failed
	case 4:
		return payment.Status_Cancel
	case 5:
		return payment.Status_Refund
	case 6:
		return payment.Status_Close
	}
	return payment.Status_Status_UNKNOWN
}

type RefundRequest struct {
	PayOrderId   string `json:"payOrderId,omitempty"`
	ExtParam     string `json:"extParam,omitempty"`
	MchOrderNo   string `json:"mchOrderNo,omitempty"`
	RefundReason string `json:"refundReason,omitempty"`
	ReqTime      int64  `json:"reqTime,omitempty"`
	ChannelExtra string `json:"channelExtra,omitempty"`
	AppId        string `json:"appId,omitempty"`
	MchRefundNo  string `json:"mchRefundNo,omitempty"`
	ClientIp     string `json:"clientIp,omitempty"`
	NotifyUrl    string `json:"notifyUrl,omitempty"`
	Currency     string `json:"currency,omitempty"`
	MchNo        string `json:"mchNo,omitempty"`
	RefundAmount int64  `json:"refundAmount,omitempty"`
	ApiInfo      string `json:"apiInfo,omitempty"`
}
type RefundResp struct {
	CommonResp
	Data *RefundResponse `json:"data"`
}
type RefundResponse struct {
	ChannelOrderNo string `json:"channelOrderNo"`
	MchRefundNo    string `json:"mchRefundNo"`
	PayAmount      int    `json:"payAmount"`
	RefundAmount   int64  `json:"refundAmount"`
	RefundOrderId  string `json:"refundOrderId"`
	State          int    `json:"state"`
	Error          Error
}

func GetRefundStatus(state int) refund.Status {
	switch state {
	case 0, 1:
		return refund.Status_Pending
	case 2:
		return refund.Status_Success
	case 3:
		return refund.Status_Failed
	case 4:
		return refund.Status_Closed
	}
	return refund.Status_Status_UNKNOWN
}

type RefundQuery struct {
	MchNo       string `json:"mchNo"`
	ApiInfo     string `json:"apiInfo"`
	MchRefundNo string `json:"mchRefundNo"`
	AppId       string `json:"appId"`
	ReqTime     int64  `json:"reqTime"`
}
type RefundQueryResp struct {
	CommonResp
	Data *RefundQueryResponse `json:"data"`
}
type RefundQueryResponse struct {
	AppId          string `json:"appId"`
	ChannelOrderNo string `json:"channelOrderNo"`
	CreatedAt      int64  `json:"createdAt"`
	Currency       string `json:"currency"`
	ExtParam       string `json:"extParam"`
	MchNo          string `json:"mchNo"`
	MchRefundNo    string `json:"mchRefundNo"`
	PayAmount      int64  `json:"payAmount"`
	PayOrderId     string `json:"payOrderId"`
	RefundAmount   int64  `json:"refundAmount"`
	RefundOrderId  string `json:"refundOrderId"`
	State          int    `json:"state"`
	SuccessTime    int64  `json:"successTime"`
	Error          Error
}
