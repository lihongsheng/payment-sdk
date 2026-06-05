package enum

import (
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"strings"
)

const (
	JSAPI  = "JSAPI"
	LETPAY = "LETPAY"
	FWC    = "FWC"

	OrderTypeALIPAY = "ALIPAY" // (统一下单、条码支付、服务窗支付)
	OrderTypeALIAPP = "ALIAPP" // (统一下单、条码支付、服务窗支付)
	OrderTypeALIH5  = "ALIH5"  // (统一下单、条码支付、服务窗支付)
	OrderTypeWXH5   = "WXH5"   // (统一下单、条码支付、服务窗支付)
	OrderTypeWECHAT = "WECHAT" // (统一下单、条码支付、服务窗支付)
	OrderTypeWXAPP  = "WXAPP"  // (统一下单、条码支付、服务窗支付)
	//SUCCESS—支付成功
	//REFUND—已退款
	//NOTPAY—未支付
	//CLOSED—已关闭
	//REVOKED—已撤销（刷卡支付）
	//USERPAYING--用户支付中
	//PAYERROR--支付失败(其他原因，如银行返回失败)
	PAY_STATUS_SUCCESS    = "SUCCESS"
	PAY_STATUS_REFUND     = "REFUND"
	PAY_STATUS_NOTPAY     = "NOTPAY"
	PAY_STATUS_CLOSED     = "CLOSED"
	PAY_STATUS_REVOKED    = "REVOKED"
	PAY_STATUS_USERPAYING = "USERPAYING"
	PAY_STATUS_PAYERROR   = "PAYERROR"
)

func GetOrderType(payment enum.Payment, product enum.PaymentProduct) string {
	switch payment {
	case enum.Payment_Wechat:
		switch product {
		case enum.PaymentProduct_JSAPI, enum.PaymentProduct_LITE:
			return OrderTypeWECHAT
		case enum.PaymentProduct_APP:
			return OrderTypeWXAPP
		case enum.PaymentProduct_H5:
			return OrderTypeWXH5
		}
	case enum.Payment_Alipay:
		switch product {
		case enum.PaymentProduct_JSAPI, enum.PaymentProduct_LITE:
			return OrderTypeALIPAY
		case enum.PaymentProduct_APP:
			return OrderTypeALIAPP
		case enum.PaymentProduct_H5:
			return OrderTypeALIH5
		}
	}
	return OrderTypeWECHAT
}

func GetPaymentStatus(status string) enum.Status {
	status = strings.ToUpper(status)
	switch status {
	case PAY_STATUS_SUCCESS:
		return enum.Status_Success
	case PAY_STATUS_REFUND:
		return enum.Status_Refund
	case PAY_STATUS_NOTPAY:
		return enum.Status_Cancel
	case PAY_STATUS_CLOSED:
		return enum.Status_Close
	case PAY_STATUS_REVOKED:
		return enum.Status_Cancel
	case PAY_STATUS_USERPAYING:
		return enum.Status_Pending
	case PAY_STATUS_PAYERROR:
		return enum.Status_Failed
	}
	return enum.Status_Failed
}

const (
	XmlRandomStr = "randomStr"
)

const (
	// 个人账户
	CleanType_Personal = "01"
	// 企业
	CleanType_Enterprise = "02"
	// 个体工商户
	CleanType_Individual = "03"
)

const (
	// 身份证
	CertTp_IDCard = "0"
	// 其他
	CertTp_Other = "1"
	// 组织机构代码证
	CertTp_OrgCode = "1"
)

const (
	POST_COMMON_PARAM                 = "req"
	POST_ENCRYPT_COMMON_PARAM_MCN     = "mchntCd"
	POST_ENCRYPT_COMMON_PARAM_MESSAGE = "message"
	SKIP_COMMON_PARAM_SIGN            = "signature"
)

const (
	SuccessCode = "000000"
)

// 验证方式
const (
	CheckType_Mobile = "1"
	CheckType_Url    = "2"
	CheckType_Wechat = "3"
)

// 用户状态
// 枚举值：
// 01失效
// 02生效
// 04删除
type UserStatus string

const (
	UserStatus_Invalid = "01"
	UserStatus_Active  = "02"
	UserStatus_Delete  = "04"
)
