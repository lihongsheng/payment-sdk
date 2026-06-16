package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

type App struct {
	*Api
	client app.AppApiService
}

func NewApp(api *client.Api) (iface.Pay, error) {
	api2, err := NewApi(api)
	if err != nil {
		return nil, err
	}
	return &App{
		Api:    api2,
		client: app.AppApiService{Client: api.Client},
	}, nil
}

func (h *App) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	resp, result, err := h.client.Prepay(ctx, h.buildPayParams(req))
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.PrepayId == nil || *resp.PrepayId == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return H5Url")
	}
	var actionParams = map[string]string{}
	actionParams["appId"] = h.C.Merchant.AppID
	actionParams["partnerId"] = h.C.Merchant.MchID
	actionParams["prepayId"] = until.StringPoint(resp.PrepayId)
	actionParams["timeStamp"] = fmt.Sprintf("%d", time.Now().Unix())
	actionParams["nonceStr"] = tools.GenerateRandomDigits(10)
	actionParams["packageValue"] = "Sign=WXPay"
	signParams := fmt.Sprintf("%s\n%s\n%s\n%s\n", actionParams["appId"], actionParams["timeStamp"], actionParams["nonceStr"], actionParams["prepayId"])
	sign, _ := utils.SignSHA256WithRSA(signParams, h.PrivateKey)
	actionParams["sign"] = sign
	return &dto.PayResponse{
		PaymentProduct: enum.PaymentProduct_APP.String(),
		Action: dto.Action{
			Action:     action.Action_Prepay.String(),
			Parameters: actionParams,
			Url:        "",
		},
		OrderNo:   req.Order.OrderNo,
		PayAmount: dto.Amount{},
		Status:    enum.Status_Pending,
	}, nil
}

func (h *App) buildPayParams(req *dto.PayOrder) app.PrepayRequest {
	var t *time.Time
	if req.TimeExpire > 0 {
		t = core.Time(time.Unix(req.TimeExpire, 0))
	}
	amount := &app.Amount{
		Total: core.Int64(req.Order.PayAmount.Total),
	}
	if req.Order.PayAmount.Currency != "" {
		amount.Currency = core.String(req.Order.PayAmount.Currency)
	}
	resp := app.PrepayRequest{
		Appid:       core.String(h.C.Merchant.AppID),
		Mchid:       core.String(h.C.Merchant.MchID),
		OutTradeNo:  core.String(req.Order.OrderNo),
		TimeExpire:  t,
		Attach:      core.String(req.PassBackParams),
		NotifyUrl:   core.String(req.NotifyUrl),
		Description: core.String(req.Order.Subject),
		Amount:      amount,
	}
	if req.SettleInfo != nil {
		resp.SettleInfo = &app.SettleInfo{
			ProfitSharing: core.Bool(req.SettleInfo.ProfitSharing),
		}
	}
	if req.SceneInfo != nil {
		resp.SceneInfo = &app.SceneInfo{
			PayerClientIp: core.String(req.SceneInfo.ClientIp),
			// DeviceId:      core.String(req.SceneInfo.DeviceID),
		}
		if req.SceneInfo.Store.Id != "" {
			resp.SceneInfo.StoreInfo = &app.StoreInfo{
				Id: core.String(req.SceneInfo.Store.Id),
			}
		}
		if req.SceneInfo.DeviceID != "" {
			resp.SceneInfo.DeviceId = core.String(req.SceneInfo.DeviceID)
		}
	}
	return resp
}

func (h *App) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	var resp *payments.Transaction
	var result *core.APIResult
	var err error
	if req.OrderNo != "" {
		resp, result, err = h.client.QueryOrderByOutTradeNo(ctx, app.QueryOrderByOutTradeNoRequest{OutTradeNo: core.String(req.OrderNo), Mchid: core.String(h.C.Merchant.MchID)})
	} else if req.TradeNo != "" {
		resp, result, err = h.client.QueryOrderById(ctx, app.QueryOrderByIdRequest{TransactionId: core.String(req.TradeNo), Mchid: core.String(h.C.Merchant.MchID)})
	} else {
		return nil, errors2.ErrorParamError("order_no or trade_no is required")
	}
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "response is nil")
	}
	status := until.PaymentStatus[until.StringPoint(resp.TradeState)]
	if status == enum.Status_Status_UNKNOWN {
		logc.Error(ctx, "wxPayErrStatus", logc.Field("resp", resp))
		return nil, until.ErrorHandler(ctx, result, err, "status is unknown")
	}
	var successTime time.Time
	if resp.SuccessTime != nil && *resp.SuccessTime != "" {
		successTime, _ = time.Parse(time.RFC3339, *resp.SuccessTime)
	}
	originBy, _ := json.Marshal(resp)
	return &dto.PayDetail{
		OrderNo: until.StringPoint(resp.OutTradeNo),
		TradeNo: until.StringPoint(resp.TransactionId),
		PayAmount: dto.Amount{
			Currency: until.StringPoint(resp.Amount.Currency),
			Total:    until.Int64Point(resp.Amount.Total),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_H5.String(),
		SuccessTime:    successTime.Unix(),
		OriginResponse: string(originBy),
	}, nil
}

func (h *App) Close(ctx context.Context, req dto.CloseQuery) error {
	if req.OrderNo == "" {
		return errors2.ErrorParamError("order_no is required")
	}
	result, err := h.client.CloseOrder(ctx, app.CloseOrderRequest{
		Mchid:      core.String(h.C.Merchant.MchID),
		OutTradeNo: core.String(req.OrderNo),
	})
	if err != nil {
		return until.ErrorHandler(ctx, result, err, "")
	}
	if result.Response.StatusCode != 204 {
		return until.ErrorHandler(ctx, result, err, "")
	}
	return nil
}
