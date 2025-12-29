package refund

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"
	"strings"
	"time"
)

const (
	refundPath  = "/aggregatePay/commonRefund"
	refundQuery = "/aggregatePay/refundQuery"
)

type Refund struct {
	*fuiou.Api
	paymentProduct enum.PaymentProduct
	payment        enum.Payment
}

func NewRefund(conf config.Config, product enum.PaymentProduct, payment enum.Payment) (iface.Refund, error) {
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
	return &Refund{
		Api:            api,
		paymentProduct: product,
		payment:        payment,
	}, nil
}

func (r *Refund) Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error) {
	reqParam := r.buildRefundRequest(req)
	if err := reqParam.Validate(); err != nil {
		return nil, err
	}
	client := r.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := client.SetContext(ctx).SetBody(reqParam).Post(r.C.ApiHost + refundPath)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &RefundResponse{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "fuiou-Pay-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if err != nil {
		return nil, errors2.ErrorSystemError("parse result error").WithCause(err)
	}
	if !resp.IsSuccess() {
		return nil, errors2.ErrorSystemError("refund is error;err:%s", resp.ResultMsg).WithCause(errors.New(fmt.Sprintf("code:%s;Msg:%s", resp.ResultCode, resp.ResultMsg)))
	}
	reservedRefundAmt, _ := strconv.Atoi(resp.ReservedRefundAmt)
	return &dto.RefundDetail{
		TradeRefundNo:       resp.RefundId,
		RefundNo:            req.RefundNo,
		TradeNo:             "",
		OrderNo:             req.OrderNo,
		Channel:             refund.RefundChannel_ORIGINAL,
		UserReceivedAccount: "",
		SuccessTime:         time.Time{},
		CreateTime:          time.Time{},
		Status:              refund.Status_Pending,
		FundsAccount:        "",
		Amount: dto.Amount{
			Total: int64(reservedRefundAmt),
		},
		OriginResponse: string(by),
	}, nil
}

func (r *Refund) buildRefundRequest(req *dto.RefundRequest) *RefundRequest {
	result := &RefundRequest{
		Version:            "",
		MchntCd:            r.C.MchID,
		RandomStr:          tools.GenerateRandomDigits(4),
		MchntOrderNo:       enum2.GenOrder(r.Extra.OrderPrefix, req.OrderNo),
		RefundOrderNo:      enum2.GenOrder(r.Extra.OrderPrefix, req.RefundNo),
		TermId:             tools.GenerateRandomDigits(4),
		TermIp:             "",
		OrderType:          "",
		TotalAmt:           req.OrderAmount.Total,
		RefundAmt:          req.Amount.Total,
		ReservedFyTermId:   "",
		ReservedDeviceInfo: "",
		Sign:               "",
	}
	if r.C.Version == "" {
		result.Version = enum2.Version
	} else {
		result.Version = r.C.Version
	}

	if r.payment == enum.Payment_Wxpay {
		result.OrderType = "WECHAT"
	}
	if r.payment == enum.Payment_Alipay {
		result.OrderType = "ALIPAY"
	}
	result.GenSign(r.C.APIKey)
	return result
}

func (r *Refund) Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error) {
	reqParam := r.buildRefundQueryRequest(req)
	client := r.Client.Client.R().SetHeader("Content-Type", "application/json")
	result, err := client.SetContext(ctx).SetBody(reqParam).Post(r.C.ApiHost + refundQuery)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("Query timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("order query request error").WithCause(err)
	}
	by := result.Body()
	resp := &RefundQueryResponse{}
	err = json.Unmarshal(by, resp)
	logc.Info(ctx, "fuiou-refund-Query-result", logc.Field("Response", string(by)), logc.Field("req", reqParam))
	if !resp.IsSuccess() {
		return nil, errors2.ErrorSystemError("query is error;err:%s", resp.ResultMsg).WithCause(errors.New(fmt.Sprintf("code:%s;Msg:%s", resp.ResultCode, resp.ResultMsg)))
	}
	reservedRefundAmt, _ := strconv.Atoi(resp.ReservedRefundAmt)
	re := &dto.RefundDetail{
		RefundNo:      req.RefundNo,
		OrderNo:       enum2.ParseOrder(r.Extra.OrderPrefix, resp.MchntOrderNo),
		TradeRefundNo: resp.RefundId,
		TradeNo:       resp.TransactionId,
		Amount: dto.Amount{
			Currency: "CNY",
			Total:    int64(reservedRefundAmt),
		},
		Channel:        refund.RefundChannel_ORIGINAL,
		Status:         resp.GetRefundStatus(),
		OriginResponse: string(by),
	}
	return re, nil
}

func (r *Refund) buildRefundQueryRequest(req dto.RefundQuery) *RefundQueryRequest {
	result := &RefundQueryRequest{
		Version:       "",
		MchntCd:       r.C.MchID,
		RandomStr:     tools.GenerateRandomDigits(4),
		RefundOrderNo: enum2.GenOrder(r.Extra.OrderPrefix, req.RefundNo),
		TermId:        tools.GenerateRandomDigits(4),
		Sign:          "",
	}
	if r.C.Version == "" {
		result.Version = enum2.Version
	} else {
		result.Version = r.C.Version
	}
	result.GenSign(r.C.APIKey)
	return result
}
