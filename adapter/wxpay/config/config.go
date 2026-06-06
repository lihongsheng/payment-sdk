package config

import "github.com/lihongsheng/payment-sdk/errors"

type Config struct {
	Merchant Merchant `json:"merchant"`
	Cert     Cert     `json:"cert"`
	Service  Service  `json:"service"`
}

type Merchant struct {
	// 微信app id
	AppID string `json:"app_id"`
	// 微信商户id
	MchID string `json:"mch_id"`
	// 应用 secret 用于微信用户登录等
	AppSecret string `json:"app_secret"`
}

type Cert struct {
	// 微信api v3 秘钥
	APISecret string `json:"api_secret"`
	// 微信私钥或者其他平台私钥
	RsaPrivate string `json:"rsa_private"`
	// 微信私钥证书序列号
	RsaPrivateNumber string `json:"rsa_private_number"`
	// 微信公钥证书序列号
	RsaPublicNumber string `json:"rsa_public_number"`
	// 微信公钥
	RsaPublic string `json:"rsa_public"`
}

type Service struct {
	// 转账相关：支付分服务id
	ScoreServiceID string `json:"score_service_id"`
}

func (c Config) Validate() error {
	if c.Merchant.AppID == "" {
		return errors.ErrorParamError("微信: 应用ID is empty")
	}
	if c.Cert.APISecret == "" {
		return errors.ErrorParamError("微信: apiV3 秘钥不可为空")
	}
	if c.Merchant.AppSecret == "" {
		return errors.ErrorParamError("微信: 应用 secret is empty")
	}
	if c.Cert.RsaPrivate == "" {
		return errors.ErrorParamError("微信: 私钥 is empty")
	}
	if c.Cert.RsaPrivateNumber == "" {
		return errors.ErrorParamError("微信: 私钥正式序列号不可为空")
	}
	if c.Cert.RsaPublic != "" && c.Cert.RsaPublicNumber == "" {
		return errors.ErrorParamError("微信: 公钥和公钥证书编码需要同时填写")
	}
	if c.Cert.RsaPublic == "" && c.Cert.RsaPublicNumber != "" {
		return errors.ErrorParamError("微信: 公钥和公钥证书编码需要同时填写")
	}
	return nil
}