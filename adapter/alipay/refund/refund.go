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
)

type Refund struct {
	Client *client.Client
	conf   config.Config
}

func NewRefund(conf config.Config) (iface.Refund, error) {
	newClient, err := client.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &Refund{
		Client: newClient,
		conf:   conf,
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
	return nil, errors.ErrorNoSupport("not support refund callback", nil)
}
