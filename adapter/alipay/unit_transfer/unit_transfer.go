package unit_transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/client"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/util"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/errors"
	"time"
)

type Transfer struct {
	Client *client.Client
	conf   config.Config
}

func NewTransfer(conf config.Config) (iface.UnitTransfer, error) {
	newClient, err := client.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &Transfer{
		Client: newClient,
		conf:   conf,
	}, nil
}

func (t *Transfer) Transfer(ctx context.Context, req *dto.UintTransferRequest) (*dto.UintTransferResponse, error) {
	reqParam := model.UnitTransferRequest{
		OutBizNo:       req.TransferNo,
		TransAmount:    req.TransferAmount.ToFloatString(),
		BizScene:       req.SceneId,
		ProductCode:    "TRANS_ACCOUNT_NO_PWD",
		OrderTitle:     req.Subject,
		PayeeInfo:      model.PayeeInfo{},
		Remark:         req.Remark,
		BusinessParams: "",
	}
	if req.User.UnionID != "" {
		reqParam.PayeeInfo.Identity = req.User.UnionID
		reqParam.PayeeInfo.IdentityType = "ALIPAY_USER_ID"
	} else if req.User.OpenID != "" {
		reqParam.PayeeInfo.Identity = req.User.OpenID
		reqParam.PayeeInfo.IdentityType = "ALIPAY_LOGON_ID"
	} else if req.User.Phone != "" {
		if req.User.UserName == "" {
			return nil, errors.ErrorParamError("phone, userName must not empty")
		}
		reqParam.PayeeInfo.Identity = req.User.Phone
		reqParam.PayeeInfo.IdentityType = "ALIPAY_LOGON_ID"
		reqParam.PayeeInfo.Name = req.User.UserName
	}
	if req.SceneId == "" {
		reqParam.BizScene = "DIRECT_TRANSFER"
	}
	commonParam := t.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_FUND_TRANS_UNI_TRANSFER
	commonParam["app_cert_sn"] = t.conf.Cert.CertificateSerialNumber
	commonParam["alipay_root_cert_sn"] = t.conf.Cert.PublicKeyID
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic", err)
		}
	}()
	resp, err := t.Client.DoPost(commonParam, reqParam, nil)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, err
	}
	body := resp.Body()
	fmt.Println(string(body))
	var response model.UnitTransferResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayFundTransUniTransferResponse.Code != "" && response.AlipayFundTransUniTransferResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if response.AlipayFundTransUniTransferResponse.SubCode == "" && response.AlipayFundTransUniTransferResponse.OutBizNo != "" {
		respTrue = true
	}
	if respTrue {
		tt := time.Time{}
		if response.AlipayFundTransUniTransferResponse.TransDate != "" {
			tt = util.ParseBeijingDateTime(response.AlipayFundTransUniTransferResponse.TransDate)
		}
		re := &dto.UintTransferResponse{
			Action: nil,
			//CreateTime:      tt,
			Status:          util.GetUnitTransferStatus(response.AlipayFundTransUniTransferResponse.Status),
			ThirdTransferNo: response.AlipayFundTransUniTransferResponse.OrderId,
			TransferNo:      response.AlipayFundTransUniTransferResponse.OutBizNo,
		}
		if tt.IsZero() {
			re.CreateTime = &tt
		}
		return re, nil
	}
	reBy, _ := json.Marshal(reqParam)
	return nil, errors.ErrorSystemError("not return OutBizNo;"+string(body)+";"+string(reBy), nil)
}
func (t *Transfer) Cancel(ctx context.Context, req dto.UnitTransferCancelRequest) (*dto.UintTransferCancelResponse, error) {
	return nil, errors.ErrorNoSupport("unit transfer cancel")
}
func (t *Transfer) Query(ctx context.Context, req dto.UintTransferQueryRequest) (*dto.UintTransferDetailResponse, error) {
	reqParam := model.UnitTransferQueryRequest{
		ProductCode: "",
		BizScene:    "",
		OutBizNo:    req.TransferNo,
		OrderId:     req.ThirdTransferNo,
	}
	commonParam := t.Client.GetCommonRequestParams()
	commonParam[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_FUND_TRANS_COMMON_QUERY
	resp, err := t.Client.DoPost(commonParam, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.UnitTransferQueryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.ErrorSystemError("json.Unmarshal error").WithCause(err)
	}
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubCode+":"+response.ErrorResponse.SubMsg, nil)
	}
	respTrue := false
	if response.AlipayFundTransCommonQueryResponse.Code != "" && response.AlipayFundTransCommonQueryResponse.Code == enum.RESPONSE_SUCCESS_CODE {
		respTrue = true
	}
	if response.AlipayFundTransCommonQueryResponse.SubCode == "" && response.AlipayFundTransCommonQueryResponse.OutBizNo != "" {
		respTrue = true
	}
	if !respTrue {
		return nil, errors.ErrorSystemError("not return OutBizNo;"+string(body), nil)
	}
	amount, _ := util.AmountToCents(response.AlipayFundTransCommonQueryResponse.TransAmount)
	re := &dto.UintTransferDetailResponse{
		TransferNo:      req.TransferNo,
		ThirdTransferNo: response.AlipayFundTransCommonQueryResponse.OrderId,
		TransferAmount: dto.Amount{
			Total:    int64(amount),
			Currency: "CNY",
		},
		User: dto.User{
			UnionID: response.AlipayFundTransCommonQueryResponse.ReceiverUserId,
			OpenID:  response.AlipayFundTransCommonQueryResponse.ReceiverOpenId,
		},
		Status: util.GetUnitTransferStatus(response.AlipayFundTransCommonQueryResponse.Status),
	}

	return re, nil
}
