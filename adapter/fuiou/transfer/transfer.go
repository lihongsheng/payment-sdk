package transfer

import (
	"context"
	"encoding/xml"
	error2 "errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"github.com/lihongsheng/payment-sdk/errors"
	"time"
)

const (
	transferPath      = "/balanceAllocate.fuiou"
	transferQueryPath = "/queryBookkeepingTrade.fuiou"
)

type Transfer struct {
	*fuiou.Api
}

func NewTransfer(conf config.Config) (*Transfer, error) {
	api, err := fuiou.NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &Transfer{
		Api: api,
	}, nil
}

func (t *Transfer) Transfer(ctx context.Context, req *dto.BatchTransferRequest) (*dto.BatchTransferResponse, error) {
	request := t.buildTransferParams(req)
	if err := request.Validate(); err != nil {
		return nil, err
	}
	resp, err := t.Client.PostReqFrom(ctx, transferPath, request, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var user = &model.TransferResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return nil, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	tResp := &dto.BatchTransferResponse{
		TransferNO:      req.TransferNO,
		ThirdTransferNO: user.BatchNo,
		//CreateTime:      time.Time{},
		TransferStatus: transfer.TransferStatus_TransferStatus_ACCEPTED,
		OriginResponse: string(body),
	}
	return tResp, nil
}

func (t *Transfer) buildTransferParams(req *dto.BatchTransferRequest) *model.TransferRequest {
	r := &model.TransferRequest{
		MchntCd:       t.C.MchID,
		TraceNo:       req.TransferNO,
		AccountIn:     req.SubAccountIn,
		AccountInList: make([]model.TransferAccountInListItem, 0, len(req.Details)),
	}
	for _, detail := range req.Details {
		allocateType := transfer.AllocateType_Auto
		r.AccountInList = append(r.AccountInList, model.TransferAccountInListItem{
			AccountIn:    detail.UserID,
			AllocateAmt:  int(detail.Amount.Total),
			AllocateType: string(allocateType),
			CleanBankNo:  detail.BankNO,
			InvoiceType:  "",
			Tax:          nil,
			Remark:       detail.Remark,
		})
	}
	return r
}

func (t *Transfer) Query(ctx context.Context, req *dto.BatchTransferQueryRequest) (*dto.BatchTransferQueryResponse, error) {
	if req.StartTime.IsZero() || req.EndTime.IsZero() {
		return nil, error2.New("start_time or end_time is empty")
	}
	request := &model.TransferQueryRequest{
		TraceNo:             fmt.Sprintf("%d", time.Now().UnixMilli()),
		MchntCd:             t.C.MchID,
		TradeType:           "10",
		BatchNo:             req.ThirdTransferNO,
		SrcFasSsn:           "",
		MchntCdTraceNo:      req.TransferNO,
		MchntCdChildTraceNo: "",
		StartDate:           req.StartTime.Format("20060102"),
		EndDate:             req.EndTime.Format("20060102"),
		PageNo:              fmt.Sprintf("%d", req.PageNum),
		PageSize:            fmt.Sprintf("%d", req.PageSize),
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}
	resp, err := t.Client.PostReqFrom(ctx, transferQueryPath, request, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var user = &model.TransferQueryResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return nil, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	tResp := &dto.BatchTransferQueryResponse{
		TransferNO:      req.TransferNO,
		ThirdTransferNO: req.ThirdTransferNO,
		//CreateTime:      time.Time{},
		Total:          int64(user.TotalNum),
		Details:        make([]dto.BatchTransferQueryDetail, 0, len(user.List)),
		TransferStatus: transfer.TransferStatus_TransferStatus_PROCESSING,
		OriginResponse: string(body),
	}
	existsProcessing := false
	for _, detail := range user.List {
		tResp.ThirdTransferNO = detail.BatchNo
		status, remark := t.getTransferStatus(detail.Status)
		if status == transfer.DetailStatus_DetailStatus_PROCESSING {
			existsProcessing = true
		}
		tResp.Details = append(tResp.Details, dto.BatchTransferQueryDetail{
			UserID:   detail.AccountIn,
			UserName: "",
			Amount: dto.Amount{
				Total: int64(detail.TxnAmt),
			},
			Remark:         detail.Remark,
			TradeNO:        detail.SrcFasSsn,
			TransferStatus: status,
			ErrorRemark:    remark + detail.ErrorMsg,
		})
	}
	// 批次处理完成
	if !existsProcessing {
		tResp.TransferStatus = transfer.TransferStatus_TransferStatus_FINISHED
	}
	return tResp, nil
}

func (t *Transfer) getTransferStatus(s string) (transfer.DetailStatus, string) {
	return util.GetTransferStatus(s)
}
