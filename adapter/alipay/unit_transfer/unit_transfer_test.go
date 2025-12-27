package unit_transfer

import (
	"context"
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

// 69E2FEC0170AC67B
func Test_Transfer(t *testing.T) {
	c := config.Config{
		AppID: "2021005199606982",
		MchID: "",
		// 02941eef3187dddf3d3b83462e1dfcf6
		APIKey: "",
		Cert: config.Cert{
			CertificateSerialNumber: "5875781f2c5b439ed3999d3c21aa5562",
			CertPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC3fcEnCC2/djBnEYFDp41ZWuM4T49Pfr6p0gHXtfHGrO9CqegJhFPd4HT4BmVWEGgsHBMSiTce4b4VsrOQG0j7Lr2It0ziPpBwJUyQti5zZUHZoNl6FDVJIUo4p1DOxoqmYqj4iK/AgZVu0YiVNnbuKGWRoLt6B2++O05lX63qgAKxitqeRvGca73vbeiavUEhVaDVYbUBZOvLOuH/FDgLAHETOjXm59+Q7OEdT/cV+tltNyn/zTCUbql3I8MZQQkFcjOylREm+Jg7nGrK0o/9Bjw413LuQwe4Y0uXO4r7mUwV6QTS6L3L+8QzcBzpOUj1oC+iNM5Xe7b6/f8RssntAgMBAAECggEBAJg0xXM2MpSeWno0qBBRvUuAu/iq7krX7Rp7LLIdp8fRDcszm1nBnbvwA7b1vzuL8o2cYCnBDqscQIxJJvAD4h8R34C35BvMIA4WspNbu95XAS+gYKYGe0kFR0cFa1+Ki9qG/njjET7Tfvxk7TIw8qiNc+s/9CO+ZW/cTCSFIIPicYGwpIw1s15o8j/W9OMsxt7+RCiPdj0Kgu6SUAvwZ5hTCKFPO1sJJYU6DJJDCLSdKXrcez+fQKsWsimbHV1G3Jk4PFzaGruk9c+SUgUnPnRRAH/kxSa6bijNJjK/5apIXL0SIOfqztWV0ciG6evPu60L8nADb2dfKAzKudxtfnUCgYEA53Clnz3wlBStGqaOsR0FxngzS84Pe2uCguwcB1BYiVfY7OaN/a4YHm2MdaNfZI1sfJAbHiJJw2cbcb51iNOlWI6KMLqpXJzz2GLhhf8BlZuVQwtY/ogRTFOCrejs+3LarZR1ct8b56193wCh90Q4ZysCLxZLdvV0Ey56vMJpc8MCgYEAyvaE+iRdsdYj2g/OaBfKRxB3BQWlg/aO8169g7f7/yCoDkMZGZaLI/TiyEp6ja9J+OYt74aIwrAOlhXZOMIpkxuPa+opzRpYExPCAG1FUV5xtXYomTLHGZpY/VURu0HVx0GMvritzgDllTol7U+hHaF/ziGZiz637o+AvJD2YI8CgYBU18XPd6xvDJlc0Lw0j3gjhsL/Qh4I16OzjQzFXZ3nU23xgp+A5MZwuSYpped9fB7OFyHpzRYPbQfxjamWrEfMDAz6yiL2EY+OvskADziseKmiN1G7lXQJ7/9S87WUhElIIslfgAzBxKcFDj1R4Q9wOqMFUf3+MQMJWRujgP2ZcwKBgCaiVF+1KmyeoYZxuc2Qsb1jQfSfxYjgFwf2gcaf8AT5d2P3m8CGKog2pjCQFNIBiugpEJdmmYPNtYkWMZofQ4CwH1XgZgKXwCopeaVRJ3+8pZJwInw+8S1LdyMJ3W0ys/nQ3XS0fMkY3JrSTcPlg3q7cjOPq8WYs7RidpPuWCe7AoGBAM7j6/hIktBmsCsjZgXQOvH+I4RpXVg6finIRDEjjPPlpujCQZuIUL8/HXznyoT/yuqTFFOLTLOSrgASwNp175WlgCpZ+RNktsczMP/PfZTqzQWyYf/gzRH9OjdD8E+LTYtH89gfiO9XtdxYxE0ke8Bz7CQLFbFfH8vrTelkJq2n
-----END RSA PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsyA5/k4SyfoHhZCPIr/VjBG1vRhOwPLoOddf5WCH6168h/el+BGMiNOdczbP11HmrfpJWsXKDR7W/T/RMwX032RY2ah7HTJG1Auh5RXu4SoKmZYq84nvT0i0N8hp222E1fECts1NaWSfQP8dZcbFTd5HEWwbBmdFR/ExWl2u3wZB0OFHgIz4rfVV3qQnqFLzPqYIV23QfkHS0cVrlZU+GjSO7BVLRozOPD+0O1N9cuq8gjqSYDbR83VTAG27hfJel3prg5FTDAeLSi1eyUIrNRB3LZjaj5X90gyRHY4nNOx7Jh01YTuZR/FSHVHVdGdkw3h131FqmJpB2xAdc/qnVwIDAQAB
-----END PUBLIC KEY-----`,
			// 687b59193f3f462dd5336e5abf83c5d8_
			PublicKeyID: "02941eef3187dddf3d3b83462e1dfcf6",
		},
		Proxy:   config.Proxy{},
		ApiHost: "",
		Version: "",
	}
	pay, err := NewTransfer(c)
	assert.NoError(t, err)
	ctx := context.Background()
	p, err := pay.Transfer(ctx, &dto.UintTransferRequest{
		TransferNo: fmt.Sprintf("%d", time.Now().Unix()),
		TransferAmount: dto.Amount{
			Total: 1,
		},
		User: dto.User{
			OpenID:   "",
			Phone:    "15321910669",
			UserName: "李红生",
		},
		NotifyUrl: "https://api-cabinet.test.jianxindianzi.com//public/v1/callback/payment/20/Test11110001",
		Remark:    "测试1分钱转账",
		Subject:   "测试支付宝转账",
	})
	assert.NoError(t, err)
	t.Log(p)
}

// 从证书文件获取序列号（十六进制字符串，不含前缀）
func Test_GetCertSN(t *testing.T) {
	ttt := tools.Md5("CN=iTrusChina Class 2 Root CA - G3,OU=China Trust Network,O=iTrusChina,C=CN95790491425698859646659490580052089075020051350")
	t.Log(ttt)
}

func Test_GetCertSNTwo(t *testing.T) {
	certContent := `-----BEGIN CERTIFICATE-----
MIIEpDCCA4ygAwIBAgIQICURBP8Ihx94HwZwP58TYzANBgkqhkiG9w0BAQsFADCBgjELMAkGA1UE
BhMCQ04xFjAUBgNVBAoMDUFudCBGaW5hbmNpYWwxIDAeBgNVBAsMF0NlcnRpZmljYXRpb24gQXV0
aG9yaXR5MTkwNwYDVQQDDDBBbnQgRmluYW5jaWFsIENlcnRpZmljYXRpb24gQXV0aG9yaXR5IENs
YXNzIDEgUjEwHhcNMjUxMTA0MDY0MzQ0WhcNMzAxMTAzMDY0MzQ0WjBrMQswCQYDVQQGEwJDTjEw
MC4GA1UECgwn5rex5Zyz5biC5YiG55+l5pm66IO956eR5oqA5pyJ6ZmQ5YWs5Y+4MQ8wDQYDVQQL
DAZBbGlwYXkxGTAXBgNVBAMMEDIwODgwNzAyNTYyMzUwMTEwggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQC3fcEnCC2/djBnEYFDp41ZWuM4T49Pfr6p0gHXtfHGrO9CqegJhFPd4HT4BmVW
EGgsHBMSiTce4b4VsrOQG0j7Lr2It0ziPpBwJUyQti5zZUHZoNl6FDVJIUo4p1DOxoqmYqj4iK/A
gZVu0YiVNnbuKGWRoLt6B2++O05lX63qgAKxitqeRvGca73vbeiavUEhVaDVYbUBZOvLOuH/FDgL
AHETOjXm59+Q7OEdT/cV+tltNyn/zTCUbql3I8MZQQkFcjOylREm+Jg7nGrK0o/9Bjw413LuQwe4
Y0uXO4r7mUwV6QTS6L3L+8QzcBzpOUj1oC+iNM5Xe7b6/f8RssntAgMBAAGjggEqMIIBJjAfBgNV
HSMEGDAWgBRxB+IEYRbk5fJl6zEPyeD0PJrVkTAdBgNVHQ4EFgQUdCulKoYSg7NU9EAGzVXDYkAh
jnAwQAYDVR0gBDkwNzA1BgdggRwBbgEBMCowKAYIKwYBBQUHAgEWHGh0dHA6Ly9jYS5hbGlwYXku
Y29tL2Nwcy5wZGYwDgYDVR0PAQH/BAQDAgbAMDAGA1UdHwQpMCcwJaAjoCGGH2h0dHA6Ly9jYS5h
bGlwYXkuY29tL2NybDEwNy5jcmwwYAYIKwYBBQUHAQEEVDBSMCgGCCsGAQUFBzAChhxodHRwOi8v
Y2EuYWxpcGF5LmNvbS9jYTYuY2VyMCYGCCsGAQUFBzABhhpodHRwOi8vY2EuYWxpcGF5LmNvbTo4
MzQwLzANBgkqhkiG9w0BAQsFAAOCAQEA0Hc39EdulFrJI6RalXNzZHHPTZ4S8kdGFHBVQq/2rN/I
BgPjgbgdbmpzRp42G4WhtiAtqDu53mWMl/M0jZpKmFx0e5AwNVWIwPWDwb4IqrJe3eRGupsC8S/V
Bo0xwHlr3M6ggesmHYB8iCWIBUhfswVMUktNIf7TcRPB7D1QH1DlipZlpn4Yvfhl+eJcEDDxyEXY
Hk2NoYitg7OrXFcbDf3OdHRBzRmYHHz3KnUuGPwKd8Xy1KOGYxGDwd2M85vAXGDn9xiobLyewRSq
hgy0K/saXpUYmVfwOjoLArykVEBtSVg6lQHo6LbNjfCaOFSWH/kYHhbApE/NTMU0rrGZAg==
-----END CERTIFICATE-----`
	// 1. 解析 PEM 格式证书
	fmt.Println("-------------------------------------")

	block, _ := pem.Decode([]byte(certContent))
	if block == nil || block.Type != "CERTIFICATE" {
		t.Error(fmt.Errorf("无效的 PEM 证书格式"))
	}

	// 2. 解析 X.509 证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Error(fmt.Errorf("解析 X.509 证书失败: %v", err))
	}

	// 3. 获取签发机构名称（去除空格，与 Java 逻辑一致）
	issuerName := cert.Issuer.String()
	issuerName = strings.ReplaceAll(issuerName, " ", "") // 移除所有空格

	// 4. 获取序列号（转为十六进制字符串）
	serialNumber := cert.SerialNumber.Text(16)

	// 5. 拼接签发机构和序列号，计算 MD5 哈希
	combined := issuerName + serialNumber
	fmt.Println(combined)
	md5Hash := md5.Sum([]byte(combined))

	// 6. 将 MD5 哈希转为小写十六进制字符串
	t.Log(hex.EncodeToString(md5Hash[:]))
}
