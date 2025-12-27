package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"time"
)

// 定义税筹相关的子结构体
type Tax struct {
	TagId              int    `xml:"tagId,omitempty"`              // 标签id
	TaskDescription    string `xml:"taskDescription,omitempty"`    // 任务描述
	AcceptanceStandard string `xml:"acceptanceStandard,omitempty"` // 任务验收标准
}

// 定义入账方列表的item子结构体
type TransferAccountInListItem struct {
	AccountIn    string `xml:"accountIn,omitempty"`    // 用户编号
	AllocateAmt  int    `xml:"allocateAmt,omitempty"`  // 分账金额
	AllocateType string `xml:"allocateType,omitempty"` // 分账方式
	CleanBankNo  string `xml:"cleanBankNo,omitempty"`  // 入账卡号
	InvoiceType  string `xml:"invoiceType,omitempty"`  // 发票形式
	Tax          *Tax   `xml:"tax,omitempty"`          // 税筹
	Remark       string `xml:"remark,omitempty"`       // 打款附言
}

func (u *TransferAccountInListItem) Validate() error {
	if u.AccountIn == "" {
		return fmt.Errorf("入账方不能为空")
	}
	if u.AllocateAmt <= 0 {
		return fmt.Errorf("分账金额不能小于等于0")
	}
	if u.AllocateType == "" {
		u.AllocateType = string(transfer.AllocateType_Auto)
	}
	if u.CleanBankNo == "" {
		return fmt.Errorf("入账卡号不能为空")
	}
	return nil
}

// 定义主请求结构体
type TransferRequest struct {
	XMLName       xml.Name                    `xml:"xml"`
	TraceNo       string                      `xml:"traceNo,omitempty" json:"traceNo"`                   // 商户流水号
	MchntCd       string                      `xml:"mchntCd,omitempty" json:"mchntCd"`                   // 商户号
	AccountIn     string                      `xml:"accountIn,omitempty" json:"accountIn"`               // 商户自有子账户编号
	AccountInList []TransferAccountInListItem `xml:"accountInlist>item,omitempty"  json:"accountInlist"` // 入账方列表
	Signature     string                      `xml:"signature,omitempty" json:"signature"`               // 签名
}

func (u *TransferRequest) Validate() error {
	if u.TraceNo == "" {
		u.TraceNo = fmt.Sprintf("%d", time.Now().UnixMilli())
	}
	if u.MchntCd == "" {
		return fmt.Errorf("商户号不能为空")
	}
	if len(u.AccountInList) == 0 {
		return fmt.Errorf("入账方列表不能为空")
	}
	return nil
}

func (u *TransferRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"randomStr",
		"accountInlist",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	return sign
}

