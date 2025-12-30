package callback

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client/mchtransfer"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/model"
	transfer2 "github.com/lihongsheng/payment-sdk/adapter/wxpay/transfer"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/enum/refund"
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	PAYMENT_EVENT_SUCCESS = "TRANSACTION.SUCCESS"
)
const (
	REFUND_EVENT_SUCCESS  = "REFUND.SUCCESS"
	REFUND_EVENT_ABNORMAL = "REFUND.ABNORMAL"
	REFUND_EVENT_CLOSED   = "REFUND.CLOSED"
)

type Callback struct {
	Id           string           `json:"id"`
	CreateTime   time.Time        `json:"create_time"`
	ResourceType string           `json:"resource_type"`
	EventType    string           `json:"event_type"`
	Summary      string           `json:"summary"`
	Resource     CallbackResource `json:"resource"`
}

type CallbackResource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

func (c Callback) Validate() error {
	//switch c.EventType {
	//case PAYMENT_EVENT_SUCCESS, e:
	//
	//}
	return nil
}

const (
	// Wechatpay-Serial
	//
	//验签的微信支付平台证书序列号/微信支付公钥ID
	//
	//Wechatpay-Signature
	//
	//验签的签名值
	//
	//Wechatpay-Timestamp
	//
	//验签的时间戳
	//
	//Wechatpay-Nonce
	//
	//验签的随机字符串

	Wechatpay_Serial    = "Wechatpay-Serial"
	Wechatpay_Signature = "Wechatpay-Signature"
	Wechatpay_Timestamp = "Wechatpay-Timestamp"
	Wechatpay_Nonce     = "Wechatpay-Nonce"
)

//
//func CallbackPaymentParse(ctx context.Context, conf config.Config, req *http.Request) (*dto.CallbackPayDetail, error) {
//	pubKeyID := conf.Cert.PublicKeyID
//	mchAPIv3Key := conf.APIKey
//	publicKey, err := utils.LoadPublicKey(conf.Cert.PublicKey)
//	if err != nil {
//		return nil, errors.New("wxpay load merchant PublicKey key errors")
//	}
//	no, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAPubkeyVerifier(pubKeyID, *publicKey))
//	if err != nil {
//		return nil, errors2.ErrorSystemError("wxpay new RSA NotifyHandler errors").WithCause(err)
//	}
//	var resp = &payments.Transaction{}
//	_, err = no.ParseNotifyRequest(ctx, req, resp)
//	if err != nil {
//		_, proErr := ProcessBody(mchAPIv3Key, req, resp)
//		if proErr != nil {
//			return nil, err
//		}
//		//return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
//	}
//	if resp.TransactionId == nil || (resp.TransactionId != nil && *resp.TransactionId == "") {
//		return nil, errors2.ErrorSystemError("wxpay parse notify not find TransactionId").WithCause(err)
//	}
//	status := until.PaymentStatus[until.StringPoint(resp.TradeState)]
//	if status == enum.Status_Status_UNKNOWN {
//		logc.Error(ctx, "wxPayErrStatus", logc.Field("resp", resp))
//		return nil, until.ErrorHandler(ctx, nil, err, "status is unknown")
//	}
//	var successTime time.Time
//	if resp.SuccessTime == nil && *resp.SuccessTime != "" {
//		successTime, _ = time.Parse(time.RFC3339, *resp.SuccessTime)
//	}
//	originBy, _ := json.Marshal(resp)
//	return &dto.CallbackPayDetail{
//		OrderNo: until.StringPoint(resp.OutTradeNo),
//		TradeNo: until.StringPoint(resp.TransactionId),
//		PayAmount: dto.Amount{
//			Currency: until.StringPoint(resp.Amount.Currency),
//			Total:    until.Int64Point(resp.Amount.Total),
//		},
//		Status:         status,
//		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
//		SuccessTime:    successTime.Unix(),
//		OriginResponse: string(originBy),
//	}, nil
//}

