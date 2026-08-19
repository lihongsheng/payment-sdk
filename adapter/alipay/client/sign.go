package client

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/errors"
	"net/url"
	"sort"
	"strings"
)

var (
	skipSign = map[string]struct{}{
		enum.COMMON_PARAM_SING_NAME: {},
		//enum.COMMON_PARAM_NOTIFY_URL_NAME:     {},
		//enum.COMMON_PARAM_APP_AUTH_TOKEN_NAME: {},
	}
)

type Sign struct {
	conf       config.Config
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewSign(conf config.Config) (*Sign, error) {
	private, err := loadPravate(conf.Cert.RsaPrivate)
	if err != nil {
		return nil, err
	}
	public, err := loadPublic(conf.Cert.RsaPublic)
	if err != nil {
		return nil, err
	}
	return &Sign{
		conf:       conf,
		PrivateKey: private,
		PublicKey:  public,
	}, nil
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
	return s.RsaSign(signStr, s.conf.Cert.RsaPrivate)
}

func loadPravate(privateKeyPEM string) (*rsa.PrivateKey, error) {
	// 解析PEM格式的私钥
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, errors.ErrorParamError("解析私钥失败")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.ErrorParamError("解析私钥内容失败: %v", err)
	}
	// 使用私钥签名
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.ErrorParamError("私钥不是RSA类型")
	}
	return privateKey, nil
}

func loadPublic(publicKeyPEM string) (*rsa.PublicKey, error) {
	// 解析PEM格式的公钥
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, errors.ErrorParamError("解析公钥失败")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.ErrorParamError("解析公钥内容失败: %v", err)
	}
	// 类型断言，确保是RSA公钥
	public, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.ErrorParamError("公钥不是RSA类型")
	}
	return public, nil
}

// RsaSign
// 使用私钥对数据进行SHA256withRSA签名
func (s *Sign) RsaSign(data string, privateKeyPEM string) (signatureBase64 string, err error) {
	// 计算SHA256哈希
	hash := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", errors.ErrorParamError("签名失败: %v", err)
	}
	// 签名结果进行Base64编码，方便传输
	return base64.StdEncoding.EncodeToString(signature), nil
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
	// 验证签名
	err = rsa.VerifyPKCS1v15(s.PublicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false, errors.ErrorParamError("签名验证失败: %v", err)
	}

	return true, nil
}

// GenerateSignString 生成待签名字符串（通用版，剔除sign、sign_type）
// params: 通知返回的所有参数（url.Values 格式）
// 返回值: 按规则处理后的待签名字符串
func (s *Sign) GenerateSignString(params url.Values) (string, string, error) {
	if len(params) == 0 || len(params["sign"]) == 0 || params["sign"][0] == "" {
		return "", "", errors.ErrorParamError("未找到签名字符串")
	}
	signValue := params["sign"][0]
	// 防御性修复：Base64 中不存在空格，所有空格都是 + 被误解析的
	signValue = strings.ReplaceAll(signValue, " ", "+")

	// 1. 过滤参数：剔除 sign、sign_type，保留其他非空参数
	filteredParams := make(url.Values)
	for k, v := range params {
		// 剔除 sign、sign_type 参数
		if k == "sign" || k == "sign_type" {
			continue
		}
		// 过滤空值参数（若业务需保留空值，可删除此判断）
		if len(v) > 0 && v[0] != "" {
			filteredParams[k] = v
		}
	}

	// 2. 提取参数名并按字典序排序
	var keys []string
	for k := range filteredParams {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Go原生字典序排序（ASCII升序）

	// 3. URL解码参数值 + 拼接待签名字符串
	var signStrBuilder strings.Builder
	for i, k := range keys {
		decodedValue := filteredParams.Get(k)
		// 拼接：多个参数用&分隔，格式为 key=value
		if i > 0 {
			signStrBuilder.WriteString("&")
		}
		signStrBuilder.WriteString(k)
		signStrBuilder.WriteString("=")
		signStrBuilder.WriteString(decodedValue)
	}
	return signStrBuilder.String(), signValue, nil
}
