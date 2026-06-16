package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/config/proxy"
	"github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/log"
	"github.com/lihongsheng/payment-sdk/tools"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HostName = "https://openapi.alipay.com/gateway.do"
)

var skipQueryParams = map[string]struct{}{
	enum.COMMON_PARAM_NOTIFY_URL_NAME:     {},
	enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME: {},
}

type Client struct {
	C      config.Config
	Client *resty.Client
	Sign   *Sign
}

func NewClient(conf config.Config, proxy *proxy.Proxy) (*Client, error) {
	cf, err := initConfig(conf)
	if err != nil {
		return nil, err
	}
	conf = cf
	client, err := providerClient(proxy)
	if err != nil {
		return nil, err
	}
	sign, err := NewSign(conf)
	if err != nil {
		return nil, err
	}
	return &Client{
		C:      conf,
		Client: client,
		Sign:   sign,
	}, nil
}

func initConfig(conf config.Config) (config.Config, error) {
	var err error
	if conf.Cert.RsaRootCrt != "" {
		conf.Cert.RsaRootCertSN, err = tools.GetCertInfo(conf.Cert.RsaRootCrt)
		if err != nil {
			return config.Config{}, err
		}
		conf.Cert.RsaPublic, err = tools.ParseCerToPublicKeyPEM(conf.Cert.RsaRootCrt)
		if err != nil {
			return config.Config{}, err
		}
	}
	if conf.Cert.RsaAppCrt != "" {
		conf.Cert.RsaAppCertSN, err = tools.GetCertInfo(conf.Cert.RsaAppCrt)
		if err != nil {
			return config.Config{}, err
		}
	}
	return conf, nil
}

func providerClient(proxy *proxy.Proxy) (*resty.Client, error) {
	client := resty.New()
	if proxy != nil && proxy.Host != "" {
		u, err := url.Parse(fmt.Sprintf("http://%s:%d", proxy.Host, proxy.Port))
		if err != nil {
			return nil, err
		}
		if proxy.UserName != "" && proxy.Password != "" {
			u.User = url.UserPassword(proxy.UserName, proxy.Password)
		}
		if proxy.UserName != "" {
			u.User = url.User(proxy.UserName)
		}
		client.SetProxy(u.String())
	}
	return client, nil
}

func (c *Client) GetCommonRequestParams() map[string]string {
	return map[string]string{
		enum.COMMON_PARAM_APP_ID_NAME:    c.C.Merchant.AppID,
		enum.COMMON_PARAM_FORMAT_NAME:    "json",
		enum.COMMON_PARAM_CHARSET_NAME:   "UTF-8",
		enum.COMMON_PARAM_SIGN_TYPE_NAME: "RSA2",
		enum.COMMON_PARAM_VERSION_NAME:   "1.0",
		enum.COMMON_PARAM_TIMESTAMP_NAME: time.Now().Format(time.DateTime),
	}
}

func (c *Client) DoPost(ctx context.Context, commonParams map[string]string, body any, header map[string]string) (*resty.Response, error) {
	method := commonParams[enum.COMMON_PARAM_METHOD_NAME]
	log.Info(ctx, "alipay request start",
		log.F(log.FieldKeyChannel, "alipay"),
		log.F(log.FieldKeyMethod, method),
		log.F(log.FieldKeyRequest, body),
	)

	reqUrl, bodyParams, err := c.PageExecute(commonParams, body)
	if err != nil {
		log.Error(ctx, "alipay request sign failed",
			log.F(log.FieldKeyChannel, "alipay"),
			log.F(log.FieldKeyMethod, method),
			log.F(log.FieldKeyError, err.Error()),
		)
		return nil, err
	}
	req := c.Client.R()
	if header != nil {
		req.SetHeaders(header)
	}
	if req.Header.Get("Content-Type") == "" {
		req.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	}
	if req.Header.Get("Accept") == "" {
		req.SetHeader("Accept", "application/json")
	}
	req.SetQueryParamsFromValues(reqUrl.Query())
	req.SetFormData(bodyParams)

	resp, err := req.SetContext(ctx).Post(HostName)
	if err != nil {
		log.Error(ctx, "alipay request failed",
			log.F(log.FieldKeyChannel, "alipay"),
			log.F(log.FieldKeyMethod, method),
			log.F(log.FieldKeyError, err.Error()),
		)
		return nil, err
	}

	log.Info(ctx, "alipay request end",
		log.F(log.FieldKeyChannel, "alipay"),
		log.F(log.FieldKeyMethod, method),
		log.F(log.FieldKeyResponse, string(resp.Body())),
	)
	return resp, nil
}

