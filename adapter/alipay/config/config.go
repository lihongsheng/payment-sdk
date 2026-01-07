package config

import "github.com/lihongsheng/payment-sdk/config/proxy"

type Config struct {
	// 应用ID
	AppID string `json:"app_id"`
	// 证书
	Cert Cert `json:"cert"`
	// 代理
	Proxy proxy.Proxy `json:"proxy"`
	// app_auth_token
	AppAuthToken string `json:"app_auth_token"`
}

type Cert struct {
	// 私钥 rsa 格式
	Private string `json:"private_key"`
	// 支付宝证书序列号,转账场景需要
	AppCertSN string `json:"private_number"`
	// 支付宝公钥根证书序列号,转账场景需要
	RootCertSN string `json:"public_number"`
	// 支付宝证书公钥 rsa 格式,可以通过AppCrt提取
	Public string `json:"public_key"`
	// 支付宝根证书crt
	RootCrt string `json:"root_crt"`
	// app_crt 应用crt 文件
	AppCrt string `json:"app_crt"`
}