//
//func CallbackRefundParse2(ctx context.Context, conf config.Config, req *http.Request) (*dto.RefundDetail, error) {
//  pubKeyID := conf.Cert.PublicKeyID
//  mchAPIv3Key := conf.APIKey
//  publicKey, err := utils.LoadPublicKey(conf.Cert.PublicKey)
//  if err != nil {
//    return nil, errors.New("wxpay load merchant PublicKey key errors")
//  }
//  verifiers.NewSHA256WithRSAPubkeyVerifier(pubKeyID, *publicKey)
//  body, err := ioutil.ReadAll(req.Body)
//  if err != nil {
//    return nil, errors2.ErrorSystemError("wxpay read request body errors").WithCause(err)
//  }
//
//  c, err := aes.NewCipher([]byte(mchAPIv3Key))
//  if err != nil {
//    return nil, err
//  }
//  aesgcm, err := cipher.NewGCM(c)
//  if err != nil {
//    return nil, err
//  }
//
//  v := CipherSuite{
//    signatureType: rsaSignatureType,
//    validator:     *validators.NewWechatPayNotifyValidator(verifier),
//    aeadAlgorithm: aeadAesGcmAlgorithm,
//    aead:          aesgcm,
//  }
//
//
//
//  var resp = &refunddomestic.Refund{}
//  _, err = no.ParseNotifyRequest(ctx, req, resp)
//  if err != nil {
//    return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
//  }
//  originBy, _ := json.Marshal(resp)
//  re := &dto.RefundDetail{
//    RefundNo:      until.StringPoint(resp.OutRefundNo),
//    OrderNo:       until.StringPoint(resp.OutTradeNo),
//    TradeRefundNo: until.StringPoint(resp.RefundId),
//    TradeNo:       until.StringPoint(resp.TransactionId),
//    Amount: dto.Amount{
//      Currency: until.StringPoint(resp.Amount.Currency),
//      Total:    until.Int64Point(resp.Amount.Refund),
//    },
//    Channel:             refund.RefundChannel(refund.RefundChannel_value[string(*resp.Channel)]),
//    UserReceivedAccount: until.StringPoint(resp.UserReceivedAccount),
//    CreateTime:          *resp.CreateTime,
//    FundsAccount:        string(*resp.FundsAccount),
//    Status:              until.GetRefundStatus(*resp.Status),
//    OriginResponse:      string(originBy),
//  }
//  if resp.SuccessTime != nil && !resp.SuccessTime.IsZero() {
//    re.SuccessTime = *resp.SuccessTime
//  }
//  return re, nil
//}

func CallbackRefundParse(ctx context.Context, conf config.Config, req *http.Request) (*dto.CallbackRefundDetail, error) {
	pubKeyID := conf.Cert.PublicKeyID
	mchAPIv3Key := conf.APIKey
	publicKey, err := utils.LoadPublicKey(conf.Cert.PublicKey)
	if err != nil {
		return nil, errors.New("wxpay load merchant PublicKey key errors")
	}
	no, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAPubkeyVerifier(pubKeyID, *publicKey))
	if err != nil {
		return nil, errors2.ErrorSystemError("wxpay new RSA NotifyHandler errors").WithCause(err)
	}
	var resp = &CallbackRefund{}
	_, err = no.ParseNotifyRequest(ctx, req, resp)
	if err != nil {
		_, proErr := ProcessBody(mchAPIv3Key, req, resp)
		if proErr != nil {
			return nil, err
		}
		// return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	if resp.RefundId == "" {
		return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	originBy, _ := json.Marshal(resp)
	re := &dto.CallbackRefundDetail{
		RefundNo:      resp.OutRefundNo,
		OrderNo:       resp.OutTradeNo,
		TradeRefundNo: resp.RefundId,
		TradeNo:       resp.TransactionId,
		Amount: dto.Amount{
			Currency: "CNY",
			Total:    int64(resp.Amount.Refund),
		},
		UserReceivedAccount: resp.UserReceivedAccount,
		OriginResponse:      string(originBy),
	}
	// 过期数据兼容
	// 		Channel:             refund.RefundChannel(refund.RefundChannel_value[string(*resp.Channel)]),
	// CreateTime:          *resp.CreateTime,
	// 		FundsAccount:        string(*resp.FundsAccount),

	if resp.RefundStatus != "" {
		re.Status = until.GetRefundStatus2(resp.RefundStatus)
	} else {
		re.Status = refund.Status_Status_UNKNOWN
	}
	if !resp.SuccessTime.IsZero() {
		re.SuccessTime = resp.SuccessTime
	}
	return re, nil
}

type CallbackRefund struct {
	Mchid               string       `json:"mchid"`
	OutTradeNo          string       `json:"out_trade_no"`
	TransactionId       string       `json:"transaction_id"`
	OutRefundNo         string       `json:"out_refund_no"`
	RefundId            string       `json:"refund_id"`
	RefundStatus        string       `json:"refund_status"`
	SuccessTime         time.Time    `json:"success_time"`
	Amount              RefundAmount `json:"amount"`
	UserReceivedAccount string       `json:"user_received_account"`
}

type RefundAmount struct {
	Total       int `json:"total"`
	Refund      int `json:"refund"`
	PayerTotal  int `json:"payer_total"`
	PayerRefund int `json:"payer_refund"`
}

func CallbackUnitTransferParse(ctx context.Context, conf config.Config, req *http.Request) (*dto.UintTransferDetailResponse, error) {
	pubKeyID := conf.Cert.PublicKeyID
	mchAPIv3Key := conf.APIKey
	publicKey, err := utils.LoadPublicKey(conf.Cert.PublicKey)
	if err != nil {
		return nil, errors.New("wxpay load merchant PublicKey key errors")
	}
	no, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAPubkeyVerifier(pubKeyID, *publicKey))
	if err != nil {
		return nil, errors2.ErrorSystemError("wxpay new RSA NotifyHandler errors").WithCause(err)
	}
	var resp = &model.UnitTransferCallback{}
	_, err = no.ParseNotifyRequest(ctx, req, resp)
	if err != nil {
		_, proErr := ProcessBody(mchAPIv3Key, req, resp)
		if proErr != nil {
			return nil, err
		}
		// return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	if resp.OutBillNo == "" {
		return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	originBy, _ := json.Marshal(resp)
	re := &dto.UintTransferDetailResponse{
		TransferNo:      resp.OutBillNo,
		ThirdTransferNo: resp.TransferBillNo,
		Status:          until.GetTransferStatus(resp.State),
		TransferAmount: dto.Amount{
			Total:    int64(resp.TransferAmount),
			Currency: "CNY",
		},
		User: dto.User{
			OpenID: resp.Openid,
		},
		FailReason:     resp.FailReason,
		OriginResponse: string(originBy),
	}
	return re, nil
}

