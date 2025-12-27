package refund

import (
	"fmt"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"strings"
)

type RefundRequest struct {
	Version      string `json:"version"`        // 版本号，长度8，必填
	MchntCd      string `json:"mchnt_cd"`       // 商户号，长度15，必填
	RandomStr    string `json:"random_str"`     // 随机字符串，长度32，必填
	MchntOrderNo string `json:"mchnt_order_no"` // 商户订单号，长度30，必填
	// refund_order_no
	RefundOrderNo string `json:"refund_order_no"`   // 商户退款单号(也要加订单前缀)
	TermId        string `json:"term_id,omitempty"` // 终端号，长度8，随机8字节数字字母组合，必填
	// IP
	TermIp string `json:"term_ip"` // 终端IP，长度16，必填
	// 订单类型：
	//ALIPAY(统一下单、条码支付、服务窗支付)
	//WECHAT(统一下单、条码支付、公众号支付、小程序)
	//UNIONPAY
	//WXAPP(微信app)
	//ALIAPP(支付宝app)
	//WXH5(微信h5)
	//ALIH5(支付宝h5)
	//WXBX(微信保险类)
	//ALBX(支付宝保险类)
	//WXXS(微信线上所有交易)(不支持线下类型)
	OrderType string `json:"order_type"` // 附加数据，长度100，非必填
	// 币种
	TotalAmt int64 `json:"total_amt"`
	// refund_amt
	RefundAmt          int64  `json:"refund_amt"`
	ReservedFyTermId   string `json:"reserved_fy_term_id,omitempty"` // 富友终端号，长度20，非必填
	ReservedDeviceInfo string `json:"reserved_device_info"`
	//拼接
	//
	// 拼接
	//
	//mchnt_cd+"|"
	//+ order_type +"|"
	//+ mchnt_order_no +"|"
	//+ refund_order_no+"|"
	//+ total_amt+"|"
	//+ refund_amt +"|"
	//+ term_id +"|"
	//+ random_str +"|"
	//+ version +"|"
	//+ mchnt_key
	//
	//并做md5摘要
	//其中 mchnt_key 为商户密钥，系统分配
	//
	//并做md5摘要
	//其中 mchnt_key 为商户密钥，系统分配
	Sign string `json:"sign"`
}

func (p *RefundRequest) Validate() error {
	if p.MchntCd == "" {
		return errors2.ErrorParamError("商户号不能为空")
	}
	if p.RandomStr == "" {
		return errors2.ErrorParamError("随机字符串不能为空")
	}
	if p.RefundAmt <= 0 {
		return errors2.ErrorParamError("退款金额不能小于0")
	}
	if p.OrderType == "" {
		return errors2.ErrorParamError("订单类型不能为空")
	}
	return nil
}

func (p *RefundRequest) GenSign(apiKey string) {
	signStrArr := []string{}
	signStrArr = append(signStrArr, p.MchntCd)
	signStrArr = append(signStrArr, p.OrderType)
	signStrArr = append(signStrArr, p.MchntOrderNo)
	signStrArr = append(signStrArr, p.RefundOrderNo)
	signStrArr = append(signStrArr, fmt.Sprintf("%d", p.TotalAmt))
	signStrArr = append(signStrArr, fmt.Sprintf("%d", p.RefundAmt))
	signStrArr = append(signStrArr, p.TermId)
	signStrArr = append(signStrArr, p.RandomStr)
	signStrArr = append(signStrArr, p.Version)
	signStrArr = append(signStrArr, apiKey)
	p.Sign = tools.Md5(strings.Join(signStrArr, "|"))
}

type RefundResponse struct {
	ResultCode string `json:"result_code"`
	ResultMsg  string `json:"result_msg"`
	MchntCd    string `json:"mchnt_cd"`
	TermId     string `json:"term_id,omitempty"`
	RandomStr  string `json:"random_str"`
	// transaction_id
	TransactionId string `json:"transaction_id"`
	// refund_id
	RefundId     string `json:"refund_id"`
	MchntOrderNo string `json:"mchnt_order_no"` // 商户订单号，长度30，必填
	// refund_order_no
	RefundOrderNo string `json:"refund_order_no"` // 商户退款单号(也要加订单前缀)
	//reserved_refund_amt
	ReservedRefundAmt string `json:"reserved_refund_amt"`
	Sign              string `json:"sign"`
}

func (o *RefundResponse) IsSuccess() bool {
	return o.ResultCode == "000000"
}

type RefundQueryRequest struct {
	Version   string `json:"version"`    // 版本号，长度8，必填
	MchntCd   string `json:"mchnt_cd"`   // 商户号，长度15，必填
	RandomStr string `json:"random_str"` // 随机字符串，长度32，必填
	// refund_order_no
	RefundOrderNo string `json:"refund_order_no"`   // 商户退款单号(也要加订单前缀)
	TermId        string `json:"term_id,omitempty"` // 终端号，长度8，随机8字节数字字母组合，必填
	Sign          string `json:"sign"`
}

func (p *RefundQueryRequest) GenSign(apiKey string) {
	signStrArr := []string{}
	signStrArr = append(signStrArr, p.MchntCd)
	signStrArr = append(signStrArr, p.RefundOrderNo)
	signStrArr = append(signStrArr, p.TermId)
	signStrArr = append(signStrArr, p.RandomStr)
	signStrArr = append(signStrArr, p.Version)
	signStrArr = append(signStrArr, apiKey)
	p.Sign = tools.Md5(strings.Join(signStrArr, "|"))
}

type RefundQueryResponse struct {
	ResultCode string `json:"result_code"`
	ResultMsg  string `json:"result_msg"`
	MchntCd    string `json:"mchnt_cd"`
	TermId     string `json:"term_id,omitempty"`
	RandomStr  string `json:"random_str"`
	// transaction_id
	TransactionId string `json:"transaction_id"`
	// refund_id
	RefundId     string `json:"refund_id"`
	MchntOrderNo string `json:"mchnt_order_no"` // 商户订单号，长度30，必填
	// refund_order_no
	RefundOrderNo string `json:"refund_order_no"` // 商户退款单号(也要加订单前缀)
	//reserved_refund_amt
	ReservedRefundAmt string `json:"reserved_refund_amt"`
	// 交易状态
	//SUCCESS—退款成功
	//PAYERROR--退款失败
	//USERPAYING-退款已受理(超时或状态未知，等T+1富友处理)
	TransStat string `json:"trans_stat"`
	Sign      string `json:"sign"`
}

func (o *RefundQueryResponse) IsSuccess() bool {
	return o.ResultCode == "000000"
}
func (o *RefundQueryResponse) GetRefundStatus() refund.Status {
	status := strings.ToUpper(o.TransStat)
	switch status {
	case "SUCCESS":
		return refund.Status_Success
	case "PAYERROR":
		return refund.Status_Failed
	case "USERPAYING":
		return refund.Status_Pending
	}
	return refund.Status_Status_UNKNOWN
}
