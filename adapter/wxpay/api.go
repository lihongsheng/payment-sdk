package wxpay

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"net/http"
	"net/url"

	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type Api struct {
	C          config.Config
	PrivateKey *rsa.PrivateKey
	Verifier   auth.Verifier
	Client     *core.Client
}

func (a *Api) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
	return nil, errors2.ErrorNoSupport("not support Complete")
}

func InitClient(c config.Config) (*Api, error) {
	w := &Api{C: c}
	// 使用 utils 提供的函数从私钥字符串中加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(c.Cert.Private)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("wxpay load merchant private key errors;%s", err.Error()))
	}
	w.PrivateKey = mchPrivateKey
	ctx := context.Background()
	opts := []core.ClientOption{}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	if c.Cert.Public != "" && c.Cert.PublicNumber != "" {
		publicKey, err := utils.LoadPublicKey(c.Cert.Public)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("wxpay load merchant Public key errors;%s", err.Error()))
		}
		w.Verifier = verifiers.NewSHA256WithRSAPubkeyVerifier(c.Cert.PublicNumber, *publicKey)
		opts = append(opts, option.WithWechatPayPublicKeyAuthCipher(c.MchID, c.Cert.PrivateNumber, mchPrivateKey, c.Cert.PublicNumber, publicKey))
	} else {
		opts = append(opts, option.WithWechatPayAutoAuthCipher(c.MchID, c.Cert.PrivateNumber, mchPrivateKey, c.APISecret))
		visitor, err := autoVisitor(c, w.PrivateKey)
		if err != nil {
			return nil, err
		}
		w.Verifier = visitor
	}
	if c.Proxy.Host != "" {
		opts = append(opts, proxy(c))
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new wechat pay client err:%s", err.Error()))
	}
	w.Client = client
	return w, nil
}

func autoVisitor(c config.Config, privateKey *rsa.PrivateKey) (auth.Verifier, error) {
	ctx := context.Background()
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	mgr := downloader.MgrInstance()
	if !mgr.HasDownloader(context.Background(), c.MchID) {
		err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, privateKey, c.Cert.PrivateNumber, c.MchID, c.APISecret)
		if err != nil {
			return nil, err
		}
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(c.MchID)
	return verifiers.NewSHA256WithRSAVerifier(certificateVisitor), nil
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
