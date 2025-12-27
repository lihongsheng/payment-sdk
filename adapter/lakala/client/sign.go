package client

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"time"
)

type SignParams struct {
	AppID     string
	SerialNO  string
	Timestamp string
	NonceStr  string
}

type Sign struct {
	conf       config.Config
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func NewSign(conf config.Config) (*Sign, error) {
	publicKey, err := tools.LoadPublicKey(conf.Cert.PublicKey)
	if err != nil {
		return nil, err
	}
	privateKey, err := tools.LoadPrivateKey(conf.Cert.CertPrivateKey)
	if err != nil {
		return nil, err
	}
	return &Sign{
		conf:       conf,
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

func (s *Sign) Gen(body interface{}) (string, error) {
	singParam := SignParams{
		AppID:     s.conf.AppID,
		SerialNO:  s.conf.Cert.CertificateSerialNumber,
		Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
		NonceStr:  tools.GenerateRandomDigits(12),
	}
	var bodyStr string
	switch v := body.(type) {
	case string:
		// body本身是字符串，直接使用
		bodyStr = v
	case []byte:
		// body是字节数组，转为字符串
		bodyStr = string(v)
	case nil:
		// body为空，设为空字符串
		bodyStr = ""
	default:
		// body是结构体/映射，JSON序列化
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("body序列化失败: %w", err)
		}
		bodyStr = string(jsonBytes)
	}
	signParamStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", singParam.AppID, singParam.SerialNO, singParam.Timestamp, singParam.NonceStr, bodyStr)
	fmt.Println(signParamStr)
	signStr, err := s.RsaSign(signParamStr)
	if err != nil {
		return "", err
	}
	signValue := fmt.Sprintf(`LKLAPI-SHA256withRSA appid="%s",serial_no="%s",timestamp="%s",nonce_str="%s",signature="%s"`,
		singParam.AppID, singParam.SerialNO, singParam.Timestamp, singParam.NonceStr, signStr)
	fmt.Println(signValue)
	return signValue, nil
}

// RsaSign
// 使用私钥对数据进行SHA256withRSA签名
func (s *Sign) RsaSign(data string) (signatureBase64 string, err error) {
	// 计算SHA256哈希
	hash := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", errors.ErrorParamError("签名失败: %v", err)
	}
	// 签名结果进行Base64编码，方便传输
	return base64.StdEncoding.EncodeToString(signature), nil
}

func (s *Sign) BuildCallbackSignStr(singParam SignParams, body string) string {
	signParamStr := fmt.Sprintf("%s\n%s\n%s\n", singParam.Timestamp, singParam.NonceStr, body)
	return signParamStr
}

// RsaVerify
// 使用公钥验证SHA256withRSA签名
func (s *Sign) RsaVerify(data string, signatureBase64 string) (bool, error) {
	// 解码Base64格式的签名
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, errors.ErrorParamError("解析签名失败: %v", err)
	}
	// 计算数据的SHA256哈希
	hash := sha256.Sum256([]byte(data))
	// 4. SHA256withRSA验签（对应PHP的OPENSSL_ALGO_SHA256）
	//hash := sha256.New()
	//hash.Write([]byte(data))
	//hashed := hash.Sum(nil)
	// 验证签名
	err = rsa.VerifyPKCS1v15(s.publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false, errors.ErrorParamError("签名验证失败: %v", err)
	}

	return true, nil
}
