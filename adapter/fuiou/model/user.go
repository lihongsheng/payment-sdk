package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/fuiou/enum"
	"github.com/singer-stack-lab/payment-sdk/errors"
	"time"
)

// UserCreateRequest 包含所有参数的结构体，对应顶层 XML 元素 <message>
// allocateScale=&bcpNo=&busiLicAddr=&busiLicPic=&busiLicValidateEnd=&busiLicValidateStart=&certNo=110101199001011234&certTp=0&channel=&checkType=&cleanType=01&contactCertNo=110101199001011234&contactEmail=&contactName=米昂&interBankNo=&legalCertNo=&legalCertTp=&legalImagB=&legalImagF=&legalMobile=&legalName=&legalValidateEnd=&legalValidateStart=&mchntCd=0002900F0370542&mchntCdUserId=1&mobile=13800000001&organizationType=&outAcntNm=米昂&outAcntNo=62221199001011234&protocolType=01&traceNo=1
// allocateScale=&bcpNo=&busilLicAddr=&busilLicPic=&busilLicValidateEnd=&busilLicValidateStart=&certNo=110101199001011234&certTp=0&channel=&checkType=&cleanType=01&contactCertNo=110101199001011234&contactEmail=&contactName=米昂&interBankNo=&legalCertNo=&legalCertTp=&legalImagB=&legalImagF=&legalMobile=&legalName=&legalValidateEnd=&legalValidateStart=&mchntCd=0002900F0370542&mchntCdUserId=1&mobile=13800000001&organizationType=&outAcntNm=米昂&outAcntNo=62221199001011234&protocolType=01&traceNo=1
type UserCreateRequest struct {
	XMLName              xml.Name `xml:"xml"`                                                        // 根元素标签
	TraceNo              string   `json:"traceNo" xml:"traceNo,omitempty"`                           // 商户流水号
	MchntCd              string   `json:"mchntCd" xml:"mchntCd,omitempty"`                           // 商户号
	CleanType            string   `json:"cleanType" xml:"cleanType,omitempty"`                       // 账户类型
	OutAcntNm            string   `json:"outAcntNm" xml:"outAcntNm,omitempty"`                       // 户名
	CertTp               string   `json:"certTp" xml:"certTp,omitempty"`                             // 证件类型
	CertNo               string   `json:"certNo" xml:"certNo,omitempty"`                             // 证件号码
	Channel              string   `json:"channel" xml:"channel,omitempty"`                           // 银行账户开户行
	OrganizationType     string   `json:"organizationType" xml:"organizationType,omitempty"`         // 主体类型
	BcpNo                string   `json:"bcpNo" xml:"bcpNo,omitempty"`                               // 基本账户开户许可核准号
	BusiLicValidateStart string   `json:"busiLicValidateStart" xml:"busiLicValidateStart,omitempty"` // 营业执照有效期起始时间
	BusiLicValidateEnd   string   `json:"busiLicValidateEnd" xml:"busiLicValidateEnd,omitempty"`     // 营业执照有效期结束时间
	BusiLicAddr          string   `json:"busiLicAddr" xml:"busiLicAddr,omitempty"`                   // 营业执照注册地址
	BusiLicPic           string   `json:"busiLicPic" xml:"busiLicPic,omitempty"`                     // 营业执照照片路径
	LegalName            string   `json:"legalName" xml:"legalName,omitempty"`                       // 法人姓名
	LegalMobile          string   `json:"legalMobile" xml:"legalMobile,omitempty"`                   // 法人手机号
	LegalCertTp          string   `json:"legalCertTp" xml:"legalCertTp,omitempty"`                   // 法人证件类型
	LegalCertNo          string   `json:"legalCertNo" xml:"legalCertNo,omitempty"`                   // 法人证件号
	LegalValidateStart   string   `json:"legalValidateStart" xml:"legalValidateStart,omitempty"`     // 法人证件有效期起始日
	LegalValidateEnd     string   `json:"legalValidateEnd" xml:"legalValidateEnd,omitempty"`         // 法人证件有效期到期日
	LegalImagF           string   `json:"legalImagF" xml:"legalImagF,omitempty"`                     // 法人证件国徽面照片路径
	LegalImagB           string   `json:"legalImagB" xml:"legalImagB,omitempty"`                     // 法人证件人像面照片路径
	Mobile               string   `json:"mobile" xml:"mobile,omitempty"`                             // 联系人手机号/个人用户手机号
	ContactName          string   `json:"contactName" xml:"contactName,omitempty"`                   // 联系人姓名
	ContactEmail         string   `json:"contactEmail" xml:"contactEmail,omitempty"`                 // 联系人邮箱
	ContactCertNo        string   `json:"contactCertNo" xml:"contactCertNo,omitempty"`               // 联系人身份证号
	OutAcntNoType        string   `json:"outAcntNoType" xml:"outAcntNoType,omitempty"`               // 银行账号类型
	OutAcntNo            string   `json:"outAcntNo" xml:"outAcntNo,omitempty"`                       // 银行账号
	InterBankNo          string   `json:"interBankNo" xml:"interBankNo,omitempty"`                   // 开户行行号
	AllocateScale        string   `json:"allocateScale" xml:"allocateScale,omitempty"`               // 分账比例
	ProtocolType         string   `json:"protocolType" xml:"protocolType,omitempty"`                 // 协议类型
	UseESign             string   `json:"useESign" xml:"useESign,omitempty"`                         // 是否使用电子签章
	MchntCdUserId        string   `json:"mchntCdUserId" xml:"mchntCdUserId,omitempty"`               // 商户用户id
	MiniAppReturnPath    string   `json:"miniAppReturnPath" xml:"miniAppReturnPath,omitempty"`       // 微信小程序回跳路径
	CheckType            string   `json:"checkType" xml:"checkType,omitempty"`                       // 验证类型
	ExtendInfo           string   `json:"extendInfo" xml:"extendInfo,omitempty"`
	Signature            string   `json:"signature"  xml:"signature,omitempty"`
}

