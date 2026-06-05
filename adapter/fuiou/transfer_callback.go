package fuiou

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/client"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"io"
	"net/http"
)

type TransferCallback struct {
	C    config.Config
	Sign *client.Sign
}

func NewTransferCallback(c config.Config) *TransferCallback {
	return &TransferCallback{
		C:    c,
		Sign: client.NewSign(c),
	}
}

func (c *TransferCallback) Parse(ctx context.Context, req *http.Request) (eventType enum.EventType, resp any, err error) {
	bodyBytes, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	encryptResponse := &model.CommonEncryptCallbackResponse{}
	if len(bodyBytes) > 0 {
		err := xml.Unmarshal(bodyBytes, encryptResponse)
		if err != nil {
			return enum.EventTypeUnKnown, nil, err
		}
	}
	switch encryptResponse.NotifyType {
	case string(enum.EventTypeUserChanged):
		resp, err = c.AccountInResultCallback(req)
		eventType = enum.EventTypeUserChanged
	case string(enum.EventTypeTransferChanged):
		resp, err = c.AccountInSettleResultCallback(req)
		eventType = enum.EventTypeTransferChanged
	}
	return enum.EventTypeUnKnown, nil, errors.New(fmt.Sprintf("not support event type[%s] [body%s]", encryptResponse.NotifyType, string(bodyBytes)))
}

func (c *TransferCallback) decryptResponse(req *http.Request) (*model.CommonEncryptCallbackResponse, error) {
	bodyBytes, _ := io.ReadAll(req.Body)
	encryptResponse := &model.CommonEncryptCallbackResponse{}
	if len(bodyBytes) > 0 {
		err := xml.Unmarshal(bodyBytes, encryptResponse)
		if err != nil {
			return nil, err
		}
	}
	encryptResponse.OriginBody = string(bodyBytes)
	messageGbk, err := c.Sign.DecryptByKey(encryptResponse.Message, []byte(c.C.RsaPrivate))
	if err != nil {
		return nil, err
	}
	message, _ := util.GBKToUTF8Byte(messageGbk)
	encryptResponse.Message = message
	return encryptResponse, nil
}

func (c *TransferCallback) AccountInResultCallback(req *http.Request) (*model.UserCreateCallbackRequest, error) {
	resp, err := c.decryptResponse(req)
	if err != nil {
		return nil, err
	}
	var user = &model.UserCreateCallbackRequest{}
	err = xml.Unmarshal([]byte(resp.Message), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// AccountInSettleResultCallback
// 账户结算结果通知
func (c *TransferCallback) AccountInSettleResultCallback(req *http.Request) (*dto.CallbackBatchTransferRequest, error) {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	utf8Body, _ := util.GBKToUTF8Byte(bodyBytes)
	var user = &model.TransferCallbackRequest{}
	err = xml.Unmarshal([]byte(utf8Body), user)
	if err != nil {
		return nil, err
	}
	existsProcessing := false
	status, remark := util.GetTransferStatus(user.Data.Status)
	if status == transfer.DetailStatus_DetailStatus_PROCESSING {
		existsProcessing = true
	}
	resp := &dto.CallbackBatchTransferRequest{
		TransferNO:      user.Data.OriTraceNo,
		ThirdTransferNO: user.Data.BatchNo,
		Details:         make([]dto.BatchTransferQueryDetail, 0, 1),
		TransferStatus:  transfer.TransferStatus_TransferStatus_PROCESSING,
		OriginResponse:  string(bodyBytes),
	}
	resp.Details = append(resp.Details, dto.BatchTransferQueryDetail{
		UserID:         user.Data.AccountIn,
		UserName:       "",
		TransferStatus: status,
		ErrorRemark:    user.Data.Desc,
		Amount: dto.Amount{
			Currency: "CNY",
			Total:    user.Data.Amt,
		},
		Remark:  remark,
		TradeNO: user.Data.OriTraceNo,
	})
	// 批次处理完成
	if !existsProcessing {
		resp.TransferStatus = transfer.TransferStatus_TransferStatus_FINISHED
	}
	return resp, nil
}
