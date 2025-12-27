package model

import (
	"errors"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/lakala/enum"
	"time"
)

type PaymentRequest struct {
	MerchantNo    string        `json:"merchant_no,omitempty"`
	TermNo        string        `json:"term_no,omitempty"`
	OutTradeNo    string        `json:"out_trade_no,omitempty"`
	AccountType   string        `json:"account_type,omitempty"`
	TransType     string        `json:"trans_type,omitempty"`
	TotalAmount   string        `json:"total_amount,omitempty"`
	NotifyUrl     string        `json:"notify_url,omitempty"`
	LocationInfo  *LocationInfo `json:"location_info,omitempty"`
	Subject       string        `json:"subject,omitempty"`
	Remark        string        `json:"remark,omitempty"`
	AccBusiFields interface{}   `json:"acc_busi_fields,omitempty"`
}

type PaymentQueryRequest struct {
	MerchantNo string `json:"merchant_no,omitempty"`
	TermNo     string `json:"term_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

type PaymentQueryRespBody struct {
	CommonResp
	RespData *PaymentQueryResponse `json:"resp_data"`
}

type PaymentQueryResponse struct {
	OutTradeNo         string           `json:"out_trade_no"`
	TradeNo            string           `json:"trade_no"`
	LogNo              string           `json:"log_no"`
	AccTradeNo         string           `json:"acc_trade_no"`
	AccountType        string           `json:"account_type"`
	SettleTermNo       string           `json:"settle_term_no"`
	TradeState         string           `json:"trade_state"`
	TradeStateDesc     string           `json:"trade_state_desc"`
	TotalAmount        string           `json:"total_amount"`
	PayerAmount        string           `json:"payer_amount"`
	AccSettleAmount    string           `json:"acc_settle_amount"`
	AccMdiscountAmount string           `json:"acc_mdiscount_amount"`
	AccDiscountAmount  string           `json:"acc_discount_amount"`
	TradeTime          string           `json:"trade_time"`
	UserId1            string           `json:"user_id1"`
	UserId2            string           `json:"user_id2"`
	BankType           string           `json:"bank_type"`
	AccActivityId      string           `json:"acc_activity_id"`
	UpCouponInfo       string           `json:"up_coupon_info"`
	TradeInfo          string           `json:"trade_info"`
	RefundSplitInfo    []RefundSubTrade `json:"refund_split_info,omitempty"`
	TradeMainType      string           `json:"trade_main_type"`
}

type LocationInfo struct {
	RequestIp string `json:"request_ip,omitempty"`
	Location  string `json:"location,omitempty"`
}

type CommonRequest struct {
	ReqTime string      `json:"req_time,omitempty"`
	Version string      `json:"version,omitempty"`
	ReqData interface{} `json:"req_data,omitempty"`
}

func BuildCommonReq(data interface{}) *CommonRequest {
	r := &CommonRequest{
		ReqTime: time.Now().Format("20060102150405"),
		Version: enum.Version,
		ReqData: data,
	}
	return r
}

type WxAccBusiFields struct {
	// 预下单有效时间（单位：分钟，不超过15分钟，默认5分钟）
	TimeoutExpress string `json:"timeout_express,omitempty"`
	// 子商户公众账号ID（微信分配，对应sub_appid）
	SubAppid string `json:"sub_appid,omitempty"`
	// 用户标识（子商户sub_appid下的唯一标识，对应sub_openid）
	UserID string `json:"user_id,omitempty"`
	// 商品详情（单品优惠功能字段）
	Detail string `json:"detail,omitempty"`
	// 订单优惠标记（微信平台配置的商品标记）
	GoodsTag string `json:"goods_tag,omitempty"`
	// 附加域（商户自定义数据）
	Attach string `json:"attach,omitempty"`
	// 设备号（终端设备号，PC网页/JSAPI支付传"WEB"）
	DeviceInfo string `json:"device_info,omitempty"`
	// 指定支付方式（no_credit-不能使用信用卡支付）
	LimitPay string `json:"limit_pay,omitempty"`
	// 场景信息（上报实际门店信息）
	SceneInfo string `json:"scene_info,omitempty"`
	// 限定支付（ADULT-成年人）
	LimitPayer string `json:"limit_payer,omitempty"`
}

type CommonResp struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	RespTime string `json:"resp_time"`
	//RespData json.RawMessage `json:"resp_data"`
}

func (c CommonResp) IsSuccess() bool {
	if c.Code == "BBS00000" {
		return true
	}
	return false
}

func (c CommonResp) GetError() error {
	if c.IsSuccess() {
		return nil
	}
	return errors.New(fmt.Sprintf("code-%s;msg-%s", c.Code, c.Msg))
}

type PaymentRespBody struct {
	CommonResp
	RespData *PaymentResponse `json:"resp_data"`
}

type PaymentResponse struct {
	MerchantNo       string           `json:"merchant_no,omitempty"`
	OutTradeNo       string           `json:"out_trade_no,omitempty"`
	TradeNo          string           `json:"trade_no,omitempty"`
	LogNo            string           `json:"log_no,omitempty"`
	SettleMerchantNo string           `json:"settle_merchant_no,omitempty"`
	SettleTermNo     string           `json:"settle_term_no,omitempty"`
	AccRespFields    *WxAccRespFields `json:"acc_resp_fields,omitempty"`
}

type WxAccRespFields struct {
	// 预下单ID（必填）
	PrepayId string `json:"prepay_id"`
	// 支付签名信息（必填）
	PaySign string `json:"pay_sign"`
	// 移动应用appid（必填）
	AppId string `json:"app_id"`
	// 时间戳（必填）
	TimeStamp string `json:"time_stamp"`
	// 随机字符串（必填）
	NonceStr string `json:"nonce_str"`
	// 订单详情扩展字符串（必填）
	Package string `json:"package"`
	// 签名方式（必填，支持RSA）
	SignType string `json:"sign_type"`
	// 从业机构号（必填）
	PartnerId string `json:"partner_id"`
	// 子商户号（可选）
	SubMchId string `json:"sub_mch_id,omitempty"`
}

// RefundSubTrade 退款子交易信息结构体
type RefundSubTrade struct {
	// 外部子退款交易流水号（必填，商户号下唯一）
	OutSubTradeNo string `json:"out_sub_trade_no"`
	// 商户号（必填，拉卡拉分配）
	MerchantNo string `json:"merchant_no"`
	// 终端号（必填，拉卡拉分配）
	TermNo string `json:"term_no"`
	// 申请退款金额（必填，单位分，整数型字符）
	RefundAmount string `json:"refund_amount"`
	// 拉卡拉子交易流水号（可选）
	SubTradeNo string `json:"sub_trade_no,omitempty"`
	// 对账单子流水号（可选，sub_trade_no后14位）
	SubLogNo string `json:"sub_log_no,omitempty"`
	// 子交易状态（可选，SUCCESS/FAIL）
	TradeState string `json:"trade_state,omitempty"`
	// 处理结果码（可选）
	ResultCode string `json:"result_code,omitempty"`
	// 处理描述（可选）
	ResultMsg string `json:"result_msg,omitempty"`
}
