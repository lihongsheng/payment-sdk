package yunhui

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/yunhui/client"
	"github.com/lihongsheng/payment-sdk/adapter/yunhui/enum"
	"github.com/lihongsheng/payment-sdk/adapter/yunhui/model"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ApiCallback struct {
	C     config.Config
	Extra enum.Extra
	Sign  *client.Sign
}

func NewApiCallback(c config.Config) (*ApiCallback, error) {
	c.ApiHost = strings.TrimRight(c.ApiHost, "/")
	extra := enum.Extra{}
	if c.Extra != "" {
		err := json.Unmarshal([]byte(c.Extra), &extra)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("not Unmarshal extra:[%s]", err.Error()))
		}
	}
	if extra.TermNO == "" {
		return nil, errors.New("拉卡拉支付必须配置终端号")
	}
	sign, err := client.NewSign(c)
	if err != nil {
		return nil, err
	}
	return &ApiCallback{
		C:     c,
		Extra: extra,
		Sign:  sign,
	}, nil
}

func (a *ApiCallback) ParsePayment(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	// 补验签，或者程序比对支付支付单号是否一致
	bodyBytes, err := getRequestBody(req)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, err
	}
	err = a.Sign.Verify(values)
	if err != nil {
		return nil, err
	}
	state, err := strconv.Atoi(values.Get("state"))
	amount, _ := strconv.Atoi(values.Get("amount"))
	if err != nil {
		return nil, err
	}
	re := &dto.CallbackPayDetail{
		OrderNo: values.Get("mchOrderNo"),
		TradeNo: values.Get("payOrderId"),
		PayAmount: dto.Amount{
			Currency: "CNY",
			Total:    int64(amount),
		},
		Status:         model.GetPaymentStatus(state),
		PaymentProduct: "",
		OriginResponse: string(bodyBytes),
	}

	return re, nil
}

func (a *ApiCallback) ParseRefund(ctx context.Context, req *http.Request) (*dto.CallbackRefundDetail, error) {
	// 补验签，或者程序比对支付支付单号是否一致
	bodyBytes, err := getRequestBody(req)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, err
	}
	err = a.Sign.Verify(values)
	if err != nil {
		return nil, err
	}
	state, err := strconv.Atoi(values.Get("state"))
	amount, _ := strconv.Atoi(values.Get("refundAmount"))
	re := &dto.CallbackRefundDetail{
		RefundNo:      values.Get("mchRefundNo"),
		OrderNo:       "",
		TradeRefundNo: values.Get("refundOrderId"),
		TradeNo:       "",
		Amount: dto.Amount{
			Currency: "CNY",
			Total:    int64(amount),
		},
		UserReceivedAccount: "",
		OriginResponse:      string(bodyBytes),
		Status:              model.GetRefundStatus(state),
	}
	return re, nil
}

func getRequestBody(request *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
