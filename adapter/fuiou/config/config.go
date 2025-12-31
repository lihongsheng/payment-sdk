package config

import "github.com/lihongsheng/payment-sdk/config/proxy"

type Config struct {
	// 商户id
	MchID string `json:"mch_id"`
	// 接口秘钥
	APISecret string `json:"api_secret"`
	// 富有订单前缀
	OrderPrefix string `json:"order_prefix"`
	// 富有接口地址，不填默认https://aipay-cloud.fuioupay.com
	ApiHost string `json:"api_host"`
	// 富有证书，转账的时候有用
	Cert Cert `json:"cert"`
	// 代理
	Proxy proxy.Proxy `json:"proxy"`
	// 版本，不填默认1.0
	Version string `json:"version"`
}

type Cert struct {
	// 私钥 rsa 格式
	Private string `json:"private_key"`
	// 公钥 rsa 格式
	Public string `json:"public_key"`
}
