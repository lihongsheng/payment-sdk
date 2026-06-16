package config

import (
	"github.com/lihongsheng/payment-sdk/errors"
)

type Config struct {
	Merchant Merchant `json:"merchant"`
	Cert     Cert     `json:"cert"`
	API      API      `json:"api"`
	// 用于获取微信用户 openid
	Wechat Wechat `json:"wechat"`
	// 支付宝相关
	Alipay Alipay `json:"alipay"`
}

type Merchant struct {
	// 商户id
	MchID string `json:"mch_id"`
	// 接口秘钥
	APISecret string `json:"api_secret"`
	// 富有订单前缀
	OrderPrefix string `json:"order_prefix"`
}

type Cert struct {
	// 私钥 rsa 格式
	RsaPrivate string `json:"rsa_private_key"`
	// 公钥 rsa 格式
	RsaPublic string `json:"rsa_public_key"`
}

type API struct {
	// 富有接口地址，不填默认https://aipay-cloud.fuioupay.com
	ApiHost string `json:"api_host"`
	// 版本，不填默认1.0
	Version string `json:"version"`
}

type Wechat struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type Alipay struct {
	AppID      string `json:"app_id"`
	RsaPrivate string `json:"rsa_private_key"`
	RsaRootCrt string `json:"rsa_root_crt"`
}

func (c Config) Validate() error {
	if c.Merchant.MchID == "" {
		return errors.ErrorParamError("富友: 商户id is empty")
	}
	if c.Cert.RsaPrivate == "" {
		return errors.ErrorParamError("富友: 私钥 is empty")
	}
	if c.Merchant.OrderPrefix == "" {
		return errors.ErrorParamError("富友: 订单前缀前缀不可为空")
	}
	return nil
}