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
	"time"
)

// https://opendocs.alipay.com/mini/6039ed0c_alipay.trade.create?scene=de4d6a1e0c6e423b9eefa7c3a6dcb7a5&pathHash=779dc517

type Jsapi struct {
	*Api
}

func NewJsApi(conf config.Config) (iface.Pay, error) {
	api, err := NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &Jsapi{
		api,
	}, nil
}

func (j *Jsapi) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := j.buildPayParams(req)
	commonParam := j.Client.GetCommonRequestParams()
	if req.NotifyUrl != "" {
		commonParam[enum.COMMON_PARAM_NOTIFY_URL_NAME] = req.NotifyUrl
	}
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_TRADES_CREATE
	resp, err := j.Client.DoPost(ctx, commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.JsApiPaymentResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayTradeCreateResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}

	if respTrue {
		return &dto.PayResponse{
			OrderNo: response.AlipayTradeCreateResponse.OutTradeNo,
			TradeNo: response.AlipayTradeCreateResponse.TradeNo,
			PayAmount: dto.Amount{
				Total:    req.Order.PayAmount.Total,
				Currency: req.Order.PayAmount.Currency,
			},
			Status:         payment.Status_Pending,
			PaymentProduct: payment.PaymentProduct_JSAPI.String(),
			Action: dto.Action{
				Action: action.Action_Prepay.String(),
				Parameters: map[string]string{
					"trade_no": response.AlipayTradeCreateResponse.TradeNo,
				},
				Url: "",
			},
			OriginResponse: string(body),
		}, nil
	}
	reBy, _ := json.Marshal(reqParam)
	return nil, errors.ErrorSystemError("not return trade_no;"+string(body)+";"+string(reBy), nil)
}

func (j *Jsapi) buildPayParams(req *dto.PayOrder) model.JsApiPaymentRequest {
	result := model.JsApiPaymentRequest{
		OutTradeNo:         req.Order.OrderNo,
		ProductCode:        enum.JSAPI,
		OpAppId:            j.C.AppID,
		TotalAmount:        req.Order.PayAmount.ToFloatString(),
		ExtendParams:       nil,
		DiscountableAmount: "",
		Subject:            req.Order.Subject,
		Body:               "",
		BuyerId:            req.Payer.UnionID,
		BuyerOpenId:        req.Payer.OpenID,
		TimeExpire:         "",
		PassbackParams:     req.PassBackParams,
		GoodsDetail:        nil,
	}
	if req.AlipayExtra != nil && req.AlipayExtra.ProductCode != "" {
		result.ProductCode = req.AlipayExtra.ProductCode
	}
	if req.TimeExpire > 0 {
		t := time.Unix(req.TimeExpire, 0)
		//beijingLoc, _ := time.LoadLocation("Asia/Shanghai")
		//t.In(beijingLoc).Format("2006-01-02 15:04:05")
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
