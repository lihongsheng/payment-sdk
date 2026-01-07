package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/config"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/enum"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/model"
	"github.com/lihongsheng/payment-sdk/config/proxy"
	"net/url"
	"strings"
)

type Client struct {
	C      config.Config
	Client *resty.Client
	Sign   *Sign
}

func NewClient(conf config.Config, proxy *proxy.Proxy) (*Client, error) {
	conf.ApiHost = strings.TrimRight(conf.ApiHost, "/")
	if conf.ApiHost == "" {
		conf.ApiHost = enum.ApiHost
	}
	if conf.TermNO == "" {
		return nil, errors.New("拉卡拉支付必须配置终端号")
	}
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

func (c *Client) DoPost(ctx context.Context, body interface{}, path string, header map[string]string) (*resty.Response, error) {
	req := c.Client.R()
	if header != nil {
		req.SetHeaders(header)
	}
	if req.Header.Get("Content-Type") == "" {
		req.SetHeader("Content-Type", "application/json")
	}
	if req.Header.Get("Accept") == "" {
		req.SetHeader("Accept", "application/json")
	}

	reqParam := model.BuildCommonReq(body)
	sign, err := c.Sign.Gen(reqParam)
	if err != nil {
		return nil, err
	}
	req.SetHeader("Authorization", sign)
	return req.SetContext(ctx).SetBody(reqParam).Post(path)
}

func (c *Client) DoPostV2(ctx context.Context, body interface{}, path string, header map[string]string) (*resty.Response, error) {
	req := c.Client.R()
	if header != nil {
		req.SetHeaders(header)
	}
	if req.Header.Get("Content-Type") == "" {
		req.SetHeader("Content-Type", "application/json")
	}
	if req.Header.Get("Accept") == "" {
		req.SetHeader("Accept", "application/json")
	}
	return req.SetContext(ctx).SetBody(body).Post(path)
}
