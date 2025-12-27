package prepay

import (
	"encoding/json"
	"fmt"
	"strings"

	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
)

// PaymentRequest 富友支付请求参数结构体
type PaymentRequest struct {
	Version      string `json:"version"`        // 版本号，长度8，必填
	MchntCd      string `json:"mchnt_cd"`       // 商户号，长度15，必填
	RandomStr    string `json:"random_str"`     // 随机字符串，长度32，必填
	OrderAmt     int64  `json:"order_amt"`      // 订单总金额，长度16，以分为单位，必填
	MchntOrderNo string `json:"mchnt_order_no"` // 商户订单号，长度30，必填
	ProductId    string `json:"product_id"`     // 商品标识，长度32，非必填
	TxnBeginTs   string `json:"txn_begin_ts"`   // 交易起始时间，长度14，格式yyyyMMddHHmmss，必填
	// 订单标题
	GoodsDes    string `json:"goods_des"`    // 商品描述，长度128，必填
	GoodsDetail string `json:"goods_detail"` // 商品详情，长度600，非必填
	GoodsTag    string `json:"goods_tag"`    // 商品标记，长度32，非必填
	TermId      string `json:"term_id"`      // 终端号，长度8，随机8字节数字字母组合，必填
	// IP
	TermIp  string `json:"term_ip"`  // 终端IP，长度16，必填
	AddnInf string `json:"addn_inf"` // 附加数据，长度100，非必填
	// 币种
	CurrType         string `json:"curr_type"`                     // 货币种类，长度3，默认人民币:CNY，非必填
	NotifyURL        string `json:"notify_url"`                    // 通知URL，长度256，接收富友异步通知回调地址，必填
	LimitPay         string `json:"limit_pay,omitempty"`           // 限制支付，长度32，非必填
	TradeType        string `json:"trade_type"`                    // 交易类型，长度16，必填
	Openid           string `json:"openid,omitempty"`              // 用户标识，长度128，非必填
	SubOpenid        string `json:"sub_openid,omitempty"`          // 子商户用户标识，长度128，非必填
	SubAppid         string `json:"sub_appid,omitempty"`           // 子商户公众号id，长度32，非必填
	ReservedFyTermId string `json:"reserved_fy_term_id,omitempty"` // 富友终端号，长度20，非必填
	// 时间 分钟
	ReservedExpireMinute int    `json:"reserved_expire_minute,omitempty"` // 交易关闭时间，长度8，非必填
	ReservedDeviceInfo   string `json:"reserved_device_info"`
	//拼接
	//
	//mchnt_cd+"|"
	//+ trade_type +"|"
	//+ order_amt +"|"
	//+ mchnt_order_no+"|"
	//+ txn_begin_ts+"|"
	//+ goods_des +"|"
	//+ term_id +"|"
	//+ term_ip +"|"
	//+ notify_url +"|"
	//+ random_str +"|"
	//+ version + "|"
	//+ mchnt_key
	//
	//并做md5摘要
	//其中 mchnt_key 为商户密钥，系统分配
	Sign string `json:"sign"`
}

func (p *PaymentRequest) Validate() error {
	if p.MchntCd == "" {
		return errors2.ErrorParamError("商户号不能为空")
	}
	if p.RandomStr == "" {
		return errors2.ErrorParamError("随机字符串不能为空")
	}
	if p.OrderAmt <= 0 {
		return errors2.ErrorParamError("订单金额不能小于0")
	}
	if p.TradeType == "" {
		return errors2.ErrorParamError("交易类型不能为空")
	}
	return nil
}

func (p *PaymentRequest) GenSign(apiKey string) {
	signStrArr := []string{}
	signStrArr = append(signStrArr, p.MchntCd)
	signStrArr = append(signStrArr, p.TradeType)
	signStrArr = append(signStrArr, fmt.Sprintf("%d", p.OrderAmt))
	signStrArr = append(signStrArr, p.MchntOrderNo)
	signStrArr = append(signStrArr, p.TxnBeginTs)
	signStrArr = append(signStrArr, p.GoodsDes)
	signStrArr = append(signStrArr, p.TermId)
	signStrArr = append(signStrArr, p.TermIp)
	signStrArr = append(signStrArr, p.NotifyURL)
	signStrArr = append(signStrArr, p.RandomStr)
	signStrArr = append(signStrArr, p.Version)
	signStrArr = append(signStrArr, apiKey)
	p.Sign = tools.Md5(strings.Join(signStrArr, "|"))
}

type DeviceInfo struct {
	// 场景类型
	Type string `json:"type"`
	// 应用名称
	AppName string `json:"app_name,omitempty"`
	// 网站URL
	AppUrl string `json:"app_url,omitempty"`
}

func (d DeviceInfo) String() string {
	by, _ := json.Marshal(d)
	return string(by)
}

