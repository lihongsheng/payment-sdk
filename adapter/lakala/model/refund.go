package model

type RefundRequest struct {
	MerchantNo       string        `json:"merchant_no"`
	TermNo           string        `json:"term_no"`
	OutTradeNo       string        `json:"out_trade_no"`
	RefundAmount     string        `json:"refund_amount"`
	RefundReason     string        `json:"refund_reason"`
	OriginOutTradeNo string        `json:"origin_out_trade_no"`
	LocationInfo     *LocationInfo `json:"location_info"`
	OriginTradeNo    string        `json:"origin_trade_no"`
}

type RefundResponseBody struct {
	CommonResp
	RespData *RefundResponse `json:"resp_data"`
}

type RefundResponse struct {
	OutTradeNo       string `json:"out_trade_no"`
	TradeNo          string `json:"trade_no"`
	LogNo            string `json:"log_no"`
	AccTradeNo       string `json:"acc_trade_no"`
	AccountType      string `json:"account_type"`
	TotalAmount      string `json:"total_amount"`
	RefundAmount     string `json:"refund_amount"`
	PayerAmount      string `json:"payer_amount"`
	TradeTime        string `json:"trade_time"`
	OriginTradeNo    string `json:"origin_trade_no"`
	OriginOutTradeNo string `json:"origin_out_trade_no"`
	UpIssAddnData    string `json:"up_iss_addn_data"`
	UpCouponInfo     string `json:"up_coupon_info"`
	TradeInfo        string `json:"trade_info"`
}
