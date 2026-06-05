package payment

import (
	"context"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/client"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	"github.com/lihongsheng/payment-sdk/enum/payment"
)

type Qrcode struct {
	*Api
}

func NewQrcode(api *client.Client) (iface.Pay, error) {
	api2, err := NewApi(api)
	if err != nil {
		return nil, err
	}
	return &Qrcode{
		api2,
	}, nil
}

func (h *Qrcode) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := h.buildPaymentRequest(req)
	commonParam := h.Client.GetCommonRequestParams()
	if req.NotifyUrl != "" {
		commonParam[enum.COMMON_PARAM_NOTIFY_URL_NAME] = req.NotifyUrl
	}
	if req.RedirectUrl != "" {
		commonParam[enum.COMMON_PARAM_RETURN_URL_NAME] = req.RedirectUrl
	}
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_TRADES_PAGE_PAY
	resp, _, err := h.Client.GetPageExecute(commonParam, reqParam)
	if err != nil {
		return nil, err
	}

	re := &dto.PayResponse{
		OrderNo: req.Order.OrderNo,
		TradeNo: "",
		PayAmount: dto.Amount{
			Total:    req.Order.PayAmount.Total,
			Currency: req.Order.PayAmount.Currency,
		},
		Status:         payment.Status_Pending,
		PaymentProduct: payment.PaymentProduct_Qrcode.String(),
		Action: dto.Action{
			Action:         action.Action_Redirect.String(),
			Parameters:     nil,
			Url:            resp.String(),
			RedirectMethod: "GET",
		},
		OriginResponse: "",
	}
	return re, nil
}

func (h *Qrcode) buildPaymentRequest(req *dto.PayOrder) model.QrCodePaymentRequest {
	result := model.QrCodePaymentRequest{
		OutTradeNo:      req.Order.OrderNo,
		ProductCode:     enum.FAST_INSTANT_TRADE_PAY,
		TotalAmount:     req.Order.PayAmount.ToFloatString(),
		ExtendParams:    nil,
		Subject:         req.Order.Subject,
		GoodsDetail:     nil,
		QrPayMode:       "2",
		IntegrationType: "PCWEB",
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
