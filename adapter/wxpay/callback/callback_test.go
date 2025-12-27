package callback

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/stretchr/testify/assert"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestCallbackRefundParse(t *testing.T) {
	h := http.Request{}
	u, _ := url.Parse("https://api-cabinet.test.jianxindianzi.com/callback/v1/tenant/payment/5/175870989522000055545")
	h.URL = u
	body := `{"id":"4b652507-f4e0-5400-9bd4-160b351fc535","create_time":"2025-09-25T22:41:16+08:00","resource_type":"encrypt-resource","event_type":"REFUND.SUCCESS","summary":"退款成功","resource":{"original_type":"refund","algorithm":"AEAD_AES_256_GCM","ciphertext":"f77TZCTe+XW8oS4A3Ww/xMmQsfrca9Ic4b623TpiJJd7texdsUElebrpprlTvnygAMS0cJIICXoy9bh83zdhSBPDPYT14GItj9f+KoxwVYmlYcMfpM/2ftn7jOa2xJOGdn77Iy4JfmmeFkyO9ThAwbaSl3biNQnX6crMqYrT2t3HTTmBR/m6G4SpPBiiehqXcInyRsS4+1+hvFAHrDO2ZxxCKt9GOejwTG1/GuaCMe+XqxFmKLmXatDVO/tgDSXK8SiTnQjEUyHAMxAk09SFgVInAXFIzJ2x2XuH+uL6TMhJ+BV3cU2w+ULo9RqdTwQfHlybQxnajHZtb4a9aQK1przAYqWzaMLut11BzFdOlFZUbmSSUmeJSjX6GRR4sdBepO9gM/cS2O7c3j46W6L7S52twYAeu+846tI837f/4HQfc4IU95PE0yOHDcN6ygtHwc6r1yi+lu2sSt6Xls1OD52gX6bVgztCbYKlEvizT+QRBgWebnFa","associated_data":"refund","nonce":"Mw40GMRVWziD"}}`
	h.Header = http.Header{
		"Wechatpay-Signature-Type": []string{"WECHATPAY2-SHA256-RSA2048"},
		"Wechatpay-Nonce":          []string{"4BmaCwj5dHJbwZc656V2NSzZIoshrZ7s"},
		"Wechatpay-Timestamp":      []string{"1758847410"},
		"Wechatpay-Serial":         []string{"PUB_KEY_ID_0117134125362025092400111793000600"},
		"Content-Length":           []string{fmt.Sprintf("%d", len(body))},
		"Wechatpay-Signature":      []string{"DTdzqVmDfNWSPeNkQTG1RMa1luSLYycMh26GCWe0nTASuIJdDObf4ES1TCphi++Unv35amw7IdyL3Vf/Iaytcz3R0B6yrBv5DLRri6hQWmG9EFhNKp5Hd0XvdvLebI67NIwKTXa3S5QHpxlQAJFw+3pNx+qjHmLqufx/afp//Xj8fUxvL3IM2XXDYad+Gkcmlo/+zCaeavXJR83LzYXKDgvD8yl0gKV1OwHhPUykrYJARdEa28Vf2QjkYaAnwzfKCdh2gO7i1kgUy/zsG2xsypNrFjJ72/BET94CjPhF0CB3vISAX6mrlo3/Oz9jnKzvMG9m3Fff9b6nLF26eNPNtQ=="},
	}

	cccc := config.Config{
		AppID: "wx160b7dc2f3438a1b",
		MchID: "1713412536",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/nNSOAa3HZ4LrDjM4vwgQWYbxq2Pzt132lp8EUowLO/LM3sX36hwwO0GlOj4qd3Fm9hJy/9txJRg7h1QB0a0j4Y8FYUvAmUAS20vQztAAGYde6JpqvNQ6cl5DDDWit+L3s4ZtEIzpvssMadtsc4C+a7AE7F807aWFdsWnux01Xl4lDLtC2rVh5u+aN0Nl06mKXnniRmfjlfZtclg8BrpowTys+fEUlZFD5xbU2gVtLt4RhFtZVPjQnfZ10GPMAz35DRoOmDu3ZXKk0wOUErBtqd4WeJ23f/pPAB5m++9H+E/BG/Gi689iFLFaffvvJEwZkOqV0viLZdrd2bH8E/FjAgMBAAECggEAe7D/vVP0HF8Dsj0Ob7lRuUwxwlwDP9bE/2On7yBiavYd/IZqgWlNHQ2DiOeaLcvBFtgOfNIRlG5/wB3R6wKxpBH9Q1nVjtTe+c06meaHeyj/rBK3a+PNlJUzqFB/ZzURfRkU0971OAcECFVlYhMFrubRT7xOkVo/mXJckbRGXKYY2/xe0G4UkQJM6F3KdCjJdtSXv7zYfqIISFHSKm9vdKs8Q0PUy+6Rt9gljWvrz/THmKQktnhZKf7bZSmDGXef1a3ZW8sdj44PUA67bPf139SK/5I7j27qpiWfMor7+5BCIuQIlbhTknpHvCOwABHFrCLg749KVev5bsTFLDENEQKBgQD1zlqGjCJWiWnogwkqmH5dphNxCk60AzKrXtHnDYLZX4cYQaiVmDvvB0T0CUTNzVyoTmRDoLwdtrjIokqM7RAXoKxZ7G46mI5x8IgLMcETq0KIGSaENj2DxM6coevgwtNJk8FAOG0+3skpbiXhTuEacgjnJ5WqW9BbyEZHHUDIfQKBgQDHjx8oeTX8ggMfKHIhjWa890ToQyE/mcq89j1uJHj/7r47to6nuDZCxXmSmVNAlKvcSxvAeMyn6x6k+Srmq6sABX29yatPLR90OFQlnLdTvEDPK1BpGpGWX+HRNVKUc/U4y0zmSt2VfoDgaaemzvd1lc9pFlnnXpZiQIms60CnXwKBgEe7KVW8TUT9osd0fddNWwsPLPs+68rCaCX0bMLFgZrXsr/UYVMOcucFMw0YK1j3hgOjpMTLgjoVmYULP0Ay6hBLFiDDy0MUQ/ViIQFLSrHnt2mqFUBd58OtSjIRWpljoW8GTE3maZMARqntd+ZxM2WZQ5nZRmbJlltCbafRFJetAoGBAIToTVgnYk1KScn2pgyyoDo6dSo7i2lQhDZVyZQRtoS9/PTIITqS9ZCC9PUuKMRaQBv36gPGcIdlkINPb8MxkjHxdk1wgye4ZbqByYlDVtXuCzvvHR7jExOTyFINsXItyKSKwiyer/Vgy3Sq6X2vWiB2Ji1XNYli9cV6Njd0dxsBAoGBAMUokGtzsFtUfkbItB3wK62sCwdT3kkvOJ+8mhXpWFE5wvVJzakCRE4pLX7V5+Ne8aWE8M40UX3KQHCEZyXovTvyuERk8oyTtaox+3RNXRSWZfWZBp0iD+yeh0tLvKBWDxipO+RbDP99edd1zzGUHw/MljWmsTmKtjfgjwHnWi49
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "59400142A58A56A0AF191C85A2C91CF434AEF32A",
			PublicKeyID:             "PUB_KEY_ID_0117134125362025092400111793000600",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6gb565rrPWBa8b+pYNY/JkYQLdkIPN4h71UsTBDD3tuua9x14s9Lk0NFCPi5khMksXUxTuYCW1vB/DcKCpwZOimAvj96sjBH8kO4JZrblTuDpRlNJ6vlL8NNMT9DPvb8ylTEJI6wmU3LfAhR9I8o67wwPRpT/uUAW6Ikz1fULhKwapCWF3oYp6JDiV5eBNatLzKnLGoW/xa4Guu8wlOjcUjs69JibJePn5PTWpFj5F9dZRYsfFYasrc4GKX547kEqSlJfHyTjOm0HiXvD/MLhtxzDudcvrxJOX7bCe4wb874JGRwAnBrbAwrxTwA1BwFdy0s8yF+JeCe4h6peWX2WQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "1bde81c041fbaPF2K0d21ccAa0147c8f",
		Proxy:  config.Proxy{},
	}
	h.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(body)))
	ctx := context.Background()
	rrrr, err := CallbackRefundParse(ctx, cccc, &h)
	if err != nil {
		t.Log(err.Error())
	}
	assert.Nil(t, err)

	t.Log(rrrr)
}

