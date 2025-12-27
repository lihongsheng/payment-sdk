package client

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/alipay/enum"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/errors"
	"sort"
)

var (
	skipSign = map[string]struct{}{
		enum.COMMON_PARAM_SING_NAME: {},
		//enum.COMMON_PARAM_NOTIFY_URL_NAME:     {},
		//enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME: {},
	}
)

type Sign struct {
	conf config.Config
}

func NewSign(conf config.Config) *Sign {
	return &Sign{
		conf: conf,
	}
}

func (s *Sign) Sign(signParams map[string]string, body map[string]string) (string, error) {
	if signParams[enum.COMMON_PARAM_METHOD_NAME] == "" {
		return "", errors.ErrorParamError("method is empty")
	}
	for k, v := range body {
		signParams[k] = v
	}
	// 1. 提取map的所有key到切片
	keys := make([]string, 0, len(signParams))
	for k := range signParams {
		if _, exists := skipSign[k]; exists {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, k := range keys {
		signStr += k + "=" + signParams[k] + "&"
	}
	signStr = signStr[:len(signStr)-1]
	fmt.Println(signStr)
	return s.RsaSign(signStr, s.conf.Cert.CertPrivateKey)
}

// RsaSign
// 使用私钥对数据进行SHA256withRSA签名
func (s *Sign) RsaSign(data string, privateKeyPEM string) (signatureBase64 string, err error) {
	// 解析PEM格式的私钥
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", errors.ErrorParamError("解析私钥失败")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", errors.ErrorParamError("解析私钥内容失败: %v", err)
	}
	// 计算SHA256哈希
	hash := sha256.Sum256([]byte(data))
	// 使用私钥签名
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return "", errors.ErrorParamError("私钥不是RSA类型")
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", errors.ErrorParamError("签名失败: %v", err)
	}
	// 签名结果进行Base64编码，方便传输
	return base64.StdEncoding.EncodeToString(signature), nil
}

//func (s *Sign) Verify(body map[string]string) (bool, error) {
//	keys := make([]string, 0, len(signParams))
//	for k := range signParams {
//		if _, exists := skipSign[k]; exists {
//			continue
//		}
//		keys = append(keys, k)
//	}
//	sort.Strings(keys)
//}

// RsaVerify
// 使用公钥验证SHA256withRSA签名
func (s *Sign) RsaVerify(data string, signatureBase64 string, publicKeyPEM string) (bool, error) {
	// 解析PEM格式的公钥
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return false, errors.ErrorParamError("解析公钥失败")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, errors.ErrorParamError("解析公钥内容失败: %v", err)
	}

	// 类型断言，确保是RSA公钥
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("公钥不是RSA类型")
	}

	// 解码Base64格式的签名
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, errors.ErrorParamError("解析签名失败: %v", err)
	}

	// 计算数据的SHA256哈希
	hash := sha256.Sum256([]byte(data))

	// 验证签名
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false, errors.ErrorParamError("签名验证失败: %v", err)
	}

	return true, nil
}
