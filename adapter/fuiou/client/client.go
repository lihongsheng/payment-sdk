package client

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/lakala/enum"
	"github.com/lihongsheng/payment-sdk/config/proxy"
	"github.com/lihongsheng/payment-sdk/log"
	"net/url"
	"strings"
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
	C      config.Config
	Client *resty.Client
	Sign   *Sign
}

func NewClient(conf config.Config, proxy *proxy.Proxy) (*Client, error) {
	if conf.API.ApiHost == "" {
		conf.API.ApiHost = enum2.ApiHost
	} else {
		conf.API.ApiHost = strings.TrimRight(conf.API.ApiHost, "/")
	}
	if conf.API.Version == "" {
		conf.API.Version = enum2.Version
	}
	client, err := providerClient(conf, proxy)
	if err != nil {
		return nil, err
	}
	return &Client{
		C:      conf,
		Client: client,
		Sign:   NewSign(conf),
	}, nil
}

func providerClient(c config.Config, proxy *proxy.Proxy) (*resty.Client, error) {
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

// 非加密post
func (c *Client) PostReqFrom(ctx context.Context, path string, req Req, header map[string]string) (*resty.Response, error) {
	log.Info(ctx, "fuiou request start",
		log.F(log.FieldKeyChannel, "fuiou"),
		log.F(log.FieldKeyMethod, path),
	)

	r := c.Client.R().SetContext(ctx)
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=GBK")
	if header != nil {
		r.SetHeaders(header)
	}
	origSign := req.GenerateSign()
	sign, err := c.Sign.Sign(origSign, skipQueryParams)
	if err != nil {
		log.Error(ctx, "fuiou request sign failed",
			log.F(log.FieldKeyChannel, "fuiou"),
			log.F(log.FieldKeyMethod, path),
			log.F(log.FieldKeyError, err.Error()),
		)
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
	resp, err := r.SetFormDataFromValues(values).Post(c.C.API.ApiHost + path)
	if err != nil {
		log.Error(ctx, "fuiou request failed",
			log.F(log.FieldKeyChannel, "fuiou"),
			log.F(log.FieldKeyMethod, path),
			log.F(log.FieldKeyError, err.Error()),
		)
		return nil, err
	}
	body := resp.Body()
	if len(body) > 0 {
		utf8Body, _ := util.URLDecodeGBK(string(body))
		resp.SetBody([]byte(utf8Body))
		log.Info(ctx, "fuiou request end",
			log.F(log.FieldKeyChannel, "fuiou"),
			log.F(log.FieldKeyMethod, path),
			log.F(log.FieldKeyResponse, string(utf8Body)),
		)
	}
	return resp, nil
}

// PostEncryptFrom 转账使用
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
	encryptStr, err := c.Sign.EncryptByPublicKey([]byte(xmlGbk), []byte(c.C.Cert.RsaPublic))
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	values.Add(enum.POST_ENCRYPT_COMMON_PARAM_MCN, c.C.Merchant.MchID)
	values.Add(enum.POST_ENCRYPT_COMMON_PARAM_MESSAGE, encryptStr)
	resp, err := r.SetFormDataFromValues(values).Post(c.C.API.ApiHost + path)
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
	messageGbk, err := c.Sign.DecryptByKey(encryptResponse.Message, []byte(c.C.Cert.RsaPrivate))
	if err != nil {
		return nil, err
	}
	message, _ := util.GBKToUTF8Byte(messageGbk)
	encryptResponse.Message = message
	return encryptResponse, nil
}
