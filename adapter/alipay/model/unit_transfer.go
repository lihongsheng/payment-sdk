package model

type UnitTransferRequest struct {
	OutBizNo       string    `json:"out_biz_no"`
	TransAmount    string    `json:"trans_amount"`
	BizScene       string    `json:"biz_scene"`
	ProductCode    string    `json:"product_code"`
	OrderTitle     string    `json:"order_title"`
	PayeeInfo      PayeeInfo `json:"payee_info"`
	Remark         string    `json:"remark"`
	BusinessParams string    `json:"business_params"`
}

type PayeeInfo struct {
	Identity     string `json:"identity"`
	Name         string `json:"name"`
	IdentityType string `json:"identity_type"`
}

type UnitTransferResponse struct {
	AlipayFundTransUniTransferResponse AlipayFundTransUniTransferResponse `json:"alipay_fund_trans_uni_transfer_response"`
	ErrorResponse                      *ErrorResponse                     `json:"error_response"`
	Sign                               string                             `json:"sign"`
}

type AlipayFundTransUniTransferResponse struct {
	Code           string `json:"code"`
	Msg            string `json:"msg"`
	OutBizNo       string `json:"out_biz_no"`
	OrderId        string `json:"order_id"`
	PayFundOrderId string `json:"pay_fund_order_id"`
	Status         string `json:"status"`
	TransDate      string `json:"trans_date"`
	SubCode        string `json:"sub_code"`
	// sub_msg
	SubMsg string `json:"sub_msg"`
}

type UnitTransferQueryRequest struct {
	ProductCode string `json:"product_code"`
	BizScene    string `json:"biz_scene"`
	OutBizNo    string `json:"out_biz_no"`
	OrderId     string `json:"order_id"`
}

type UnitTransferQueryResponse struct {
	AlipayFundTransCommonQueryResponse AlipayFundTransCommonQueryResponse `json:"alipay_fund_trans_common_query_response"`
	ErrorResponse                      *ErrorResponse                     `json:"error_response"`
	Sign                               string                             `json:"sign"`
}

type AlipayFundTransCommonQueryResponse struct {
	Code                 string `json:"code"`
	Msg                  string `json:"msg"`
	OrderId              string `json:"order_id"`
	InflowSettleSerialNo string `json:"inflow_settle_serial_no"`
	PayFundOrderId       string `json:"pay_fund_order_id"`
	OutBizNo             string `json:"out_biz_no"`
	TransAmount          string `json:"trans_amount"`
	Status               string `json:"status"`
	PayDate              string `json:"pay_date"`
	SubStatus            string `json:"sub_status"`
	ArrivalTimeEnd       string `json:"arrival_time_end"`
	OrderFee             string `json:"order_fee"`
	ErrorCode            string `json:"error_code"`
	FailReason           string `json:"fail_reason"`
	SubOrderErrorCode    string `json:"sub_order_error_code"`
	SubOrderFailReason   string `json:"sub_order_fail_reason"`
	SubOrderStatus       string `json:"sub_order_status"`
	SettleSerialNo       string `json:"settle_serial_no"`
	ReceiverOpenId       string `json:"receiver_open_id"`
	ReceiverUserId       string `json:"receiver_user_id"`
	FailInstReason       string `json:"fail_inst_reason"`
	FailInstName         string `json:"fail_inst_name"`
	FailInstErrorCode    string `json:"fail_inst_error_code"`
	SubCode              string `json:"sub_code"`
	// sub_msg
	SubMsg string `json:"sub_msg"`
}
