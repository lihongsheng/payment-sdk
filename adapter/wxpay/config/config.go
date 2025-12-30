package config

type Config struct {
	// 微信app id
	AppID string `json:"app_id"`
	// 微信商户id
	MchID string `json:"mch_id"`
	// 微信api v3 秘钥
	APISecret string `json:"api_secret"`
	// 证书相关
	Cert Cert `json:"cert"`
	// 转账相关
	ScoreServiceID string `json:"score_service_id"`
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
