package model

import (
	"github.com/lihongsheng/payment-sdk/errors"
	"time"
)

type RefundRequest struct {
	OutTradeNo              string                    `json:"out_trade_no,omitempty"`
	TradeNo                 string                    `json:"trade_no,omitempty"`
	RefundAmount            string                    `json:"refund_amount,omitempty"`
	RefundReason            string                    `json:"refund_reason,omitempty"`
	OutRequestNo            string                    `json:"out_request_no,omitempty"`
	RefundGoodsDetail       []RefundGoodsDetail       `json:"refund_goods_detail,omitempty"`
	RefundRoyaltyParameters []RefundRoyaltyParameters `json:"refund_royalty_parameters,omitempty"`
	QueryOptions            []string                  `json:"query_options"`
	RelatedSettleConfirmNo  string                    `json:"related_settle_confirm_no"`
}

func (r RefundRequest) Validate() error {
	if r.OutTradeNo == "" && r.TradeNo == "" {
		return errors.ErrorParamError("out_trade_no or trade_no must not empty", nil)
	}
	if r.RefundAmount == "" {
		return errors.ErrorParamError("refund_amount must not empty", nil)
	}
	return nil
}

type RefundGoodsDetail struct {
	OutSkuId             string   `json:"out_sku_id"`
	OutItemId            string   `json:"out_item_id"`
	GoodsId              string   `json:"goods_id"`
	RefundAmount         string   `json:"refund_amount"`
	OutCertificateNoList []string `json:"out_certificate_no_list"`
}

type RefundRoyaltyParameters struct {
	Amount       string `json:"amount"`
	TransIn      string `json:"trans_in"`
	RoyaltyType  string `json:"royalty_type"`
	TransOut     string `json:"trans_out"`
	TransOutType string `json:"trans_out_type"`
	RoyaltyScene string `json:"royalty_scene"`
	TransInType  string `json:"trans_in_type"`
	TransInName  string `json:"trans_in_name"`
	Desc         string `json:"desc"`
}

type RefundResponse struct {
	AlipayTradeRefundResponse TradeRefund    `json:"alipay_trade_refund_response"`
	ErrorResponse             *ErrorResponse `json:"error_response"`
	Sign                      string         `json:"sign"`
}

type TradeRefund struct {
	Code                    string                `json:"code"`
	Msg                     string                `json:"msg"`
	SubCode                 string                `json:"sub_code"`
	SubMsg                  string                `json:"sub_msg"`
	TradeNo                 string                `json:"trade_no"`
	OutTradeNo              string                `json:"out_trade_no"`
	BuyerLogonId            string                `json:"buyer_logon_id"`
	FundChange              string                `json:"fund_change"`
	RefundFee               string                `json:"refund_fee"`
	RefundDetailItemList    []RefundDetailItem    `json:"refund_detail_item_list"`
	StoreName               string                `json:"store_name"`
	BuyerUserId             string                `json:"buyer_user_id"`
	BuyerOpenId             string                `json:"buyer_open_id"`
	SendBackFee             string                `json:"send_back_fee"`
	RefundHybAmount         string                `json:"refund_hyb_amount"`
	RefundChargeInfoList    []RefundChargeInfo    `json:"refund_charge_info_list"`
	RefundVoucherDetailList []RefundVoucherDetail `json:"refund_voucher_detail_list"`
	PreAuthCancelFee        string                `json:"pre_auth_cancel_fee"`
	GmtRefundPay            string                `json:"gmt_refund_pay"`
}

func (t TradeRefund) GetRefundSuccessTime() time.Time {
	if t.GmtRefundPay != "" {
		t, _ := time.Parse(time.DateTime, t.GmtRefundPay)
		return t
	}
	return time.Time{}
}

