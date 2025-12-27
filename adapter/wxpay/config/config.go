package config

type Config struct {
	AppID          string
	MchID          string
	APIKey         string
	Cert           Cert
	ScoreServiceID string `json:"score_service_id"`
}

type Cert struct {
	// 微信私钥或者其他平台私钥
	Private string `json:"private_key"`
	// 微信证书序列号
	PrivateNumber string `json:"private_number"`
	// 微信公钥ID
	PublicNumber string `json:"public_number"`
	// 微信公钥
	PublicKey string `json:"public_key"`
}
