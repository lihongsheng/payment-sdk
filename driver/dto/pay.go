package dto

import (
	"fmt"
	"time"

	enum "github.com/lihongsheng/payment-sdk/enum/payment"
)

type Amount struct {
	// 分单位，一元为100
	Total    int64
	Currency string
}

func (a Amount) ToFloatString() string {
	return fmt.Sprintf("%.2f", float64(a.Total)/100)
}

type Payer struct {
	// 用户在商户的标识
	UnionID string
	// 用户在商户下某个应用的标识；微信，支付宝，富有都需要
	OpenID string
	// 应用标识;富有需要
	AppID string
}

type Goods struct {
	// 商品名称
	Name string
	// 商品SKU
	Sku string
	// 商品价格，单位分
	Price int64
	// 商品数量
	Quantity int
	//
	Desc string
}

type Order struct {
	OrderNo string
	// 订单金额
	Amount Amount
	// 实付金额
	PayAmount Amount
	// 订单商品
	Goods []Goods
	// 订单名称
	Subject string
	// 订单描述
	Desc string
	//
	CreateAt time.Time
}

type ApplicationInfo struct {
	AppName        string
	Url            string
	IOSPackage     string
	AndroidPackage string
	Type           string
}

type SceneInfo struct {
	// 客户端IP
	ClientIp string
	// 设备ID
	DeviceID string
	// 门店ID
	Store  Store
	H5Info ApplicationInfo
}

type PayOrder struct {
	Order Order
	Payer Payer
	// 支付跳转地址
	RedirectUrl string
	// 订单超时时间
	TimeExpire int64
	// 支付回调地址
	NotifyUrl string
	// 透传参数 如果请求时传递了该参数，异步通知时将该参数原样返回。
	PassbackParams string
	SettleInfo     *SettleInfo
	SceneInfo      *SceneInfo
	RiskFund       *RiskFund
	AlipayExtra    *AlipayPaymentExtra
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
	OrderNo string
	// 交易单号
	TradeNo string
	// 实付金额
	PayAmount Amount
	// Pending | Success | Fail
	Status enum.Status
	// H5 | JSAPI | NATIVE | APP | 扫码
	PaymentProduct string
	Action         Action
	OriginResponse string
}

type SettleInfo struct {
	// 是否分账
	ProfitSharing bool
}

type Action struct {
	// Redirect | Qrcode | Prepay
	Action string
	// Method     string
	Url        string
	Parameters map[string]string
}

type Query struct {
	// 订单号
	OrderNo string
	// 交易单号
	TradeNo string
}

type PayDetail struct {
	// 实付金额
	PayAmount Amount
	// 订单号
	OrderNo string
	// 交易单号
	TradeNo     string
	Status      enum.Status
	SuccessTime int64
	// H5 | JSAPI | NATIVE | APP | 扫码
	PaymentProduct string
	// 支付发卡行，不一定都有
	BankType       string
	OriginResponse string
}

type CloseQuery struct {
	// 订单号
	OrderNo string
	// 交易单号
	TradeNo string
}

type Store struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AreaCode string `json:"area_code"`
	Address  string `json:"address"`
}
