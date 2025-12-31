package config

import "github.com/lihongsheng/payment-sdk/config/proxy"

type Config struct {
	// 微信商户id
	MchID string `json:"mch_id"`
	// 微信api v3 秘钥
	APISecret string `json:"api_secret"`
	// 富有订单前缀
	OrderPrefix string `json:"order_prefix"`
	// 富有接口地址
	ApiHost string `json:"api_host"`
	Cert    Cert   `json:"cert"`
	// 代理
	Proxy   proxy.Proxy `json:"proxy"`
	Version string      `json:"version"`
}

type Cert struct {
	// 微信私钥或者其他平台私钥
	Private string `json:"private_key"`
	// 微信私钥证书序列号
	PrivateNumber string `json:"private_number"`
	// 微信公钥证书序列号
	PublicNumber string `json:"public_number"`
	// 微信公钥
	PublicKey string `json:"public_key"`
}