func (c *Client) PageExecute(commonParams map[string]string, body any) (u *url.URL, bodyParams map[string]string, err error) {
	req, err := url.Parse(HostName)
	if err != nil {
		return nil, nil, err
	}
	switch body.(type) {
	case map[string]string:
		bodyParams = body.(map[string]string)
	default:
		by, err := json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
		byStr := string(by)
		bodyParams[enum.COMMON_PARAM_Biz_NAME] = byStr
	}
	sign, err := c.Sign.Sign(commonParams, bodyParams)
	if err != nil {
		return nil, nil, err
	}
	queryParams := url.Values{}
	queryParams.Add(enum.COMMON_PARAM_SING_NAME, sign)
	for k, v := range commonParams {
		if _, ok := skipQueryParams[k]; ok {
			continue
		}
		queryParams.Add(k, v)
	}
	req.RawQuery = queryParams.Encode()
	if commonParams[enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME] != "" {
		bodyParams[enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME] = commonParams[enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME]
	}
	if commonParams[enum.COMMON_PARAM_NOTIFY_URL_NAME] != "" {
		bodyParams[enum.COMMON_PARAM_NOTIFY_URL_NAME] = commonParams[enum.COMMON_PARAM_NOTIFY_URL_NAME]
	}
	if commonParams[enum.COMMON_PARAM_RETURN_URL_NAME] != "" {
		bodyParams[enum.COMMON_PARAM_RETURN_URL_NAME] = commonParams[enum.COMMON_PARAM_RETURN_URL_NAME]
	}
	return req, bodyParams, nil
}

// BuildRequestForm 生成支付宝自动提交表单（和你PHP逻辑完全一致）
func (c *Client) BuildRequestForm(paraMap map[string]string) string {
	var sHtml strings.Builder

	// 1. 拼接 form 开头
	sHtml.WriteString(fmt.Sprintf(
		"<form id='alipaysubmit' name='alipaysubmit' action='%s?charset=%s' method='POST'>",
		HostName,
		"utf-8",
	))

	// 2. 遍历参数，生成 hidden input
	for key, val := range paraMap {
		// 跳过空值
		if c.checkEmpty(val) {
			continue
		}

		// 转义单引号 ' → &apos;（和PHP逻辑一样）
		val = strings.ReplaceAll(val, "'", "&apos;")

		// HTML 安全转义（Go推荐，防止XSS，不影响支付宝）
		keyEscaped := html.EscapeString(key)
		valEscaped := html.EscapeString(val)

		// 拼接 input
		sHtml.WriteString(fmt.Sprintf(
			"<input type='hidden' name='%s' value='%s'/>",
			keyEscaped,
			valEscaped,
		))
	}

	// 3. 提交按钮（隐藏）
	sHtml.WriteString("<input type='submit' value='ok' style='display:none;'></form>")

	// 4. JS 自动提交
	sHtml.WriteString("<script>document.forms['alipaysubmit'].submit();</script>")

	return sHtml.String()
}

// checkEmpty 判断是否为空（对应PHP checkEmpty）
func (c *Client) checkEmpty(val string) bool {
	return val == "" || val == " " || len(val) == 0
}

