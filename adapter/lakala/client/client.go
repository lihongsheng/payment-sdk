package client

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/model"
	"github.com/lihongsheng/payment-sdk/config"
	"net/url"
)

type Client struct {
	Conf   config.Config
	Client *resty.Client
	Sign   *Sign
}

func NewClient(conf config.Config) (*Client, error) {
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