func TestCallbackPaymentParse(t *testing.T) {
	apiV3Key := "dajowidao23127821bkb3i213123uiib"
	c, err := aes.NewCipher([]byte(apiV3Key))
	if err != nil {
		t.Log(err.Error())
	}
	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		t.Log(err.Error())
	}
	noce := "wH6I3SmUqHWX"
	ciphertext := "xDqz9ebZ7FduE7Gcsaup7dstXkzFJ+59g0ox6c+WlQv7Mi02LrSsaFyZuwiSWO8kRsuMsjYMsGtH1dpRDOerG2nDDjR60XMSh3Ls3UMo6Iwp+YkvKw5mOERgECjx+sEmLr2PVHo7HQvTrAEUlG5Omm5UeySUYeNKWxAl9WoaX67TYqpYLQHp4pHCPsxDihEO9oitb12NXAwozuFWRgcK2QaTta1tE0zjMRruGxokbLigvrkOsh7atNKjfDaK8sHBuRrPNjXUzVCvRJqMXgkUyXvbRGex0bUdOHR8HmVpccbCnjx+BzuopRUZjZHT8QIsKloBKwGaWL7i6+65Z0NrIvKODhdgFwAguNS5hj5xIA3QF35hy0vHj+npHMlYXi8ctoMGy93vWrwgLeX+OIAjaCYpFPCGYbqvI3GC6ZG8vqRCmr0h8X+AWzKHj9qj1VS/It0OCSQwBp5MxvczV0vDCWD+GyseejsMCLeYNcHkpuw8AaeLclvlfd1BvNNwMVba723jiUPzsLr8jPtSDn2lelgynbG1jt6zEppBsICrS0lNLifkIgKw6eNPTnYbehflzVqdLrTqcu44NS8EPsngfSHtUeQ="
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	associated_data := "transaction"
	if err != nil {
		t.Log(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, []byte(noce), data, []byte(associated_data))

	fmt.Println(string(plaintext))

}

