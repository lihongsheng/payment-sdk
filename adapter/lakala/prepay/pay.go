package prepay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	enum1 "github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/tools"
	"net/http"

	"github.com/lihongsheng/payment-sdk/adapter/lakala/client"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/config"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/lakala/enum"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/model"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"net/url"
	"strconv"
)

const (
	PayMethod   = "/api/v3/labs/trans/preorder"
	QueryMethod = "/api/v3/labs/query/tradequery"
)

type Pay struct {
	*client.Client
	paymentProduct enum.PaymentProduct
	payment        enum.Payment
}

func NewPay(conf config.Config, product enum.PaymentProduct, payment enum.Payment) (iface.Pay, error) {
	api, err := client.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &Pay{
		Client:         api,
		paymentProduct: product,
		payment:        payment,
	}, nil
}

func (p *Pay) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := p.buildPayParams(req)
	result, err := p.Client.DoPost(ctx, reqParam, p.C.ApiHost+PayMethod, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.PaymentRespBody{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "lakala-Pay-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	var paymentResp = resp.RespData
	if paymentResp == nil {
		return nil, fmt.Errorf("解析PaymentResponse失败: %w, 原始数据: %s", err, string(by))
	}
	if paymentResp.AccRespFields == nil {
		return nil, errors.New(fmt.Sprintf("无法获取支付结果%s", string(by)))
	}
	wxResp := paymentResp.AccRespFields
	re := &dto.PayResponse{
		OrderNo: req.Order.OrderNo,
		TradeNo: paymentResp.TradeNo,
		PayAmount: dto.Amount{
			Currency: req.Order.PayAmount.Currency,
			Total:    req.Order.PayAmount.Total,
		},
		Status:         enum.Status_Pending,
		PaymentProduct: p.paymentProduct.String(),
		OriginResponse: string(by),
	}
	var actionParams = map[string]string{}
	actionParams["appId"] = wxResp.AppId
	actionParams["timeStamp"] = wxResp.TimeStamp
	actionParams["nonceStr"] = wxResp.NonceStr
	actionParams["package"] = wxResp.Package
	actionParams["signType"] = wxResp.SignType
	actionParams["paySign"] = wxResp.PaySign
	params, err := url.ParseQuery(wxResp.Package)
	if err != nil {
		return nil, errors2.ErrorSystemError("parse SdkPackage error").WithCause(err)
	}
	prepayID := params[action.PrepayID]
	if len(prepayID) == 0 {
		return nil, errors2.ErrorSystemError("lakala not return prepayID error")
	}
	re.Action = dto.Action{
		Action:     action.Action_Prepay.String(),
		Parameters: actionParams,
	}
	return re, nil
}

// 目前只支持微信
func (p *Pay) buildPayParams(req *dto.PayOrder) *model.PaymentRequest {
	r := &model.PaymentRequest{
		MerchantNo:  p.C.MchID,
		TermNo:      p.C.TermNO,
		OutTradeNo:  req.Order.OrderNo,
		AccountType: enum2.PaymentMap[p.payment],
		TransType:   enum2.ProductMap[p.paymentProduct],
		TotalAmount: fmt.Sprintf("%d", req.Order.PayAmount.Total),
		NotifyUrl:   req.NotifyUrl,
		Subject:     req.Order.Subject,
		Remark:      "",
		AccBusiFields: &model.WxAccBusiFields{
			TimeoutExpress: "",
			SubAppid:       req.Payer.AppID,
			UserID:         req.Payer.OpenID,
		},
	}
	if req.SceneInfo != nil {
		r.LocationInfo = &model.LocationInfo{
			RequestIp: req.SceneInfo.ClientIp,
		}
	}
	return r
}

func (p *Pay) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	reqParam := model.PaymentQueryRequest{
		MerchantNo: p.C.MchID,
		TermNo:     p.C.TermNO,
		OutTradeNo: req.OrderNo,
		TradeNo:    req.TradeNo,
	}
	result, err := p.Client.DoPost(ctx, reqParam, p.C.ApiHost+QueryMethod, nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.PaymentQueryRespBody{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "lakala-Pay-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	var paymentResp = resp.RespData
	if paymentResp == nil {
		return nil, fmt.Errorf("解析PaymentResponse失败: %w, 原始数据: %s", err, string(by))
	}
	totalAmount := int64(0)
	ta, _ := strconv.Atoi(paymentResp.TotalAmount)
	totalAmount = int64(ta)
	return &dto.PayDetail{
		OrderNo: paymentResp.OutTradeNo,
		TradeNo: paymentResp.TradeNo,
		PayAmount: dto.Amount{
			Currency: "CNY",
			Total:    totalAmount,
		},
		Status:         enum2.GetPaymentStatus(paymentResp.TradeState),
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		//SuccessTime:    successTime.Unix(),
		OriginResponse: string(by),
	}, nil
}

func (p *Pay) Close(ctx context.Context, req dto.CloseQuery) error {
	return errors2.ErrorNoSupport("not support Close")
}

func (p *Pay) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors2.ErrorNoSupport("not support Complete")
}

func (p *Pay) Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	originBy, err := tools.GetRequestBody(req)
	if err != nil {
		return nil, err
	}
	resp := &model.PaymentCallbackRequest{}
	err = json.Unmarshal(originBy, resp)
	if err != nil {
		return nil, err
	}
	signParam, sign, err := p.Sign.ParseSign(req)
	if err != nil {
		return nil, err
	}
	signStr := p.Sign.BuildCallbackSignStr(signParam, string(originBy))
	signResp, err := p.Sign.RsaVerify(signStr, sign)
	if err != nil {
		return nil, errors.New("验证签名失败：" + err.Error())
	}
	if !signResp {
		return nil, errors.New("验证签名失败：" + string(originBy))
	}
	status := enum.Status_Status_UNKNOWN
	orderAmt, _ := strconv.Atoi(resp.TotalAmount)
	status = enum2.GetPaymentStatus(resp.TradeStatus)
	return &dto.CallbackPayDetail{
		OrderNo: resp.OutTradeNo,
		TradeNo: resp.TradeNo,
		PayAmount: dto.Amount{
			Currency: "CNY",
			Total:    int64(orderAmt),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		OriginResponse: string(originBy),
		EventAction:    enum1.Event_PAYMENT,
	}, nil
}
