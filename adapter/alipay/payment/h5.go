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
	"time"
)

type H5 struct {
	*Api
}

func NewH5(conf config.Config) (iface.Pay, error) {
	api, err := NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &H5{
		api,
	}, nil
}

func (h *H5) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := h.buildFacePaymentRequest(req)
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

	re := &dto.PayResponse{
		OrderNo: req.Order.OrderNo,
		TradeNo: "",
		PayAmount: dto.Amount{
			Total:    req.Order.PayAmount.Total,
			Currency: req.Order.PayAmount.Currency,
		},
		Status:         payment.Status_Pending,
		PaymentProduct: payment.PaymentProduct_H5.String(),
		Action: dto.Action{
			Action:         action.Action_Redirect.String(),
			Parameters:     params,
			Url:            resp.String(),
			RedirectMethod: "POST",
		},
		OriginResponse: "",
	}
	return re, nil
}

func (h *H5) buildFacePaymentRequest(req *dto.PayOrder) model.FacePaymentRequest {
	result := model.FacePaymentRequest{
		OutTradeNo:   req.Order.OrderNo,
		ProductCode:  enum.QUICK_WAP_WAY,
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

func (h *H5) buildH5PaymentRequest(req *dto.PayOrder) model.H5PaymentRequest {
	result := model.H5PaymentRequest{
		OutTradeNo:         req.Order.OrderNo,
		ProductCode:        enum.QUICK_WAP_WAY,
		TotalAmount:        req.Order.PayAmount.ToFloatString(),
		ExtendParams:       nil,
		DiscountableAmount: "",
		Subject:            req.Order.Subject,
		TimeExpire:         "",
		PassbackParams:     req.PassbackParams,
		GoodsDetail:        nil,
	}
	if req.TimeExpire > 0 {
		t := time.Unix(req.TimeExpire, 0)
		result.TimeExpire = t.Format(time.DateTime)
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
