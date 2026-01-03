package config

import "github.com/lihongsheng/payment-sdk/config/proxy"

type Config struct {
	// 应用ID
	AppID string `json:"app_id"`
	// 证书
	Cert Cert `json:"cert"`
	// 代理
	Proxy proxy.Proxy `json:"proxy"`
}

type Cert struct {
	// 私钥 rsa 格式
	Private string `json:"private_key"`
	// 私钥证书序列号
	PrivateNumber string `json:"private_number"`
	// 公钥证书序列号
	PublicNumber string `json:"public_number"`
	// 公钥 rsa 格式
	Public string `json:"public_key"`
}