func (u *TransferRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *TransferRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

type TransferResponse struct {
	XMLName      xml.Name `xml:"xml"`
	TraceNo      string   `xml:"traceNo,omitempty"`   // 商户流水号
	MchntCd      string   `xml:"mchntCd,omitempty"`   // 商户号
	RespCode     string   `xml:"respCode,omitempty"`  // 响应码
	RespDesc     string   `xml:"respDesc,omitempty"`  // 响应码描述
	Signature    string   `xml:"signature,omitempty"` // 签名
	BatchNo      string   `xml:"batchNo,omitempty"`
	AllocateDate string   `xml:"allocateDate,omitempty"`
}

type TransferQueryRequest struct {
	XMLName             xml.Name `xml:"request"`
	TraceNo             string   `xml:"traceNo,omitempty" json:"traceNo"`                         // 商户流水号
	MchntCd             string   `xml:"mchntCd,omitempty" json:"mchntCd"`                         // 商户号
	TradeType           string   `xml:"tradeType,omitempty" json:"tradeType"`                     // 交易类型
	BatchNo             string   `xml:"batchNo,omitempty" json:"batchNo"`                         // 富友批次号
	SrcFasSsn           string   `xml:"srcFasSsn,omitempty" json:"srcFasSsn"`                     // 富友交易参考号
	MchntCdTraceNo      string   `xml:"mchntCdTraceNo,omitempty" json:"mchntCdTraceNo,"`          // 源交易商户流水号
	MchntCdChildTraceNo string   `xml:"mchntCdChildTraceNo,omitempty" json:"mchntCdChildTraceNo"` // 源交易商户子流水号
	StartDate           string   `xml:"startDate,omitempty" json:"startDate"`                     // 起始日期
	EndDate             string   `xml:"endDate,omitempty" json:"endDate"`                         // 截止日期
	PageNo              string   `xml:"pageNo,omitempty" json:"pageNo"`                           // 页码
	PageSize            string   `xml:"pageSize,omitempty" json:"pageSize"`                       // 分页大小
	Signature           string   `xml:"signature,omitempty" json:"signature"`                     // 签名
}

func (u *TransferQueryRequest) Validate() error {
	if u.TraceNo == "" {
		u.TraceNo = fmt.Sprintf("%d", time.Now().UnixMilli())
	}
	if u.MchntCd == "" {
		return fmt.Errorf("商户号不能为空")
	}
	if u.PageNo == "" {
		u.PageNo = "1"
	}
	if u.PageSize == "" {
		u.PageSize = "10"
	}
	return nil
}

func (u *TransferQueryRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"randomStr",
		"batchNo",
		"srcFasSsn",
		"mchntCdTraceNo",
		"mchntCdChildTraceNo",
		"tradeType",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	return sign
}

func (u *TransferQueryRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *TransferQueryRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

// 定义List内的item子结构体
type ListItem struct {
	TradeTime           string `xml:"tradeTime,omitempty"`           // 交易时间
	BatchNo             string `xml:"batchNo,omitempty"`             // 富友批次号
	MchntCdTraceNo      string `xml:"mchntCdTraceNo,omitempty"`      // 源交易商户流水号
	MchntCdChildTraceNo string `xml:"mchntCdChildTraceNo,omitempty"` // 源交易商户子流水号
	SrcFasSsn           string `xml:"srcFasSsn,omitempty"`           // 富友交易参考号
	TxnAmt              int    `xml:"txnAmt,omitempty"`              // 交易金额
	AccountOut          string `xml:"accountOut,omitempty"`          // 转出用户编号或商户号
	AccountIn           string `xml:"accountIn,omitempty"`           // 转入用户编号或商户号
	Status              string `xml:"status,omitempty"`              // 状态
	ErrorMsg            string `xml:"errorMsg,omitempty"`            // 失败原因
	TradeType           string `xml:"tradeType,omitempty"`           // 交易类型
	Remark              string `xml:"remark,omitempty"`              // 备注
}

// 定义主请求/响应结构体
type TransferQueryResponse struct {
	XMLName   xml.Name   `xml:"xml"`
	TraceNo   string     `xml:"traceNo,omitempty"`   // 商户流水号
	MchntCd   string     `xml:"mchntCd,omitempty"`   // 商户号
	StartDate string     `xml:"startDate,omitempty"` // 起始日期
	EndDate   string     `xml:"endDate,omitempty"`   // 截止日期
	TotalNum  int        `xml:"totalNum,omitempty"`  // 查询结果总数
	RespCode  string     `xml:"respCode,omitempty"`  // 接口返回码
	RespDesc  string     `xml:"respDesc,omitempty"`  // 接口返回码描述
	Signature string     `xml:"signature,omitempty"` // 签名
	List      []ListItem `xml:"List>item,omitempty"` // 列表数据
}

//type Request struct {
//	XMLName            xml.Name `xml:"request"`
//	TraceNo            string   `xml:"traceNo,omitempty"`            // 商户流水号
//	MchntCd            string   `xml:"mchntCd,omitempty"`            // 商户号
//	SubaccountIn       string   `xml:"subaccountIn,omitempty"`       // 用户编号
//	MchntCdConcentrate string   `xml:"mchntCdConcentrate,omitempty"` // 被归集商户号
//	TradeSsn           string   `xml:"tradeSsn,omitempty"`           // 交易参考号
//	StartDate          string   `xml:"startDate,omitempty"`          // 起始日期
//	EndDate            string   `xml:"endDate,omitempty"`            // 截止日期
//	PageNo             int      `xml:"pageNo,omitempty"`             // 页码
//	PageSize           int      `xml:"pageSize,omitempty"`           // 分页大小
//	BusiCd             string   `xml:"busiCd,omitempty"`             // 业务类型
//	Signature          string   `xml:"signature,omitempty"`          // 签名
//}

// 定义data节点的结构体
type TransferCallbackData struct {
	MchntCd          string `xml:"mchntCd,omitempty"`          // 商户号
	AccountIn        string `xml:"accountIn,omitempty"`        // 用户编号
	BatchNo          string `xml:"batchNo,omitempty"`          // 富友批次号
	SettleType       string `xml:"settleType,omitempty"`       // 结算方式
	OutAcctNo        string `xml:"outAcctNo,omitempty"`        // 入账卡号
	Amt              int64  `xml:"amt,omitempty"`              // 金额
	Fee              int64  `xml:"fee,omitempty"`              // 手续费
	GenerationTaxFee int64  `xml:"generationTaxFee,omitempty"` // 代税服务费
	TaxFee           int64  `xml:"taxFee,omitempty"`           // 税费
	ArrivalAmt       int64  `xml:"arrivalAmt,omitempty"`       // 到账金额
	InvoiceType      string `xml:"invoiceType,omitempty"`      // 发票形式
	Status           string `xml:"status,omitempty"`           // 分账状态
	AllocateStatus   string `xml:"allocateStatus,omitempty"`   // 结算状态
	Desc             string `xml:"desc,omitempty"`             // 通知描述
	OriTraceNo       string `xml:"oriTraceNo,omitempty"`       // 源商户流水号
}

// 定义主请求结构体
type TransferCallbackRequest struct {
	XMLName    xml.Name             `xml:"request"`
	NotifyType string               `xml:"notifyType,omitempty"` // 回调类型
	Data       TransferCallbackData `xml:"data,omitempty"`       // data节点
	Signature  string               `xml:"signature,omitempty"`  // 签名
}
