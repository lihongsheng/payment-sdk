package config

import (
	"github.com/lihongsheng/payment-sdk/config/proxy"
	"github.com/lihongsheng/payment-sdk/errors"
)

type Config struct {
	// 应用ID
	AppID string `json:"app_id"`
	// 证书
	// 私钥 rsa 格式
	RsaPrivate string `json:"rsa_private_key"`
	// 支付宝证书序列号,转账场景需要
	RsaAppCertSN string `json:"rsa_private_number"`
	// 支付宝公钥根证书序列号,转账场景需要
	RsaRootCertSN string `json:"rsa_public_number"`
	// 支付宝证书公钥 rsa 格式,可以通过AppCrt提取
	RsaPublic string `json:"rsa_public_key"`
	// 支付宝根证书crt
	RsaRootCrt string `json:"rsa_root_crt"`
	// app_crt 应用crt 文件
	RsaAppCrt string `json:"rsa_app_crt"`
	// 代理
	Proxy proxy.Proxy `json:"proxy"`
	// app_auth_token
	AppAuthToken string `json:"app_auth_token"`
}

func (c Config) Validate() error {
	if c.AppID == "" {
		return errors.ErrorParamError("支付宝: 应用ID is empty")
	}
	if c.RsaPrivate == "" {
		return errors.ErrorParamError("支付宝: 私钥 is empty")
	}
	if c.RsaAppCertSN == "" && (c.RsaPublic == "" || c.RsaAppCertSN == "") {
		return errors.ErrorParamError("支付宝: 应用crt 或者公钥不可同时为空")
	}
	return nil
}
