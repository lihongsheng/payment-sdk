package wxpay

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"net/http"
	"net/url"

	"github.com/lihongsheng/payment-sdk/config"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

//var apies = sync.Map{}

type Api struct {
	C          config.Config
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Client     *core.Client
}

func (a *Api) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors2.ErrorNoSupport("not support Complete")
}

func InitClient(c config.Config) (*Api, error) {
	w := &Api{C: c}
	// 使用 utils 提供的函数从私钥字符串中加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(c.Cert.CertPrivateKey)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("wxpay load merchant private key errors;%s", err.Error()))
	}
	w.PrivateKey = mchPrivateKey
	publicKey, err := utils.LoadPublicKey(c.Cert.PublicKey)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("wxpay load merchant PublicKey key errors;%s", err.Error()))
	}
	w.PublicKey = publicKey
	ctx := context.Background()
	opts := []core.ClientOption{}
	if c.Proxy.Host != "" {
		opts = append(opts, proxy(c))
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	//opts = append(opts, option.WithWechatPayAutoAuthCipher(c.MchID, c.CertificateSerialNumber, mchPrivateKey, c.APIKey))
	opts = append(opts, option.WithWechatPayPublicKeyAuthCipher(c.MchID, c.Cert.CertificateSerialNumber, mchPrivateKey, c.Cert.PublicKeyID, publicKey))
	//opts = append(opts, option.WithMerchantCredential(c.MchID, c.CertificateSerialNumber, mchPrivateKey))
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new wechat pay client err:%s", err.Error()))
	}
	w.Client = client
	//apies.Store(key, w)
	return w, nil
}

type WithProxyOption struct {
	C config.Config
}

func (w *WithProxyOption) Apply(settings *core.DialSettings) error {
	settings.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (u *url.URL, err error) {
				u, err = url.Parse(fmt.Sprintf("http://%s:%d", w.C.Proxy.Host, w.C.Proxy.Port))
				if err != nil {
					return nil, err
				}
				if w.C.Proxy.UserName != "" && w.C.Proxy.Password != "" {
					u.User = url.UserPassword(w.C.Proxy.UserName, w.C.Proxy.Password)
				}
				if w.C.Proxy.UserName != "" {
					u.User = url.User(w.C.Proxy.UserName)
				}
				return u, nil
			},
		},
	}
	return nil
}
func proxy(c config.Config) core.ClientOption {
	return &WithProxyOption{C: c}
}
