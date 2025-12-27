package client

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/singer-stack-lab/payment-sdk/adapter/fuiou/util"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/errors"
	"sort"
)

type Sign struct {
	C config.Config
}

func NewSign(c config.Config) *Sign {
	return &Sign{
		C: c,
	}
}

func (s *Sign) Sign(signParams map[string]string, filter map[string]struct{}) (sign string, err error) {
	keys := make([]string, 0, len(signParams))
	for k, _ := range signParams {
		if _, ok := filter[k]; ok {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, k := range keys {
		v := signParams[k]
		signStr += k + "=" + v + "&"
	}
	signStr = signStr[:len(signStr)-1]
	//fmt.Println("signStr", signStr)
	// 需要转码为GBK
	signBy, _ := util.GbkEncode(signStr)
	sign, err = s.RsaSign(signBy, s.C.Cert.CertPrivateKey)
	return sign, err
}

// RsaSign
// 使用商户私钥对数据进行Md5withRSA签名
func (s *Sign) RsaSign(data []byte, privateKeyPEM string) (signatureBase64 string, err error) {
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
	hasher := md5.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	// 使用私钥签名
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return "", errors.ErrorParamError("私钥不是RSA类型")
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hash[:])
	if err != nil {
		return "", errors.ErrorParamError("签名失败: %v", err)
	}
	// 签名结果进行Base64编码，方便传输
	return base64.StdEncoding.EncodeToString(signature), nil
}

// EncryptByPublicKey 使用RSA公钥加密数据
// publicKeyBytes 是PEM格式的公钥字节数组
// srcData 是要加密的原始数据
func (s *Sign) EncryptByPublicKey(srcData []byte, publicKeyBytes []byte) (string, error) {
	// 解析PEM格式的公钥
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return "", errors.ErrorParamError("failed to decode public key PEM")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言为*rsa.PublicKey
	publicKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return "", errors.ErrorParamError("not a RSA public key")
	}
	// 获取RSA公钥的大小（字节数）
	keySize := publicKey.Size()
	// 计算最大明文块大小（RSA加密中，PKCS#1填充会占用11字节）
	maxBlockSize := keySize - 11
	var encryptedData []byte
	// 分段加密
	for i := 0; i < len(srcData); i += maxBlockSize {
		end := i + maxBlockSize
		if end > len(srcData) {
			end = len(srcData)
		}
		// 加密当前块
		blockEncrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, srcData[i:end])
		if err != nil {
			return "", err
		}
		// 将加密后的块追加到结果中
		encryptedData = append(encryptedData, blockEncrypted...)
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// DecryptByKey 使用RSA密钥（私钥或公钥）解密数据
// keyBytes 是PEM格式的密钥字节数组（私钥或公钥）
// srcData 是要解密的加密数据
func (s *Sign) DecryptByKey(srcDataStr string, keyBytes []byte) ([]byte, error) {
	srcData, err := base64.StdEncoding.DecodeString(srcDataStr)
	if err != nil {
		return nil, errors.ErrorParamError("base64 decode err").WithCause(err)
	}
	// 解析PEM格式的密钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.ErrorParamError("failed to decode key PEM")
	}
	// 尝试解析私钥（优先私钥，因为RSA通常用私钥解密）

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.ErrorParamError("解析私钥内容失败: %v", err)
	}
	private, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.ErrorParamError("私钥不是RSA类型")
	}
	// 私钥解密（常规场景）
	return s.decryptByPrivateKey(srcData, private)
}

// decryptByPrivateKey 使用RSA私钥解密
func (s *Sign) decryptByPrivateKey(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	keySize := privateKey.Size()
	var plaintext []byte
	// 分段解密（RSA解密的块大小等于密钥长度）
	for i := 0; i < len(ciphertext); i += keySize {
		end := i + keySize
		if end > len(ciphertext) {
			end = len(ciphertext)
		}
		// 使用PKCS#1 v1.5填充模式解密
		blockDecrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext[i:end])
		if err != nil {
			return nil, err
		}
		plaintext = append(plaintext, blockDecrypted...)
	}
	return plaintext, nil
}