func TestProcessBody(t *testing.T) {
	body := `{"id":"b0cd41a4-8510-5c16-86e2-cf3276dc9470","create_time":"2025-11-19T10:56:02+08:00","resource_type":"encrypt-resource","event_type":"TRANSACTION.SUCCESS","summary":"支付成功","resource":{"original_type":"transaction","algorithm":"AEAD_AES_256_GCM","ciphertext":"xDqz9ebZ7FduE7Gcsaup7dstXkzFJ+59g0ox6c+WlQv7Mi02LrSsaFyZuwiSWO8kRsuMsjYMsGtH1dpRDOerG2nDDjR60XMSh3Ls3UMo6Iwp+YkvKw5mOERgECjx+sEmLr2PVHo7HQvTrAEUlG5Omm5UeySUYeNKWxAl9WoaX67TYqpYLQHp4pHCPsxDihEO9oitb12NXAwozuFWRgcK2QaTta1tE0zjMRruGxokbLigvrkOsh7atNKjfDaK8sHBuRrPNjXUzVCvRJqMXgkUyXvbRGex0bUdOHR8HmVpccbCnjx+BzuopRUZjZHT8QIsKloBKwGaWL7i6+65Z0NrIvKODhdgFwAguNS5hj5xIA3QF35hy0vHj+npHMlYXi8ctoMGy93vWrwgLeX+OIAjaCYpFPCGYbqvI3GC6ZG8vqRCmr0h8X+AWzKHj9qj1VS/It0OCSQwBp5MxvczV0vDCWD+GyseejsMCLeYNcHkpuw8AaeLclvlfd1BvNNwMVba723jiUPzsLr8jPtSDn2lelgynbG1jt6zEppBsICrS0lNLifkIgKw6eNPTnYbehflzVqdLrTqcu44NS8EPsngfSHtUeQ=","associated_data":"transaction","nonce":"wH6I3SmUqHWX"}}`
	httpReq := &http.Request{Body: ioutil.NopCloser(bytes.NewBuffer([]byte(body)))}
	mchAPIv3Key := "dajowidao23127821bkb3i213123uiib"
	var resp = &payments.Transaction{}
	_, proErr := ProcessBody(mchAPIv3Key, httpReq, resp)
	if proErr != nil {
		t.Log(proErr.Error())
	}
	if resp.TransactionId == nil || (resp.TransactionId != nil && *resp.TransactionId == "") {
		t.Log("wxpay parse notify not find TransactionId")
	} else {
		t.Log(*resp.TransactionId)
	}
}

