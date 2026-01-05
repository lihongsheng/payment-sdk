package payment

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/errors"
)

type Qrcode struct {
	*Api
}

func NewQrcode(conf config.Config) (iface.Pay, error) {
	api, err := NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &Qrcode{
		api,
	}, nil
}

func (h *Qrcode) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := h.buildPaymentRequest(req)
	commonParam := h.Client.GetCommonRequestParams()
	if req.NotifyUrl != "" {
		commonParam[enum.COMMON_PARAM_NOTIFY_URL_NAME] = req.NotifyUrl
	}
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_H5_TRADES_CREATE
	resp, err := h.Client.DoPost(ctx, commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.PreCreateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayTradePreCreateResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if respTrue {
		return nil, errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	re := &dto.PayResponse{
		OrderNo: req.Order.OrderNo,
		TradeNo: response.AlipayTradePreCreateResponse.OutTradeNo,
		PayAmount: dto.Amount{
			Total:    req.Order.PayAmount.Total,
			Currency: req.Order.PayAmount.Currency,
		},
		Status:         payment.Status_Pending,
		PaymentProduct: payment.PaymentProduct_H5.String(),
		Action: dto.Action{
			Action:         action.Action_Qrcode.String(),
			Parameters:     map[string]string{},
			Url:            resp.String(),
			RedirectMethod: response.AlipayTradePreCreateResponse.QrCode,
		},
		OriginResponse: string(body),
	}
	return re, nil
}

func (h *Qrcode) buildPaymentRequest(req *dto.PayOrder) model.QrCodePaymentRequest {
	result := model.QrCodePaymentRequest{
		OutTradeNo:   req.Order.OrderNo,
		ProductCode:  enum.QR_CODE_OFFLINE,
		TotalAmount:  req.Order.PayAmount.ToFloatString(),
		ExtendParams: nil,
		Subject:      req.Order.Subject,
		GoodsDetail:  nil,
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
