package payment

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/driver/iface"

	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client/payscore"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"

	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/enum/action"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"time"
)

// 支付分先享模式

type AfterPay struct {
	*CallbackMethod
	client payscore.ScoreApiService
}

func NewAfterPay(conf config.Config) (iface.Pay, error) {
	api, err := NewCallback(conf)
	if err != nil {
		return nil, err
	}
	svc := payscore.ScoreApiService{Client: api.Client}
	return &AfterPay{
		CallbackMethod: api,
		client:         svc,
	}, nil
}

func (a *AfterPay) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	resp, result, err := a.client.Prepay(ctx, a.BuildPayParams(req))
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.Package == nil || *resp.Package == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return Package")
	}
	var actionParams = map[string]string{}
	actionParams["package"] = *resp.Package
	return &dto.PayResponse{
		PaymentProduct: payment.PaymentProduct_AFTER_PAY.String(),
		Action: dto.Action{
			Action:     action.Action_Prepay.String(),
			Parameters: actionParams,
			Url:        "",
		},
		OrderNo:   req.Order.OrderNo,
		PayAmount: dto.Amount{},
		TradeNo:   until.StringPoint(resp.OrderId),
		Status:    payment.Status_Pending,
	}, nil
}

func (a *AfterPay) BuildPayParams(req *dto.PayOrder) payscore.CreateServiceOrderRequest {
	r := payscore.CreateServiceOrderRequest{
		OutOrderNo:          core.String(req.Order.OrderNo),
		Appid:               core.String(a.C.AppID),
		ServiceId:           core.String(a.C.ScoreServiceID),
		ServiceIntroduction: core.String(req.Order.Subject),
		PostPayments:        nil,
		PostDiscounts:       nil,
		TimeRange: &payscore.TimeRange{
			StartTime: core.String(req.Order.CreateAt.Format(time.DateTime)),
			EndTime:   nil,
		},
		Location: nil,
		RiskFund: &payscore.RiskFund{
			Name:        core.String(string(enum.RISK_FUND_TYPE_ESTIMATE_ORDER_COST)),
			Amount:      core.Int64(req.Order.PayAmount.Total),
			Description: nil,
		},
		Attach:          nil,
		NotifyUrl:       core.String(req.NotifyUrl),
		NeedUserConfirm: nil,
		Device:          nil,
	}
	if req.RiskFund != nil {
		r.RiskFund.Amount = core.Int64(req.RiskFund.Amount)
		r.RiskFund.Description = core.String(req.RiskFund.Description)
		r.RiskFund.Name = core.String(req.RiskFund.Name)
	}
	return r
}

func (a *AfterPay) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	reqParams := payscore.GetServiceOrderRequest{
		ServiceId: core.String(a.C.ScoreServiceID),
		Appid:     core.String(a.C.AppID),
	}
	if req.TradeNo != "" {
		reqParams.QueryId = core.String(req.TradeNo)
	}
	if req.OrderNo != "" {
		reqParams.OutOrderNo = core.String(req.OrderNo)
	}
	resp, result, err := a.client.GetServiceOrder(ctx, reqParams)
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.OutOrderNo == nil || *resp.OutOrderNo == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return OutOrderNo")
	}
	originBy, _ := json.Marshal(resp)
	re := &dto.PayDetail{
		OrderNo: until.StringPoint(resp.OutOrderNo),
		TradeNo: until.StringPoint(resp.OrderId),
		PayAmount: dto.Amount{
			Currency: payment.Currency_CNY.String(),
			Total:    until.Int64Point(resp.TotalAmount),
		},
		Status:         0,
		PaymentProduct: payment.PaymentProduct_AFTER_PAY.String(),
		SuccessTime:    0,
		OriginResponse: string(originBy),
	}
	collectionStatus := ""
	if resp.Collection != nil && resp.Collection.State != nil {
		collectionStatus = *resp.Collection.State
	}
	re.Status = until.GetScorePaymentStatus(until.StringPoint(resp.State), until.StringPoint(resp.StateDescription), collectionStatus)
	return re, nil
}

func (a *AfterPay) Close(ctx context.Context, req dto.CloseQuery) error {
	reqParams := payscore.CancelServiceOrderRequest{
		ServiceId: core.String(a.C.ScoreServiceID),
		Appid:     core.String(a.C.AppID),
	}
	if req.OrderNo != "" {
		reqParams.OutOrderNo = core.String(req.OrderNo)
	}
	resp, result, err := a.client.CancelServiceOrder(ctx, reqParams)
	if err != nil {
		return until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.OutOrderNo == nil || *resp.OutOrderNo == "" {
		return until.ErrorHandler(ctx, result, err, "not return OutOrderNo")
	}
	return nil
}

func (a *AfterPay) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParams := payscore.CompleteServiceOrderRequest{
		OutOrderNo: core.String(req.Order.OrderNo),
		Appid:      core.String(a.C.AppID),
		ServiceId:  core.String(a.C.ScoreServiceID),
		PostPayments: []payscore.Payment{{
			Name: core.String(req.Order.Subject),
		}},
		PostDiscounts: nil,
		TotalAmount:   core.Int64(req.Order.PayAmount.Total),
		TimeRange:     nil,
		Location:      nil,
		ProfitSharing: nil,
		GoodsTag:      nil,
		Device:        nil,
	}
	resp, result, err := a.client.Complete(ctx, reqParams)
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.OutOrderNo == nil || *resp.OutOrderNo == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return OutOrderNo")
	}
	originBy, _ := json.Marshal(resp)
	re := &dto.PayResponse{
		OrderNo: until.StringPoint(resp.OutOrderNo),
		TradeNo: until.StringPoint(resp.OrderId),
		PayAmount: dto.Amount{
			Currency: payment.Currency_CNY.String(),
			Total:    until.Int64Point(resp.TotalAmount),
		},
		Status:         0,
		PaymentProduct: payment.PaymentProduct_AFTER_PAY.String(),
		OriginResponse: string(originBy),
	}
	re.Status = until.GetScorePaymentStatus(until.StringPoint(resp.State), until.StringPoint(resp.StateDescription), "")
	return re, nil
}