// PaymentResponse 支付响应参数结构体（续）
type PaymentResponse struct {
	// 已有字段（承接上一部分）
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	MchntCd      string `json:"mchnt_cd"`
	TermId       string `json:"term_id,omitempty"`
	RandomStr    string `json:"random_str"`
	SubMerId     string `json:"sub_mer_id,omitempty"`
	SessionId    string `json:"session_id,omitempty"`
	QrCode       string `json:"qr_code,omitempty"`
	SubAppid     string `json:"sub_appid,omitempty"`
	SubOpenid    string `json:"sub_openid,omitempty"`
	SdkAppid     string `json:"sdk_appid,omitempty"`
	SdkTimestamp string `json:"sdk_timestamp,omitempty"`
	SdkNoncestr  string `json:"sdk_noncestr,omitempty"`
	// prepay_id=wx121627586241613c7d311e003087612412
	SdkPackage string `json:"sdk_package,omitempty"`

	// 新增字段（当前部分）
	SdkSigntype        string `json:"sdk_signtype,omitempty"`          // 签名方式，长度32，微信支付时返回
	SdkPaysign         string `json:"sdk_paysign,omitempty"`           // 签名，长度64，微信支付时返回
	SdkPartnerid       string `json:"sdk_partnerid,omitempty"`         // trade_type为APP时返回，长度32，非必填
	ReservedFyOrderNo  string `json:"reserved_fy_order_no,omitempty"`  // 富友生成的订单号，长度30，非必填
	ReservedFySettleDt string `json:"reserved_fy_settle_dt,omitempty"` // 富友清算日，长度8，非必填
	// 支付宝交易流水号
	ReservedTransactionId  string `json:"reserved_transaction_id,omitempty"`   // 渠道交易流水号，长度32，trade_type为FWC时返回，非必填
	ReservedFyTraceNo      string `json:"reserved_fy_trace_no"`                // 追踪号，长度12，必填
	ReservedPayInfo        string `json:"reserved_pay_info,omitempty"`         // 支付参数，不定长，非必填
	ReservedChannelOrderId string `json:"reserved_channel_order_id,omitempty"` // 银行交易号，长度32，非必填
	ReservedAddnInf        string `json:"reserved_addn_inf,omitempty"`         // 附加数据，长度50，非必填
	Sign                   string `json:"sign"`
}

func (p PaymentResponse) IsSuccess() bool {
	return p.ResultCode == "000000"
}

// OrderRequest 订单请求参数结构体
type OrderRequest struct {
	Version      string `json:"version"`        // 版本号，长度8，必填
	MchntCd      string `json:"mchnt_cd"`       // 商户号，长度15，必填
	RandomStr    string `json:"random_str"`     // 随机字符串，长度32，必填
	OrderType    string `json:"order_type"`     // 订单类型，长度20，必填，可选值如ALIPAY、WECHAT等
	MchntOrderNo string `json:"mchnt_order_no"` // 商户订单号，长度30，必填
	TermId       string `json:"term_id"`        // 终端号，长度8，必填
	// mchnt_cd+"|"
	//+ order_type +"|"
	//+ mchnt_order_no+"|"
	//+ term_id +"|"
	//+ random_str +"|"
	//+ version + "|"
	//+ mchnt_key
	Sign string `json:"sign"` // 签名，必填
}

func (o *OrderRequest) GenSign(apiKey string) {
	signStrArr := make([]string, 0, 7)
	signStrArr = append(signStrArr, o.MchntCd)
	signStrArr = append(signStrArr, o.OrderType)
	signStrArr = append(signStrArr, o.MchntOrderNo)
	signStrArr = append(signStrArr, o.TermId)
	signStrArr = append(signStrArr, o.RandomStr)
	signStrArr = append(signStrArr, o.Version)
	signStrArr = append(signStrArr, apiKey)
	o.Sign = tools.Md5(strings.Join(signStrArr, "|"))
}

