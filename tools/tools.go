package tools

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

// isTimeActive 判断时间是否在一个范围内
func isTimeActive(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}

func UnixToTime(t int64) time.Time {
	return time.Unix(t, 0)
}

func EndTime(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d+1, 0, 0, 0, 0, time.Local)
}

func StartTime(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func Md5(content string) (md string) {
	h := md5.New()
	_, _ = io.WriteString(h, content)
	md = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func IsURI(fl string) bool {
	if !strings.Contains(fl, "http://") && !strings.Contains(fl, "https://") {
		fl = "http://" + fl
	}
	s := fl
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func GenerateID() string {
	return uuid.New().String()
}

const time33Hash = uint64(5381)

func Time33(str string) uint64 {
	hash := time33Hash
	for _, char := range str {
		tmp, _ := strconv.ParseInt(fmt.Sprintf("%d", char), 10, 0)
		hash = (hash << 5) + uint64(tmp)
	}
	return hash & 0x7FFFFFFF
}

func HmacSha256(key []byte, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(data)
	return mac.Sum(nil)
}

// GenerateQRToBase64 生成二维码并转为 Base64 字符串
// content: 二维码内容（如URL、文本）
// size: 二维码尺寸（像素）
// level: 纠错级别（Low/Medium/High/Highest）
// 返回值: Base64 字符串（带 Data URI 前缀，可直接用于网页 <img> 标签）
func GenerateQRToBase64(content string, size int, level qrcode.RecoveryLevel) (string, error) {
	// 1. 生成二维码图片（内存中，不落地）
	qrImg, err := qrcode.Encode(content, level, size)
	if err != nil {
		return "", err
	}
	// 2. 将二进制图片转为 Base64 编码
	base64Str := base64.StdEncoding.EncodeToString(qrImg)
	// 3. 拼接 Data URI 前缀（网页中可直接用 <img src="xxx"> 展示）
	// 若仅需纯 Base64 字符串，可去掉 "data:image/png;base64," 前缀
	return base64Str, nil
}

// GetCertInfo
// 解析证书内容，获取序列号和签发机构并拼接
func GetCertInfo(certContent string) (string, error) {
	// 1. 解析 PEM 格式证书
	block, _ := pem.Decode([]byte(certContent))
	if block == nil || block.Type != "CERTIFICATE" {
		return "", fmt.Errorf("无效的 PEM 证书格式")
	}
	// 2. 解析 X.509 证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("解析 X.509 证书失败: %v", err)
	}
	// 3. 获取序列号（十进制字符串，与 Java 的 cert.getSerialNumber().toString() 一致）
	serialNumber := cert.SerialNumber.String()
	// 4. 获取签发机构名称（与 Java 的 cert.getIssuerX500Principal().getName() 格式一致）
	// 注意：Go 的 cert.Issuer.String() 会自动按 RFC 格式排序（如 C=CN,O=xxx,CN=xxx），与 Java 一致
	issuerName := cert.Issuer.String()
	// 5. 拼接签发机构和序列号（与 Java 输出一致）
	combined := issuerName + serialNumber
	return Md5(combined), nil
}

// GetCertSerial 从证书文件中提取纯序列号（十六进制，无前缀）
func GetCertSerial(certData string) (string, error) {
	// 1. 解析 PEM 格式
	block, _ := pem.Decode([]byte(certData))
	if block == nil || block.Type != "CERTIFICATE" {
		return "", fmt.Errorf("无效的 PEM 证书")
	}
	// 2. 解析 X.509 证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("解析 X.509 证书失败: %v", err)
	}
	// 3. 序列号转为十六进制字符串（大写，无前缀）
	return cert.SerialNumber.Text(16), nil
}

// ParseCerToPublicKeyPEM 解析 cer 文件字符串（PEM/DER 格式），返回 PEM 格式公钥字符串
func ParseCerToPublicKeyPEM(cerContent string) (string, error) {
	cerBytes := []byte(cerContent)

	// 尝试解析 PEM 格式证书
	var cert *x509.Certificate
	block, _ := pem.Decode(cerBytes)
	if block == nil || block.Type != "CERTIFICATE" {
		return "", fmt.Errorf("无效的 crt 证书")
	}
	if block != nil && block.Type == "CERTIFICATE" {
		// PEM 格式证书
		var err error
		cert, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			return "", fmt.Errorf("解析 PEM 证书失败: %w", err)
		}
	} else {

	}

	// 提取公钥并转为 PKIX 格式
	pubBytes, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return "", fmt.Errorf("公钥编码失败: %w", err)
	}

	// 编码为 PEM 格式
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}
	publicKeyPEM := string(pem.EncodeToMemory(pemBlock))
	if publicKeyPEM == "" {
		return "", errors.New("PEM 编码公钥失败")
	}

	return publicKeyPEM, nil
}

// LoadPublicKey 智能加载公钥：支持cer证书字符串或PEM公钥字符串
func LoadPublicKey(input string) (*rsa.PublicKey, error) {
	inputBytes := []byte(input)
	// 尝试解析为PEM格式公钥（优先判断）
	publicKey, _ := loadPublicKey(input)
	if publicKey != nil {
		return publicKey, nil
	}
	// 尝试解析为CER证书（PEM/DER格式）
	var cert *x509.Certificate
	// 先尝试PEM格式证书
	certBlock, _ := pem.Decode(inputBytes)
	if certBlock == nil || certBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("无效的 crt 证书")
	}
	var err error
	cert, err = x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析PEM格式CER证书失败: %w", err)
	}
	// 类型断言，确保是RSA公钥
	rsaPublicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("公钥不是RSA类型")
	}
	return rsaPublicKey, nil
}

func loadPublicKey(input string) (*rsa.PublicKey, error) {
	inputBytes := []byte(input)
	// 尝试解析为PEM格式公钥（优先判断）
	pubBlock, _ := pem.Decode(inputBytes)
	if pubBlock != nil {
		switch pubBlock.Type {
		case "PUBLIC KEY":
			// PKIX格式公钥
			pubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
			if err != nil {
				return nil, fmt.Errorf("解析PEM公钥失败: %w", err)
			}
			// 类型断言，确保是RSA公钥
			rsaPublicKey, ok := pubKey.(*rsa.PublicKey)
			if !ok {
				return nil, fmt.Errorf("公钥不是RSA类型")
			}
			return rsaPublicKey, nil
		case "RSA PUBLIC KEY":
			// RSA格式公钥
			rsaPubKey, err := x509.ParsePKCS1PublicKey(pubBlock.Bytes)
			if err != nil {
				return nil, fmt.Errorf("解析RSA PEM公钥失败: %w", err)
			}
			return rsaPubKey, nil
		}
	}
	return nil, errors.New("parse ras public key fail")
}

func LoadPrivateKey(input string) (*rsa.PrivateKey, error) {
	// 解析PEM格式的私钥
	block, _ := pem.Decode([]byte(input))
	if block == nil {
		return nil, errors.New("解析私钥失败")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析私钥内容失败: %v", err))
	}
	// 使用私钥签名
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("私钥不是RSA类型")
	}
	return privateKey, nil
}
