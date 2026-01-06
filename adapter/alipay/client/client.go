package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
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
	Conf   config.Config
	Client *resty.Client
	Sign   *Sign
}

func NewClient(conf config.Config) (*Client, error) {
	cf, err := initConfig(conf)
	if err != nil {
		return nil, err
	}
	conf = cf
	client, err := providerClient(conf)
	if err != nil {
		return nil, err
	}
	sign, err := NewSign(conf)
	if err != nil {
		return nil, err
	}
	return &Client{
		Conf:   conf,
		Client: client,
		Sign:   sign,
	}, nil
}

func initConfig(conf config.Config) (config.Config, error) {
	var err error
	if conf.Cert.RootCrt != "" {
		conf.Cert.RootCertSN, err = tools.GetCertInfo(conf.Cert.RootCrt)
		if err != nil {
			return config.Config{}, err
		}
		conf.Cert.Public, err = tools.ParseCerToPublicKeyPEM(conf.Cert.RootCrt)
		if err != nil {
			return config.Config{}, err
		}
	}
	if conf.Cert.AppCrt != "" {
		conf.Cert.AppCertSN, err = tools.GetCertInfo(conf.Cert.AppCrt)
		if err != nil {
			return config.Config{}, err
		}
	}
	return conf, nil
}

func providerClient(c config.Config) (*resty.Client, error) {
	client := resty.New()
	if c.Proxy.Host != "" {
		u, err := url.Parse(fmt.Sprintf("http://%s:%d", c.Proxy.Host, c.Proxy.Port))
		if err != nil {
			return nil, err
		}
		if c.Proxy.UserName != "" && c.Proxy.Password != "" {
			u.User = url.UserPassword(c.Proxy.UserName, c.Proxy.Password)
		}
		if c.Proxy.UserName != "" {
			u.User = url.User(c.Proxy.UserName)
		}
		client.SetProxy(u.String())
	}
	return client, nil
}

func (c *Client) GetCommonRequestParams() map[string]string {
	return map[string]string{
		enum.COMMON_PARAM_APP_ID_NAME:    c.Conf.AppID,
		enum.COMMON_PARAM_FORMAT_NAME:    "json",
		enum.COMMON_PARAM_CHARSET_NAME:   "UTF-8",
		enum.COMMON_PARAM_SIGN_TYPE_NAME: "RSA2",
		enum.COMMON_PARAM_VERSION_NAME:   "1.0",
		enum.COMMON_PARAM_TIMESTAMP_NAME: time.Now().Format(time.DateTime),
	}
}

func (c *Client) DoPost(ctx context.Context, commonParams map[string]string, body any, header map[string]string) (*resty.Response, error) {
	reqUrl, bodyParams, err := c.PageExecute(commonParams, body)
	if err != nil {
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
	return req.SetContext(ctx).Post(HostName)
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
