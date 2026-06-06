package client

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/lakala/config"
	"github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/tools"
	"net/http"
	"regexp"
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
	publicKey, err := tools.LoadPublicKey(conf.Cert.RsaPublic)
	if err != nil {
		return nil, err
	}
	privateKey, err := tools.LoadPrivateKey(conf.Cert.RsaPrivate)
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
		AppID:     s.conf.Merchant.AppID,
		SerialNO:  s.conf.Cert.RsaPrivateNumber,
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

func (s *Sign) ParseSign(req *http.Request) (SignParams, string, error) {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return SignParams{}, "", errors.ErrorParamError("not find Authorization")
	}
	timestamp, nonceStr, signature, exists := ExtractLKLAuthParams(auth)
	if !exists {
		return SignParams{}, "", errors.ErrorParamError("not find Authorization")
	}
	signParams := SignParams{
		Timestamp: timestamp,
		NonceStr:  nonceStr,
	}
	return signParams, signature, nil
}

// ExtractLKLAuthParams 对齐PHP逻辑提取timestamp、nonce_str、signature
// 参数：authorization - 原始授权字符串（如"LKLAPI-SHA256withRSA timestamp=\"1765848504\",nonce_str=\"bGfguYYXFzbO\",signature=\"xxx\""）
// 返回：timestamp, nonceStr, signature, 提取成功标识
func ExtractLKLAuthParams(authorization string) (timestamp, nonceStr, signature string, ok bool) {
	// 完全对齐PHP的正则表达式：
	// PHP: /timestamp="(\d+)",nonce_str="(\w+)",signature="([^"]+)"/
	// Go中需转义反斜杠，规则完全一致：
	// \d+ → 匹配纯数字（timestamp）
	// \w+ → 匹配字母/数字/下划线（nonce_str）
	// [^"]+ → 匹配除双引号外的所有字符（signature，兼容特殊字符）
	pattern := `timestamp="(\d+)",nonce_str="(\w+)",signature="([^"]+)"`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(authorization)
	// 注意：FindStringSubmatch匹配成功时，matches[0]是完整匹配串，1/2/3是捕获组（对齐PHP的$matches[1/2/3]）
	if len(matches) < 4 { // 匹配成功时至少有4个元素（0:完整串,1:timestamp,2:nonce_str,3:signature）
		return "", "", "", false
	}
	// 按PHP的捕获组顺序赋值
	timestamp = matches[1]
	nonceStr = matches[2]
	signature = matches[3]
	return timestamp, nonceStr, signature, true
}
