package client

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	pro "github.com/lihongsheng/payment-sdk/config/proxy"
	"github.com/lihongsheng/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"net/http"
	"net/url"
)

type Api struct {
	C          config.Config
	PrivateKey *rsa.PrivateKey
	Verifier   auth.Verifier
	Client     *core.Client
}

func InitClient(c config.Config, proxyInfo *pro.Proxy) (*Api, error) {
	w := &Api{C: c}
	// 使用 utils 提供的函数从私钥字符串中加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(c.RsaPrivate)
	if err != nil {
		return nil, errors.ErrorParamError("wxpay load merchant private key errors;%s", err.Error())
	}
	w.PrivateKey = mchPrivateKey
	ctx := context.Background()
	opts := []core.ClientOption{}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	if c.RsaPublic != "" && c.RsaPublicNumber != "" {
		publicKey, err := utils.LoadPublicKey(c.RsaPublic)
		if err != nil {
			return nil, errors.ErrorParamError("wxpay load merchant Public key errors;%s", err.Error())
		}
		w.Verifier = verifiers.NewSHA256WithRSAPubkeyVerifier(c.RsaPublicNumber, *publicKey)
		opts = append(opts, option.WithWechatPayPublicKeyAuthCipher(c.MchID, c.RsaPrivateNumber, mchPrivateKey, c.RsaPublicNumber, publicKey))
	} else {
		opts = append(opts, option.WithWechatPayAutoAuthCipher(c.MchID, c.RsaPrivateNumber, mchPrivateKey, c.APISecret))
		visitor, err := autoVisitor(c, w.PrivateKey)
		if err != nil {
			return nil, err
		}
		w.Verifier = visitor
	}
	if proxyInfo != nil && proxyInfo.Host != "" {
		opts = append(opts, proxy(*proxyInfo))
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.ErrorSystemError("new wechat pay client err:%s", err.Error())
	}
	w.Client = client
	return w, nil
}

func autoVisitor(c config.Config, privateKey *rsa.PrivateKey) (auth.Verifier, error) {
	ctx := context.Background()
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	mgr := downloader.MgrInstance()
	if !mgr.HasDownloader(context.Background(), c.MchID) {
		err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, privateKey, c.RsaPrivateNumber, c.MchID, c.APISecret)
		if err != nil {
			return nil, err
		}
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(c.MchID)
	return verifiers.NewSHA256WithRSAVerifier(certificateVisitor), nil
}

type WithProxyOption struct {
	Proxy pro.Proxy
}

func (w *WithProxyOption) Apply(settings *core.DialSettings) error {
	settings.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (u *url.URL, err error) {
				u, err = url.Parse(fmt.Sprintf("http://%s:%d", w.Proxy.Host, w.Proxy.Port))
				if err != nil {
					return nil, err
				}
				if w.Proxy.UserName != "" && w.Proxy.Password != "" {
					u.User = url.UserPassword(w.Proxy.UserName, w.Proxy.Password)
				}
				if w.Proxy.UserName != "" {
					u.User = url.User(w.Proxy.UserName)
				}
				return u, nil
			},
		},
	}
	return nil
}
func proxy(proxy pro.Proxy) core.ClientOption {
	return &WithProxyOption{Proxy: proxy}
}
