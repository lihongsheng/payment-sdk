package refund

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/yunhui"
	"github.com/lihongsheng/payment-sdk/adapter/yunhui/model"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

const (
	refundPath = "/api/refund/refundOrderApi"
	queryPath  = "/api/refund/queryApi"
)

type Refund struct {
	*yunhui.Api
	paymentProduct enum.PaymentProduct
	payment        enum.Payment
}

func NewRefund(conf config.Config, product enum.PaymentProduct, payment enum.Payment) (iface.Refund, error) {
	api, err := yunhui.NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &Refund{
		Api:            api,
		paymentProduct: product,
		payment:        payment,
	}, nil
}

func (r *Refund) Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error) {
	reqParams := r.buildParams(req)
	result, err := r.Client.DoPost(ctx, reqParams, r.C.ApiHost+refundPath, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.RefundResp{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParams)
	logc.Info(ctx, "yunhui-refund-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.CommonResp.Error() != nil {
		return nil, resp.CommonResp.Error()
	}
	if resp.Data == nil {
		return nil, errors.New(fmt.Sprintf("not find body:%s", string(by)))
	}
	if resp.Data != nil && resp.Data.Error.Error() != nil {
		return nil, resp.Data.Error.Error()
	}
	re := &dto.RefundDetail{
		RefundNo:      req.RefundNo,
		OrderNo:       req.OrderNo,
		TradeRefundNo: resp.Data.RefundOrderId,
		TradeNo:       "",
		Amount: dto.Amount{
			Currency: "CNY",
			Total:    resp.Data.RefundAmount,
		},
		Channel:        refund.RefundChannel_ORIGINAL,
		Status:         model.GetRefundStatus(resp.Data.State),
		OriginResponse: string(by),
	}
	return re, nil
}

func (r *Refund) buildParams(req *dto.RefundRequest) *model.RefundRequest {
	repParams := &model.RefundRequest{
		PayOrderId:   "",
		ExtParam:     "",
		MchOrderNo:   req.OrderNo,
		RefundReason: req.Reason,
		ReqTime:      time.Now().UnixMilli(),
		ChannelExtra: "",
		AppId:        r.C.AppID,
		MchRefundNo:  req.RefundNo,
		ClientIp:     "",
		NotifyUrl:    req.NotifyUrl,
		Currency:     "cny",
		MchNo:        r.C.MchID,
		RefundAmount: req.Amount.Total,
		ApiInfo:      r.Extra.TermNO,
	}
	return repParams
}

func (r *Refund) Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error) {
	reqParams := model.RefundQuery{
		MchNo:       r.C.MchID,
		ApiInfo:     r.Extra.TermNO,
		MchRefundNo: req.RefundNo,
		AppId:       r.C.AppID,
		ReqTime:     time.Now().UnixMilli(),
	}
	result, err := r.Client.DoPost(ctx, reqParams, r.C.ApiHost+queryPath, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.RefundQueryResp{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParams)
	logc.Info(ctx, "yunhui-refund-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.CommonResp.Error() != nil {
		return nil, resp.CommonResp.Error()
	}
	if resp.Data == nil {
		return nil, errors.New(fmt.Sprintf("not find body:%s", string(by)))
	}
	if resp.Data != nil && resp.Data.Error.Error() != nil {
		return nil, resp.Data.Error.Error()
	}
	re := &dto.RefundDetail{
		RefundNo:      req.RefundNo,
		OrderNo:       req.OrderNo,
		TradeRefundNo: resp.Data.RefundOrderId,
		TradeNo:       "",
		Amount: dto.Amount{
			Currency: "CNY",
			Total:    resp.Data.RefundAmount,
		},
		Channel:        refund.RefundChannel_ORIGINAL,
		Status:         model.GetRefundStatus(resp.Data.State),
		OriginResponse: string(by),
	}
	return re, nil
}
