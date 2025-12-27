package alipay

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/alipay/model"
	"github.com/singer-stack-lab/payment-sdk/adapter/alipay/util"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	enum1 "github.com/singer-stack-lab/payment-sdk/enum"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
	"github.com/singer-stack-lab/payment-sdk/enum/refund"
	"github.com/singer-stack-lab/payment-sdk/errors"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func CallbackPaymentParse(ctx context.Context, conf config.Config, req *http.Request) (*dto.CallbackPayDetail, error) {
	// 补验签，或者程序比对支付支付单号是否一致
	bodyBytes, _ := io.ReadAll(req.Body)
	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, err
	}

	sign, signValue, err := GenerateSignString(values)
	if err != nil {
		return nil, err
	}
	verifg, err := RsaVerify(sign, signValue, conf.Cert.PublicKey)
	if err != nil {
		return nil, err
	}
	if !verifg {
		return nil, errors.ErrorSignError("签名验证失败："+string(bodyBytes), nil)
	}
	resp := buildCallbackParams(values)
	re := &dto.CallbackPayDetail{
		OrderNo: resp.OutTradeNo,
		TradeNo: resp.TradeNo,
		PayAmount: dto.Amount{
			Currency: enum.Currency_CNY.String(),
			Total:    resp.GetTotalAmount(),
		},
		Status:         util.PaymentStatus(resp.TradeStatus),
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		SuccessTime:    resp.GetGmtPayment().Unix(),
		OriginResponse: string(bodyBytes),
	}
	if resp.IsRefund() {
		re.EventAction = enum1.Event_REFUND
		re.EventRefund = &dto.EventRefundActionParams{
			RefundNo: resp.OutBizNo,
			OrderNo:  resp.OutTradeNo,
		}
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return re, nil
}

func buildCallbackParams(values url.Values) model.AlipayNotifyBody {
	return model.AlipayNotifyBody{
		NotifyTime:        values.Get("notify_time"),
		NotifyType:        values.Get("notify_type"),
		NotifyId:          values.Get("notify_id"),
		SignType:          values.Get("sign_type"),
		Sign:              values.Get("sign"),
		TradeNo:           values.Get("trade_no"),
		AppId:             values.Get("app_id"),
		AuthAppId:         values.Get("auth_app_id"),
		OutTradeNo:        values.Get("out_trade_no"),
		OutBizNo:          values.Get("out_biz_no"),
		BuyerId:           values.Get("buyer_id"),
		BuyerLogonId:      values.Get("buyer_logon_id"),
		SellerId:          values.Get("seller_id"),
		SellerEmail:       values.Get("seller_email"),
		TradeStatus:       values.Get("trade_status"),
		TotalAmount:       values.Get("total_amount"),
		ReceiptAmount:     values.Get("receipt_amount"),
		InvoiceAmount:     values.Get("invoice_amount"),
		BuyerPayAmount:    values.Get("buyer_pay_amount"),
		PointAmount:       values.Get("point_amount"),
		RefundFee:         values.Get("refund_fee"),
		SendBackFee:       values.Get("send_back_fee"),
		Subject:           values.Get("subject"),
		Body:              values.Get("body"),
		GmtCreate:         values.Get("gmt_create"),
		GmtPayment:        values.Get("gmt_payment"),
		GmtRefund:         values.Get("gmt_refund"),
		GmtClose:          values.Get("gmt_close"),
		FundBillList:      values.Get("fund_bill_list"),
		VoucherDetailList: values.Get("voucher_detail_list"),
		BizSettleMode:     values.Get("biz_settle_mode"),
	}
}

func CallbackRefundParse(ctx context.Context, conf config.Config, req *http.Request) (*dto.CallbackRefundDetail, error) {
	bodyBytes, _ := io.ReadAll(req.Body)
	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, err
	}
	resp := buildCallbackParams(values)
	re := &dto.CallbackRefundDetail{
		RefundNo:      resp.OutBizNo,
		OrderNo:       resp.OutTradeNo,
		TradeRefundNo: resp.TradeNo,
		TradeNo:       resp.TradeNo,
		Amount: dto.Amount{
			Currency: "CNY",
			Total:    resp.GetRefundFee(),
		},
		UserReceivedAccount: "",
		OriginResponse:      string(bodyBytes),
	}
	if resp.IsRefund() {
		re.Status = refund.Status_Success
	} else {
		re.Status = refund.Status_Status_UNKNOWN
	}
	return re, nil
}

// GenerateSignString 生成待签名字符串（通用版，剔除sign、sign_type）
// params: 通知返回的所有参数（url.Values 格式）
// 返回值: 按规则处理后的待签名字符串
func GenerateSignString(params url.Values) (string, string, error) {
	if len(params) == 0 || len(params["sign"]) == 0 || params["sign"][0] == "" {
		return "", "", errors.ErrorParamError("未找到签名字符串", nil)
	}
	// 1. 过滤参数：剔除 sign、sign_type，保留其他非空参数
	filteredParams := make(url.Values)
	for k, v := range params {
		// 剔除 sign、sign_type 参数
		if k == "sign" || k == "sign_type" {
			continue
		}
		// 过滤空值参数（若业务需保留空值，可删除此判断）
		if len(v) > 0 && v[0] != "" {
			filteredParams[k] = v
		}
	}

	// 2. 提取参数名并按字典序排序
	var keys []string
	for k := range filteredParams {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Go原生字典序排序（ASCII升序）

	// 3. URL解码参数值 + 拼接待签名字符串
	var signStrBuilder strings.Builder
	for i, k := range keys {
		decodedValue := filteredParams.Get(k)
		// 拼接：多个参数用&分隔，格式为 key=value
		if i > 0 {
			signStrBuilder.WriteString("&")
		}
		signStrBuilder.WriteString(k)
		signStrBuilder.WriteString("=")
		signStrBuilder.WriteString(decodedValue)
	}
	return signStrBuilder.String(), params["sign"][0], nil
}

func RsaVerify(data string, signatureBase64 string, publicKeyPEM string) (bool, error) {
	// 解析PEM格式的公钥
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return false, errors.ErrorParamError("解析公钥失败")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, errors.ErrorParamError("解析公钥内容失败: %v", err)
	}
	// 类型断言，确保是RSA公钥
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("公钥不是RSA类型")
	}
	// 解码Base64格式的签名
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, errors.ErrorParamError("解析签名失败: %v", err)
	}
	// 计算数据的SHA256哈希
	hash := sha256.Sum256([]byte(data))
	// 验证签名
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false, errors.ErrorParamError("签名验证失败: %v", err)
	}

	return true, nil
}
