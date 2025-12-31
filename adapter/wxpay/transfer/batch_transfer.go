package transfer

import (
	"context"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/transferbatch"
	"time"
)

// https://pay.weixin.qq.com/doc/v3/merchant/4012458841 老版
// https://pay.weixin.qq.com/doc/v3/merchant/4012712115 新版

type BatchTransfer struct {
	*wxpay.Api
	client transferbatch.TransferBatchApiService
	detail transferbatch.TransferDetailApiService
}

func NewTransfer(conf config.Config) (*BatchTransfer, error) {
	api, err := wxpay.InitClient(conf)
	if err != nil {
		return nil, err
	}
	svc := transferbatch.TransferBatchApiService{Client: api.Client}
	detail := transferbatch.TransferDetailApiService{Client: api.Client}
	return &BatchTransfer{
		Api:    api,
		client: svc,
		detail: detail,
	}, nil
}

func (t *BatchTransfer) Transfer(ctx context.Context, req *dto.BatchTransferRequest) (*dto.BatchTransferResponse, error) {
	if err := t.validateTransferParams(req); err != nil {
		return nil, err
	}
	transferRequest := t.buildTransferParams(req)
	resp, result, err := t.client.InitiateBatchTransfer(ctx, transferRequest)
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "not return result")
	}
	if resp.BatchId == nil {
		return nil, until.ErrorHandler(ctx, result, err, "batch_id is nil")
	}
	tt := time.Now()
	if resp.CreateTime != nil {
		tt = *resp.CreateTime
	}
	status := transfer.TransferStatus_TransferStatus_ACCEPTED
	if resp.BatchStatus != nil {
		status = until.GetBatchStatus(*resp.BatchStatus)
	}
	return &dto.BatchTransferResponse{
		TransferNO:      req.TransferNO,
		ThirdTransferNO: until.StringPoint(resp.BatchId),
		CreateTime:      tt,
		TransferStatus:  status,
	}, nil
}

func (t *BatchTransfer) validateTransferParams(req *dto.BatchTransferRequest) error {
	if req.Amount.Total <= 0 {
		return errors.ErrorParamError("total amount must be greater than 0")
	}
	if req.TransferNO == "" {
		return errors.ErrorParamError("transfer no must not empty")
	}
	if req.Subject == "" {
		return errors.ErrorParamError("subject must not empty")
	}
	if req.Remark == "" {
		return errors.ErrorParamError("remark must not empty")
	}
	for _, detail := range req.Details {
		if detail.TradeNO == "" {
			return errors.ErrorParamError("detail no must not empty")
		}
		if detail.UserID == "" {
			return errors.ErrorParamError("open id must not empty")
		}
		if detail.Amount.Total <= 0 {
			return errors.ErrorParamError("amount must greater than zero")
		}
		if detail.Remark == "" {
			return errors.ErrorParamError("remark must not empty")
		}
		if detail.UserName == "" && detail.Amount.Total >= 200000 {
			return errors.ErrorParamError("amount > 2000,user name must not empty")
		}
	}
	return nil
}

func (t *BatchTransfer) buildTransferParams(req *dto.BatchTransferRequest) transferbatch.InitiateBatchTransferRequest {
	result := transferbatch.InitiateBatchTransferRequest{
		Appid:              core.String(t.C.AppID),
		OutBatchNo:         core.String(req.TransferNO),
		BatchName:          core.String(req.Subject),
		BatchRemark:        core.String(req.Remark),
		TotalAmount:        core.Int64(req.Amount.Total),
		TotalNum:           core.Int64(int64(len(req.Details))),
		TransferDetailList: make([]transferbatch.TransferDetailInput, 0, len(req.Details)),
		TransferSceneId:    core.String(req.SceneID),
		NotifyUrl:          core.String(req.NotifyUrl),
	}
	for _, v := range req.Details {
		tmp := transferbatch.TransferDetailInput{
			Openid:         core.String(v.UserID),
			OutDetailNo:    core.String(v.TradeNO),
			TransferAmount: core.Int64(v.Amount.Total),
			TransferRemark: core.String(v.Remark),
		}
		// 大于两千必须传入用户姓名
		if v.Amount.Total >= 200000 {
			tmp.UserName = core.String(v.UserName)
		}
		result.TransferDetailList = append(result.TransferDetailList, tmp)
	}
	return result
}

func (t *BatchTransfer) Query(ctx context.Context, req dto.BatchTransferQueryRequest) (*dto.BatchTransferQueryResponse, error) {
	var resp *transferbatch.TransferBatchEntity
	var result *core.APIResult
	var err error
	limit := int64(20)
	if req.PageSize > 20 {
		limit = req.PageSize
	}
	if req.ThirdTransferNO != "" {
		reqParam := transferbatch.GetTransferBatchByNoRequest{
			BatchId:         core.String(req.ThirdTransferNO),
			NeedQueryDetail: core.Bool(true),
			Limit:           core.Int64(limit),
			Offset:          core.Int64((req.PageNum - 1) * limit),
			DetailStatus:    core.String("ALL"),
		}
		resp, result, err = t.client.GetTransferBatchByNo(ctx, reqParam)
		fmt.Println("||||||||||||||||||||", resp, err)
	} else if req.TransferNO != "" {
		reqParam := transferbatch.GetTransferBatchByNoRequest{
			BatchId:         core.String(req.TransferNO),
			NeedQueryDetail: core.Bool(true),
			Limit:           core.Int64(limit),
			Offset:          core.Int64((req.PageNum - 1) * limit),
			DetailStatus:    core.String("ALL"),
		}
		resp, result, err = t.client.GetTransferBatchByNo(ctx, reqParam)
	} else {
		return nil, errors.ErrorParamError("transfer no or third transfer no must not empty")
	}

	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "not return result")
	}
	status := transfer.TransferStatus_TransferStatus_ACCEPTED
	if resp.TransferBatch.BatchStatus != nil {
		status = until.GetBatchStatus(*resp.TransferBatch.BatchStatus)
	}
	originBy := ""
	if result != nil && result.Response != nil {
		body, _ := tools.GetResponseBody(result.Response)
		originBy = string(body)
	}
	re := &dto.BatchTransferQueryResponse{
		TransferNO:      until.StringPoint(resp.TransferBatch.OutBatchNo),
		ThirdTransferNO: until.StringPoint(resp.TransferBatch.BatchId),
		TransferStatus:  status,
		Total:           until.Int64Point(resp.TransferBatch.TotalNum),
		OriginResponse:  originBy,
		Details:         make([]dto.BatchTransferQueryDetail, 0, len(resp.TransferDetailList)),
	}
	for _, item := range resp.TransferDetailList {
		re.Details = append(re.Details, dto.BatchTransferQueryDetail{
			UserID:          "",
			UserName:        "",
			TradeNO:         until.StringPoint(item.OutDetailNo),
			ThirdTransferNO: until.StringPoint(item.DetailId),
			Amount:          dto.Amount{},
			Remark:          "",
			TransferStatus:  until.GetTransferDetailStatus(until.StringPoint(item.DetailStatus)),
			ErrorRemark:     "",
		})
	}
	return re, nil
}
