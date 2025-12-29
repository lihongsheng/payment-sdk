package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	"github.com/lihongsheng/payment-sdk/tools"
	"time"

	"github.com/lihongsheng/payment-sdk/adapter/wxpay"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"github.com/zeromicro/go-zero/core/logc"
)

type Jsapi struct {
	*wxpay.Api
	client jsapi.JsapiApiService
}

func NewJsApi(conf config.Config) (iface.Pay, error) {
	api, err := wxpay.InitClient(conf)
	if err != nil {
		return nil, err
	}
	svc := jsapi.JsapiApiService{Client: api.Client}
	return &Jsapi{
		Api:    api,
		client: svc,
	}, nil
}

func (j *Jsapi) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	resp, result, err := j.client.Prepay(ctx, j.buildPayParams(req))
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil || resp.PrepayId == nil || *resp.PrepayId == "" {
		return nil, until.ErrorHandler(ctx, result, err, "not return PrepayId")
	}
	var actionParams = map[string]string{}
	actionParams["appId"] = j.C.AppID
	actionParams["timeStamp"] = fmt.Sprintf("%d", time.Now().Unix())
	actionParams["nonceStr"] = tools.GenerateRandomDigits(10)
	actionParams["package"] = fmt.Sprintf("prepay_id=%s", *resp.PrepayId)
	actionParams["signType"] = "RSA"
	signParams := fmt.Sprintf("%s\n%s\n%s\n%s\n", actionParams["appId"], actionParams["timeStamp"], actionParams["nonceStr"], actionParams["package"])
	//signParams := fmt.Sprintf("%s\n%s\n%s", actionParams["timeStamp"], actionParams["nonceStr"], actionParams["package"])
	sign, _ := utils.SignSHA256WithRSA(signParams, j.PrivateKey)
	actionParams["paySign"] = sign
	return &dto.PayResponse{
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		Action: dto.Action{
			Action:     action.Action_Prepay.String(),
			Parameters: actionParams,
			Url:        "",
		},
		OrderNo:   req.Order.OrderNo,
		PayAmount: dto.Amount{},
		Status:    enum.Status_Pending,
	}, nil
}

//
//// 使用RSA私钥和SHA256签名数据
//func signData(privateKey *rsa.PrivateKey, data []byte) (string, error) {
//  // 计算SHA256哈希
//  hash := sha256.Sum256(data)
//  // 使用私钥签名
//  signature, err := untils.SignSHA256WithRSA(rand.Reader, privateKey, crypto.SHA256, hash[:])
//  if err != nil {
//    return "", err
//  }
//  // 返回Base64编码的签名
//  return base64.StdEncoding.EncodeToString(signature), nil
//}

func (j *Jsapi) buildPayParams(req *dto.PayOrder) jsapi.PrepayRequest {
	var t *time.Time
	if req.TimeExpire > 0 {
		t = core.Time(time.Unix(req.TimeExpire, 0))
	}
	amount := &jsapi.Amount{
		Total: core.Int64(req.Order.PayAmount.Total),
	}
	if req.Order.PayAmount.Currency != "" {
		amount.Currency = core.String(req.Order.PayAmount.Currency)
	}
	resp := jsapi.PrepayRequest{
		Appid:       core.String(j.C.AppID),
		Mchid:       core.String(j.C.MchID),
		OutTradeNo:  core.String(req.Order.OrderNo),
		TimeExpire:  t,
		Attach:      core.String(req.PassbackParams),
		NotifyUrl:   core.String(req.NotifyUrl),
		Description: core.String(req.Order.Subject),
		Amount:      amount,
		Payer: &jsapi.Payer{
			Openid: core.String(req.Payer.OpenID),
		},
	}
	if req.SettleInfo != nil {
		resp.SettleInfo = &jsapi.SettleInfo{
			ProfitSharing: core.Bool(req.SettleInfo.ProfitSharing),
		}
	}
	if req.SceneInfo != nil {
		resp.SceneInfo = &jsapi.SceneInfo{
			PayerClientIp: core.String(req.SceneInfo.ClientIp),
			DeviceId:      core.String(req.SceneInfo.DeviceID),
		}
		if req.SceneInfo.Store.Id != "" {
			resp.SceneInfo.StoreInfo = &jsapi.StoreInfo{
				Id: core.String(req.SceneInfo.Store.Id),
			}
		}
	}
	return resp
}
func (j *Jsapi) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	var resp *payments.Transaction
	var result *core.APIResult
	var err error
	if req.OrderNo != "" {
		resp, result, err = j.client.QueryOrderByOutTradeNo(ctx, jsapi.QueryOrderByOutTradeNoRequest{OutTradeNo: core.String(req.OrderNo), Mchid: core.String(j.C.MchID)})
	} else if req.TradeNo != "" {
		resp, result, err = j.client.QueryOrderById(ctx, jsapi.QueryOrderByIdRequest{TransactionId: core.String(req.TradeNo), Mchid: core.String(j.C.MchID)})
	} else {
		return nil, errors2.ErrorParamError("order_no or trade_no is required")
	}
	if err != nil {
		return nil, until.ErrorHandler(ctx, result, err, "")
	}
	if resp == nil {
		return nil, until.ErrorHandler(ctx, result, err, "response is nil")
	}
	status := until.PaymentStatus[until.StringPoint(resp.TradeState)]
	if status == enum.Status_Status_UNKNOWN {
		logc.Error(ctx, "wxPayErrStatus", logc.Field("resp", resp))
		return nil, until.ErrorHandler(ctx, result, err, "status is unknown")
	}
	var successTime time.Time
	if resp.SuccessTime == nil && *resp.SuccessTime != "" {
		successTime, _ = time.Parse(time.RFC3339, *resp.SuccessTime)
	}
	originBy, _ := json.Marshal(resp)
	return &dto.PayDetail{
		OrderNo: until.StringPoint(resp.OutTradeNo),
		TradeNo: until.StringPoint(resp.TransactionId),
		PayAmount: dto.Amount{
			Currency: until.StringPoint(resp.Amount.Currency),
			Total:    until.Int64Point(resp.Amount.Total),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		SuccessTime:    successTime.Unix(),
		OriginResponse: string(originBy),
	}, nil
}

func (j *Jsapi) Close(ctx context.Context, req dto.CloseQuery) error {
	if req.OrderNo == "" {
		return errors2.ErrorParamError("order_no is required")
	}
	result, err := j.client.CloseOrder(ctx, jsapi.CloseOrderRequest{
		Mchid:      core.String(j.C.MchID),
		OutTradeNo: core.String(req.OrderNo),
	})
	if err != nil {
		return until.ErrorHandler(ctx, result, err, "")
	}
	if result.Response.StatusCode != 204 {
		return until.ErrorHandler(ctx, result, err, "")
	}
	return nil
}
