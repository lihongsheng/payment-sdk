package config

import "github.com/lihongsheng/payment-sdk/errors"

type Config struct {
	//app id
	AppID string `json:"app_id"`
	// 商户号
	MchID string `json:"mch_id"`
	// 私钥 rsa 格式
	RsaPrivate string `json:"rsa_private_key"`
	// 私钥证书序列号
	RsaPrivateNumber string `json:"rsa_private_number"`
	// 公钥 rsa 格式
	RsaPublic string `json:"rsa_public_key"`
	// 终端号
	TermNO string `json:"term_no"`
	// api 默认地址 https://s2.lakala.com
	ApiHost string `json:"api_host"`
}

func (c Config) Validate() error {
	if c.AppID == "" {
		return errors.ErrorParamError("拉卡拉: 应用ID is empty")
	}
	if c.TermNO == "" {
		return errors.ErrorParamError("拉卡拉: 终端号 is empty")
	}
	if c.RsaPrivate == "" {
		return errors.ErrorParamError("拉卡拉: 私钥 is empty")
	}
	if c.RsaPrivateNumber == "" {
		return errors.ErrorParamError("拉卡拉: 私钥正式序列号不可为空")
	}
	if c.RsaPublic == "" {
		return errors.ErrorParamError("拉卡拉: 公钥不可为空")
	}
	return nil
}
