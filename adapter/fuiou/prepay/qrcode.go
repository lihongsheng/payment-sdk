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
	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

const (
	qrcodeMethodPath = "/aggregatePay/preCreate"
)

type Qrcode struct {
	*Api
}

func NewQrcode(api *client.Client, product enum.PaymentProduct, payment enum.Payment) (iface.Pay, error) {
	api2, err := NewApi(api, product, payment)
	if err != nil {
		return nil, err
	}
	return &Qrcode{
		api2,
	}, nil
}

func (p *Qrcode) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := p.buildPayParams(req)
	if err := reqParam.Validate(); err != nil {
		return nil, err
	}
	r := p.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := r.SetContext(ctx).SetBody(reqParam).Post(p.C.ApiHost + qrcodeMethodPath)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &QrcodePaymentResponse{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "fuiou-qrcode-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
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
			Action: action.Action_Qrcode.String(),
			Url:    resp.QrCode,
		},
	}

	return re, nil
}
func (p *Qrcode) buildPayParams(req *dto.PayOrder) *QrcodePaymentRequest {
	result := &QrcodePaymentRequest{
		MchntCd:              p.C.MchID,
		RandomStr:            tools.GenerateRandomDigits(4),
		OrderAmt:             req.Order.PayAmount.Total,
		MchntOrderNo:         enum2.GenOrder(p.C.OrderPrefix, req.Order.OrderNo),
		TermId:               tools.GenerateRandomDigits(4),
		GoodsDes:             req.Order.Subject,
		GoodsDetail:          "",
		GoodsTag:             "",
		AddnInf:              req.PassBackParams,
		CurrType:             req.Order.PayAmount.Currency,
		NotifyURL:            req.NotifyUrl,
		ReservedFyTermId:     "",
		ReservedExpireMinute: 0,
		Sign:                 "",
		Version:              p.C.Version,
	}
	if req.SceneInfo != nil {
		result.TermIp = req.SceneInfo.ClientIp
	}

	if req.Order.CreateAt.IsZero() {
		result.TxnBeginTs = time.Now().Format("20060102150405")
	} else {
		result.TxnBeginTs = req.Order.CreateAt.Format("20060102150405")
	}

	if p.payment == enum.Payment_Wechat {
		result.OrderType = enum2.OrderTypeALIPAY
	}
	if p.payment == enum.Payment_Alipay {
		result.OrderType = enum2.OrderTypeWECHAT
	}
	result.GenSign(p.C.APISecret)
	return result
}
