package alipay

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/client"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/util"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/errors"
	"time"
)

type Api struct {
	Client *client.Client
	C      config.Config
}

func NewApi(conf config.Config) (*Api, error) {
	client, err := client.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &Api{
		C:      conf,
		Client: client,
	}, nil
}

func (a *Api) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors.ErrorNoSupport("not support Complete")
}

func (a *Api) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	reqParam := model.PaymentQueryRequest{
		OutTradeNo: req.OrderNo,
		TradeNo:    req.TradeNo,
		QueryOptions: []string{
			"trade_settle_info",
		},
	}
	commonParam := a.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_TRADES_QUERY
	resp, err := a.Client.DoPost(commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.PaymentQueryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error")
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayTradeQueryResponse.Code != "" && response.AlipayTradeQueryResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if response.AlipayTradeQueryResponse.SubCode == "" && response.AlipayTradeQueryResponse.TradeNo != "" {
		respTrue = true
	}
	if !respTrue {
		return nil, errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	amount, _ := util.AmountToCents(response.AlipayTradeQueryResponse.TotalAmount)
	successTime := int64(0)
	if response.AlipayTradeQueryResponse.SendPayDate != "" {
		t, _ := time.Parse(time.DateTime, response.AlipayTradeQueryResponse.SendPayDate)
		successTime = t.Unix()
	}
	return &dto.PayDetail{
		OrderNo: response.AlipayTradeQueryResponse.OutTradeNo,
		TradeNo: response.AlipayTradeQueryResponse.TradeNo,
		PayAmount: dto.Amount{
			Total:    int64(amount),
			Currency: response.AlipayTradeQueryResponse.TransCurrency,
		},
		Status:         util.PaymentStatus(response.AlipayTradeQueryResponse.TradeStatus),
		PaymentProduct: payment.PaymentProduct_JSAPI.String(),
		SuccessTime:    successTime,
		OriginResponse: string(body),
	}, nil
}

func (a *Api) Close(ctx context.Context, req dto.CloseQuery) error {
	reqParam := model.PaymentCloseRequest{
		OutTradeNo: req.OrderNo,
	}
	commonParam := a.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_TRADES_CLOSE
	resp, err := a.Client.DoPost(commonParam, reqParam, nil)
	if err != nil {
		return err
	}
	body := resp.Body()
	var response model.PaymentCloseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	respTrue := false
	if response.TradeCloseResponse.Code != "" && response.TradeCloseResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if response.TradeCloseResponse.SubCode == "" && response.TradeCloseResponse.TradeNo != "" {
		respTrue = true
	}
	if !respTrue {
		return errors.ErrorSystemError("not return trade_no;"+string(body), nil)
	}
	return nil
}
