package fuiou

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"io/ioutil"
	"net/http"
	"strconv"
)

type APICallback struct {
	C     config.Config
	Extra enum2.Extra
}

func NewAPICallback(c config.Config) (*APICallback, error) {
	extra := enum2.Extra{}
	if c.Extra != "" {
		err := json.Unmarshal([]byte(c.Extra), &extra)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("not Unmarshal extra:[%s]", err.Error()))
		}
	}
	return &APICallback{
		C:     c,
		Extra: extra,
	}, nil
}

func (a *APICallback) CallbackPaymentParse(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	originBy, err := getRequestBody(req)
	if err != nil {
		return nil, err
	}
	resp := &model.PaymentCallback{}
	err = json.Unmarshal(originBy, resp)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, errors.New(fmt.Sprintf("%s - %s", resp.ResultCode, resp.ResultMsg))
	}

	if resp.Sign != resp.GenSign(a.C.APIKey) {
		return nil, errors.New("签名错误：" + string(originBy))
	}
	// 外部需要在比对下金额
	status := enum.Status_Status_UNKNOWN
	orderAmt, _ := strconv.Atoi(resp.OrderAmt)
	if orderAmt > 0 {
		status = enum.Status_Success
	}
	return &dto.CallbackPayDetail{
		OrderNo: enum2.ParseOrder(a.Extra.OrderPrefix, resp.MchntOrderNo),
		TradeNo: resp.TransactionId,
		PayAmount: dto.Amount{
			Currency: "CNY",
			Total:    int64(orderAmt),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		OriginResponse: string(originBy),
	}, nil
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
