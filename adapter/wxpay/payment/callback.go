package payment

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"time"
)

type CallbackMethod struct {
	*wxpay.Api
}

func NewCallback(conf config.Config) (*CallbackMethod, error) {
	api, err := wxpay.InitClient(conf)
	if err != nil {
		return nil, err
	}
	return &CallbackMethod{
		Api: api,
	}, nil
}

func (c *CallbackMethod) Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	no, err := notify.NewRSANotifyHandler(c.C.APISecret, c.Verifier)
	if err != nil {
		return nil, errors2.ErrorSystemError("wxpay new RSA NotifyHandler errors").WithCause(err)
	}
	var resp = &payments.Transaction{}
	_, err = no.ParseNotifyRequest(ctx, req, resp)
	if err != nil {
		_, proErr := until.ProcessBody(c.C.APISecret, req, resp)
		if proErr != nil {
			return nil, err
		}
		//return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	if resp.TransactionId == nil || (resp.TransactionId != nil && *resp.TransactionId == "") {
		return nil, errors2.ErrorSystemError("wxpay parse notify not find TransactionId").WithCause(err)
	}
	status := until.PaymentStatus[until.StringPoint(resp.TradeState)]
	if status == enum.Status_Status_UNKNOWN {
		logc.Error(ctx, "wxPayErrStatus", logc.Field("resp", resp))
		return nil, until.ErrorHandler(ctx, nil, err, "status is unknown")
	}
	var successTime time.Time
	if resp.SuccessTime == nil && *resp.SuccessTime != "" {
		successTime, _ = time.Parse(time.RFC3339, *resp.SuccessTime)
	}
	originBy, _ := json.Marshal(resp)
	return &dto.CallbackPayDetail{
		OrderNo: until.StringPoint(resp.OutTradeNo),
		TradeNo: until.StringPoint(resp.TransactionId),
		PayAmount: dto.Amount{
			Currency: until.StringPoint(resp.Amount.Currency),
			Total:    until.Int64Point(resp.Amount.Total),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		SuccessTime:    successTime.Unix(),
		OriginResponse: string(originBy),
		EventAction:    enum1.Event_PAYMENT,
	}, nil
}
