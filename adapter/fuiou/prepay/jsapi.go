package prepay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/client"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	enum3 "github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"net/url"
	"time"

	errors2 "github.com/lihongsheng/payment-sdk/errors"

	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/zeromicro/go-zero/core/logc"
)

// link http://fundwx.fuiou.com/doc/#/aggregatePay/api?id=_26-%e5%85%ac%e4%bc%97%e5%8f%b7%e6%9c%8d%e5%8a%a1%e7%aa%97%e7%bb%9f%e4%b8%80%e4%b8%8b%e5%8d%95%e6%8e%a5%e5%8f%a3
const (
	payMethodPath      = "/aggregatePay/wxPreCreate"
	payQueryMethodPath = "/aggregatePay/commonQuery"
	payCloseMethodPath = "/aggregatePay/closeOrder"
	PrePayment         = "/aggregatePay/preCreate"
)

type Jsapi struct {
	*Api
}

func NewJsApi(api *client.Client, product enum.PaymentProduct, payment enum.Payment) (iface.Pay, error) {
	api2, err := NewApi(api, product, payment)
	if err != nil {
		return nil, err
	}
	return &Jsapi{
		api2,
	}, nil
}

func (p *Jsapi) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := p.buildPayParams(req)
	if err := reqParam.Validate(); err != nil {
		return nil, err
	}
	r := p.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := r.SetContext(ctx).SetBody(reqParam).Post(p.C.ApiHost + payMethodPath)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &JsApiPaymentResponse{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "fuiou-Jsapi-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if err != nil {
		return nil, errors2.ErrorSystemError("parse result error").WithCause(err)
	}
	if !resp.IsSuccess() {
		return nil, errors2.ErrorSystemError("pay is error;err:%s", resp.ResultMsg).WithCause(errors.New(fmt.Sprintf("code:%s;Msg:%s", resp.ResultCode, resp.ResultMsg)))
	}
	re := &dto.PayResponse{
		OrderNo: p.C.OrderPrefix + req.Order.OrderNo,
		TradeNo: resp.ReservedFyOrderNo,
		PayAmount: dto.Amount{
			Currency: req.Order.PayAmount.Currency,
			Total:    req.Order.PayAmount.Total,
		},
		Status:         enum.Status_Pending,
		PaymentProduct: p.paymentProduct.String(),
		OriginResponse: string(by),
		Action: dto.Action{
			Action: action.Action_Prepay.String(),
		},
	}
	switch reqParam.TradeType {
	case enum2.FWC: // 支付宝服务窗
		re.Action = dto.Action{
			Action: action.Action_Redirect.String(),
			Parameters: map[string]string{
				"trade_no": resp.ReservedTransactionId,
			},
		}
	case enum2.JSAPI, enum2.LETPAY:
		if p.payment == enum.Payment_Wechat {
			var actionParams = map[string]string{}
			actionParams["appId"] = resp.SdkAppid
			actionParams["timeStamp"] = resp.SdkTimestamp
			actionParams["nonceStr"] = resp.SdkNoncestr
			actionParams["package"] = resp.SdkPackage
			actionParams["signType"] = resp.SdkSigntype
			actionParams["paySign"] = resp.SdkPaysign
			params, err := url.ParseQuery(resp.SdkPackage)
			if err != nil {
				return nil, errors2.ErrorSystemError("parse SdkPackage error").WithCause(err)
			}
			prepayID := params[action.PrepayID]
			if len(prepayID) == 0 {
				return nil, errors2.ErrorSystemError("fuiou not return prepayID error")
			}
			re.Action.Parameters = actionParams
		}
	}
	return re, nil
}
func (p *Jsapi) buildPayParams(req *dto.PayOrder) *JsApiPaymentRequest {
	result := &JsApiPaymentRequest{
		MchntCd:              p.C.MchID,
		RandomStr:            tools.GenerateRandomDigits(4),
		OrderAmt:             req.Order.PayAmount.Total,
		MchntOrderNo:         enum2.GenOrder(p.C.OrderPrefix, req.Order.OrderNo),
		ProductId:            "",
		TermId:               tools.GenerateRandomDigits(4),
		GoodsDes:             req.Order.Subject,
		GoodsDetail:          "",
		GoodsTag:             "",
		AddnInf:              req.PassBackParams,
		CurrType:             req.Order.PayAmount.Currency,
		NotifyURL:            req.NotifyUrl,
		LimitPay:             "",
		TradeType:            "",
		Openid:               "",
		SubOpenid:            req.Payer.OpenID,
		SubAppid:             req.Payer.AppID,
		ReservedFyTermId:     "",
		ReservedExpireMinute: 0,
		//ReservedDeviceInfo:   DeviceInfo{},
		Sign:    "",
		Version: p.C.Version,
	}
	if req.SceneInfo != nil {
		result.TermIp = req.SceneInfo.ClientIp
		if req.SceneInfo.DeviceID == "" {
			result.TermId = tools.GenerateRandomDigits(4)
		}
		if req.SceneInfo.ApplicationInfo.AppName != "" {
			deviceInfo := DeviceInfo{
				Type:    string(req.SceneInfo.Device),
				AppName: req.SceneInfo.ApplicationInfo.AppName,
				AppUrl:  req.SceneInfo.ApplicationInfo.Url,
			}
			if req.SceneInfo.Device == enum3.Device_H5 {
				deviceInfo.Type = string(enum3.Device_H5)
			}
			if req.SceneInfo.System == enum3.Android {
				deviceInfo.Type = string(enum3.Android)
			} else if req.SceneInfo.System == enum3.IOS {
				deviceInfo.Type = "IOS"
			}

			if req.SceneInfo.ApplicationInfo.AppPackage != "" {
				deviceInfo.Type = "IOS"
				deviceInfo.AppUrl = req.SceneInfo.ApplicationInfo.AppPackage
				if req.SceneInfo.System == enum3.Android {
					deviceInfo.Type = string(enum3.Android)
				}
			}
			result.ReservedDeviceInfo = deviceInfo.String()
		}
	}

	if req.Order.CreateAt.IsZero() {
		result.TxnBeginTs = time.Now().Format("20060102150405")
	} else {
		result.TxnBeginTs = req.Order.CreateAt.Format("20060102150405")
	}

	if p.payment == enum.Payment_Wechat {
		result.TradeType = enum2.WxPaymentProductMap[p.paymentProduct]
	}
	if p.payment == enum.Payment_Alipay {
		result.TradeType = enum2.AliPaymentProductMap[p.paymentProduct]
	}
	result.GenSign(p.C.APISecret)
	return result
}
