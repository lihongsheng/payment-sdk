package prepay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/fuiou"
	"net/url"
	"strconv"
	"strings"
	"time"

	enum2 "github.com/singer-stack-lab/payment-sdk/adapter/fuiou/enum"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	"github.com/singer-stack-lab/payment-sdk/enum/action"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"

	errors2 "github.com/singer-stack-lab/payment-sdk/errors"

	"github.com/singer-stack-lab/payment-sdk/tools"
	"github.com/zeromicro/go-zero/core/logc"
)

// link http://fundwx.fuiou.com/doc/#/aggregatePay/api?id=_26-%e5%85%ac%e4%bc%97%e5%8f%b7%e6%9c%8d%e5%8a%a1%e7%aa%97%e7%bb%9f%e4%b8%80%e4%b8%8b%e5%8d%95%e6%8e%a5%e5%8f%a3
const (
	payMethodPath      = "/aggregatePay/wxPreCreate"
	payQueryMethodPath = "/aggregatePay/commonQuery"
	payCloseMethodPath = "/aggregatePay/closeOrder"
	PrePayment         = "/aggregatePay/preCreate"
)

type Pay struct {
	*fuiou.Api
	paymentProduct enum.PaymentProduct
	payment        enum.Payment
}

func NewPay(conf config.Config, product enum.PaymentProduct, payment enum.Payment) (driver.Pay, error) {
	api, err := fuiou.NewApi(conf)
	if err != nil {
		return nil, err
	}
	if api.Extra.OrderPrefix == "" {
		return nil, errors.New("富有订单前缀需要配置")
	}
	if _, exists := enum2.WxPaymentProductMap[product]; !exists {
		return nil, errors2.ErrorNoSupport("product [%s] is not exists", product.String())
	}
	if conf.ApiHost == "" {
		conf.ApiHost = enum2.ApiHost
	} else {
		conf.ApiHost = strings.TrimRight(conf.ApiHost, "/")
	}
	return &Pay{
		Api:            api,
		paymentProduct: product,
		payment:        payment,
	}, nil
}

