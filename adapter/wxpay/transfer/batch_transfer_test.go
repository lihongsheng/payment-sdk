package transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatchTransfer_Transfer(t *testing.T) {

	// 微信小程序支付商户ID	1680469374
	//微信小程序支付应用ID	wxb3f684c3ffeeaec9
	//微信小程序支付APIV3密钥	725e9T6908A69349b27d48c78b797820
	//微信小程序支付证书序列号	2336990880F1AAB5342B6E789CA5F56DF7FFFB1B
	//微信小程序支付私钥	-----BEGIN PRIVATE KEY----- MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDAh/ZCKm/wwRDo Bi98z27gafuTqu5uz+a5vb+fE+RdmUMbxb4l9CtlftSDpyKZtAgdttcLc4DPWF2W X+Kr8VRN6eeXLy+OlidO9Ha7Gp4nDfcC99ZIUOAAUP6YAvOHXvmqTngb1D8EApul AnCzl+RlE6SCtr9n4Q+xlXiNxPeSONyIhA/S2/DVkDfBH47ch5H18MJry3/gLazI Vi1s3ysPMOgr7XHiEya02E0sBVY/6Fmh/nA7EwVkbNWU2x6dzsfniFTguPTTJMvv PooI8/koZFgGOF3Oozo+guo92OmIHt/6WXNXJUbHejti5hsjg8vNg5sWqR0t/Z6w q9JrNZLxAgMBAAECggEAbimdm3BiyrP9s3VOmLCsNZYO0AifEyK8Vw0eZqxgH7pB LtPvoBHc/t8aRBZO8vmTJ0VmOHQngPyD2DYXIeEdJtOuos/p+0EpTsEMarSpR5ly ++zJXGmCMkwl0N7nb07Ftk/d47aHNAk1+62wxOIlWjxqBi0AyjR54VewOhY4vh3x 3yTstCXbflvCN+8npDuykgFxxfbx6KWWbbzeFCA7rk7rM2IhOZ4D2D5zJBvOqagw 5XcSCMngmlvg0/JZeIXizYAOlgBF/bhRIzI0HBNLQVG4MMXCo6b/4/E7IZbn/xCD 7wEyTkDTzTGa53FdEEn1FKE7RGoAh7RUbfdGOGE0PQKBgQDzDGrbSO+Et1Vc73TL vnlIbtzyUf9WyfJm9vQ7e4tH7vVQMPx95huqVdsuydbopbhwwLHeVvehR6qp874D xYRNX58wT9XrsPx7cYgzPfQbhivarKKP5EGOkwPJPL+g5vc0TblvCUbfKG5heZCP sTw48sJvFuo0NmMXEfJ7Ow4oewKBgQDKymdfqdlPQnASq+s5Y9Eq88R+RXiE0K7b isi8+uPVJ4wNKvkBLOniMtSnQE6QTMyKh07MgeBGUz9B65Kqy/PPP9xxvjtTzjgi GWAwXptaD/1fWzJXWBIrEO3YpBNsGh3MisNUecH2t/0jgSsUCJYPNVhAMDwpoFSE /xnDAAPUgwKBgQDa3Wu8XEmUMaPlNJzwfG2rFGpSrBRLiz+GCQyWJAKgnEa8F+DH Xte64+WklI26osEch+mkVpxX17mPobaSnuMs6dboInx46b/zSaFkw31267PCD4O0 H9AJG8plBRdrRZGGwBckLi5r9nSvwlCzlN6uSa1dbD0rF27T741z+xIc0QKBgBVB /nAix+rtFf5/ExtjAUw1XYW3Fpklnw7Yj45v5m4zxRSkSpH6+VGC7pPRw+qQgmiL IpGBk9cLZvgsD6Xj110EXsF1pZZ1aaD5dAHeCP756h+S2cYaL11gWu541KhiXHlS moNCeDi6tMiCA9gHxmM1vliMNeJPMkz2yOaOG9itAoGARLOt2b0FCdBcY0Owd4pt a/w3+fmcrZkkNU+MymqOpF/JBDlQA55iF2KyWxlPRDWqh5C9mqGeHN2khwMH1cga ipXtzHuemeE9QjxcM0mmtZZ0JrHiWNrZmC4OKEV1SkYdjh+Ui36KdCWF+czIxyXb Z+tOyObCqpBZv1M/C31/QnM= -----END PRIVATE KEY-----

	var cert = &config.Cert{}
	certStr := `{"secret": "LFJEGAVFJHGVQH566654564646546546", "public_key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA53hj5a7GD5WZ4fjc9P5h\nNEeTZks2sRyrAumssHWhxSaWIlzXllIyM/RTP3eQANjK++PH8IXEQb7F6uWWZntc\n6I/b1YTFWz7Dpiv1AC9jgIwZ6duKRS+2LMLL1XekQ13rdEGkM2aL7UF9rKjWBuRo\n8SNEKQPsPeWlOkhsIHrmsRI4SkbEJ4GmUsxkvLd0yr/bqPC+TN3xeydIC3mgYvRP\n1nhXFUVtva+ROOX1PA9lf6iLz+pchuZT+om+I3/RZSvBqUkRfUCiSpf2l5gfBgId\nG1WSv+JHQOcZJyK0iynGYbWReTvU6fwAxOUq4T4SbfkmUpNFx1y0fEaYowjmC8J6\n7wIDAQAB\n-----END PUBLIC KEY-----", "public_key_id": "PUB_KEY_ID_0117303941692025110700112342004204", "cert_private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDGNMDYkcLImgUf\nVxfCP6x9tvsB2GH80929QvObgFkZivrO7Lcq1hHjEKO2aNxoSFhlD17XHiq9TY7P\nXBKRRXsbww+0kJ0yvCgVZgVIuxJ6Xl3UIiuVLUOFZK1iLclmKdQ88gEr+W44G1ee\nolLwVFYFkyOpqwU46+FGMpNO61tClCaryhEV0jWuIeaoQFRmkfvdrHW4FtfUNFpG\nPYSZfHtNxwSGIk0Ox0rGZwZUdirLRZjaqilIK3/jsr1I5JYdyoaAOfdmAaCpr0f6\n5w9hB/gwSflgAbGb73lYNJRZNgRpnX9ZUYKnk338ICk3Gw3V9HEzLgOnwTSKR36H\nglofCLqRAgMBAAECggEASzcdr5Gw3ztDkl8UjwxmwtY7brnUhYNI9FeB5pLQPRc2\nkmvWHpVZ+FcKKWCls5Uwpjks/mleQrQ564q/KQ266BAf15/BZ81rUKOEgdMr8e/5\nn1TQQS4KH4bTHRzO/swE1JPpyew/3V7S40oqOzVl53us3ugddTpeJKtwp1Q9L4gq\nJJ0FJ4PUHr6bNQ116LVGzhxDxT4qizn48j7d1JUyEjG33/wOgDfGWUSefH2Q2ZwL\ntggXHPJP6dBcI3UUXbSQoIds2ipMYqmc5CMhy21xObTvo9nKCUI/xts4viJ2/zK6\n4YKx6ZCoKGYec/TMu/qkr++7xTu+qvnDllME8Ik3gQKBgQDjATRb3CqCcWFB6WRh\nRFVufCNpyh5dzFKyKDugoSgBED4nS4uE8QZH36dbcOyMNVEtHuKSamGXCi2x09Nf\nsQOLQGPup67q9vv56NBNkibXATaWas7pbmT9B4MW6Bv9ByBSgS56AHSnWyjOD650\nC6/Fa/auaGo3r4RUptTIY+O8uQKBgQDfhd25S8L2Dpa/13QuXGdBPF1cspkj5puD\nzKY3MEZo6Y/swSrRdb18MIXSCLJY8sCB3SH5rG1HkcVhOY7gtoKw1iRAMEf6babF\npJDG+D8k223KgKkzKHP5Kg/NrRm6nnuBvikNGNEgCAWtBybDKX1hqKke4z0CA7Sf\nkqrb/jxwmQKBgD58fmUq3ai6fQMfs7nyjXG0Sis8r88yBzFzUbaNpe1lAzbd3LHj\nhs8SCYdqNjMCGi5JaiTTk7l328wveufEWi1itB9lmQikpAfOxkgUCwz0EIqnK/2l\nnbbo8nTDv7CO3Z7YYGrE5VeMCFdwiZz3+pJlfanUpChf8BU9NyVSGcZBAoGAK3xR\nrJDutwwTi/MQqUxU0j46M6STYoakzrlrxOThbduyom7aM7HiUVznS/thJyjjBuDM\nkVRYVkonykh2YYVgW6LtnodGGZRnk5/2gp8dOcBu1ay+PjOqjFkAhhUdIk9e29jx\nB5lCZibpY8Y2ZlWWDP/RFy9CWTf7Vegk0XPeslECgYAEbMrGRw72obPg8eWXqsTr\nEJ5OghNcC2Xm8ZWy3onYqHPq6qUzIuJf98zuR+fyoRxbfNyrYF3LYeylSTjMSSzw\nqHdSMwxgar44Cvw02evWizMgXCO5bkigeBIpGwTTx551LfGUsivbStUOyFxNNjmU\nOODeiuJSy7DEFpDhlTvupQ==\n-----END PRIVATE KEY-----", "certificate_serial_number": "3DF1CB427C4D3F65E4F92D03EB42A1A697F6D010"}`
	json.Unmarshal([]byte(certStr), &cert)
	cccc := config.Config{
		AppID: "wxb3f684c3ffeeaec9",
		MchID: "1680469374",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY----- 
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDAh/ZCKm/wwRDo Bi98z27gafuTqu5uz+a5vb+fE+RdmUMbxb4l9CtlftSDpyKZtAgdttcLc4DPWF2W X+Kr8VRN6eeXLy+OlidO9Ha7Gp4nDfcC99ZIUOAAUP6YAvOHXvmqTngb1D8EApul AnCzl+RlE6SCtr9n4Q+xlXiNxPeSONyIhA/S2/DVkDfBH47ch5H18MJry3/gLazI Vi1s3ysPMOgr7XHiEya02E0sBVY/6Fmh/nA7EwVkbNWU2x6dzsfniFTguPTTJMvv PooI8/koZFgGOF3Oozo+guo92OmIHt/6WXNXJUbHejti5hsjg8vNg5sWqR0t/Z6w q9JrNZLxAgMBAAECggEAbimdm3BiyrP9s3VOmLCsNZYO0AifEyK8Vw0eZqxgH7pB LtPvoBHc/t8aRBZO8vmTJ0VmOHQngPyD2DYXIeEdJtOuos/p+0EpTsEMarSpR5ly ++zJXGmCMkwl0N7nb07Ftk/d47aHNAk1+62wxOIlWjxqBi0AyjR54VewOhY4vh3x 3yTstCXbflvCN+8npDuykgFxxfbx6KWWbbzeFCA7rk7rM2IhOZ4D2D5zJBvOqagw 5XcSCMngmlvg0/JZeIXizYAOlgBF/bhRIzI0HBNLQVG4MMXCo6b/4/E7IZbn/xCD 7wEyTkDTzTGa53FdEEn1FKE7RGoAh7RUbfdGOGE0PQKBgQDzDGrbSO+Et1Vc73TL vnlIbtzyUf9WyfJm9vQ7e4tH7vVQMPx95huqVdsuydbopbhwwLHeVvehR6qp874D xYRNX58wT9XrsPx7cYgzPfQbhivarKKP5EGOkwPJPL+g5vc0TblvCUbfKG5heZCP sTw48sJvFuo0NmMXEfJ7Ow4oewKBgQDKymdfqdlPQnASq+s5Y9Eq88R+RXiE0K7b isi8+uPVJ4wNKvkBLOniMtSnQE6QTMyKh07MgeBGUz9B65Kqy/PPP9xxvjtTzjgi GWAwXptaD/1fWzJXWBIrEO3YpBNsGh3MisNUecH2t/0jgSsUCJYPNVhAMDwpoFSE /xnDAAPUgwKBgQDa3Wu8XEmUMaPlNJzwfG2rFGpSrBRLiz+GCQyWJAKgnEa8F+DH Xte64+WklI26osEch+mkVpxX17mPobaSnuMs6dboInx46b/zSaFkw31267PCD4O0 H9AJG8plBRdrRZGGwBckLi5r9nSvwlCzlN6uSa1dbD0rF27T741z+xIc0QKBgBVB /nAix+rtFf5/ExtjAUw1XYW3Fpklnw7Yj45v5m4zxRSkSpH6+VGC7pPRw+qQgmiL IpGBk9cLZvgsD6Xj110EXsF1pZZ1aaD5dAHeCP756h+S2cYaL11gWu541KhiXHlS moNCeDi6tMiCA9gHxmM1vliMNeJPMkz2yOaOG9itAoGARLOt2b0FCdBcY0Owd4pt a/w3+fmcrZkkNU+MymqOpF/JBDlQA55iF2KyWxlPRDWqh5C9mqGeHN2khwMH1cga ipXtzHuemeE9QjxcM0mmtZZ0JrHiWNrZmC4OKEV1SkYdjh+Ui36KdCWF+czIxyXb Z+tOyObCqpBZv1M/C31/QnM=
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: `2336990880F1AAB5342B6E789CA5F56DF7FFFB1B`,
			PublicKeyID:             "1efa732737927e914ba34ac3bb8fd911150461cc",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7H7HjP7os+Dl74bzKv5J
UDxwjkFRCk2N+wtvN5c58+GafOm2yVVP6AIDmjUcdRZjgjU8Pa3jRycSvIQg6waK
Cc42G/njLth9mq0n9l4FgUVjMJxcyV1vbwQ1s+WvIvhIUOgPWVmLE/lohJNsFdBy
X9uVfh2PJfLIwNrkiE7N6AhDWQ09gQV8p2cFrvel/s37xTnIuRsaIX46zcJZCZtY
beAjoHoziL4a4NDhQfqF3FbmgHRxdqITyslnFryANEpTn9fp9yY1KrHjwWx85pBB
8mK3D2pk/sRQgzWADzLtyAosSGcDlTSBqxAY7V8OCGUHaTIF1g/kDFRng+G6KFqq
EQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "725e9T6908A69349b27d48c78b797820",
		Proxy:  config.Proxy{},
	}
	t.Log(cert.CertPrivateKey)
	client, err := NewTransfer(cccc)
	assert.NoError(t, err)
	ctx := context.Background()
	req := &dto.BatchTransferRequest{
		TransferNO: "200000020251110170368514500",
		Subject:    "测试",
		Remark:     "测试",
		Amount: dto.Amount{
			Total: 1,
		},
		SceneID:   "1000",
		NotifyUrl: "https://backend.sslab.saburobox.com/public/v1/callback/batch/transfer/200000020251110170368514500",
		Extend:    nil,
		Details: []dto.BatchTransferDetailRequest{
			{
				UserID:  "olymevsNYgJ9UxG1XRtQrlecKvsE",
				Amount:  dto.Amount{Total: 1},
				TradeNO: "200000020251110170368514500",
				Remark:  "ces",
			},
		},
		SubAccountIn: "",
	}
	resp, err := client.Transfer(ctx, req)
	t.Log(fmt.Sprintf("%+v", resp))
	if err != nil {
		t.Log(err.Error())
	}
}
