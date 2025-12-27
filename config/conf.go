package config

type Config struct {
	AppID   string
	MchID   string
	APIKey  string
	Cert    Cert
	Proxy   Proxy
	ApiHost string
	Version string
	Extra   string
}

type Cert struct {
	// 微信私钥或者其他平台私钥
	CertPrivateKey string `json:"cert_private_key"`
	// 微信证书序列号
	CertificateSerialNumber string `json:"certificate_serial_number"`
	// 微信公钥ID
	PublicKeyID string `json:"public_key_id"`
	// 微信公钥
	PublicKey      string `json:"public_key"`
	ScoreServiceID string `json:"score_service_id"`
}

type Proxy struct {
	Host     string
	Port     int
	UserName string
	Password string
}
