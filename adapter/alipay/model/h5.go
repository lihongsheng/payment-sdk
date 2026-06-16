package model

type H5PaymentRequest struct {
	OutTradeNo  string `json:"out_trade_no,omitempty"`
	TotalAmount string `json:"total_amount,omitempty"`
	Subject     string `json:"subject,omitempty"`
	ProductCode string `json:"product_code,omitempty"`
	AuthToken   string `json:"auth_token,omitempty"`
	// quit_url
	QuitUrl string `json:"quit_url,omitempty"`
	// goods_detail
	GoodsDetail []*GoodDetails `json:"goods_detail,omitempty"`
	TimeExpire  string         `json:"time_expire,omitempty"`

	ExtendParams *H5PaymentExtendParam `json:"extend_params,omitempty"`
	// business_params
	BusinessParams string `json:"business_params,omitempty"`
	// passback_params  url 参数，异步通知时候返回
	PassbackParams     string `json:"passback_params,omitempty"`
	DiscountableAmount string `json:"discountable_amount,omitempty"`
	// merchant_order_no
	MerchantOrderNo string `json:"merchant_order_no,omitempty"`
	// ext_user_info
}

type H5PaymentExtendParam struct {
	SysServiceProviderID string `json:"sys_service_provider_id,omitempty"` // 系统商编号，可选
	HbFqNum              string `json:"hb_fq_num,omitempty"`               // 花呗分期期数，可选
	HbFqSellerPercent    string `json:"hb_fq_seller_percent,omitempty"`    // 花呗分期卖家承担手续费比例，可选
	IndustryRefluxInfo   string `json:"industry_reflux_info,omitempty"`    // 行业数据回流信息，可选
	CardType             string `json:"card_type,omitempty"`               // 卡类型，可选
	RoyaltyFreeze        string `json:"royalty_freeze,omitempty"`          // 是否进行资金冻结，可选
}

//type H5PaymentResponse struct {
//	AlipayTradeCreateResponse AlipayTradeCreateResponse `json:"alipay_trade_create_response"`
//	ErrorResponse             *ErrorResponse            `json:"error_response"`
//	Sign                      string                    `json:"sign"`
//	// pageRedirectionData
//	PageRedirectionData string `json:"pageRedirectionData,omitempty"`
//}

type FacePaymentRequest struct {
	OutTradeNo     string              `json:"out_trade_no"`
	TotalAmount    string              `json:"total_amount"`
	Subject        string              `json:"subject"`
	AuthCode       string              `json:"auth_code"`
	Scene          string              `json:"scene"`
	ProductCode    string              `json:"product_code"`
	SellerId       string              `json:"seller_id"`
	GoodsDetail    []*GoodDetails      `json:"goods_detail"`
	ExtendParams   *PaymentExtendParam `json:"extend_params"`
	BusinessParams *BusinessParams     `json:"business_params"`
	PromoParams    *FacePromoParams    `json:"promo_params"`
	StoreId        string              `json:"store_id"`
	OperatorId     string              `json:"operator_id"`
	TerminalId     string              `json:"terminal_id"`
	QueryOptions   []string            `json:"query_options"`
}

type PaymentExtendParam struct {
	SysServiceProviderId string `json:"sys_service_provider_id"`
	CardType             string `json:"card_type"`
}

type BusinessParams struct {
	McCreateTradeIp string `json:"mc_create_trade_ip"`
}

type FacePromoParams struct {
	ActualOrderTime string `json:"actual_order_time"`
}

type FacePaymentResponse struct {
	AlipayTradePayResponse AlipayTradePayResponse `json:"alipay_trade_pay_response"`
	ErrorResponse          *ErrorResponse         `json:"error_response"`
	Sign                   string                 `json:"sign"`
}

type AlipayTradePayResponse struct {
	Code                string              `json:"code"`
	Msg                 string              `json:"msg"`
	SubCode             string              `json:"sub_code"`
	SubMsg              string              `json:"sub_msg"`
	TradeNo             string              `json:"trade_no"`
	OutTradeNo          string              `json:"out_trade_no"`
	BuyerLogonId        string              `json:"buyer_logon_id"`
	TotalAmount         string              `json:"total_amount"`
	ReceiptAmount       string              `json:"receipt_amount"`
	BuyerPayAmount      string              `json:"buyer_pay_amount"`
	PointAmount         string              `json:"point_amount"`
	InvoiceAmount       string              `json:"invoice_amount"`
	GmtPayment          string              `json:"gmt_payment"`
	FundBillList        []FundBill          `json:"fund_bill_list"`
	StoreName           string              `json:"store_name"`
	DiscountGoodsDetail string              `json:"discount_goods_detail"`
	BuyerOpenId         string              `json:"buyer_open_id"`
	BuyerUserId         string              `json:"buyer_user_id"`
	VoucherDetailList   []VoucherDetailList `json:"voucher_detail_list"`
	MdiscountAmount     string              `json:"mdiscount_amount"`
	DiscountAmount      string              `json:"discount_amount"`
}

type VoucherDetailList struct {
	Id                         string `json:"id"`
	Name                       string `json:"name"`
	Type                       string `json:"type"`
	Amount                     string `json:"amount"`
	MerchantContribute         string `json:"merchant_contribute"`
	OtherContribute            string `json:"other_contribute"`
	Memo                       string `json:"memo"`
	TemplateId                 string `json:"template_id"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute"`
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute"`
	PurchaseAntContribute      string `json:"purchase_ant_contribute"`
}

type QrCodePaymentRequest struct {
	OutTradeNo      string              `json:"out_trade_no"`
	TotalAmount     string              `json:"total_amount"`
	Subject         string              `json:"subject"`
	ProductCode     string              `json:"product_code"`
	SellerId        string              `json:"seller_id"`
	GoodsDetail     []*GoodDetails      `json:"goods_detail"`
	ExtendParams    *PaymentExtendParam `json:"extend_params"`
	BusinessParams  *BusinessParams     `json:"business_params"`
	StoreId         string              `json:"store_id,omitempty"`
	QrPayMode       string              `json:"qr_pay_mode,omitempty"`
	IntegrationType string              `json:"integration_type,omitempty"`
}
