package lakala

import (
	"bytes"
	"context"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestExtractLKLAuthParams(t *testing.T) {
	auth := `LKLAPI-SHA256withRSA timestamp="1765848504",nonce_str="bGfguYYXFzbO",signature="I1UJq1xfo+BEqKMsd5hzqxwrHUOFUDfp6IYs3Y3FBVgYwbOm32Z6tr0yD8yWfBkUbi+z6DSH34mrGk9eVb6shr8sy1kNX2dlf8U6l1UeYDTX+a4Irf6HYu9QlTUhB+dpROFR3bxecK20NPYsazZD8dk7XMU4ic/kpgxbXN/M0g9OL4LLgHJc4ac4TpXv/jcOzUbii/l0emBPuYqMfs7hGm8p6W//eiXPs1BynkT7Vv5WHwOhDxlO9fr7qdskg0NiacFoFXL0t42pYQnFYeTRezH6BY95Qk7lOVGO8GU13TVU+s87CxHADQMzoZ2V585FWGuM0lIS5RfhUU7XnAMp6A=="`
	timestamp, nonceStr, signature, exists := ExtractLKLAuthParams(auth)
	assert.Equal(t, exists, true)
	t.Log(timestamp, nonceStr, signature)
}

// {"acc_discount_amount":"0","acc_mdiscount_amount":"0","acc_settle_amount":"30","acc_trade_no":"4200002968202512168521236797","account_type":"WECHAT","bank_type":"CITIC_DEBIT","card_type":"00","gb_amount":"","log_no":"66202326851350","merchant_no":"822584073920DEH","notify_url":"https://proxy.test.jianxindianzi.com/public/v1/callback/payment/37/176584849551401225676","out_trade_no":"176584849551401225676","payer_amount":"30","promotion_detail":"","qb_amount":"","remark":"","sub_mch_id":"843442957","total_amount":"30","trade_no":"20251216110113130266202326851350","trade_req_date":"20251216","trade_state":"SUCCESS","trade_status":"SUCCESS","trade_time":"20251216092824","user_id1":"olymevsNYgJ9UxG1XRtQrlecKvsE","user_id2":"oVxsc1fhZTVRe2mAiN7dLMFRiO78"}

func TestAPICallback_CallbackPaymentParse(t *testing.T) {
	c := config.Config{
		AppID:  "OP10002115",
		MchID:  "822584073920DEH",
		APIKey: "",
		Cert: config.Cert{
			CertificateSerialNumber: "00dfba8194c41b84cf",
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDbxyDQUdBATtMM
+/CDqATnRuSPfpsGKWU9wVGon34q0xNB+vuz+4O3IXATAKEpZZLj6MMPFSmtWaUV
QVhLOqGiHMZyN028XwEuuyyM2f9kr84wm1YsyuWzAoRZkjQ/hkVCPCeZd78u0WJ0
YeIhRsjGcUcuh7O59gN2lRDzV5UVWf8LTIKXfwDKsCJX1Epex/+cyxzYEynPB1nL
QIAQ1OVCZVbxN7m+Fw/TLhaGJd5Tx9IGlSTkVcIKts0KA8OY3Ro0EqtqeHoZ8eq0
MGpouxmh3XzqVjyl+7EZt+IoTZ10R/Ed3XYONF9JQvDducHKOX4P8U7I0ym/Yb0Q
4+6YYBgRAgMBAAECggEAED2v3BUfqZDpob0Acgo5iom/nCcD97mZZK3jhe17WljM
xIRyk0NT4XWUHaNfRXrfFv59Y6DxuoC0ZVS13KFRjnRH6erSUMhIgxaL3UDC0cL4
Hrlr7dV0kfzuoNvgBo26koF1f67Mrv4EI4uUNVdQwPFgDD0099oJOXscjI79Ul14
H7SksF5d7dk2Mx4s3fjZV6UgNaGELZW5WVpRiufu+jMfvCsE+4/5Tp8Sco5zK+C2
kAD/3YGOSAc26vQxivO7Yg4worCcYUfltrIX3SQzG2yBpnbK2ZQYcWZY374qp0Xp
leRlXWLJnhUL8PpR6NPZq0Us2SUQxMbW0QaePjTGAQKBgQDuZltdmGrwYD5dsxho
HkJ7kMFwjWIfoSFXazgU6IgLlYWef8GFdwKmmlkLIT4Lk/m1b8/1Uw1qVK19ClWG
6660iBB6IBA5kxa7gU99ZQed2/SYeVpLHIw0ITJzEX29y3OkyxsaRscPmNUHXBgj
u7R7i5eAtO2xU52o18FiOHRGsQKBgQDsANMZa2KZGY+8snUofrMeAGJGEuLYAo8B
a2pSCYK4eOhJhyWw1mKLcZVlzdZyXlZ3eEsWKweMOQyyyOl3cupLyCcNWjMEv4Ha
AqC5drh99dI2KgXdUybAf9p4i9dpY42URpo5boI+W6j0jYLOXN7YYaFRatr0s+Mu
m3vriBf/YQKBgELwlsMHIy/vtlNVEItbw8sycD6MVHsRIW2Me6jTSjAGgghpUwuI
yUPCnzIS2Xsix8D8bmYyNdgfgr9TgYRq9RlYA1hnXGbuODnaK1nIXoUi1+FgYcwp
bezNTX8l8Cq0z/n71dZg/VAR1+9DGrwd3qW6IoZPR1a9Zc2dF33e4DdhAoGBAJ1s
n2PhYc/GYT75u3TbrxdgIi2kA3Ubn9DOmglHFs9+t1P0touTNgDWL1XNTDLWAs+G
im+rHEnI9FN9+V4YZXlPdd1OQaH1LOUDw7pzGvXKuAIxXeAYy0y0/EJU5cgDBDnY
LqAIuxBli/o1Ov/0qyGjXjw1DwETzYMVbD/cdEWBAoGAcngSWwC6rKE7nyfJdbhW
jUlIqljUFYTHWCt+BKabaa05/tiwM4tP/CGi9Ik0NH0HeENQBj/NBuGFk07ezYob
qJxdIULDA+7FULTMxGEdYMkNxms9NagmSJkeJgDANSiD4/eYsomKrRW+uwdrSOF4
UBw+2oCczJC4S01uBmg76/g=
-----END PRIVATE KEY-----`,
			PublicKey: `-----BEGIN CERTIFICATE-----
MIIEMTCCAxmgAwIBAgIGAXUrc4b4MA0GCSqGSIb3DQEBCwUAMHYxCzAJBgNVBAYT
AkNOMRAwDgYDVQQIDAdCZWlKaW5nMRAwDgYDVQQHDAdCZWlKaW5nMRcwFQYDVQQK
DA5MYWthbGEgQ28uLEx0ZDEqMCgGA1UEAwwhTGFrYWxhIE9yZ2FuaXphdGlvbiBW
YWxpZGF0aW9uIENBMB4XDTIwMTAxNTA4NDk1MloXDTMwMTAxMzA4NDk1MlowZTEL
MAkGA1UEBhMCQ04xEDAOBgNVBAgMB0JlaUppbmcxEDAOBgNVBAcMB0JlaUppbmcx
FzAVBgNVBAoMDkxha2FsYSBDby4sTHRkMRkwFwYDVQQDDBBBUElHVy5MQUtBTEEu
Q09NMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwAXZw9lupWcFXouC
Nhm0DQT47Zf4KOIRF8rqT8Ps3pYzT8odROJ8rq4P+lciGrg29czpqrRM22yQktFr
itvcM7JlE6jFbGH3rycnvGvhRYU/j1N9k0ozm8oVwmKX357/OtGzNivBECGSnU9L
Bkp4Nm9M1K4cOwEuZ0xsQEthZjQYF0mDpnlWmVJL5i1Lq834atN2qrb/mzMHBNtD
JnqRV7rPL39lKpe7LJiitsC2JuW1UbWZZU1NNwA/rz2d83C+KD1DLJ0+sMYY2Q3T
OQ4BPAowDEwOH7XAXrHM/0kRm+ZeIFlwevEGIQWmMt1Ogz+AW4Iq0slINc4wOINK
vH9tHwIDAQABo4HVMIHSMIGSBgNVHSMEgYowgYeAFCnH4DkZPR6CZxRn/kIqVsMo
dJHpoWekZTBjMQswCQYDVQQGEwJDTjEQMA4GA1UECAwHQmVpSmluZzEQMA4GA1UE
BwwHQmVpSmluZzEXMBUGA1UECgwOTGFrYWxhIENvLixMdGQxFzAVBgNVBAMMDkxh
a2FsYSBSb290IENBggYBaiUALIowHQYDVR0OBBYEFIya0Yc4OSBer55JLyA0AYe9
m8mTMAwGA1UdEwEB/wQCMAAwDgYDVR0PAQH/BAQDAgeAMA0GCSqGSIb3DQEBCwUA
A4IBAQCBEwOlk3mXigNv94Drn3dcaY2ml/y+8yNpAIuUhuBE00WFoqEX5lOatFy5
fzdXuC12lBVQ8SjSm3aH7k2X0eXqDzkOHiur2ZBRKmJ++J4TeenuSUOjSIbQK/DT
vxaqFUjYwFSVCyizpy7wfU4wKt+jOuFb9LyULJ9lkM1dV9Kh7Lmd9+nlJYYuPEPU
LJkkVZqSALSiiJudXnTwlISjZTXEAkJpdIlMw+hvPTAkoG95B95M+OV/uLbItGK+
qT4+RHWo8EbBDPQYo6J4QYHOxRlfMoGBMyrz6XDt7ELLmT7ld4aE02w6KQPfK3gq
kLDT+/STozvaNmXzBJh7J6KqxJBH
-----END CERTIFICATE-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://s2.lakala.com",
		//ApiHost: "https://s2.lakala.com/",
		Version: "",
		Extra:   `{"term_no":"L3170815"}`,
	}
	client, err := NewAPICallback(c)
	assert.NoError(t, err)
	ctx := context.Background()
	req := &http.Request{
		Header: http.Header{},
	}
	auth := `LKLAPI-SHA256withRSA timestamp="1765848988",nonce_str="nBF21CEp8LDf",signature="GNTRUWTnY84Wp3U09K9lm4MVF1hU1++POXZtjmb634OPCnEpiHBRDIMzvXqknNyFoDBYbsWDmmwkDyGoFv4Zr/V5oKn70QbcAFNypy3R8tr/Iq3im6H80dzJti4SDrIlp3JAgp8UkqB0qPPQT0QmYACxVfLyJc+lC5fBAgfW1iaDVOfy7uO03s99TMDUlRn4Hl56zt2OBDivEe4YuCIOTyBnic+sJkxjIEJZN+8/tOHBblsX5+3VIXKTu8cr9V/DjcXX3rs1yh4CQ0ypmEec2cCdnxDvMF+eIPHyaJeZiAy2Dh03kOHGgJmvCILeJMgEDSjavZ17Gi5NQxMFn+tKWA=="`
	req.Header.Set("Authorization", auth)
	body := `{"acc_discount_amount":"0","acc_mdiscount_amount":"0","acc_settle_amount":"30","acc_trade_no":"4200002968202512168521236797","account_type":"WECHAT","bank_type":"CITIC_DEBIT","card_type":"00","gb_amount":"","log_no":"66202326851350","merchant_no":"822584073920DEH","notify_url":"https://proxy.test.jianxindianzi.com/public/v1/callback/payment/37/176584849551401225676","out_trade_no":"176584849551401225676","payer_amount":"30","promotion_detail":"","qb_amount":"","remark":"","sub_mch_id":"843442957","total_amount":"30","trade_no":"20251216110113130266202326851350","trade_req_date":"20251216","trade_state":"SUCCESS","trade_status":"SUCCESS","trade_time":"20251216092824","user_id1":"olymevsNYgJ9UxG1XRtQrlecKvsE","user_id2":"oVxsc1fhZTVRe2mAiN7dLMFRiO78"}`

	req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(body)))
	result, err := client.CallbackPaymentParse(ctx, req)
	assert.NoError(t, err)
	t.Log(result)
}
