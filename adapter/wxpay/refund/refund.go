package refund

import (
	"context"
	"encoding/json"
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay"
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay/until"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
	"github.com/singer-stack-lab/payment-sdk/enum/refund"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

type Refund struct {
	*wxpay.Api
	client refunddomestic.RefundsApiService
}

func NewRefund(conf config.Config) (driver.Refund, error) {
	api, err := wxpay.InitClient(conf)
	if err != nil {
		return nil, err
	}
	return &Refund{
		Api:    api,
		client: refunddomestic.RefundsApiService{Client: api.Client},
	}, nil
}

func (r *Refund) Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error) {
	resp, result, err := r.client.Create(ctx, r.buildRefundRequest(req))
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.RefundId == nil || *resp.RefundId == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return RefundId")
	}
	originBy, _ := json.Marshal(resp)
	re := &dto.RefundDetail{
		RefundNo:      until.StringPoint(resp.OutRefundNo),
		OrderNo:       until.StringPoint(resp.OutTradeNo),
		TradeRefundNo: until.StringPoint(resp.RefundId),
		TradeNo:       until.StringPoint(resp.TransactionId),
		Amount: dto.Amount{
			Currency: until.StringPoint(resp.Amount.Currency),
			Total:    until.Int64Point(resp.Amount.Refund),
		},
		Channel:             refund.RefundChannel(refund.RefundChannel_value[string(*resp.Channel)]),
		UserReceivedAccount: until.StringPoint(resp.UserReceivedAccount),
		CreateTime:          *resp.CreateTime,
		FundsAccount:        string(*resp.FundsAccount),
		Status:              until.GetRefundStatus(*resp.Status),
		OriginResponse:      string(originBy),
	}
	if resp.SuccessTime != nil && !resp.SuccessTime.IsZero() {
		re.SuccessTime = *resp.SuccessTime
	}
	return re, nil
}

func (r *Refund) buildRefundRequest(req *dto.RefundRequest) refunddomestic.CreateRequest {
	result := refunddomestic.CreateRequest{
		SubMchid:      nil,
		TransactionId: core.String(req.TradeNo),
		OutTradeNo:    core.String(req.OrderNo),
		OutRefundNo:   core.String(req.RefundNo),
		Reason:        core.String(req.Reason),
		NotifyUrl:     core.String(req.NotifyUrl),
		Amount: &refunddomestic.AmountReq{
			Refund:   core.Int64(req.Amount.Total),
			Total:    core.Int64(req.OrderAmount.Total),
			Currency: core.String(enum.Currency_CNY.String()),
		},
		GoodsDetail: nil,
	}
	if req.TradeNo == "" {
		result.TransactionId = nil
	}
	if req.OrderNo == "" {
		result.OutTradeNo = nil
	}
	if len(req.Goods) > 0 {
		result.GoodsDetail = make([]refunddomestic.GoodsDetail, len(req.Goods))
		for i, v := range req.Goods {
			result.GoodsDetail[i] = refunddomestic.GoodsDetail{
				MerchantGoodsId: core.String(v.Sku),
				UnitPrice:       core.Int64(v.Price),
				RefundAmount:    core.Int64(v.Price * int64(v.Quantity)),
				RefundQuantity:  core.Int64(int64(v.Quantity)),
			}
		}
	}
	if req.Amount.Currency != "" {
		result.Amount.Currency = core.String(req.Amount.Currency)
	}
	return result
}

func (r *Refund) Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error) {
	resp, result, err := r.client.QueryByOutRefundNo(ctx, refunddomestic.QueryByOutRefundNoRequest{
		OutRefundNo: core.String(req.RefundNo),
	})
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.RefundId == nil || *resp.RefundId == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return RefundId")
	}
	originBy, _ := json.Marshal(resp)
	re := &dto.RefundDetail{
		RefundNo:      until.StringPoint(resp.OutRefundNo),
		OrderNo:       until.StringPoint(resp.OutTradeNo),
		TradeRefundNo: until.StringPoint(resp.RefundId),
		TradeNo:       until.StringPoint(resp.TransactionId),
		Amount: dto.Amount{
			Currency: until.StringPoint(resp.Amount.Currency),
			Total:    until.Int64Point(resp.Amount.Refund),
		},
		Channel:             refund.RefundChannel(refund.RefundChannel_value[string(*resp.Channel)]),
		UserReceivedAccount: until.StringPoint(resp.UserReceivedAccount),
		CreateTime:          *resp.CreateTime,
		FundsAccount:        string(*resp.FundsAccount),
		Status:              until.GetRefundStatus(*resp.Status),
		OriginResponse:      string(originBy),
	}
	if resp.SuccessTime != nil || !resp.SuccessTime.IsZero() {
		re.SuccessTime = *resp.SuccessTime
	}
	return re, nil

}
