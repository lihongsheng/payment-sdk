package prepay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/client"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"strconv"
	"time"
)

type Api struct {
	*client.Client
	paymentProduct enum.PaymentProduct
	payment        enum.Payment
}

func NewApi(api *client.Client, product enum.PaymentProduct, payment enum.Payment) (*Api, error) {
	if _, exists := enum2.WxPaymentProductMap[product]; !exists {
		return nil, errors2.ErrorNoSupport("product [%s] is not exists", product.String())
	}
	return &Api{
		Client:         api,
		paymentProduct: product,
		payment:        payment,
	}, nil
}

func (p *Api) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	reqParam := p.buildQueryParams(req)
	r := p.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := r.SetContext(ctx).SetBody(reqParam).Post(p.C.API.ApiHost + payQueryMethodPath)
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
func (p *Api) buildQueryParams(req dto.Query) OrderRequest {
	result := &OrderRequest{
		Version:      p.C.API.Version,
		MchntCd:      p.C.Merchant.MchID,
		RandomStr:    tools.GenerateRandomDigits(4),
		OrderType:    enum2.GetOrderType(p.payment, p.paymentProduct),
		MchntOrderNo: enum2.GenOrder(p.C.Merchant.OrderPrefix, req.OrderNo),
		TermId:       tools.GenerateRandomDigits(4),
		Sign:         "",
	}
	result.GenSign(p.C.Merchant.APISecret)
	return *result
}

func (p *Api) Close(ctx context.Context, req dto.CloseQuery) error {
	reqParam := p.buildCloseParams(req)
	if err := reqParam.Validate(); err != nil {
		return err
	}
	r := p.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := r.SetContext(ctx).SetBody(reqParam).Post(p.C.API.ApiHost + payCloseMethodPath)
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

func (p *Api) buildCloseParams(req dto.CloseQuery) CloseOrderRequest {
	result := &CloseOrderRequest{
		Version:      p.C.API.Version,
		MchntCd:      p.C.Merchant.MchID,
		RandomStr:    tools.GenerateRandomDigits(4),
		OrderType:    enum2.GetOrderType(p.payment, p.paymentProduct),
		MchntOrderNo: req.OrderNo,
		TermId:       tools.GenerateRandomDigits(4),
		Sign:         "",
	}
	result.GenSign(p.C.Merchant.APISecret)
	return *result
}

func (p *Api) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors2.ErrorNoSupport("not support Complete")
}

func (p *Api) Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	originBy, err := util.GetRequestBody(req)
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

	if resp.Sign != resp.GenSign(p.C.Merchant.APISecret) {
		return nil, errors.New("签名错误：" + string(originBy))
	}
	// 外部需要在比对下金额
	status := enum.Status_Status_UNKNOWN
	orderAmt, _ := strconv.Atoi(resp.OrderAmt)
	if orderAmt > 0 {
		status = enum.Status_Success
	}
	return &dto.CallbackPayDetail{
		OrderNo: enum2.ParseOrder(p.C.Merchant.OrderPrefix, resp.MchntOrderNo),
		TradeNo: resp.TransactionId,
		PayAmount: dto.Amount{
			Currency: "CNY",
			Total:    int64(orderAmt),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		OriginResponse: string(originBy),
		Response:       "1",
		EventAction:    enum1.Event_PAYMENT,
	}, nil
}
