package dto

import (
	"fmt"
	"time"

	enum "github.com/lihongsheng/payment-sdk/enum/payment"
)

type Amount struct {
	// 分单位，一元为100
	Total    int64  `json:"total"`
	Currency string `json:"currency"`
}

func (a Amount) ToFloatString() string {
	return fmt.Sprintf("%.2f", float64(a.Total)/100)
}

type Payer struct {
	// 用户在商户的标识
	UnionID string `json:"union_id"`
	// 用户在商户下某个应用的标识；微信，支付宝，富有都需要
	OpenID string `json:"open_id"`
	// 应用标识;富有需要
	AppID string `json:"app_id"`
}

type Goods struct {
	// 商品名称
	Name string `json:"name"`
	// 商品SKU
	Sku string `json:"sku"`
	// 商品价格，单位分
	Price int64 `json:"price"`
	// 商品数量
	Quantity int `json:"quantity"`
	//
	Desc string `json:"desc"`
}

type Order struct {
	OrderNo string `json:"orderNo"`
	// 订单金额
	Amount Amount `json:"amount"`
	// 实付金额
	PayAmount Amount `json:"pay_amount"`
	// 订单商品
	Goods []Goods `json:"goods"`
	// 订单名称
	Subject string `json:"subject"`
	// 订单描述
	Desc string `json:"desc"`
	// 订单创建时间
	CreateAt time.Time `json:"create_at"`
}

type ApplicationInfo struct {
	AppName        string `json:"app_name"`
	Url            string `json:"url"`
	IOSPackage     string `json:"ios_package"`
	AndroidPackage string `json:"android_package"`
	Type           string `json:"type"`
}

type SceneInfo struct {
	// 客户端IP
	ClientIp string `json:"client_ip"`
	// 设备ID
	DeviceID string `json:"device_id"`
	// 门店ID
	Store  Store           `json:"store"`
	H5Info ApplicationInfo `json:"h5_info"`
}

type PayOrder struct {
	Order Order `json:"order"`
	Payer Payer `json:"payer"`
	// 支付跳转地址
	RedirectUrl string `json:"redirect_url"`
	// 订单超时时间
	TimeExpire int64 `json:"time_expire"`
	// 支付回调地址
	NotifyUrl string `json:"notify_url"`
	// 透传参数 如果请求时传递了该参数，异步通知时将该参数原样返回。
	PassbackParams string              `json:"passback_params"`
	SettleInfo     *SettleInfo         `json:"settle_info"`
	SceneInfo      *SceneInfo          `json:"scene_info"`
	RiskFund       *RiskFund           `json:"risk_fund"`
	AlipayExtra    *AlipayPaymentExtra `json:"alipay_extra"`
}

type AlipayPaymentExtra struct {
	// 产品码。
	//商家和支付宝签约的产品码。
	//当面付场景下，如果签约的是当面付快捷版，则传 OFFLINE_PAYMENT;
	//其它支付宝当面付产品传 FACE_TO_FACE_PAYMENT；
	//不传则默认使用FACE_TO_FACE_PAYMENT。
	//【示例值】FACE_TO_FACE_PAYMENT
	ProductCode string `json:"product_code,omitempty"`
}

type RiskFund struct {
	Name        string `json:"name,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}

type PayResponse struct {
	// 订单号
	OrderNo string `json:"order_no"`
	// 交易单号
	TradeNo string `json:"trade_no"`
	// 实付金额
	PayAmount Amount `json:"pay_amount"`
	// Pending | Success | Fail
	Status enum.Status `json:"status"`
	// H5 | JSAPI | NATIVE | APP | 扫码
	PaymentProduct string `json:"payment_product"`
	Action         Action `json:"action"`
	OriginResponse string `json:"origin_response"`
}

type SettleInfo struct {
	// 是否分账
	ProfitSharing bool `json:"profit_sharing"`
}

type Action struct {
	// Redirect | Qrcode | Prepay
	Action string `json:"action"`
	// Method     string
	Url        string            `json:"url"`
	Parameters map[string]string `json:"parameters"`
	// GET | POST
	RedirectMethod string `json:"redirect_method"`
}

type Query struct {
	// 订单号
	OrderNo string `json:"order_no"`
	// 交易单号
	TradeNo string `json:"trade_no"`
}

type PayDetail struct {
	// 实付金额
	PayAmount Amount `json:"pay_amount"`
	// 订单号
	OrderNo string `json:"order_no"`
	// 交易单号
	TradeNo     string      `json:"trade_no"`
	Status      enum.Status `json:"status"`
	SuccessTime int64       `json:"success_time"`
	// H5 | JSAPI | NATIVE | APP | 扫码
	PaymentProduct string `json:"payment_product"`
	// 支付发卡行，不一定都有
	BankType       string `json:"bank_type"`
	OriginResponse string `json:"origin_response"`
}

type CloseQuery struct {
	// 订单号
	OrderNo string `json:"order_no"`
	// 交易单号
	TradeNo string `json:"trade_no"`
}

type Store struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AreaCode string `json:"area_code"`
	Address  string `json:"address"`
}
