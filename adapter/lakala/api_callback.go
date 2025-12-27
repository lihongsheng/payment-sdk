package lakala

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/lakala/client"
	"github.com/singer-stack-lab/payment-sdk/adapter/lakala/enum"
	"github.com/singer-stack-lab/payment-sdk/adapter/lakala/model"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	"github.com/singer-stack-lab/payment-sdk/enum/payment"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type APICallback struct {
	C    config.Config
	Sign *client.Sign
}

func NewAPICallback(c config.Config) (*APICallback, error) {
	sign, err := client.NewSign(c)
	if err != nil {
		return nil, err
	}
	return &APICallback{
		C:    c,
		Sign: sign,
	}, nil
}

func (a *APICallback) CallbackPaymentParse(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
	originBy, err := getRequestBody(req)
	if err != nil {
		return nil, err
	}
	resp := &model.PaymentCallbackRequest{}
	err = json.Unmarshal(originBy, resp)
	if err != nil {
		return nil, err
	}
	signParam, sign, err := a.parseSign(req)
	if err != nil {
		return nil, err
	}
	signStr := a.Sign.BuildCallbackSignStr(signParam, string(originBy))
	signResp, err := a.Sign.RsaVerify(signStr, sign)
	if err != nil {
		return nil, errors.New("验证签名失败：" + err.Error())
	}
	if !signResp {
		return nil, errors.New("验证签名失败：" + string(originBy))
	}
	status := payment.Status_Status_UNKNOWN
	orderAmt, _ := strconv.Atoi(resp.TotalAmount)
	status = enum.GetPaymentStatus(resp.TradeStatus)
	return &dto.CallbackPayDetail{
		OrderNo: resp.OutTradeNo,
		TradeNo: resp.TradeNo,
		PayAmount: dto.Amount{
			Currency: "CNY",
			Total:    int64(orderAmt),
		},
		Status:         status,
		PaymentProduct: payment.PaymentProduct_JSAPI.String(),
		OriginResponse: string(originBy),
	}, nil
}

func (a *APICallback) parseSign(req *http.Request) (client.SignParams, string, error) {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return client.SignParams{}, "", errors.New("not find Authorization")
	}
	timestamp, nonceStr, signature, exists := ExtractLKLAuthParams(auth)
	if !exists {
		return client.SignParams{}, "", errors.New("not find Authorization")
	}
	signParams := client.SignParams{
		Timestamp: timestamp,
		NonceStr:  nonceStr,
	}
	return signParams, signature, nil
}

// ExtractLKLAuthParams 对齐PHP逻辑提取timestamp、nonce_str、signature
// 参数：authorization - 原始授权字符串（如"LKLAPI-SHA256withRSA timestamp=\"1765848504\",nonce_str=\"bGfguYYXFzbO\",signature=\"xxx\""）
// 返回：timestamp, nonceStr, signature, 提取成功标识
func ExtractLKLAuthParams(authorization string) (timestamp, nonceStr, signature string, ok bool) {
	// 完全对齐PHP的正则表达式：
	// PHP: /timestamp="(\d+)",nonce_str="(\w+)",signature="([^"]+)"/
	// Go中需转义反斜杠，规则完全一致：
	// \d+ → 匹配纯数字（timestamp）
	// \w+ → 匹配字母/数字/下划线（nonce_str）
	// [^"]+ → 匹配除双引号外的所有字符（signature，兼容特殊字符）
	pattern := `timestamp="(\d+)",nonce_str="(\w+)",signature="([^"]+)"`
	re := regexp.MustCompile(pattern)

	// 执行匹配（对应PHP的preg_match）
	matches := re.FindStringSubmatch(authorization)

	// PHP中count($matches)==0返回false，Go中len(matches)==0同理
	// 注意：FindStringSubmatch匹配成功时，matches[0]是完整匹配串，1/2/3是捕获组（对齐PHP的$matches[1/2/3]）
	if len(matches) < 4 { // 匹配成功时至少有4个元素（0:完整串,1:timestamp,2:nonce_str,3:signature）
		return "", "", "", false
	}

	// 按PHP的捕获组顺序赋值
	timestamp = matches[1]
	nonceStr = matches[2]
	signature = matches[3]

	return timestamp, nonceStr, signature, true
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