// CallbackComplaintParse 投诉回调
func CallbackComplaintParse(ctx context.Context, conf config.Config, req *http.Request) (*model.ComplaintCallbackRequest, error) {
	pubKeyID := conf.Cert.PublicKeyID
	mchAPIv3Key := conf.APIKey
	publicKey, err := utils.LoadPublicKey(conf.Cert.PublicKey)
	if err != nil {
		return nil, errors.New("wxpay load merchant PublicKey key errors")
	}
	no, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAPubkeyVerifier(pubKeyID, *publicKey))
	if err != nil {
		return nil, errors2.ErrorSystemError("wxpay new RSA NotifyHandler errors").WithCause(err)
	}
	var resp = &model.ComplaintCallbackRequest{}
	_, err = no.ParseNotifyRequest(ctx, req, resp)
	if err != nil {
		_, proErr := ProcessBody(mchAPIv3Key, req, resp)
		if proErr != nil {
			return nil, err
		}
		// return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	if resp.ComplaintId == "" {
		return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	return resp, nil
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

func ProcessBody(apiV3Key string, request *http.Request, content interface{}) (*notify.Request, error) {
	body, err := getRequestBody(request)
	if err != nil {
		return nil, err
	}
	ret := new(notify.Request)
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, fmt.Errorf("parse request body error: %v", err)
	}
	aead, err := getAes(apiV3Key)
	if err != nil {
		return nil, err
	}
	plaintext, err := doAEADOpen(
		aead,
		ret.Resource.Nonce,
		ret.Resource.Ciphertext,
		ret.Resource.AssociatedData,
	)
	if err != nil {
		return ret, fmt.Errorf("%s decrypt error: %v", ret.Resource.Algorithm, err)
	}
	ret.Resource.Plaintext = plaintext
	if err = json.Unmarshal([]byte(plaintext), &content); err != nil {
		return ret, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}
	return ret, nil
}

func doAEADOpen(c cipher.AEAD, nonce, ciphertext, additionalData string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := c.Open(
		nil,
		[]byte(nonce),
		data,
		[]byte(additionalData),
	)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func getAes(apiV3Key string) (cipher.AEAD, error) {
	c, err := aes.NewCipher([]byte(apiV3Key))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	return aesgcm, nil
}

func CallbackBatchTransferParse(ctx context.Context, conf config.Config, req *http.Request) (*dto.CallbackBatchTransferRequest, error) {
	pubKeyID := conf.Cert.PublicKeyID
	mchAPIv3Key := conf.APIKey
	publicKey, err := utils.LoadPublicKey(conf.Cert.PublicKey)
	if err != nil {
		return nil, errors.New("wxpay load merchant PublicKey key errors")
	}
	no, err := notify.NewRSANotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAPubkeyVerifier(pubKeyID, *publicKey))
	if err != nil {
		return nil, errors2.ErrorSystemError("wxpay new RSA NotifyHandler errors").WithCause(err)
	}
	var resp = &mchtransfer.CallBackRequest{}
	_, err = no.ParseNotifyRequest(ctx, req, resp)
	if err != nil {
		_, proErr := ProcessBody(mchAPIv3Key, req, resp)
		if proErr != nil {
			return nil, err
		}
		// return nil, errors2.ErrorSystemError("wxpay parse notify request errors").WithCause(err)
	}
	if resp.BatchId == "" {
		return nil, errors2.ErrorParamError("wechat is not return batchID")
	}
	result := &dto.CallbackBatchTransferRequest{
		TransferNO:      resp.OutBatchNo,
		ThirdTransferNO: resp.BatchId,
		TransferStatus:  resp.GetTransferStatus(),
		Details:         make([]dto.BatchTransferQueryDetail, 0),
	}
	// 请求获取明细列表
	if result.TransferStatus == transfer.TransferStatus_TransferStatus_FINISHED {
		transfer, err := transfer2.NewTransfer(conf)
		if err != nil {
			return nil, err
		}
		queryReq := dto.BatchTransferQueryRequest{
			ThirdTransferNO: resp.BatchId,
			PageNum:         1,
			PageSize:        int64(resp.TotalNum),
		}
		queryResp, err := transfer.Query(ctx, queryReq)
		if err != nil {
			return nil, err
		}
		fmt.Println("***********************", err)
		fmt.Println(queryResp)
		result.Details = queryResp.Details
	}
	return result, nil
}
