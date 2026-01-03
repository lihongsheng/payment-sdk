package config

import "github.com/lihongsheng/payment-sdk/config/proxy"

type Config struct {
	//app id
	AppID string `json:"app_id"`
	// 商户号
	MchID string `json:"mch_id"`
	// api  秘钥
	APISecret string `json:"api_secret"`
	// 证书
	Cert Cert `json:"cert"`
	// 代理
	Proxy proxy.Proxy `json:"proxy"`
	// 终端号
	TermNO string `json:"term_no"`
	// api 地址 https://s2.lakala.com
	ApiHost string `json:"api_host"`
}

type Cert struct {
	// 私钥 rsa 格式
	Private string `json:"private_key"`
	// 私钥证书序列号
	PrivateNumber string `json:"private_number"`
	// 公钥 rsa 格式
	Public string `json:"public_key"`
}
