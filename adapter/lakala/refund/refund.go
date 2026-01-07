package refund

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/client"

	"github.com/lihongsheng/payment-sdk/adapter/lakala/model"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"strconv"
	"time"
)

const (
	refundMethod = "/api/v3/labs/relation/refund"
	QueryMethod  = "/api/v3/labs/query/tradequery"
)

type Refund struct {
	*client.Client
}

func NewRefund(api *client.Client) (iface.Refund, error) {
	return &Refund{
		api,
	}, nil
}

func (r *Refund) Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error) {
	reqParam := r.buildRefundQueryRequest(req)
	result, err := r.Client.DoPost(ctx, reqParam, r.C.ApiHost+refundMethod, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("Refund timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.RefundResponseBody{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "lakala-Pay-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	refundResp := resp.RespData
	if refundResp == nil {
		return nil, errors.New(fmt.Sprintf("无法获取退款结果%s", string(by)))
	}
	reservedRefundAmt, _ := strconv.Atoi(refundResp.RefundAmount)
	return &dto.RefundDetail{
		TradeRefundNo:       refundResp.TradeNo,
		RefundNo:            refundResp.OutTradeNo,
		TradeNo:             refundResp.OriginTradeNo,
		OrderNo:             req.OrderNo,
		Channel:             refund.RefundChannel_ORIGINAL,
		UserReceivedAccount: "",
		SuccessTime:         time.Time{},
		CreateTime:          time.Time{},
		Status:              refund.Status_Pending,
		FundsAccount:        "",
		Amount: dto.Amount{
			Total: int64(reservedRefundAmt),
		},
		OriginResponse: string(by),
	}, nil
}

func (r *Refund) buildRefundQueryRequest(req *dto.RefundRequest) *model.RefundRequest {
	rr := &model.RefundRequest{
		MerchantNo:       r.C.MchID,
		TermNo:           r.C.TermNO,
		OutTradeNo:       req.RefundNo,
		RefundAmount:     fmt.Sprintf("%d", req.Amount.Total),
		RefundReason:     req.Reason,
		OriginOutTradeNo: req.OrderNo,
		LocationInfo: &model.LocationInfo{
			RequestIp: "127.0.0.1",
			Location:  "",
		},
	}
	return rr
}

// Query 不确认 OutTradeNo 是否退款单号,还是订单号，此处代码需要有真实拉卡拉账号测试
func (r *Refund) Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error) {
	if req.RefundNo == "" {
		return nil, errors.New("refund_no is  empty")
	}
	reqParam := model.PaymentQueryRequest{
		MerchantNo: r.C.MchID,
		TermNo:     r.C.TermNO,
		OutTradeNo: req.RefundNo,
	}
	result, err := r.Client.DoPost(ctx, reqParam, r.C.ApiHost+QueryMethod, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("Refund query timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.PaymentQueryRespBody{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "lakala-Pay-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	paymentResp := resp.RespData
	if paymentResp == nil || paymentResp.TradeMainType != "REFUND" {
		return nil, errors.New(fmt.Sprintf("无法获取支付结果%s", string(by)))
	}
	total, _ := strconv.Atoi(paymentResp.TotalAmount)
	detail := &dto.RefundDetail{
		TradeRefundNo:       paymentResp.TradeNo,
		RefundNo:            paymentResp.OutTradeNo,
		TradeNo:             "",
		OrderNo:             req.OrderNo,
		Channel:             refund.RefundChannel_ORIGINAL,
		UserReceivedAccount: "",
		SuccessTime:         time.Time{},
		CreateTime:          time.Time{},
		Status:              r.GetRefundStatus(paymentResp.TradeState),
		FundsAccount:        "",
		Amount: dto.Amount{
			Total: int64(total),
		},
		OriginResponse: "",
	}

	if detail == nil {
		return nil, errors.New("not find refund")
	}
	return detail, nil
}

func (r *Refund) GetRefundStatus(status string) refund.Status {
	switch status {
	case "SUCCESS":
		return refund.Status_Success
	case "FAIL":
		return refund.Status_Failed
	}
	return refund.Status_Status_UNKNOWN
}

func (r *Refund) Callback(ctx context.Context, req *http.Request) (*dto.CallbackRefundDetail, error) {
	return nil, errors2.ErrorNoSupport("not support refund callback", nil)
}

func (r *Refund) IsSupportCallback() bool {
	return false
}
