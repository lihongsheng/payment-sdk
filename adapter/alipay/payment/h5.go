package payment

import (
	"context"
	"encoding/json"
	"fmt"
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

type H5 struct {
	*CallbackMethod
}

func NewH5(conf config.Config) (iface.Pay, error) {
	api, err := NewCallback(conf)
	if err != nil {
		return nil, err
	}
	return &H5{
		api,
	}, nil
}

func (h *H5) Pay(ctcx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := h.buildFacePaymentRequest(req)
	commonParam := h.Client.GetCommonRequestParams()
	if req.NotifyUrl != "" {
		commonParam[enum.COMMON_PARAM_NOTIFY_URL_NAME] = req.NotifyUrl
	}
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_TRADES_PAY
	resp, err := h.Client.DoPost(commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	fmt.Println(string(body))
	var response model.FacePaymentResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayTradePayResponse.Code != "" && response.AlipayTradePayResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if response.AlipayTradePayResponse.SubCode == "" && response.AlipayTradePayResponse.TradeNo != "" {
		respTrue = true
	}
	if !respTrue {
		return nil, errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	re := &dto.PayResponse{
		OrderNo: req.Order.OrderNo,
		TradeNo: response.AlipayTradePayResponse.TradeNo,
		PayAmount: dto.Amount{
			Total:    req.Order.PayAmount.Total,
			Currency: req.Order.PayAmount.Currency,
		},
		Status:         payment.Status_Pending,
		PaymentProduct: payment.PaymentProduct_H5.String(),
		Action: dto.Action{
			Action: action.Action_Prepay.String(),
			Parameters: map[string]string{
				"trade_no": response.AlipayTradePayResponse.TradeNo,
			},
			Url: "",
		},
		OriginResponse: string(body),
	}
	return re, nil
}

func (h *H5) buildFacePaymentRequest(req *dto.PayOrder) model.FacePaymentRequest {
	result := model.FacePaymentRequest{
		OutTradeNo:   req.Order.OrderNo,
		ProductCode:  enum.FACE_TO_FACE_PAYMENT,
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

//func (h *H5) Pay(ctcx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
//	reqParam := h.buildH5PaymentRequest(req)
//	commonParam := h.Client.GetCommonRequestParams()
//	if req.NotifyUrl != "" {
//		commonParam[enum.COMMON_PARAM_NOTIFY_URL_NAME] = req.NotifyUrl
//	}
//	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_H5_TRADES_CREATE
// 不用发起请求，需要自行拼接下边的参数参数,返给前端
// <form name="punchout_form" method="post" action="https://openapi-sandbox.dl.alipaydev.com/gateway.do?charset=UTF8&method=alipay.trade.wap.pay&sign=DB%2FbC%2Bg8P9Pv6rBcPEl92eamGikPbfUzi4IGWHZrXe8NtcNDbC4bIklULdqLAGjwdCsdWKiJIEALRS72CL5tzDMAJctCyEjN15EPbMnHaIVoxG1dV3lc4bAahqv8FZ0GlUYJbwevhR0PXZashoXFDIGjk8Qt3OuzeQa8tTt%2FoSgWzsRp6ANBdyEqral6py2KI7PghI%2Bjo4VyfijGFtqevay44J7Vmu78l0TvgVUqOZsDz98%2FABNq%2FH%2Ft2MKGbSWE9G44Rfj5W7YGlqHjjCtovKygKXnEGX5lJy2AgrOEydi8R2Mqp7xKZPsL4bki%2FfY4XVpHSrU7NergUhLbIVL5bQ%3D%3D&app_id=9021000133622619&sign_type=RSA2&timestamp=2025-10-15+18%3A01%3A09&alipay_sdk=alipay-sdk-java-4.35.171.ALL&format=json"> <input type="hidden" name="biz_content" value="{&quot;time_expire&quot;:&quot;2025-10-16 10:00:00&quot;,&quot;out_trade_no&quot;:&quot;70501111111S001111119&quot;,&quot;total_amount&quot;:&quot;0.01&quot;,&quot;subject&quot;:&quot;大乐透&quot;,&quot;product_code&quot;:&quot;QUICK_WAP_WAY&quot;,&quot;seller_id&quot;:&quot;2088102147948060&quot;}"> <input type="submit" value="立即支付" style="display:none" > </form> <script>document.forms[0].submit();</script>

//	resp, err := h.Client.DoPost(commonParam, reqParam, nil)
//	if err != nil {
//		return nil, err
//	}
//	body := resp.Body()
//	fmt.Println(string(body))
//	f, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)
//	f.Write(body)
//	f.Close()
//	//var response model.H5PaymentResponse
//	//err = json.Unmarshal(body, &response)
//	//if err != nil {
//	//	return nil, errors.ErrorSystemError("json.Unmarshal error")
//	//}
//	//if response.ErrorResponse != nil {
//	//	return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
//	//}
//	//respTrue := false
//	//if response.AlipayTradeCreateResponse.Code != "" && response.AlipayTradeCreateResponse.Code == enum.RESPONSE_SUCCESS_CODE {
//	//	respTrue = true
//	//}
//	//if response.AlipayTradeCreateResponse.SubCode == "" && response.AlipayTradeCreateResponse.TradeNo != "" {
//	//	respTrue = true
//	//}
//	//if respTrue {
//	return &dto.PayResponse{
//		OrderNo: req.Order.OrderNo,
//		TradeNo: "",
//		PayAmount: dto.Amount{
//			Total:    req.Order.PayAmount.Total,
//			Currency: req.Order.PayAmount.Currency,
//		},
//		Status:         payment.Status_Pending,
//		PaymentProduct: payment.PaymentProduct_H5.String(),
//		Action: dto.Action{
//			Action: action.Action_Prepay.String(),
//			Parameters: map[string]string{
//				"package": string(body),
//			},
//			Url: "",
//		},
//		OriginResponse: string(body),
//	}, nil
//	//}
//	//return nil, errors.ErrorSystemError("not return trade_no;"+string(body), nil)
//}

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