func (p *Pay) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
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
	resp := &PaymentResponse{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "fuiou-Pay-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if err != nil {
		return nil, errors2.ErrorSystemError("parse result error").WithCause(err)
	}
	if !resp.IsSuccess() {
		return nil, errors2.ErrorSystemError("pay is error;err:%s", resp.ResultMsg).WithCause(errors.New(fmt.Sprintf("code:%s;Msg:%s", resp.ResultCode, resp.ResultMsg)))
	}
	re := &dto.PayResponse{
		OrderNo: p.Extra.OrderPrefix + req.Order.OrderNo,
		TradeNo: resp.ReservedFyOrderNo,
		PayAmount: dto.Amount{
			Currency: req.Order.PayAmount.Currency,
			Total:    req.Order.PayAmount.Total,
		},
		Status:         enum.Status_Pending,
		PaymentProduct: p.paymentProduct.String(),
		OriginResponse: string(by),
	}
	switch reqParam.TradeType {
	case enum2.FWC:
		re.Action = dto.Action{
			Action: action.Action_Redirect.String(),
			Parameters: map[string]string{
				action.PrepayID: resp.ReservedTransactionId,
			},
		}
	case enum2.JSAPI, enum2.LETPAY:
		if p.payment == enum.Payment_Wxpay {
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
			re.Action = dto.Action{
				Action:     action.Action_Prepay.String(),
				Parameters: actionParams,
			}
		}
	}
	return re, nil
}
func (p *Pay) buildPayParams(req *dto.PayOrder) *PaymentRequest {
	result := &PaymentRequest{
		MchntCd:              p.C.MchID,
		RandomStr:            tools.GenerateRandomDigits(4),
		OrderAmt:             req.Order.PayAmount.Total,
		MchntOrderNo:         enum2.GenOrder(p.Extra.OrderPrefix, req.Order.OrderNo),
		ProductId:            "",
		TermId:               tools.GenerateRandomDigits(4),
		GoodsDes:             req.Order.Subject,
		GoodsDetail:          "",
		GoodsTag:             "",
		AddnInf:              req.PassbackParams,
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
		Sign: "",
	}
	if req.SceneInfo != nil {
		result.TermIp = req.SceneInfo.ClientIp
		if req.SceneInfo.DeviceID == "" {
			result.TermId = tools.GenerateRandomDigits(4)
		}
		if req.SceneInfo.H5Info.Type != "" {
			deviceInfo := DeviceInfo{
				Type:    req.SceneInfo.H5Info.Type,
				AppName: req.SceneInfo.H5Info.AppName,
				AppUrl:  req.SceneInfo.H5Info.Url,
			}
			if req.SceneInfo.H5Info.IOSPackage != "" {
				deviceInfo.AppUrl = req.SceneInfo.H5Info.IOSPackage
			}
			if req.SceneInfo.H5Info.AndroidPackage != "" {
				deviceInfo.AppUrl = req.SceneInfo.H5Info.AndroidPackage
			}
			result.ReservedDeviceInfo = deviceInfo.String()
		}
	}

	if req.Order.CreateAt.IsZero() {
		result.TxnBeginTs = time.Now().Format("20060102150405")
	} else {
		result.TxnBeginTs = req.Order.CreateAt.Format("20060102150405")
	}
	if p.C.Version == "" {
		result.Version = enum2.Version
	} else {
		result.Version = p.C.Version
	}

	if p.payment == enum.Payment_Wxpay {
		result.TradeType = enum2.WxPaymentProductMap[p.paymentProduct]
	}
	if p.payment == enum.Payment_Alipay {
		result.TradeType = enum2.AliPaymentProductMap[p.paymentProduct]
	}
	result.GenSign(p.C.APIKey)
	return result
}
func (p *Pay) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	reqParam := p.buildQueryParams(req)
	r := p.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := r.SetContext(ctx).SetBody(reqParam).Post(p.C.ApiHost + payQueryMethodPath)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("Query timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("order query request error").WithCause(err)
	}
	by := result.Body()
	resp := &OrderDetail{}
	err = json.Unmarshal(by, resp)
	if err != nil {
		return nil, err
	}
	logc.Info(ctx, "fuiou-Query-result", logc.Field("Response", string(by)), logc.Field("req", reqParam))
	if !resp.IsSuccess() {
		return nil, errors2.ErrorSystemError("query is error;err:%s", resp.ResultMsg).WithCause(errors.New(fmt.Sprintf("code:%s;Msg:%s", resp.ResultCode, resp.ResultMsg)))
	}
	ts, _ := time.Parse("20060102150405", resp.ReservedTxnFinTs)
	total, _ := strconv.Atoi(resp.OrderAmt)
	return &dto.PayDetail{
		OrderNo: req.OrderNo,
		TradeNo: resp.TransactionId,
		PayAmount: dto.Amount{
			//Currency: enum.Currency_CNY.String(),
			Total: int64(total),
		},
		Status:         enum2.GetPaymentStatus(resp.TransStat),
		PaymentProduct: p.paymentProduct.String(),
		SuccessTime:    ts.Unix(),
		OriginResponse: string(by),
	}, nil
}
func (p *Pay) buildQueryParams(req dto.Query) OrderRequest {
	result := &OrderRequest{
		Version:      "",
		MchntCd:      p.C.MchID,
		RandomStr:    tools.GenerateRandomDigits(4),
		OrderType:    enum2.GetOrderType(p.payment, p.paymentProduct),
		MchntOrderNo: enum2.GenOrder(p.Extra.OrderPrefix, req.OrderNo),
		TermId:       tools.GenerateRandomDigits(4),
		Sign:         "",
	}
	if p.C.Version == "" {
		result.Version = enum2.Version
	} else {
		result.Version = p.C.Version
	}
	result.GenSign(p.C.APIKey)
	return *result
}

func (p *Pay) Close(ctx context.Context, req dto.CloseQuery) error {
	reqParam := p.buildCloseParams(req)
	if err := reqParam.Validate(); err != nil {
		return err
	}
	r := p.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := r.SetContext(ctx).SetBody(reqParam).Post(p.C.ApiHost + payCloseMethodPath)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return errors2.ErrorTimeOut("Close timeout").WithCause(err)
	}
	if err != nil {
		return errors2.ErrorSystemError("order close request error").WithCause(err)
	}
	by := result.Body()
	resp := &CloseOrderResponse{}
	err = json.Unmarshal(by, resp)
	logc.Info(ctx, "fuiou-Close-result", logc.Field("Response", string(by)), logc.Field("req", reqParam))
	if !resp.IsSuccess() {
		return errors2.ErrorSystemError("close is error;err:%s", resp.ResultMsg).WithCause(errors.New(fmt.Sprintf("code:%s;Msg:%s", resp.ResultCode, resp.ResultMsg)))
	}
	return nil
}

func (p *Pay) buildCloseParams(req dto.CloseQuery) CloseOrderRequest {
	result := &CloseOrderRequest{
		Version:      "",
		MchntCd:      p.C.MchID,
		RandomStr:    tools.GenerateRandomDigits(4),
		OrderType:    enum2.GetOrderType(p.payment, p.paymentProduct),
		MchntOrderNo: req.OrderNo,
		TermId:       tools.GenerateRandomDigits(4),
		Sign:         "",
	}
	if p.C.Version == "" {
		result.Version = enum2.Version
	} else {
		result.Version = p.C.Version
	}
	result.GenSign(p.C.APIKey)
	return *result
}

func (p *Pay) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors2.ErrorNoSupport("not support Complete")
}
