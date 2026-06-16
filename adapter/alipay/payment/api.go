package payment

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/client"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/util"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/errors"
	"net/http"
	"net/url"
	"time"
)

type Api struct {
	Client *client.Client
}

func NewApi(api *client.Client) (*Api, error) {
	return &Api{
		Client: api,
	}, nil
}

func (a *Api) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors.ErrorNoSupport("not support Complete")
}

func (a *Api) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	reqParam := model.PaymentQueryRequest{
		OutTradeNo: req.OrderNo,
		TradeNo:    req.TradeNo,
		QueryOptions: []string{
			"trade_settle_info",
		},
	}
	commonParam := a.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_TRADES_QUERY
	resp, err := a.Client.DoPost(ctx, commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.PaymentQueryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error")
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayTradeQueryResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if !respTrue {
		return nil, errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	amount, _ := util.AmountToCents(response.AlipayTradeQueryResponse.TotalAmount)
	successTime := int64(0)
	if response.AlipayTradeQueryResponse.SendPayDate != "" {
		t, _ := time.Parse(time.DateTime, response.AlipayTradeQueryResponse.SendPayDate)
		successTime = t.Unix()
	}
	return &dto.PayDetail{
		OrderNo: response.AlipayTradeQueryResponse.OutTradeNo,
		TradeNo: response.AlipayTradeQueryResponse.TradeNo,
		PayAmount: dto.Amount{
			Total:    int64(amount),
			Currency: response.AlipayTradeQueryResponse.TransCurrency,
		},
		Status:         util.PaymentStatus(response.AlipayTradeQueryResponse.TradeStatus),
		PaymentProduct: payment.PaymentProduct_JSAPI.String(),
		SuccessTime:    successTime,
		OriginResponse: string(body),
	}, nil
}

func (a *Api) Close(ctx context.Context, req dto.CloseQuery) error {
	reqParam := model.PaymentCloseRequest{
		OutTradeNo: req.OrderNo,
	}
	commonParam := a.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_TRADES_CLOSE
	resp, err := a.Client.DoPost(ctx, commonParam, reqParam, nil)
	if err != nil {
		return err
	}
	body := resp.Body()
	var response model.PaymentCloseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	respTrue := false
	if response.TradeCloseResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if !respTrue {
		return errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	return nil
}

func (c *Api) Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	values, bodyBytes, err := c.Client.VerifyCallback(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := buildCallbackParams(values)
	re := &dto.CallbackPayDetail{
		OrderNo: resp.OutTradeNo,
		TradeNo: resp.TradeNo,
		PayAmount: dto.Amount{
			Currency: payment.Currency_CNY.String(),
			Total:    resp.GetTotalAmount(),
		},
		Status:         util.PaymentStatus(resp.TradeStatus),
		SuccessTime:    resp.GetGmtPayment().Unix(),
		OriginResponse: string(bodyBytes),
		EventAction:    enum1.Event_PAYMENT,
	}
	if resp.IsRefund() {
		re.EventAction = enum1.Event_REFUND
		re.EventRefund = &dto.EventRefundActionParams{
			RefundNo: resp.OutBizNo,
			OrderNo:  resp.OutTradeNo,
		}
	}
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
