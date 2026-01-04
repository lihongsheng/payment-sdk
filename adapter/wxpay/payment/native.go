package payment

import (
	"context"
	"encoding/json"
	tools2 "github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"time"

	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/skip2/go-qrcode"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

type Native struct {
	*Api
	client native.NativeApiService
}

func NewNative(conf config.Config) (iface.Pay, error) {
	api, err := NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &Native{
		Api:    api,
		client: native.NativeApiService{Client: api.Client},
	}, nil
}

func (n *Native) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	resp, result, err := n.client.Prepay(ctx, n.buildPayParmams(req))
	if err != nil {
		return nil, tools2.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.CodeUrl == nil || *resp.CodeUrl == "" {
		return nil, tools2.ErrorHandler(ctx, result, err, "not return NativeUrl")
	}
	qrCode, _ := tools.GenerateQRToBase64(*resp.CodeUrl, 256, qrcode.Medium)
	return &dto.PayResponse{
		PaymentProduct: enum.PaymentProduct_Native.String(),
		Action: dto.Action{
			Action: action.Action_Qrcode.String(),
			Parameters: map[string]string{
				"qrcode": qrCode,
			},
			Url: tools2.StringPoint(resp.CodeUrl),
		},
		OrderNo: req.Order.OrderNo,
		PayAmount: dto.Amount{
			Currency: req.Order.PayAmount.Currency,
			Total:    req.Order.PayAmount.Total,
		},
		Status: enum.Status_Pending,
	}, nil
}

func (n *Native) buildPayParmams(req *dto.PayOrder) native.PrepayRequest {
	var t *time.Time
	if req.TimeExpire > 0 {
		t = core.Time(time.Unix(req.TimeExpire, 0))
	}
	amount := &native.Amount{
		Total: core.Int64(req.Order.PayAmount.Total),
	}
	if req.Order.PayAmount.Currency != "" {
		amount.Currency = core.String(req.Order.PayAmount.Currency)
	}
	resp := native.PrepayRequest{
		Appid:       core.String(n.C.AppID),
		Mchid:       core.String(n.C.MchID),
		OutTradeNo:  core.String(req.Order.OrderNo),
		TimeExpire:  t,
		Attach:      core.String(req.PassbackParams),
		NotifyUrl:   core.String(req.NotifyUrl),
		Description: core.String(req.Order.Subject),
		Amount:      amount,
	}
	if req.SettleInfo != nil {
		resp.SettleInfo = &native.SettleInfo{
			ProfitSharing: core.Bool(req.SettleInfo.ProfitSharing),
		}
	}
	if req.SceneInfo != nil {
		resp.SceneInfo = &native.SceneInfo{
			PayerClientIp: core.String(req.SceneInfo.ClientIp),
			DeviceId:      core.String(req.SceneInfo.DeviceID),
		}
		if req.SceneInfo.Store.Id != "" {
			resp.SceneInfo.StoreInfo = &native.StoreInfo{
				Id: core.String(req.SceneInfo.Store.Id),
			}
		}
	}
	return resp
}

func (n *Native) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	var resp *payments.Transaction
	var result *core.APIResult
	var err error
	if req.OrderNo != "" {
		resp, result, err = n.client.QueryOrderByOutTradeNo(ctx, native.QueryOrderByOutTradeNoRequest{OutTradeNo: core.String(req.OrderNo), Mchid: core.String(n.C.MchID)})
	} else if req.TradeNo != "" {
		resp, result, err = n.client.QueryOrderById(ctx, native.QueryOrderByIdRequest{TransactionId: core.String(req.TradeNo), Mchid: core.String(n.C.MchID)})
	} else {
		return nil, errors2.ErrorParamError("order_no or trade_no is required")
	}
	if err != nil {
		return nil, tools2.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, tools2.ErrorHandler(ctx, result, err, "response is nil")
	}
	status := tools2.PaymentStatus[tools2.StringPoint(resp.TradeState)]
	if status == enum.Status_Status_UNKNOWN {
		logc.Error(ctx, "wxPayErrStatus", logc.Field("resp", resp))
		return nil, tools2.ErrorHandler(ctx, result, err, "status is unknown")
	}
	var successTime time.Time
	if resp.SuccessTime == nil && *resp.SuccessTime != "" {
		successTime, _ = time.Parse(time.RFC3339, *resp.SuccessTime)
	}
	originBy, _ := json.Marshal(resp)
	return &dto.PayDetail{
		OrderNo: tools2.StringPoint(resp.OutTradeNo),
		TradeNo: tools2.StringPoint(resp.TransactionId),
		PayAmount: dto.Amount{
			Currency: tools2.StringPoint(resp.Amount.Currency),
			Total:    tools2.Int64Point(resp.Amount.Total),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_Native.String(),
		SuccessTime:    successTime.Unix(),
		OriginResponse: string(originBy),
	}, nil
}

func (n *Native) Close(ctx context.Context, req dto.CloseQuery) error {
	if req.OrderNo == "" {
		return errors2.ErrorParamError("order_no is required")
	}
	result, err := n.client.CloseOrder(ctx, native.CloseOrderRequest{
		Mchid:      core.String(n.C.MchID),
		OutTradeNo: core.String(req.OrderNo),
	})
	if err != nil {
		return tools2.ErrorHandler(ctx, result, err, "")
	}
	if result.Response.StatusCode != 204 {
		return tools2.ErrorHandler(ctx, result, err, "")
	}
	return nil
}
