package unit_transfer

import (
	"context"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client/mchtransfer"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

type Transfer struct {
	*wxpay.Api
	client mchtransfer.TransferApiService
}

func NewTransfer(conf config.Config) (iface.UnitTransfer, error) {
	api, err := wxpay.InitClient(conf)
	if err != nil {
		return nil, err
	}
	svc := mchtransfer.TransferApiService{Client: api.Client}
	return &Transfer{
		Api:    api,
		client: svc,
	}, nil
}

func (t *Transfer) Transfer(ctx context.Context, req *dto.UintTransferRequest) (*dto.UintTransferResponse, error) {
	if err := t.validateTransferParams(req); err != nil {
		return nil, err
	}
	reqParams := t.BuildTransfer(req)
	resp, result, err := t.client.TransferBill(ctx, reqParams)
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "response is nil")
	}
	if resp.TransferBillNo == "" {
		return nil, until.ErrorHandler(ctx, result, err, "transfer_bill_no is nil")
	}
	r := &dto.UintTransferResponse{
		TransferNo:      req.TransferNo,
		Status:          until.GetTransferStatus(resp.State),
		ThirdTransferNo: resp.TransferBillNo,
		CreateTime:      resp.CreateTime,
		Action: &dto.Action{
			Action: action.Action_Prepay.String(),
			Url:    "",
			Parameters: map[string]string{
				"package": resp.PackageInfo,
			},
		},
	}

	return r, nil
}

func (t *Transfer) validateTransferParams(req *dto.UintTransferRequest) error {
	if req == nil {
		return errors.ErrorParamError("request can not be empty")
	}
	if req.TransferNo == "" {
		return errors.ErrorParamError("transfer_no can not be empty")
	}
	if req.TransferAmount.Total <= 0 {
		return errors.ErrorParamError("total amount must be greater than 0")
	}
	if req.User.OpenID == "" {
		return errors.ErrorParamError("user openid can not be empty")
	}
	if req.Remark == "" {
		return errors.ErrorParamError("remark can not be empty")
	}
	if len(req.SceneReport) < 1 {
		return errors.ErrorParamError("scene report can not be empty")
	}
	return nil
}

func (t *Transfer) BuildTransfer(req *dto.UintTransferRequest) mchtransfer.TransferBillRequest {
	r := mchtransfer.TransferBillRequest{
		Appid:                    t.C.AppID,
		NotifyUrl:                req.NotifyUrl,
		Openid:                   req.User.OpenID,
		UserName:                 req.User.UserName,
		OutBillNo:                req.TransferNo,
		TransferAmount:           req.TransferAmount.Total,
		TransferRemark:           req.Remark,
		TransferSceneId:          req.SceneId,
		TransferSceneReportInfos: make([]mchtransfer.TransferSceneReportInfo, 0, len(req.SceneReport)),
	}
	for _, v := range req.SceneReport {
		r.TransferSceneReportInfos = append(r.TransferSceneReportInfos, mchtransfer.TransferSceneReportInfo{
			InfoContent: v.Content,
			InfoType:    v.Type,
		})
	}
	return r
}

func (t *Transfer) Cancel(ctx context.Context, req dto.UnitTransferCancelRequest) (*dto.UintTransferCancelResponse, error) {
	if req.TransferNo == "" {
		return nil, errors.ErrorParamError("third_transfer_no can not be empty")
	}
	resp, result, err := t.client.Cancel(ctx, mchtransfer.TransferBillCancelRequest{
		OutBillNo: req.TransferNo,
	})
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "response is nil")
	}
	if resp.OutBillNo == "" {
		return nil, until.ErrorHandler(ctx, result, err, "transfer_bill_no is nil")
	}
	r := &dto.UintTransferCancelResponse{
		TransferNo:      req.TransferNo,
		ThirdTransferNo: resp.TransferBillNo,
		Status:          until.GetTransferStatus(resp.State),
	}
	return r, nil
}

func (t *Transfer) Query(ctx context.Context, req dto.UintTransferQueryRequest) (*dto.UintTransferDetailResponse, error) {
	var resp *mchtransfer.TransferBillDetailResponse
	var result *core.APIResult
	var err error
	if req.ThirdTransferNo != "" {
		resp, result, err = t.client.GetTransferBillByNo(ctx, mchtransfer.GetTransferBilByNORequest{
			TransferBillNo: req.ThirdTransferNo,
		})
	} else if req.TransferNo != "" {
		resp, result, err = t.client.GetTransferBillByOutBillNo(ctx, mchtransfer.GetTransferBilByOutBillNORequest{
			OutBillNo: req.TransferNo,
		})
	} else {
		return nil, errors.ErrorParamError("transfer_no or third_transfer_no can not be empty")
	}
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "response is nil")
	}
	r := &dto.UintTransferDetailResponse{
		TransferNo:      req.TransferNo,
		ThirdTransferNo: resp.TransferBillNo,
		TransferAmount: dto.Amount{
			Total:    resp.TransferAmount,
			Currency: enum.Currency_CNY.String(),
		},
		User: dto.User{
			UnionID:  "",
			OpenID:   resp.Openid,
			UserName: resp.UserName,
		},
		Status:      until.GetTransferStatus(resp.State),
		CreateTime:  resp.CreateTime,
		ArrivalTime: resp.UpdateTime,
	}
	return r, nil
}