func (c *Client) GetPageExecute(commonParams map[string]string, body any) (u *url.URL, bodyParams map[string]string, err error) {
	req, err := url.Parse(HostName)
	if err != nil {
		return nil, nil, err
	}
	switch body.(type) {
	case map[string]string:
		bodyParams = body.(map[string]string)
	default:
		by, err := json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
		byStr := string(by)
		bodyParams[enum.COMMON_PARAM_Biz_NAME] = byStr
	}
	sign, err := c.Sign.Sign(commonParams, bodyParams)
	if err != nil {
		return nil, nil, err
	}
	queryParams := url.Values{}
	queryParams.Add(enum.COMMON_PARAM_SING_NAME, sign)
	for k, v := range commonParams {
		if _, ok := skipQueryParams[k]; ok {
			continue
		}
		queryParams.Add(k, v)
	}
	req.RawQuery = queryParams.Encode()
	if commonParams[enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME] != "" {
		bodyParams[enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME] = commonParams[enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME]
	}
	if commonParams[enum.COMMON_PARAM_NOTIFY_URL_NAME] != "" {
		bodyParams[enum.COMMON_PARAM_NOTIFY_URL_NAME] = commonParams[enum.COMMON_PARAM_NOTIFY_URL_NAME]
	}
	if commonParams[enum.COMMON_PARAM_RETURN_URL_NAME] != "" {
		bodyParams[enum.COMMON_PARAM_RETURN_URL_NAME] = commonParams[enum.COMMON_PARAM_RETURN_URL_NAME]
	}
	return req, bodyParams, nil
}

func (c *Client) GetResponseSignContent(body string, method string) (string, error) {
	responseNode := strings.Replace(method, ".", "_", -1) + enum.RESPONSE_LAST_PREFIX
	index := strings.Index(body, responseNode) + len(responseNode) + 2
	errorIndex := strings.Index(body, enum.RESPONSE_ERROR) + len(enum.RESPONSE_ERROR) + 2
	lastIndex := strings.Index(body, `"`+enum.COMMON_PARAM_SING_NAME+`"`)
	if lastIndex < 0 {
		return "", errors.ErrorParamError("获取响应签名内容失败,未见签名字段")
	}
	if !(lastIndex > index || lastIndex > errorIndex) {
		return "", errors.ErrorParamError("获取响应签名内容失败,未见响应内容")
	}
	if index > 0 {
		return body[index : lastIndex-1], nil
	} else if errorIndex > 0 {
		return body[errorIndex : lastIndex-1], nil
	}
	return "", errors.ErrorParamError("获取响应签名内容失败")
}

// VerifyCallback 统一验签回调请求
// 返回: 解析后的参数 map, 原始 body 字节, 错误
func (c *Client) VerifyCallback(ctx context.Context, req *http.Request) (url.Values, []byte, error) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error(ctx, "alipay callback read body failed",
			log.F(log.FieldKeyChannel, "alipay"),
			log.F(log.FieldKeyError, err.Error()),
		)
		return nil, nil, errors.ErrorParamError("read request body err: %v", err)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		log.Error(ctx, "alipay callback parse body failed",
			log.F(log.FieldKeyChannel, "alipay"),
			log.F(log.FieldKeyError, err.Error()),
		)
		return nil, nil, errors.ErrorParamError("parse request body err: %v", err)
	}

	signStr, signValue, err := c.Sign.GenerateSignString(values)
	if err != nil {
		log.Error(ctx, "alipay callback generate sign string failed",
			log.F(log.FieldKeyChannel, "alipay"),
			log.F(log.FieldKeyError, err.Error()),
		)
		return nil, nil, err
	}

	verify, err := c.Sign.RsaVerify(signStr, signValue)
	if err != nil || !verify {
		log.Error(ctx, "alipay callback verify sign failed",
			log.F(log.FieldKeyChannel, "alipay"),
			log.F("sign_str", signStr),
			log.F("sign_value", signValue),
			log.F(log.FieldKeyError, err),
		)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, errors.ErrorSignError("签名验证失败："+string(bodyBytes), nil)
	}

	log.Info(ctx, "alipay callback verify success",
		log.F(log.FieldKeyChannel, "alipay"),
		log.F(log.FieldKeyOrderNo, values.Get("out_trade_no")),
		log.F(log.FieldKeyTradeNo, values.Get("trade_no")),
	)

	return values, bodyBytes, nil
}