// OrderDetail 订单查询响应报文结构体
type OrderDetail struct {
	ResultCode string `json:"result_code"`       // 响应代码，长度16，必填
	ResultMsg  string `json:"result_msg"`        // 中文描述，长度128，必填
	MchntCd    string `json:"mchnt_cd"`          // 商户代码，长度15，必填
	TermId     string `json:"term_id,omitempty"` // 终端号，长度8，非必填
	RandomStr  string `json:"random_str"`        // 随机字符串，长度32，必填
	//订单类型：
	//ALIPAY (统一下单、条码支付、服务窗支付)
	//WECHAT(统一下单、条码支付，公众号支付)
	//UNIONPAY
	//WXAPP(微信app)
	//ALIAPP(支付宝app)
	//WXH5(微信h5)
	//ALIH5(支付宝h5)
	//WXBX(微信保险类)
	//ALBX(支付宝保险类)
	//【公众号相关的交易为空，其他都有】
	OrderType          string `json:"order_type"`            // 订单类型，长度20，必填，可选值如ALIPAY、WECHAT等
	OrderAmt           string `json:"order_amt"`             // 订单金额，长度16，以分为单位的整数，必填
	BuyerId            string `json:"buyer_id,omitempty"`    // 买家渠道号，长度32，非必填
	TransactionId      string `json:"transaction_id"`        // 渠道交易流水号，长度32，必填
	AddnInf            string `json:"addn_inf,omitempty"`    // 附加数据，长度100，非必填
	ReservedFySettleDt string `json:"reserved_fy_settle_dt"` // 富友清算日，长度8，必填
	MchntOrderNo       string `json:"mchnt_order_no"`        // 商户订单号，长度30，必填
	//SUCCESS—支付成功
	//REFUND—已退款
	//NOTPAY—未支付
	//CLOSED—已关闭
	//REVOKED—已撤销（刷卡支付）
	//USERPAYING--用户支付中
	//PAYERROR--支付失败(其他原因，如银行返回失败)
	TransStat string `json:"trans_stat"` // 交易状态，长度32，必填，可选值如SUCCESS、REFUND等
	// 新增字段（当前部分）
	ReservedCouponFee      string `json:"reserved_coupon_fee,omitempty"`       // 优惠金额，长度10，以分为单位，非必填
	ReservedBuyerLogonId   string `json:"reserved_buyer_logon_id,omitempty"`   // 买家登录账号，长度128，非必填
	ReservedFundBillList   string `json:"reserved_fund_bill_list,omitempty"`   // 渠道信息，不定长，非必填
	ReservedFyTraceNo      string `json:"reserved_fy_trace_no"`                // 富友系统内部追踪号，长度12，必填
	ReservedChannelOrderId string `json:"reserved_channel_order_id,omitempty"` // 银行交易号，长度32，非必填
	ReservedFyTermId       string `json:"reserved_fy_term_id,omitempty"`       // 富友终端号，长度20，非必填
	ReservedIsCredit       string `json:"reserved_is_credit,omitempty"`        // 信用卡标识，长度8，非必填，1表示信用卡/花呗，0表示其他，不填表示未知
	ReservedBankType       string `json:"reserved_bank_type,omitempty"`        // 付款方式，长度16，非必填
	ReservedTxnFinTs       string `json:"reserved_txn_fin_ts,omitempty"`       // 支付完成时间，长度14，格式yyyyMMddHHmmss，非必填
}

func (o *OrderDetail) IsSuccess() bool {
	return o.ResultCode == "000000"
}

// CloseOrderRequest 关闭订单请求报文结构体
type CloseOrderRequest struct {
	Version      string `json:"version"`        // 版本号，长度8，必填
	MchntCd      string `json:"mchnt_cd"`       // 商户号，长度15，必填
	RandomStr    string `json:"random_str"`     // 随机字符串，长度32，必填
	TermId       string `json:"term_id"`        // 终端号，长度8，必填
	MchntOrderNo string `json:"mchnt_order_no"` // 商户订单号，长度30，必填
	//订单类型：
	//ALIPAY(统一下单、服务窗支付)
	//WECHAT(统一下单、公众号支付、小程序)
	//WXAPP(微信app)
	//WXH5(微信h5)
	OrderType string `json:"order_type"`          // 订单类型，长度20，必填，可选值如ALIPAY、WECHAT等
	SubAppid  string `json:"sub_appid,omitempty"` // 子商户公众号，长度32，非必填（子商户配置多个公众号时必填）
	//拼接
	//
	//mchnt_cd+"|"
	//+ mchnt_order_no +"|"
	//+ order_type +"|"
	//+ term_id +"|"
	//+ random_str +"|"
	//+ version + "|"
	//+ mchnt_key
	//
	//做md5摘要
	//其中mchnt_key为商户密钥，系统分
	Sign string `json:"sign"` // 签名，必填
}

func (c *CloseOrderRequest) GenSign(apiKey string) {
	signStrArr := make([]string, 0, 7)
	signStrArr = append(signStrArr, c.MchntCd)
	signStrArr = append(signStrArr, c.MchntOrderNo)
	signStrArr = append(signStrArr, c.OrderType)
	signStrArr = append(signStrArr, c.TermId)
	signStrArr = append(signStrArr, c.RandomStr)
	signStrArr = append(signStrArr, c.Version)
	signStrArr = append(signStrArr, apiKey)
	c.Sign = tools.Md5(strings.Join(signStrArr, "|"))
}
func (c *CloseOrderRequest) Validate() error {
	if c.MchntOrderNo == "" {
		return errors2.ErrorParamError("mchnt_order_no is required")
	}
	return nil
}

// CloseOrderResponse 关闭订单响应报文结构体
type CloseOrderResponse struct {
	ResultCode   string `json:"result_code"`          // 响应代码，长度16，必填
	ResultMsg    string `json:"result_msg"`           // 中文描述，长度128，必填
	InsCd        string `json:"ins_cd"`               // 机构号，长度20，必填
	MchntCd      string `json:"mchnt_cd"`             // 商户号，长度15，必填
	TermId       string `json:"term_id"`              // 终端号，长度8，必填
	RandomStr    string `json:"random_str"`           // 随机字符串，长度32，必填
	OrderType    string `json:"order_type,omitempty"` // 订单类型，长度20，非必填
	MchntOrderNo string `json:"mchnt_order_no"`       // 商户订单号，长度30，必填
	Sign         string `json:"sign"`                 // 签名，必填
}

func (o CloseOrderResponse) IsSuccess() bool {
	return o.ResultCode == "000000"
}