type RefundDetailItem struct {
	FundChannel string `json:"fund_channel"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
	FundType    string `json:"fund_type"`
}

type RefundChargeInfo struct {
	RefundChargeFee        string               `json:"refund_charge_fee"`
	SwitchFeeRate          string               `json:"switch_fee_rate"`
	ChargeType             string               `json:"charge_type"`
	RefundSubFeeDetailList []RefundSubFeeDetail `json:"refund_sub_fee_detail_list"`
}

type RefundSubFeeDetail struct {
	RefundChargeFee string `json:"refund_charge_fee"`
	SwitchFeeRate   string `json:"switch_fee_rate"`
}

type RefundVoucherDetail struct {
	Id                         string                  `json:"id"`
	Name                       string                  `json:"name"`
	Type                       string                  `json:"type"`
	Amount                     string                  `json:"amount"`
	MerchantContribute         string                  `json:"merchant_contribute"`
	OtherContribute            string                  `json:"other_contribute"`
	Memo                       string                  `json:"memo"`
	TemplateId                 string                  `json:"template_id"`
	OtherContributeDetail      []OtherContributeDetail `json:"other_contribute_detail"`
	PurchaseBuyerContribute    string                  `json:"purchase_buyer_contribute"`
	PurchaseMerchantContribute string                  `json:"purchase_merchant_contribute"`
	PurchaseAntContribute      string                  `json:"purchase_ant_contribute"`
}

type OtherContributeDetail struct {
	ContributeType   string `json:"contribute_type"`
	ContributeAmount string `json:"contribute_amount"`
}

type RefundQueryRequest struct {
	TradeNo      string   `json:"trade_no"`
	OutTradeNo   string   `json:"out_trade_no"`
	OutRequestNo string   `json:"out_request_no"`
	QueryOptions []string `json:"query_options"`
}

type RefundQueryResponse struct {
	AlipayTradeFastpayRefundQueryResponse TradeFastpayRefundQueryResponse `json:"alipay_trade_fastpay_refund_query_response"`
	ErrorResponse                         *ErrorResponse                  `json:"error_response"`
	Sign                                  string                          `json:"sign"`
}

type TradeFastpayRefundQueryResponse struct {
	Code                    string                `json:"code"`
	Msg                     string                `json:"msg"`
	SubCode                 string                `json:"sub_code"`
	SubMsg                  string                `json:"sub_msg"`
	TradeNo                 string                `json:"trade_no"`
	OutTradeNo              string                `json:"out_trade_no"`
	OutRequestNo            string                `json:"out_request_no"`
	TotalAmount             string                `json:"total_amount"`
	RefundAmount            string                `json:"refund_amount"`
	RefundStatus            string                `json:"refund_status"`
	RefundRoyaltys          []RefundRoyaltys      `json:"refund_royaltys"`
	GmtRefundPay            string                `json:"gmt_refund_pay"`
	RefundDetailItemList    []RefundDetailItem    `json:"refund_detail_item_list"`
	SendBackFee             string                `json:"send_back_fee"`
	DepositBackInfo         DepositBackInfo       `json:"deposit_back_info"`
	RefundHybAmount         string                `json:"refund_hyb_amount"`
	RefundChargeInfoList    []RefundChargeInfo    `json:"refund_charge_info_list"`
	DepositBackInfoList     []DepositBackInfo     `json:"deposit_back_info_list"`
	RefundVoucherDetailList []RefundVoucherDetail `json:"refund_voucher_detail_list"`
	PreAuthCancelFee        string                `json:"pre_auth_cancel_fee"`
}

func (t TradeFastpayRefundQueryResponse) RefundSuccessTime() time.Time {
	if t.GmtRefundPay != "" {
		t, _ := time.Parse(time.DateTime, t.GmtRefundPay)
		return t
	}
	return time.Time{}
}

type DepositBackInfo struct {
	HasDepositBack     string `json:"has_deposit_back"`
	DbackStatus        string `json:"dback_status"`
	DbackAmount        string `json:"dback_amount"`
	BankAckTime        string `json:"bank_ack_time"`
	EstBankReceiptTime string `json:"est_bank_receipt_time"`
}

type RefundRoyaltys struct {
	RefundAmount  string `json:"refund_amount"`
	RoyaltyType   string `json:"royalty_type"`
	ResultCode    string `json:"result_code"`
	TransOut      string `json:"trans_out"`
	TransOutEmail string `json:"trans_out_email"`
	TransIn       string `json:"trans_in"`
	TransInEmail  string `json:"trans_in_email"`
	OriTransOut   string `json:"ori_trans_out"`
	OriTransIn    string `json:"ori_trans_in"`
}
