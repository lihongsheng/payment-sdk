package config

import (
	"github.com/lihongsheng/payment-sdk/errors"
)

type Config struct {
	// 商户id
	MchID string `json:"mch_id"`
	// 接口秘钥
	APISecret string `json:"api_secret"`
	// 富有订单前缀
	OrderPrefix string `json:"order_prefix"`
	// 富有接口地址，不填默认https://aipay-cloud.fuioupay.com
	ApiHost string `json:"api_host"`
	// 私钥 rsa 格式
	RsaPrivate string `json:"rsa_private_key"`
	// 公钥 rsa 格式
	RsaPublic string `json:"rsa_public_key"`
	// 版本，不填默认1.0
	Version string `json:"version"`
	// 用于获取微信用户 openid
	WechatAppId string `json:"wechat_app_id"`
	// wechat_app_secret
	WechatAppSecret string `json:"wechat_app_secret"`
	// alipay_app_id
	AlipayAppId string `json:"alipay_app_id"`
	// alipay_rsa_private_key
	AlipayRsaPrivate string `json:"alipay_rsa_private_key"`
	// alipay_rsa_root_crt
	AlipayRsaRootCrt string `json:"alipay_rsa_root_crt"`
}

type Wechat struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func (c Config) Validate() error {
	if c.MchID == "" {
		return errors.ErrorParamError("富友: 商户id is empty")
	}
	if c.RsaPrivate == "" {
		return errors.ErrorParamError("富友: 私钥 is empty")
	}
	if c.OrderPrefix == "" {
		return errors.ErrorParamError("富友: 订单前缀前缀不可为空")
	}
	return nil
}