func TestCallbackBatchTransferParseParse(t *testing.T) {
	h := http.Request{}
	u, _ := url.Parse("https://api-cabinet.test.jianxindianzi.com/callback/v1/tenant/payment/5/175870989522000055545")
	h.URL = u
	body := `{"create_time":"2025-12-10T18:26:09+08:00","event_type":"MCHTRANSFER.BATCH.FINISHED","id":"3fe8147e-2f4a-5df4-b567-10438d787d7c","resource":{"algorithm":"AEAD_AES_256_GCM","associated_data":"mch_payment","ciphertext":"lZc+49R6ltxTWNH4PStYmYw7gMyQXQRCK18hy6IT4WOp3zO1MS4Sc2qKY2cU4BAH+kjG5ePxXvv2u0vyavT7x5NskN+YRwZxihFW2Rtbf0SIpG1Q6zg794xCJ1CnvMPbGudr3t/VH22GuY4WujsfAaL7kRaNgZtNdOG9sjbgOT3LiqoXIgTVaJ/oSgZ8YdDbn1GPWABS06y1r8ud8H0rUsRANO9O7qGvgAfbFQgEMZ9vMhNzcOisbq46sdZQA5z3NgFHQAjiSHkW2v7EKCn1wx0IOYayjfDtCtoFwwAt2E7pf+PFGoN6yagSn92vuhoqX642JqIoNilZAHuCBYdovPe2vDEIYlYNziIAUcrqlRRsCNYmloIHqVeKFBukDX7bLKsT+1g9yoKCpth4HzBkVPE=","nonce":"G331CyBcibi4","original_type":"mch_payment"},"resource_type":"encrypt-resource","summary":"商家转账批次完成通知"}`
	h.Header = http.Header{
		"Wechatpay-Signature-Type": []string{"WECHATPAY2-SHA256-RSA2048"},
		"Wechatpay-Nonce":          []string{"4BmaCwj5dHJbwZc656V2NSzZIoshrZ7s"},
		"Wechatpay-Timestamp":      []string{"1758847410"},
		"Wechatpay-Serial":         []string{"PUB_KEY_ID_0117134125362025092400111793000600"},
		"Content-Length":           []string{fmt.Sprintf("%d", len(body))},
		"Wechatpay-Signature":      []string{"DTdzqVmDfNWSPeNkQTG1RMa1luSLYycMh26GCWe0nTASuIJdDObf4ES1TCphi++Unv35amw7IdyL3Vf/Iaytcz3R0B6yrBv5DLRri6hQWmG9EFhNKp5Hd0XvdvLebI67NIwKTXa3S5QHpxlQAJFw+3pNx+qjHmLqufx/afp//Xj8fUxvL3IM2XXDYad+Gkcmlo/+zCaeavXJR83LzYXKDgvD8yl0gKV1OwHhPUykrYJARdEa28Vf2QjkYaAnwzfKCdh2gO7i1kgUy/zsG2xsypNrFjJ72/BET94CjPhF0CB3vISAX6mrlo3/Oz9jnKzvMG9m3Fff9b6nLF26eNPNtQ=="},
	}

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
			PublicKeyID:             "1EFA732737927E914BA34AC3BB8FD911150461CC",
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
	h.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(body)))
	ctx := context.Background()
	rrrr, err := CallbackBatchTransferParse(ctx, cccc, &h)
	if err != nil {
		t.Log(err.Error())
	}
	assert.Nil(t, err)

	t.Log(rrrr)
}
