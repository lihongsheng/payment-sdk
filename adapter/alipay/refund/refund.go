package refund

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/client"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/util"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	"github.com/lihongsheng/payment-sdk/errors"
	"net/http"
	"net/url"
)

type Refund struct {
	Client *client.Client
	conf   config.Config
}

func NewRefund(api *client.Client) (iface.Refund, error) {
	return &Refund{
		Client: api,
	}, nil
}

func (r *Refund) Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error) {
	reqParam := r.buildRefundParam(req)
	if err := reqParam.Validate(); err != nil {
		return nil, err
	}
	commonParam := r.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_REFUND_CREATE
	resp, err := r.Client.DoPost(ctx, commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.RefundResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error:%s", err.Error()).WithCause(err)
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayTradeRefundResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}

	if !respTrue {
		return nil, errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	re := &dto.RefundDetail{
		OrderNo:  req.OrderNo,
		TradeNo:  response.AlipayTradeRefundResponse.TradeNo,
		RefundNo: req.RefundNo,
		Amount:   req.Amount,
		Status:   refund.Status_Pending,
	}
	if !response.AlipayTradeRefundResponse.GetRefundSuccessTime().IsZero() {
		re.Status = refund.Status_Success
		re.SuccessTime = response.AlipayTradeRefundResponse.GetRefundSuccessTime()
	}
	return re, nil
}

func (r *Refund) buildRefundParam(req *dto.RefundRequest) model.RefundRequest {
	result := model.RefundRequest{
		OutTradeNo:              req.OrderNo,
		TradeNo:                 req.TradeNo,
		RefundAmount:            req.Amount.ToFloatString(),
		RefundReason:            req.Reason,
		OutRequestNo:            req.RefundNo,
		RefundGoodsDetail:       nil,
		RefundRoyaltyParameters: nil,
		QueryOptions:            nil,
		RelatedSettleConfirmNo:  "",
	}
	if req.Goods != nil {
		result.RefundGoodsDetail = make([]model.RefundGoodsDetail, 0, len(req.Goods))
		for _, v := range req.Goods {
			result.RefundGoodsDetail = append(result.RefundGoodsDetail, model.RefundGoodsDetail{
				OutSkuId:             "",
				OutItemId:            "",
				GoodsId:              v.Sku,
				RefundAmount:         util.Int64ToFloatString(v.Price),
				OutCertificateNoList: nil,
			})
		}
	}
	return result
}

func (r *Refund) Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error) {
	reqParam := model.RefundQueryRequest{
		OutTradeNo:   req.OrderNo,
		TradeNo:      req.TradeNo,
		OutRequestNo: req.RefundNo,
		QueryOptions: []string{
			"refund_detail_item_list",
			"gmt_refund_pay",
		},
	}
	commonParam := r.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_REFUND_QUERY
	resp, err := r.Client.DoPost(ctx, commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.RefundQueryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubMsg, nil).WithCause(err)
	}
	respTrue := false
	if response.AlipayTradeFastpayRefundQueryResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}

	if !respTrue {
		return nil, errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	amount, _ := util.AmountToCents(response.AlipayTradeFastpayRefundQueryResponse.RefundAmount)
	result := &dto.RefundDetail{
		OrderNo:  req.OrderNo,
		TradeNo:  response.AlipayTradeFastpayRefundQueryResponse.TradeNo,
		RefundNo: req.RefundNo,
		Amount: dto.Amount{
			Total:    int64(amount),
			Currency: "CNY",
		},
		Status: refund.Status_Failed,
	}
	if response.AlipayTradeFastpayRefundQueryResponse.RefundStatus == enum.REFUND_STATUS_SUCCESS {
		result.Status = refund.Status_Success
	}
	if response.AlipayTradeFastpayRefundQueryResponse.GmtRefundPay != "" {
		result.SuccessTime = response.AlipayTradeFastpayRefundQueryResponse.RefundSuccessTime()
	}
	return result, nil
}

func (r *Refund) Callback(ctx context.Context, req *http.Request) (*dto.CallbackRefundDetail, error) {
	values, bodyBytes, err := r.Client.VerifyCallback(ctx, req)
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
		Response:            "",
	}
	if resp.IsRefund() {
		re.Status = refund.Status_Success
	} else {
		re.Status = refund.Status_Status_UNKNOWN
	}
	return re, nil
}

func (r *Refund) IsSupportCallback() bool {
	return true
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
