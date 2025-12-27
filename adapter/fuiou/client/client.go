package client

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	"github.com/lihongsheng/payment-sdk/config"
	"net/url"
)

var skipQueryParams = map[string]struct{}{
	"signature":         {},
	"randomStr":         {},
	"outAcntNoType":     {},
	"useESign":          {},
	"miniAppReturnPath": {},
	"XMLName":           {},
}

type Req interface {
	GenerateSign() map[string]string
	Xml() (string, error)
	Sign(sign string)
}

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
	return &Client{
		Conf:   conf,
		Client: client,
		Sign:   NewSign(conf),
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

// 非加密post
func (c *Client) PostReqFrom(ctx context.Context, path string, req Req, header map[string]string) (*resty.Response, error) {
	r := c.Client.R().SetContext(ctx)
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=GBK")
	if header != nil {
		r.SetHeaders(header)
	}
	origSign := req.GenerateSign()
	sign, err := c.Sign.Sign(origSign, skipQueryParams)
	if err != nil {
		return nil, err
	}
	req.Sign(sign)
	xml, err := req.Xml()
	if err != nil {
		return nil, err
	}
	xmlGbk, _ := util.URLEncodeGBK(xml)
	values := url.Values{}
	values.Add(enum.POST_COMMON_PARAM, xmlGbk)
	resp, err := r.SetFormDataFromValues(values).Post(c.Conf.ApiHost + path)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if len(body) > 0 {
		utf8Body, _ := util.URLDecodeGBK(string(body))
		resp.SetBody([]byte(utf8Body))
	}
	return resp, nil
}

// PostEncryptFrom
// 加密post
// 私钥签名
// 公钥加密
func (c *Client) PostEncryptFrom(ctx context.Context, path string, req Req, header map[string]string) (encryptResponse *model.EncryptResponse, err error) {
	r := c.Client.R().SetContext(ctx)
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=GBK")
	if header != nil {
		r.SetHeaders(header)
	}
	signStr := req.GenerateSign()
	sign, err := c.Sign.Sign(signStr, skipQueryParams)
	if err != nil {
		return nil, err
	}
	req.Sign(sign)
	xmlStr, err := req.Xml()
	xmlGbk, _ := util.Utf8ToGbk(xmlStr)
	encryptStr, err := c.Sign.EncryptByPublicKey([]byte(xmlGbk), []byte(c.Conf.Cert.PublicKey))
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	values.Add(enum.POST_ENCRYPT_COMMON_PARAM_MCN, c.Conf.MchID)
	values.Add(enum.POST_ENCRYPT_COMMON_PARAM_MESSAGE, encryptStr)
	resp, err := r.SetFormDataFromValues(values).Post(c.Conf.ApiHost + path)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	// fmt.Println("body", string(body))
	encryptResponse = &model.EncryptResponse{}
	if len(body) > 0 {
		err = xml.Unmarshal(body, encryptResponse)
		if err != nil {
			return nil, err
		}
	}
	encryptResponse.OriginBody = string(body)
	messageGbk, err := c.Sign.DecryptByKey(encryptResponse.Message, []byte(c.Conf.Cert.CertPrivateKey))
	if err != nil {
		return nil, err
	}
	message, _ := util.GBKToUTF8Byte(messageGbk)
	encryptResponse.Message = message
	return encryptResponse, nil
}
