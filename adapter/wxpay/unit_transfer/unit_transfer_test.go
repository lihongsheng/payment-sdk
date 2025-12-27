package unit_transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransfer_Transfer(t *testing.T) {
	var cert = &config.Cert{}
	certStr := `{"secret": "LFJEGAVFJHGVQH566654564646546546", "public_key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA53hj5a7GD5WZ4fjc9P5h\nNEeTZks2sRyrAumssHWhxSaWIlzXllIyM/RTP3eQANjK++PH8IXEQb7F6uWWZntc\n6I/b1YTFWz7Dpiv1AC9jgIwZ6duKRS+2LMLL1XekQ13rdEGkM2aL7UF9rKjWBuRo\n8SNEKQPsPeWlOkhsIHrmsRI4SkbEJ4GmUsxkvLd0yr/bqPC+TN3xeydIC3mgYvRP\n1nhXFUVtva+ROOX1PA9lf6iLz+pchuZT+om+I3/RZSvBqUkRfUCiSpf2l5gfBgId\nG1WSv+JHQOcZJyK0iynGYbWReTvU6fwAxOUq4T4SbfkmUpNFx1y0fEaYowjmC8J6\n7wIDAQAB\n-----END PUBLIC KEY-----", "public_key_id": "PUB_KEY_ID_0117303941692025110700112342004204", "cert_private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDGNMDYkcLImgUf\nVxfCP6x9tvsB2GH80929QvObgFkZivrO7Lcq1hHjEKO2aNxoSFhlD17XHiq9TY7P\nXBKRRXsbww+0kJ0yvCgVZgVIuxJ6Xl3UIiuVLUOFZK1iLclmKdQ88gEr+W44G1ee\nolLwVFYFkyOpqwU46+FGMpNO61tClCaryhEV0jWuIeaoQFRmkfvdrHW4FtfUNFpG\nPYSZfHtNxwSGIk0Ox0rGZwZUdirLRZjaqilIK3/jsr1I5JYdyoaAOfdmAaCpr0f6\n5w9hB/gwSflgAbGb73lYNJRZNgRpnX9ZUYKnk338ICk3Gw3V9HEzLgOnwTSKR36H\nglofCLqRAgMBAAECggEASzcdr5Gw3ztDkl8UjwxmwtY7brnUhYNI9FeB5pLQPRc2\nkmvWHpVZ+FcKKWCls5Uwpjks/mleQrQ564q/KQ266BAf15/BZ81rUKOEgdMr8e/5\nn1TQQS4KH4bTHRzO/swE1JPpyew/3V7S40oqOzVl53us3ugddTpeJKtwp1Q9L4gq\nJJ0FJ4PUHr6bNQ116LVGzhxDxT4qizn48j7d1JUyEjG33/wOgDfGWUSefH2Q2ZwL\ntggXHPJP6dBcI3UUXbSQoIds2ipMYqmc5CMhy21xObTvo9nKCUI/xts4viJ2/zK6\n4YKx6ZCoKGYec/TMu/qkr++7xTu+qvnDllME8Ik3gQKBgQDjATRb3CqCcWFB6WRh\nRFVufCNpyh5dzFKyKDugoSgBED4nS4uE8QZH36dbcOyMNVEtHuKSamGXCi2x09Nf\nsQOLQGPup67q9vv56NBNkibXATaWas7pbmT9B4MW6Bv9ByBSgS56AHSnWyjOD650\nC6/Fa/auaGo3r4RUptTIY+O8uQKBgQDfhd25S8L2Dpa/13QuXGdBPF1cspkj5puD\nzKY3MEZo6Y/swSrRdb18MIXSCLJY8sCB3SH5rG1HkcVhOY7gtoKw1iRAMEf6babF\npJDG+D8k223KgKkzKHP5Kg/NrRm6nnuBvikNGNEgCAWtBybDKX1hqKke4z0CA7Sf\nkqrb/jxwmQKBgD58fmUq3ai6fQMfs7nyjXG0Sis8r88yBzFzUbaNpe1lAzbd3LHj\nhs8SCYdqNjMCGi5JaiTTk7l328wveufEWi1itB9lmQikpAfOxkgUCwz0EIqnK/2l\nnbbo8nTDv7CO3Z7YYGrE5VeMCFdwiZz3+pJlfanUpChf8BU9NyVSGcZBAoGAK3xR\nrJDutwwTi/MQqUxU0j46M6STYoakzrlrxOThbduyom7aM7HiUVznS/thJyjjBuDM\nkVRYVkonykh2YYVgW6LtnodGGZRnk5/2gp8dOcBu1ay+PjOqjFkAhhUdIk9e29jx\nB5lCZibpY8Y2ZlWWDP/RFy9CWTf7Vegk0XPeslECgYAEbMrGRw72obPg8eWXqsTr\nEJ5OghNcC2Xm8ZWy3onYqHPq6qUzIuJf98zuR+fyoRxbfNyrYF3LYeylSTjMSSzw\nqHdSMwxgar44Cvw02evWizMgXCO5bkigeBIpGwTTx551LfGUsivbStUOyFxNNjmU\nOODeiuJSy7DEFpDhlTvupQ==\n-----END PRIVATE KEY-----", "certificate_serial_number": "3DF1CB427C4D3F65E4F92D03EB42A1A697F6D010"}`
	json.Unmarshal([]byte(certStr), &cert)
	cccc := config.Config{
		AppID: "wx8cf29dc5f2fb40fe",
		MchID: "1730394169",
		Cert: config.Cert{
			CertPrivateKey:          cert.CertPrivateKey,
			CertificateSerialNumber: cert.CertificateSerialNumber,
			PublicKeyID:             cert.PublicKeyID,
			PublicKey:               cert.PublicKey,
		},
		APIKey: "LFJEGAVFJHGVQH566654564646546546",
		Proxy:  config.Proxy{},
	}
	c, err := NewTransfer(cccc)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	req := &dto.UintTransferRequest{
		NotifyUrl: "https://www.baidu.com",
		Remark:    "测试",
		SceneId:   "1000",
		User: dto.User{
			UnionID:  "",
			OpenID:   "olymevsNYgJ9UxG1XRtQrlecKvsE",
			UserName: "",
		},
		TransferAmount: dto.Amount{
			Total:    0,
			Currency: "",
		},
		SceneReport: []dto.SceneReport{
			{
				Type:    "USER_ID",
				Content: "123",
			},
		},
	}
	result, err := c.Transfer(ctx, req)
	assert.NoError(t, err)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(fmt.Sprintf("%+v", result))
}
