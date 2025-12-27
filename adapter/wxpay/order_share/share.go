package order_share

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay"
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay/until"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	"github.com/singer-stack-lab/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
)

type Share struct {
	*wxpay.Api
	orderClient    profitsharing.OrdersApiService
	returnOrder    profitsharing.ReturnOrdersApiService
	receiverClient profitsharing.ReceiversApiService
}

func NewShare(conf config.Config) (*Share, error) {
	api, err := wxpay.InitClient(conf)
	if err != nil {
		return nil, err
	}
	return &Share{
		Api:            api,
		orderClient:    profitsharing.OrdersApiService{Client: api.Client},
		returnOrder:    profitsharing.ReturnOrdersApiService{Client: api.Client},
		receiverClient: profitsharing.ReceiversApiService{Client: api.Client},
	}, nil
}

func (s *Share) AsyncCreateOrder(ctx context.Context, req *dto.CreateShareOrderRequest) (*dto.ShareOrderResponse, error) {
	if err := s.validateCreateParams(ctx, req); err != nil {
		return nil, err
	}
	reqParams := s.buildCreateParams(ctx, req)
	resp, result, err := s.orderClient.CreateOrder(ctx, reqParams)
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "response is nil")
	}
	if resp.OrderId == nil {
		return nil, until.ErrorHandler(ctx, result, err, "order_id is nil")
	}
	r := &dto.ShareOrderResponse{
		ShareOrderNo:      *resp.OutOrderNo,
		Status:            until.GetShareTransferStatus(resp.State),
		ThirdShareOrderNo: *resp.OrderId,
		TradeNo:           *resp.TransactionId,
	}
	return r, nil
}

func (s *Share) buildCreateParams(ctx context.Context, req *dto.CreateShareOrderRequest) profitsharing.CreateOrderRequest {
	r := profitsharing.CreateOrderRequest{
		Appid:           core.String(s.C.AppID),
		TransactionId:   core.String(req.TradeNo),
		OutOrderNo:      core.String(req.ShareOrderNo),
		UnfreezeUnsplit: core.Bool(req.UnfreezeUnsplit),
		Receivers:       make([]profitsharing.CreateOrderReceiver, 0, len(req.Users)+len(req.Mch)),
	}
	if len(req.Users) > 0 {
		for _, u := range req.Users {
			tmp := profitsharing.CreateOrderReceiver{}
			tmp.Type = core.String("PERSONAL_OPENID")
			tmp.Account = core.String(u.User.OpenID)
			tmp.Name = core.String(u.User.UserName)
			tmp.Amount = core.Int64(u.Amount.Total)
			tmp.Description = core.String(u.Remark)
			if u.User.UserName != "" {
				userName, _ := encryptOAEP(s.PublicKey, []byte(u.User.UserName))
				tmp.Name = core.String(userName)
			} else {
				tmp.Name = nil
			}
			r.Receivers = append(r.Receivers, tmp)
		}
	}
	if len(req.Mch) > 0 {
		for _, m := range req.Mch {
			tmp := profitsharing.CreateOrderReceiver{}
			tmp.Type = core.String("MERCHANT_ID")
			tmp.Account = core.String(m.Account)
			tmp.Name = core.String(m.AccountName)
			tmp.Amount = core.Int64(m.Amount.Total)
			tmp.Description = core.String(m.Remark)
			if m.AccountName != "" {
				userName, _ := encryptOAEP(s.PublicKey, []byte(m.AccountName))
				tmp.Name = core.String(userName)
			} else {
				tmp.Name = nil
			}
			r.Receivers = append(r.Receivers, tmp)
		}
	}
	return r
}
func (s *Share) validateCreateParams(ctx context.Context, req *dto.CreateShareOrderRequest) error {
	if req == nil {
		return errors.ErrorParamError("share order params is empty")
	}
	if req.TradeNo == "" {
		return errors.ErrorParamError("transaction_id is required")
	}
	if req.ShareOrderNo == "" {
		return errors.ErrorParamError("out_order_no is required")
	}
	if req.Users == nil || len(req.Users) < 1 {
		return errors.ErrorParamError("users is required")
	}
	return nil
}

func (s *Share) Query(ctx context.Context, req dto.ShareOrderQueryRequest) (*dto.ShareOrderDetailResponse, error) {
	if req.ThirdShareOrderNo == "" {
		return nil, errors.ErrorParamError("third_share_order_no is required")
	}
	if req.ShareOrderNo == "" {
		return nil, errors.ErrorParamError("share_order_no is required")
	}
	reqParams := profitsharing.QueryOrderRequest{
		TransactionId: core.String(req.ThirdShareOrderNo),
		OutOrderNo:    core.String(req.ShareOrderNo),
	}
	resp, result, err := s.orderClient.QueryOrder(ctx, reqParams)
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "response is nil")
	}
	if resp.OrderId == nil {
		return nil, until.ErrorHandler(ctx, result, err, "order_id is nil")
	}
	r := &dto.ShareOrderDetailResponse{
		ShareOrderNo:      until.StringPoint(resp.OutOrderNo),
		Status:            until.GetShareTransferStatus(resp.State),
		ThirdShareOrderNo: until.StringPoint(resp.OrderId),
		TradeNo:           until.StringPoint(resp.TransactionId),
		Receivers:         make([]dto.ShareUserReceive, 0, len(resp.Receivers)),
		MchReceivers:      make([]dto.ShareMchReceive, 0, len(resp.Receivers)),
	}
	for _, item := range resp.Receivers {
		ty := *item.Type
		switch ty {
		case profitsharing.RECEIVERTYPE_PERSONAL_OPENID:
			tmp := dto.ShareUserReceive{
				User: dto.User{
					OpenID: until.StringPoint(item.Account),
				},
				Amount: dto.Amount{
					Total: until.Int64Point(item.Amount),
				},
				Remark:       until.StringPoint(item.Description),
				Status:       0,
				ShareOrderNO: until.StringPoint(item.DetailId),
			}
			if item.FailReason != nil {
				tmp.FailReason = string(*item.FailReason)
			}
			r.Receivers = append(r.Receivers, tmp)
		case profitsharing.RECEIVERTYPE_MERCHANT_ID:
			tmp := dto.ShareMchReceive{
				Account: until.StringPoint(item.Account),
				Amount: dto.Amount{
					Total: until.Int64Point(item.Amount),
				},
				Remark:       until.StringPoint(item.Description),
				Status:       0,
				ShareOrderNO: until.StringPoint(item.DetailId),
			}
			if item.FailReason != nil {
				tmp.FailReason = string(*item.FailReason)
			}
			r.MchReceivers = append(r.MchReceivers, tmp)
		}
	}
	return r, nil
}

//func (s *Share) RollBack(ctx context.Context, req *dto.CreateShareOrderRequest) (*dto.ShareOrderResponse, error) {
//}

// RSAES-OAEP 加密
func encryptOAEP(pub *rsa.PublicKey, plaintext []byte) (string, error) {
	// 使用 SHA-256 哈希函数，空标签
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, plaintext, nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
