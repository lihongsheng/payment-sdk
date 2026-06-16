package payment

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	"net/url"
	"time"

	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
)

type H5 struct {
	*Api
	client h5.H5ApiService
}

func NewH5(api *client.Api) (iface.Pay, error) {
	api2, err := NewApi(api)
	if err != nil {
		return nil, err
	}
	return &H5{
		Api:    api2,
		client: h5.H5ApiService{Client: api.Client},
	}, nil
}

func (h *H5) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	resp, result, err := h.client.Prepay(ctx, h.buildPayParmams(req))
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.H5Url == nil || *resp.H5Url == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return H5Url")
	}
	h5url := *resp.H5Url
	if req.RedirectUrl != "" {
		u, _ := url.Parse(h5url)
		u.Query().Add("redirect_url", req.RedirectUrl)
		h5url = u.String()
	}
	return &dto.PayResponse{
		PaymentProduct: enum.PaymentProduct_H5.String(),
		Action: dto.Action{
			Action:     action.Action_Redirect.String(),
			Parameters: map[string]string{},
			Url:        h5url,
		},
		OrderNo:   req.Order.OrderNo,
		PayAmount: dto.Amount{},
		Status:    enum.Status_Pending,
	}, nil
}

func (h *H5) buildPayParmams(req *dto.PayOrder) h5.PrepayRequest {
	var t *time.Time
	if req.TimeExpire > 0 {
		t = core.Time(time.Unix(req.TimeExpire, 0))
	}
	amount := &h5.Amount{
		Total: core.Int64(req.Order.PayAmount.Total),
	}
	if req.Order.PayAmount.Currency != "" {
		amount.Currency = core.String(req.Order.PayAmount.Currency)
	}
	resp := h5.PrepayRequest{
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
		resp.SettleInfo = &h5.SettleInfo{
			ProfitSharing: core.Bool(req.SettleInfo.ProfitSharing),
		}
	}
	if req.SceneInfo != nil {
		resp.SceneInfo = &h5.SceneInfo{
			PayerClientIp: core.String(req.SceneInfo.ClientIp),
			// DeviceId:      core.String(req.SceneInfo.DeviceID),
			H5Info: &h5.H5Info{
				Type: core.String(""),
				//AppName: core.String(req.SceneInfo.ApplicationInfo.AppName),
				//AppUrl:  core.String(req.SceneInfo.ApplicationInfo.Url),
			},
		}
		if req.SceneInfo.Device == enum1.Device_H5 {
			resp.SceneInfo.H5Info.Type = core.String("Wap")
		}

		if req.SceneInfo.Store.Id != "" {
			resp.SceneInfo.StoreInfo = &h5.StoreInfo{
				Id: core.String(req.SceneInfo.Store.Id),
			}
		}
		if req.SceneInfo.DeviceID != "" {
			resp.SceneInfo.DeviceId = core.String(req.SceneInfo.DeviceID)
		}
	}
	return resp
}

func (h *H5) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	var resp *payments.Transaction
	var result *core.APIResult
	var err error
	if req.OrderNo != "" {
		resp, result, err = h.client.QueryOrderByOutTradeNo(ctx, h5.QueryOrderByOutTradeNoRequest{OutTradeNo: core.String(req.OrderNo), Mchid: core.String(h.C.Merchant.MchID)})
	} else if req.TradeNo != "" {
		resp, result, err = h.client.QueryOrderById(ctx, h5.QueryOrderByIdRequest{TransactionId: core.String(req.TradeNo), Mchid: core.String(h.C.Merchant.MchID)})
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

func (h *H5) Close(ctx context.Context, req dto.CloseQuery) error {
	if req.OrderNo == "" {
		return errors2.ErrorParamError("order_no is required")
	}
	result, err := h.client.CloseOrder(ctx, h5.CloseOrderRequest{
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
