package payment

import (
	"bytes"
	"context"
	"github.com/lihongsheng/payment-sdk/adapter/alipay"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/util"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/errors"
	"io"
	"net/http"
	"net/url"
)

type CallbackMethod struct {
	*alipay.Api
}

func NewCallback(conf config.Config) (*CallbackMethod, error) {
	api, err := alipay.NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &CallbackMethod{
		Api: api,
	}, nil
}

func (c *CallbackMethod) Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	bodyBytes, err := util.GetRequestBody(req)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, err
	}
	sign, signValue, err := c.Client.Sign.GenerateSignString(values)
	if err != nil {
		return nil, err
	}
	verifg, err := c.Client.Sign.RsaVerify(sign, signValue)
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
			Currency: payment.Currency_CNY.String(),
			Total:    resp.GetTotalAmount(),
		},
		Status: util.PaymentStatus(resp.TradeStatus),
		//  PaymentProduct: payment.PaymentProduct_JSAPI.String(),
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
