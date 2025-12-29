package prepay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/lakala_jfyui"
	"github.com/lihongsheng/payment-sdk/adapter/lakala_jfyui/model"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/action"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type JFYUI struct {
	*lakala_jfyui.Api
	paymentProduct enum.PaymentProduct
	payment        enum.Payment
}

func NewJFYUI(conf config.Config, product enum.PaymentProduct, payment enum.Payment) (iface.Pay, error) {
	api, err := lakala_jfyui.NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &JFYUI{
		Api:            api,
		paymentProduct: product,
		payment:        payment,
	}, nil
}

func (p *JFYUI) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	reqParam := p.buildPayParams(req)
	result, err := p.Client.DoPost(ctx, reqParam, "https://jfyconsole.lakala.com/order/api/cashier/pay", nil)
	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.JFCreateOrderResponse{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "lakala-Pay-result", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	// 检查响应码
	if resp.Code != 200 {
		return nil, fmt.Errorf("订单创建失败: code=%d, msg=%s", resp.Code, resp.Msg)
	}

	// 检查支付URL
	payURL := resp.Data.PayURL
	if payURL == "" {
		return nil, fmt.Errorf("未返回支付地址")
	}

	// 检查URL是否包含lakala
	if !strings.Contains(payURL, "lakala") {
		return nil, fmt.Errorf("未返回支付地址")
	}

	// 从支付URL中提取payOrderNo
	payOrderNo, err := extractPayOrderNo(payURL)
	if err != nil {
		return nil, fmt.Errorf("提取订单号失败: %v", err)
	}
	re := &dto.PayResponse{
		OrderNo: req.Order.OrderNo,
		TradeNo: payOrderNo,
		PayAmount: dto.Amount{
			Currency: req.Order.PayAmount.Currency,
			Total:    req.Order.PayAmount.Total,
		},
		Status:         enum.Status_Pending,
		PaymentProduct: p.paymentProduct.String(),
		OriginResponse: string(by),
	}
	re.Action = dto.Action{
		Action:     action.Action_RedirectNOBack.String(),
		Url:        payURL,
		Parameters: nil,
	}
	return re, nil
}

// extractPayOrderNo 从支付URL中提取payOrderNo
func extractPayOrderNo(payURL string) (string, error) {
	// 解析URL
	u, err := url.Parse(payURL)
	if err != nil {
		return "", fmt.Errorf("解析URL失败: %v", err)
	}

	// 获取查询参数
	query := u.Query()

	// URL解码
	decodedQuery, err := url.QueryUnescape(query.Encode())
	if err != nil {
		return "", fmt.Errorf("URL解码失败: %v", err)
	}

	// 解析查询字符串
	params, err := url.ParseQuery(decodedQuery)
	if err != nil {
		return "", fmt.Errorf("解析查询字符串失败: %v", err)
	}

	// 获取payOrderNo
	payOrderNo := params.Get("payOrderNo")
	if payOrderNo == "" {
		return "", fmt.Errorf("未找到payOrderNo参数")
	}

	return strings.Trim(payOrderNo, "="), nil
}

func (p *JFYUI) buildPayParams(req *dto.PayOrder) *model.JFCreateOrderRequest {
	amountStr := fmt.Sprintf("%.2f", float64(req.Order.PayAmount.Total)/100)
	r := &model.JFCreateOrderRequest{
		MerchID:     p.C.AppID,
		TradeAmount: amountStr,
		Remark:      req.Order.Subject,
		OrderTemplateData: []model.JFOrderTemplateItem{
			{
				Key:                 1747822239290,
				Type:                "number",
				Index:               0,
				Label:               "支付金额",
				Value:               amountStr,
				Origin:              "number17478222392900",
				DisplayName:         "金额类型",
				FormItemFlag:        false,
				SettingsTitle:       "金额类型设置",
				MarginLeftRight:     10,
				MarginTopBottom:     5,
				CashierTemplateName: req.Order.Subject,
				State:               true,
				Options: map[string]interface{}{
					"label":      "支付金额",
					"content":    amountStr,
					"required":   true,
					"labelAlign": "",
				},
			},
		},
	}
	return r
}
func (p *JFYUI) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
	// 构建请求数据
	reqParam := model.JFQueryOrderRequest{
		ReqTime: time.Now().Format("20060102150405"),
		Version: "1.0",
		ReqData: model.JFQueryReqData{
			ChannelID:  "95",
			PayOrderNo: req.TradeNo,
			MerchantNo: p.C.AppID,
		},
	}
	// 设置请求头
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
	}
	result, err := p.Client.DoPost(ctx, reqParam, "https://payment.lakala.com/m/ccss/counter/order/query", headers)

	if err != nil && errors.Is(context.DeadlineExceeded, err) {
		return nil, errors2.ErrorTimeOut("pay timeout").WithCause(err)
	}
	if err != nil {
		return nil, errors2.ErrorSystemError("request error").WithCause(err)
	}
	by := result.Body()
	resp := &model.JFQueryOrderResponse{}
	err = json.Unmarshal(by, resp)
	reqBy, _ := json.Marshal(reqParam)
	logc.Info(ctx, "lakala-Pay-QueryResult", logc.Field("Response", string(by)), logc.Field("REq", string(reqBy)))
	// 检查响应码（兼容字符串"000000"和整数0）
	//codeStr := fmt.Sprintf("%v", resp.Code)
	if resp.Code != "000000" {
		return nil, fmt.Errorf("查单出错: code=%v, msg=%s", resp.Code, resp.Msg)
	}
	status := enum.Status_Failed
	if resp.RespData.OrderStatus == "2" {
		status = enum.Status_Success
	}
	totalAmount, _ := strconv.Atoi(resp.RespData.TotalAmount)
	return &dto.PayDetail{
		OrderNo: req.OrderNo,
		TradeNo: resp.RespData.PayOrderNo,
		PayAmount: dto.Amount{
			Currency: "CNY",
			Total:    int64(totalAmount),
		},
		Status:         status,
		PaymentProduct: enum.PaymentProduct_JSAPI.String(),
		//SuccessTime:    successTime.Unix(),
		OriginResponse: string(by),
	}, nil
}

func (p *JFYUI) Close(ctx context.Context, req dto.CloseQuery) error {
	return errors2.ErrorNoSupport("not support Close")
}

func (p *JFYUI) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors2.ErrorNoSupport("not support Complete")
}
