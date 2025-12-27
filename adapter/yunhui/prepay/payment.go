package prepay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/yunhui"
	"github.com/singer-stack-lab/payment-sdk/adapter/yunhui/model"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	"github.com/singer-stack-lab/payment-sdk/enum/action"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
	errors2 "github.com/singer-stack-lab/payment-sdk/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

const (
	PayPath   = "/api/pay/unifiedOrderApi"
	queryPath = "/api/pay/queryApi"
)

type Payment struct {
	*yunhui.Api
	paymentProduct enum.PaymentProduct
	payment        enum.Payment
}

func NewPayment(conf config.Config, product enum.PaymentProduct, payment enum.Payment) (driver.Pay, error) {
	api, err := yunhui.NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &Payment{
		Api:            api,
		paymentProduct: product,
		payment:        payment,
	}, nil
}

func (p *Payment) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := p.BuildPayment(req)
	result, err := p.Client.DoPost(ctx, reqParam, p.C.ApiHost+PayPath, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.PaymentResp{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "yunhui-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.CommonResp.Error() != nil {
		return nil, resp.CommonResp.Error()
	}
	if resp.Data == nil {
		return nil, errors.New(fmt.Sprintf("not find body:%s", string(by)))
	}
	if resp.Data != nil && resp.Data.Error.Error() != nil {
		return nil, resp.Data.Error.Error()
	}
	re := &dto.PayResponse{
		OrderNo: resp.Data.MchOrderNo,
		TradeNo: resp.Data.PayOrderId,
		PayAmount: dto.Amount{
			Currency: req.Order.PayAmount.Currency,
			Total:    req.Order.PayAmount.Total,
		},
		Status:         enum.Status_Pending,
		PaymentProduct: p.paymentProduct.String(),
		OriginResponse: string(by),
		Action: dto.Action{
			Action:     action.Action_Redirect.String(),
			Parameters: nil,
			Url:        resp.Data.PayData,
		},
	}
	return re, nil
}

func (p *Payment) BuildPayment(req *dto.PayOrder) *model.PaymentRequest {
	r := &model.PaymentRequest{
		Amount:       req.Order.PayAmount.Total,
		MchNo:        p.C.MchID,
		WayCode:      "QR_CASHIER",
		ApiInfo:      p.Extra.TermNO,
		MchOrderNo:   req.Order.OrderNo,
		NotifyUrl:    req.NotifyUrl,
		Subject:      req.Order.Subject,
		Body:         req.Order.Subject,
		AppId:        p.C.AppID,
		Currency:     "cny",
		ChannelExtra: "",
		ReqTime:      time.Now().UnixMilli(),
		ClientIp:     "",
		ReturnUrl:    req.RedirectUrl,
	}
	if req.TimeExpire > 0 && time.Now().Unix() < req.TimeExpire {
		r.ExpiredTime = req.TimeExpire - time.Now().Unix()
	}
	r.ChannelExtra = fmt.Sprintf(`{"payDataType":"%s"}`, "codeUrl")
	//switch p.payment {
	//case enum.Payment_Alipay:
	//	r.WayCode = "ALI_JSAPI"
	//	r.ChannelExtra = fmt.Sprintf(`{"buyerUserId":"%s"}`, req.Payer.OpenID)
	//case enum.Payment_Wxpay:
	//	r.ChannelExtra = fmt.Sprintf(`{"payDataType":"%s"}`, "codeUrl")
	//	//switch p.paymentProduct {
	//	//case enum.PaymentProduct_JSAPI:
	//	//	r.WayCode = "WX_H5"
	//	//}
	//}
	return r
}

func (p *Payment) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	reqParams := model.PaymentQuery{
		MchNo:      p.C.MchID,
		ApiInfo:    p.Extra.TermNO,
		MchOrderNo: req.OrderNo,
		AppId:      p.C.AppID,
		ReqTime:    time.Now().UnixMilli(),
	}
	result, err := p.Client.DoPost(ctx, reqParams, p.C.ApiHost+queryPath, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.PaymentQueryResp{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParams)
	logc.Info(ctx, "yunhui-query-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.CommonResp.Error() != nil {
		return nil, resp.CommonResp.Error()
	}
	if resp.Data == nil {
		return nil, errors.New(fmt.Sprintf("not find body:%s", string(by)))
	}
	if resp.Data != nil && resp.Data.Error.Error() != nil {
		return nil, resp.Data.Error.Error()
	}
	return &dto.PayDetail{
		OrderNo: resp.Data.MchOrderNo,
		TradeNo: resp.Data.PayOrderId,
		PayAmount: dto.Amount{
			Total: int64(resp.Data.Amount),
		},
		Status:         model.GetPaymentStatus(resp.Data.State),
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		OriginResponse: string(by),
	}, nil
}

func (p *Payment) Close(ctx context.Context, req dto.CloseQuery) error {
	return errors2.ErrorNoSupport("not support Close")
}

func (p *Payment) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors2.ErrorNoSupport("not support Complete")
}
