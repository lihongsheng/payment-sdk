package model

import (
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/tools"
	"strings"
)

type PaymentCallback struct {
	// 已有字段（承接上一部分）
	ResultCode string `json:"result_code"`
	ResultMsg  string `json:"result_msg"`
	MchntCd    string `json:"mchnt_cd"`
	TermId     string `json:"term_id,omitempty"`
	RandomStr  string `json:"random_str"`
	// order_type
	// 订单类型：
	//ALIPAY
	//WECHAT
	//UNIONPAY(银联二维码)等
	OrderType string `json:"order_type"`
	// transaction_id
	TransactionId string `json:"transaction_id"`
	// user_id
	UserId string `json:"user_id"`
	// 9. 商户订单号
	MchntOrderNo string `json:"mchnt_order_no"`
	// 10. 订单金额
	OrderAmt string `json:"order_amt"`
	// 11. 应结订单金额
	SettleOrderAmt string `json:"settle_order_amt"`
	// 12. 货币种类
	CurrType string `json:"curr_type,omitempty"`
	// 13. 支付完成时间（格式：yyyyMMddHHmmss）
	TxnFinTs string `json:"txn_fin_ts"`
	// 14. 富友清算日
	ReservedFySettleDt string `json:"reserved_fy_settle_dt"`
	// 15. 优惠金额
	ReservedCouponFee string `json:"reserved_coupon_fee,omitempty"`
	// 16. 买家登录账号
	ReservedBuyerLogonId string `json:"reserved_buyer_logon_id,omitempty"`
	// 17. 渠道信息
	ReservedFundBillList string `json:"reserved_fund_bill_list,omitempty"`
	// 18. 富友追踪号
	ReservedFyTraceNo string `json:"reserved_fy_trace_no"`
	// 19. 银行交易号
	ReservedChannelOrderId string `json:"reserved_channel_order_id,omitempty"`
	// 20. 信用卡标识（1-信用卡/花呗；0-其他；不填-未知）
	ReservedIsCredit string `json:"reserved_is_credit,omitempty"`
	Sign             string `json:"sign"`
}

func (p PaymentCallback) IsSuccess() bool {
	return p.ResultCode == "000000"
}

func (c *PaymentCallback) GenSign(apiKey string) string {
	signStrArr := make([]string, 0, 7)
	signStrArr = append(signStrArr, c.MchntCd)
	signStrArr = append(signStrArr, c.MchntOrderNo)
	signStrArr = append(signStrArr, fmt.Sprintf("%s", c.SettleOrderAmt))
	signStrArr = append(signStrArr, fmt.Sprintf("%s", c.OrderAmt))
	signStrArr = append(signStrArr, c.TxnFinTs)
	signStrArr = append(signStrArr, c.ReservedFySettleDt)
	signStrArr = append(signStrArr, c.RandomStr)
	signStrArr = append(signStrArr, apiKey)
	return tools.Md5(strings.Join(signStrArr, "|"))
}
