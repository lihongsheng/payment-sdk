package payment

import (
	"context"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	"github.com/lihongsheng/payment-sdk/enum/payment"
)

type App struct {
	*Api
}

func NewApp(conf config.Config) (iface.Pay, error) {
	api, err := NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &App{
		api,
	}, nil
}

func (h *App) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := h.buildPaymentRequest(req)
	commonParam := h.Client.GetCommonRequestParams()
	if req.NotifyUrl != "" {
		commonParam[enum.COMMON_PARAM_NOTIFY_URL_NAME] = req.NotifyUrl
	}
	if req.RedirectUrl != "" {
		commonParam[enum.COMMON_PARAM_RETURN_URL_NAME] = req.RedirectUrl
	}
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_H5_TRADES_CREATE
	resp, params, err := h.Client.PageExecute(commonParam, reqParam)
	if err != nil {
		return nil, err
	}
	query := resp.Query()
	for k, v := range params {
		query.Add(k, v)
	}

	re := &dto.PayResponse{
		OrderNo: req.Order.OrderNo,
		TradeNo: "",
		PayAmount: dto.Amount{
			Total:    req.Order.PayAmount.Total,
			Currency: req.Order.PayAmount.Currency,
		},
		Status:         payment.Status_Pending,
		PaymentProduct: payment.PaymentProduct_APP.String(),
		Action: dto.Action{
			Action: action.Action_Prepay.String(),
			Parameters: map[string]string{
				"orderStr": query.Encode(),
			},
			Url: "",
		},
		OriginResponse: "",
	}
	return re, nil
}

func (h *App) buildPaymentRequest(req *dto.PayOrder) model.FacePaymentRequest {
	result := model.FacePaymentRequest{
		OutTradeNo:   req.Order.OrderNo,
		ProductCode:  enum.APP_WAY,
		TotalAmount:  req.Order.PayAmount.ToFloatString(),
		ExtendParams: nil,
		Subject:      req.Order.Subject,
		GoodsDetail:  nil,
		AuthCode:     req.Order.OrderNo,
		Scene:        "bar_code",
	}
	if req.AlipayExtra != nil && req.AlipayExtra.ProductCode != "" {
		result.ProductCode = req.AlipayExtra.ProductCode
	}
	if len(req.Order.Goods) > 0 {
		result.GoodsDetail = make([]*model.GoodDetails, 0)
		for _, v := range req.Order.Goods {
			p := dto.Amount{
				Total:    v.Price,
				Currency: "",
			}
			tmp := &model.GoodDetails{
				GoodsId:   v.Sku,
				GoodsName: v.Name,
				Quantity:  v.Quantity,
				Price:     p.ToFloatString(),
			}
			result.GoodsDetail = append(result.GoodsDetail, tmp)
		}
	}
	return result
}