func (u *UserCreateRequest) Validate() error {
	if u.OutAcntNo == "" {
		return errors.ErrorParamError("out_acnt_no must not empty")
	}
	if u.CleanType == enum.CleanType_Personal {
		// 身份证
		if u.CertTp != enum.CertTp_IDCard {
			return errors.ErrorParamError("cert_tp is ID_CARD, cert_no must not empty")
		}
	}
	if u.CertNo == "" {
		return errors.ErrorParamError("cert_no must not empty")
	}
	return nil
}

func (u *UserCreateRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"randomStr",
		"outAcntNoType",
		"useESign",
		"miniAppReturnPath",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	if u.ExtendInfo == "" {
		delete(sign, "extendInfo")
	}
	return sign
}

func (u *UserCreateRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *UserCreateRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

type SignResponse struct {
	XMLName  xml.Name `xml:"xml"`
	Message  string   `xml:"message,omitempty"` // 商户流水号
	MchntCd  string   `xml:"mchntCd,omitempty"` // 商户号
	RespDesc string   `xml:"respDesc"`
}

type UserCreateResponse struct {
	XMLName    xml.Name `xml:"xml"`
	TraceNo    string   `xml:"traceNo,omitempty"`    // 商户流水号
	MchntCd    string   `xml:"mchntCd,omitempty"`    // 商户号
	RespCode   string   `xml:"respCode,omitempty"`   // 返回码
	RespDesc   string   `xml:"respDesc,omitempty"`   // 返回信息
	AccountIn  string   `xml:"accountIn,omitempty"`  // 用户编号
	KsCheckUrl string   `xml:"ksCheckUrl,omitempty"` // 旷世验证h5链接
	CheckUrl   string   `xml:"checkUrl,omitempty"`   // 阿里验证h5链接
	Signature  string   `xml:"signature,omitempty"`  // 签名
}

type UserActiveRequest struct {
	XMLName           xml.Name `xml:"xml"`
	TraceNo           string   `xml:"traceNo,omitempty" json:"traceNo"`                     // 商户流水号
	MchntCd           string   `xml:"mchntCd,omitempty" json:"mchntCd"`                     // 商户号
	AccountIn         string   `xml:"accountIn,omitempty" json:"accountIn"`                 // 用户编号
	CheckType         string   `xml:"checkType,omitempty" json:"checkType"`                 // 验证类型
	MiniAppReturnPath string   `xml:"miniAppReturnPath,omitempty" json:"miniAppReturnPath"` // 微信小程序回跳路径
	Signature         string   `xml:"signature,omitempty" json:"signature"`                 // 签名
}

func (u *UserActiveRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"randomStr",
		"outAcntNoType",
		"useESign",
		"miniAppReturnPath",
		"checkType",
		"miniAppReturnPath",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	return sign
}

func (u *UserActiveRequest) Validate() error {
	if u.TraceNo == "" {
		u.TraceNo = fmt.Sprintf("%d", time.Now().UnixMilli())
	}
	if u.CheckType == "" {
		u.CheckType = enum.CheckType_Mobile
	}
	if u.AccountIn == "" {
		return errors.ErrorParamError("AccountIn must not empty")
	}
	return nil
}

func (u *UserActiveRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *UserActiveRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

type UserDeleteRequest struct {
	XMLName   xml.Name `xml:"xml"`
	TraceNo   string   `xml:"traceNo,omitempty" json:"traceNo"`     // 商户流水号
	MchntCd   string   `xml:"mchntCd,omitempty" json:"mchntCd"`     // 商户号
	AccountIn string   `xml:"accountIn,omitempty" json:"accountIn"` // 用户编号
	Signature string   `xml:"signature,omitempty" json:"signature"` // 签名
}

func (u *UserDeleteRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"randomStr",
		"outAcntNoType",
		"useESign",
		"miniAppReturnPath",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	return sign
}

func (u *UserDeleteRequest) Validate() error {
	if u.TraceNo == "" {
		u.TraceNo = fmt.Sprintf("%d", time.Now().UnixMilli())
	}
	if u.AccountIn == "" {
		return errors.ErrorParamError("AccountIn must not empty")
	}
	return nil
}

func (u *UserDeleteRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *UserDeleteRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

type UserDeleteResponse struct {
	XMLName   xml.Name `xml:"xml"`
	TraceNo   string   `xml:"traceNo,omitempty"`   // 商户流水号
	MchntCd   string   `xml:"mchntCd,omitempty"`   // 商户号
	RespCode  string   `xml:"respCode,omitempty"`  // 响应码
	RespDesc  string   `xml:"respDesc,omitempty"`  // 响应信息
	Signature string   `xml:"signature,omitempty"` // 签名
	AccountIn string   `xml:"accountIn,omitempty"` // 用户编号
}

type UserUpdateMobileRequest struct {
	XMLName   xml.Name `xml:"xml"`
	TraceNo   string   `xml:"traceNo,omitempty" json:"traceNo"`     // 商户流水号
	MchntCd   string   `xml:"mchntCd,omitempty" json:"mchntCd"`     // 商户号
	AccountIn string   `xml:"accountIn,omitempty" json:"accountIn"` // 用户编号
	CheckType string   `xml:"checkType,omitempty" json:"checkType"` // 验证类型
	Mobile    string   `xml:"mobile,omitempty" json:"mobile"`       // 手机号码
	// 01渠道合作协议
	//02销售合作协议03账户专用协议
	//04渠道合作协议(无比例版本)
	//05销售合作协议(无比例版本)09无协议（2025.9.8已上线）
	//修改手机号时，非必传
	//ProtocolType string `xml:"protocolType,omitempty" json:"protocolType"` // 协议类型
	Signature string `xml:"signature,omitempty" json:"signature"` // 签名
}

type UserUpdateResponse struct {
	XMLName   xml.Name `xml:"xml"`
	TraceNo   string   `xml:"traceNo,omitempty"`   // 商户流水号
	MchntCd   string   `xml:"mchntCd,omitempty"`   // 商户号
	RespCode  string   `xml:"respCode,omitempty"`  // 响应码
	RespDesc  string   `xml:"respDesc,omitempty"`  // 响应信息
	Signature string   `xml:"signature,omitempty"` // 签名
	AccountIn string   `xml:"accountIn,omitempty"` // 用户编号
	Url       string   `xml:"url,omitempty"`
}

type UserUpdateAllocateScaleRequest struct {
	XMLName   xml.Name `xml:"xml"`
	TraceNo   string   `xml:"traceNo,omitempty" json:"traceNo"`     // 商户流水号
	MchntCd   string   `xml:"mchntCd,omitempty" json:"mchntCd"`     // 商户号
	AccountIn string   `xml:"accountIn,omitempty" json:"accountIn"` // 用户编号
	// 01渠道合作协议
	//02销售合作协议03账户专用协议
	//04渠道合作协议(无比例版本)
	//05销售合作协议(无比例版本)09无协议（2025.9.8已上线）
	//修改手机号时，非必传
	ProtocolType  string `xml:"protocolType,omitempty" json:"protocolType"`   // 协议类型
	AllocateScale string `xml:"allocateScale,omitempty" json:"allocateScale"` // 分账比例
	Signature     string `xml:"signature,omitempty" json:"signature"`         // 签名
}

func UserUpdateAllocateScaleRequestToMobileAndTransfer(u *UserUpdateAllocateScaleRequest) *UserUpdateMobileAndTransferRequest {
	return &UserUpdateMobileAndTransferRequest{
		TraceNo:       u.TraceNo,
		MchntCd:       u.MchntCd,
		AccountIn:     u.AccountIn,
		Type:          "1",
		AllocateScale: u.AllocateScale,
		ProtocolType:  u.ProtocolType,
		UseESign:      "",
		CheckType:     "1",
		Mobile:        "",
	}
}

func UserUpdateMobileRequestToMobileAndTransfer(u *UserUpdateMobileRequest) *UserUpdateMobileAndTransferRequest {
	return &UserUpdateMobileAndTransferRequest{
		TraceNo:       u.TraceNo,
		MchntCd:       u.MchntCd,
		AccountIn:     u.AccountIn,
		Type:          "2",
		AllocateScale: "",
		ProtocolType:  "",
		UseESign:      "",
		CheckType:     u.CheckType,
		Mobile:        u.Mobile,
	}
}

type UserUpdateMobileAndTransferRequest struct {
	XMLName       xml.Name `xml:"xml"`
	TraceNo       string   `xml:"traceNo,omitempty" json:"traceNo"`             // 商户流水号
	MchntCd       string   `xml:"mchntCd,omitempty" json:"mchntCd"`             // 商户号
	AccountIn     string   `xml:"accountIn,omitempty" json:"accountIn"`         // 用户编号
	Type          string   `xml:"type,omitempty" json:"type"`                   // 修改类型
	AllocateScale string   `xml:"allocateScale,omitempty" json:"allocateScale"` // 分账比例
	ProtocolType  string   `xml:"protocolType,omitempty" json:"protocolType"`   // 协议类型
	UseESign      string   `xml:"useESign,omitempty" json:"useESign"`           // 是否使用电子签章
	CheckType     string   `xml:"checkType,omitempty" json:"checkType"`         // 验证类型
	Mobile        string   `xml:"mobile,omitempty" json:"mobile"`               // 手机号码
	Signature     string   `xml:"signature,omitempty" json:"signature"`         // 签名
}

func (u *UserUpdateMobileAndTransferRequest) Validate() error {
	if u.TraceNo == "" {
		u.TraceNo = fmt.Sprintf("%d", time.Now().UnixMilli())
	}
	if u.Type == "2" {
		if u.AccountIn == "" && u.Mobile == "" {
			return errors.ErrorParamError("AccountIn and Mobile must not empty")
		}
	}

	if u.Type == "1" {
		if u.AccountIn == "" {
			return errors.ErrorParamError("AccountIn and Mobile must not empty")
		}
		if u.ProtocolType == "" {
			return errors.ErrorParamError("ProtocolType must not empty")
		}
		if u.AllocateScale == "" {
			return errors.ErrorParamError("AllocateScale must not empty")
		}
	}

	return nil
}

func (u *UserUpdateMobileAndTransferRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"type",
		//"allocateScale",
		//"protocolType",
		"useESign",
		//"checkType",
		//"mobile",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	return sign
}

func (u *UserUpdateMobileAndTransferRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *UserUpdateMobileAndTransferRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

type UserQueryRequest struct {
	XMLName xml.Name `xml:"xml"`
	// 请求流水号
	TraceNo string `json:"traceNo" xml:"traceNo,omitempty"` // 商户流水号
	// 商户号
	MchntCd string `json:"mchntCd" xml:"mchntCd,omitempty"` // 商户号
	// 开户银行请求流水号，非必填
	MchntTraceNo string `json:"mchntTraceNo" xml:"mchntTraceNo,omitempty"` // 开户请求流水号
	OutAcntNm    string `json:"outAcntNm" xml:"outAcntNm,omitempty"`       // 户名
	Mobile       string `json:"mobile" xml:"mobile,omitempty"`             // 手机号码
	OutAcntNo    string `json:"outAcntNo" xml:"outAcntNo,omitempty"`       // 银行账号
	// 富有的用户编号
	AccountIn string `json:"accountIn" xml:"accountIn,omitempty"` // 用户编号
	Signature string `json:"signature" xml:"signature,omitempty"` // 签名
}

func (u *UserQueryRequest) Validate() error {
	if u.TraceNo == "" {
		u.TraceNo = fmt.Sprintf("%d", time.Now().UnixMilli())
	}
	if u.AccountIn == "" && u.Mobile == "" {
		return errors.ErrorParamError("out_acnt_nm must not empty")
	}
	return nil
}

func (u *UserQueryRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"randomStr",
		"mchntTraceNo",
		"idNo",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	return sign
}

func (u *UserQueryRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *UserQueryRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

// 用户列表中的item子结构体
type AccountInListItem struct {
	MchntTraceNo    string `xml:"mchntTraceNo,omitempty"`    // 开户请求流水号
	AccountIn       string `xml:"accountIn,omitempty"`       // 用户编号
	RelationMchntCd string `xml:"relationMchntCd,omitempty"` // 关联商户
	AccountType     string `xml:"accountType,omitempty"`     // 账户类型
	InterBankNo     string `xml:"interBankNo,omitempty"`     // 开户行号
	OutAcctNm       string `xml:"outAcctNm,omitempty"`       // 户名
	Mobile          string `xml:"mobile,omitempty"`          // 手机号码
	IdNo            string `xml:"idNo,omitempty"`            // 证件号
	OutAcctNo       string `xml:"outAcctNo,omitempty"`       // 开户银行账号
	AllocateScale   int    `xml:"allocateScale,omitempty"`   // 分账比例
	Status          string `xml:"status,omitempty"`          // 状态
	BankAcctNo      string `xml:"bankAcctNo,omitempty"`      // 银行收款账户号
	BankInterNo     string `xml:"bankInterNo,omitempty"`     // 银行账户联行号
	BankNm          string `xml:"bankNm,omitempty"`          // 开户银行名称
	DepositAccount  string `xml:"depositAccount,omitempty"`  // 富友收款记账账户号
	DepositName     string `xml:"depositName,omitempty"`     // 富友收款记账账户名称
	IssInsName      string `xml:"issInsName,omitempty"`      // 开户行名称
	IssCityName     string `xml:"issCityName,omitempty"`     // 开户省市区
	SubBranchName   string `xml:"subBranchName,omitempty"`   // 支行名称
	FyInterBankNo   string `xml:"fyInterBankNo,omitempty"`   // 开户行号
}

// 主结构体
type UserQueryResponse struct {
	XMLName       xml.Name            `xml:"xml"`
	TraceNo       string              `xml:"traceNo,omitempty"`            // 商户流水号
	MchntCd       string              `xml:"mchntCd,omitempty"`            // 商户号
	RespCode      string              `xml:"respCode,omitempty"`           // 返回码
	RespDesc      string              `xml:"respDesc,omitempty"`           // 返回信息
	AccountInList []AccountInListItem `xml:"accountInList>item,omitempty"` // 用户列表
	Signature     string              `xml:"signature,omitempty"`          // 签名
}

type UserUnsetBankAccountRequest struct {
	TraceNo   string `xml:"traceNo,omitempty" json:"traceNo"`     // 商户流水号
	MchntCd   string `xml:"mchntCd,omitempty" json:"mchntCd"`     // 商户号
	AccountIn string `xml:"accountIn,omitempty" json:"accountIn"` // 用户编号
	// 个体工商户用户必传
	//枚举值：
	//01法人对私卡
	//02企业对公户（个体工商户绑定富友VA账户时，类型传02）
	OutAcntNoType     string `xml:"outAcntNoType,omitempty" json:"outAcntNoType"`         // 银行账号类型
	OutAcntNm         string `xml:"outAcntNm,omitempty" json:"outAcntNm"`                 // 开户户名
	CertNo            string `xml:"certNo,omitempty" json:"certNo"`                       // 证件号
	ShxyNo            string `xml:"shxyNo,omitempty" json:"shxyNo"`                       // 统一社会信用代码
	InterBankNo       string `xml:"interBankNo,omitempty" json:"interBankNo"`             // 开户行号
	OutAcntNo         string `xml:"outAcntNo,omitempty" json:"outAcntNo"`                 // 银行账号
	MiniAppReturnPath string `xml:"miniAppReturnPath,omitempty" json:"miniAppReturnPath"` // 小程序回调地址
	//枚举值：
	//1 短信模式
	//2 返回url
	//3 返回微信小程序可以使用的url
	//不传默认1
	CheckType string `xml:"checkType,omitempty" json:"checkType"` // 验证类型
	//小程序人脸识别回跳路径，在checkType为3时如果不传，默认跳转/pages/index/index首页
	Mobile string `xml:"mobile,omitempty" json:"mobile"` // 手机号码
}

func UserUnsetBankAccountRequestToUserBankAccountRequest(u *UserUnsetBankAccountRequest) *UserBankAccountRequest {
	return &UserBankAccountRequest{
		TraceNo:           u.TraceNo,
		MchntCd:           u.MchntCd,
		AccountIn:         u.AccountIn,
		Type:              "2",
		OutAcntNoType:     u.OutAcntNoType,
		OutAcntNm:         u.OutAcntNm,
		CertNo:            u.CertNo,
		ShxyNo:            u.ShxyNo,
		InterBankNo:       u.InterBankNo,
		OutAcntNo:         u.OutAcntNo,
		MiniAppReturnPath: u.MiniAppReturnPath,
		CheckType:         u.CheckType,
		Mobile:            u.Mobile,
	}
}

type UserBindBankAccountRequest struct {
	TraceNo   string `xml:"traceNo,omitempty" json:"traceNo"`     // 商户流水号
	MchntCd   string `xml:"mchntCd,omitempty" json:"mchntCd"`     // 商户号
	AccountIn string `xml:"accountIn,omitempty" json:"accountIn"` // 用户编号
	// 个体工商户用户必传
	//枚举值：
	//01法人对私卡
	//02企业对公户（个体工商户绑定富友VA账户时，类型传02）
	OutAcntNoType     string `xml:"outAcntNoType,omitempty" json:"outAcntNoType"`         // 银行账号类型
	OutAcntNm         string `xml:"outAcntNm,omitempty" json:"outAcntNm"`                 // 开户户名
	CertNo            string `xml:"certNo,omitempty" json:"certNo"`                       // 证件号
	ShxyNo            string `xml:"shxyNo,omitempty" json:"shxyNo"`                       // 统一社会信用代码
	InterBankNo       string `xml:"interBankNo,omitempty" json:"interBankNo"`             // 开户行号
	OutAcntNo         string `xml:"outAcntNo,omitempty" json:"outAcntNo"`                 // 银行账号
	MiniAppReturnPath string `xml:"miniAppReturnPath,omitempty" json:"miniAppReturnPath"` // 小程序回调地址
	//枚举值：
	//1 短信模式
	//2 返回url
	//3 返回微信小程序可以使用的url
	//不传默认1
	CheckType string `xml:"checkType,omitempty" json:"checkType"` // 验证类型
	//小程序人脸识别回跳路径，在checkType为3时如果不传，默认跳转/pages/index/index首页
	Mobile string `xml:"mobile,omitempty" json:"mobile"` // 手机号码
}

func UserBindBankAccountRequestToUserBankAccountRequest(u *UserBindBankAccountRequest) *UserBankAccountRequest {
	return &UserBankAccountRequest{
		TraceNo:           u.TraceNo,
		MchntCd:           u.MchntCd,
		AccountIn:         u.AccountIn,
		Type:              "1",
		OutAcntNoType:     u.OutAcntNoType,
		OutAcntNm:         u.OutAcntNm,
		CertNo:            u.CertNo,
		ShxyNo:            u.ShxyNo,
		InterBankNo:       u.InterBankNo,
		OutAcntNo:         u.OutAcntNo,
		MiniAppReturnPath: u.MiniAppReturnPath,
		CheckType:         u.CheckType,
		Mobile:            u.Mobile,
	}
}

type UserBankAccountRequest struct {
	XMLName   xml.Name `xml:"xml"`
	TraceNo   string   `xml:"traceNo,omitempty" json:"traceNo"`     // 商户流水号
	MchntCd   string   `xml:"mchntCd,omitempty" json:"mchntCd"`     // 商户号
	AccountIn string   `xml:"accountIn,omitempty" json:"accountIn"` // 用户编号
	// 枚举值：
	//1绑卡
	//2解绑
	Type string `xml:"type,omitempty" json:"type"` // 修改类型
	// 个体工商户用户必传
	//枚举值：
	//01法人对私卡
	//02企业对公户（个体工商户绑定富友VA账户时，类型传02）
	OutAcntNoType     string `xml:"outAcntNoType,omitempty" json:"outAcntNoType"`         // 银行账号类型
	OutAcntNm         string `xml:"outAcntNm,omitempty" json:"outAcntNm"`                 // 开户户名
	CertNo            string `xml:"certNo,omitempty" json:"certNo"`                       // 证件号
	ShxyNo            string `xml:"shxyNo,omitempty" json:"shxyNo"`                       // 统一社会信用代码
	InterBankNo       string `xml:"interBankNo,omitempty" json:"interBankNo"`             // 开户行号
	OutAcntNo         string `xml:"outAcntNo,omitempty" json:"outAcntNo"`                 // 银行账号
	MiniAppReturnPath string `xml:"miniAppReturnPath,omitempty" json:"miniAppReturnPath"` // 小程序回调地址
	//枚举值：
	//1 短信模式
	//2 返回url
	//3 返回微信小程序可以使用的url
	//不传默认1
	CheckType string `xml:"checkType,omitempty" json:"checkType"` // 验证类型
	//小程序人脸识别回跳路径，在checkType为3时如果不传，默认跳转/pages/index/index首页
	Mobile    string `xml:"mobile,omitempty" json:"mobile"`       // 手机号码
	Signature string `xml:"signature,omitempty" json:"signature"` // 签名
}

type UserBankAccountResponse struct {
	XMLName   xml.Name `xml:"xml"`
	TraceNo   string   `xml:"traceNo,omitempty" json:"traceNo"`     // 商户流水号
	MchntCd   string   `xml:"mchntCd,omitempty" json:"mchntCd"`     // 商户号
	AccountIn string   `xml:"accountIn,omitempty" json:"accountIn"` // 用户编号
	// 枚举值：
	//1绑卡
	//2解绑
	Type        string `xml:"type,omitempty" json:"type"`               // 修改类型
	InterBankNo string `xml:"interBankNo,omitempty" json:"interBankNo"` // 开户行号
	RespCode    string `xml:"respCode,omitempty"`                       // 返回码
	RespDesc    string `xml:"respDesc,omitempty"`                       // 返回信息
	OutAcntNo   string `xml:"outAcntNo,omitempty" json:"outAcntNo"`     // 银行账号
	Signature   string `xml:"signature,omitempty" json:"signature"`     // 签名
	KsCheckUrl  string `xml:"ksCheckUrl,omitempty"`                     // 旷世验证h5链接
	CheckUrl    string `xml:"checkUrl,omitempty"`                       // 阿里验证h5链接
}

func (u *UserBankAccountRequest) Validate() error {
	if u.TraceNo == "" {
		u.TraceNo = fmt.Sprintf("%d", time.Now().UnixMilli())
	}
	if u.AccountIn == "" {
		return errors.ErrorParamError("AccountIn must not empty")
	}
	if u.Type == "" {
		return errors.ErrorParamError("Type must not empty")
	}
	if u.OutAcntNo == "" {
		return errors.ErrorParamError("OutAcntNo must not empty")
	}
	if u.CheckType == "" {
		u.CheckType = enum.CheckType_Mobile
	}
	return nil
}

func (u *UserBankAccountRequest) GenerateSign() map[string]string {
	var sign = make(map[string]string)
	by, _ := json.Marshal(u)
	_ = json.Unmarshal(by, &sign)
	var filter = []string{
		"signature",
		"randomStr",
		"mchntTraceNo",
		"idNo",
		"outAcntNoType",
		"outAcntNm",
		"certNo",
		"shxyNo",
		"miniAppReturnPath",
		//"mobile",
	}
	for _, v := range filter {
		delete(sign, v)
	}
	return sign
}

func (u *UserBankAccountRequest) Sign(sign string) {
	u.Signature = sign
}

func (u *UserBankAccountRequest) Xml() (string, error) {
	xmlBy, err := xml.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(xmlBy), nil
}

type CommonEncryptCallbackResponse struct {
	XMLName    xml.Name `xml:"xml"`
	NotifyType string   `xml:"notifyType,omitempty"`
	Message    string   `xml:"message,omitempty"`
	OriginBody string
}

type UserCreateCallbackRequest struct {
	XMLName         xml.Name        `xml:"xml"`
	MchntCd         string          `xml:"mchntCd,omitempty"`         // 商户号
	Type            string          `xml:"type,omitempty"`            // 业务类型
	AccountTag      string          `xml:"accountTag,omitempty"`      // 特殊账户标记
	RelationMchntCd string          `xml:"relationMchntCd,omitempty"` // 关联商户
	AccountIn       string          `xml:"accountIn,omitempty"`       // 用户编号
	OutAcctNm       string          `xml:"outAcctNm,omitempty"`       // 户名
	OutAcctCardNm   string          `xml:"outAcctCardNm,omitempty"`   // 银行卡户名
	Mobile          string          `xml:"mobile,omitempty"`          // 手机号
	OutAcctNo       string          `xml:"outAcctNo,omitempty"`       // 银行账号
	InterBankNo     string          `xml:"interBankNo,omitempty"`     // 开户行行号
	Status          enum.UserStatus `xml:"status,omitempty"`          // 状态
	MchntTraceNo    string          `xml:"mchntTraceNo,omitempty"`    // 商户流水号
	LicNo           string          `xml:"licNo,omitempty"`           // 证件号码
	AllocateScale   int             `xml:"allocateScale,omitempty"`   // 分账比例
	BankAcctNo      string          `xml:"bankAcctNo,omitempty"`      // 银行账户编号
	BankInterNo     string          `xml:"bankInterNo,omitempty"`     // 银行账户联行号
	BankNm          string          `xml:"bankNm,omitempty"`          // 银行名称
	Msg             string          `xml:"msg,omitempty"`             // 银行开户失败原因
	Signature       string          `xml:"signature,omitempty"`       // 签名
}

type EncryptResponse struct {
	XMLName    xml.Name `xml:"xml"`
	Message    string   `xml:"message,omitempty"` // 商户流水号
	MchntCd    string   `xml:"mchntCd,omitempty"` // 商户号
	RespDesc   string   `xml:"respDesc"`
	RespCode   string   `xml:"respCode"`
	OriginBody string
}
