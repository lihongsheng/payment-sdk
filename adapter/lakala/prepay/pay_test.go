package prepay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/tools"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPay_Pay(t *testing.T) {
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
	pppp, err := NewPay(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	ctx := context.Background()
	req := &dto.PayOrder{
		Order: dto.Order{
			OrderNo: tools.GetID(),
			PayAmount: dto.Amount{
				Total:    10,
				Currency: enum.Currency_CNY.String(),
			},
			Subject: "测试支付订单",
		},
		Payer: dto.Payer{
			OpenID: "olymevsNYgJ9UxG1XRtQrlecKvsE",
			AppID:  "wx7159f5cc93f4ea9f",
		},
		RedirectUrl:    "",
		TimeExpire:     time.Now().Add(time.Hour).Unix(),
		NotifyUrl:      "https://api.test.com/order",
		PassbackParams: "",
		SettleInfo:     nil,
		SceneInfo: &dto.SceneInfo{
			ClientIp: "127.0.0.1",
		},
	}

	// {"mchnt_cd":"0002900F1503036","qr_code":"","random_str":"20251126164722483439","reserved_addn_inf":"","reserved_channel_order_id":"106617641468421811000000008530","reserved_fund_state":"","reserved_fy_order_no":"","reserved_fy_settle_dt":"","reserved_fy_trace_no":"230659899430","reserved_pay_info":"{\"timeStamp\":\"1764146842\",\"package\":\"prepay_id=wx2616472230110302f0d4f5e0571ae80001\",\"paySign\":\"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==\",\"appId\":\"wxfa089da95020ba1a\",\"signType\":\"RSA\",\"nonceStr\":\"68af5752a9e0489eaf1865e451519f42\"}","reserved_transaction_id":"","result_code":"000000","result_msg":"SUCCESS","sdk_appid":"wxfa089da95020ba1a","sdk_noncestr":"68af5752a9e0489eaf1865e451519f42","sdk_package":"prepay_id=wx2616472230110302f0d4f5e0571ae80001","sdk_partnerid":"","sdk_paysign":"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==","sdk_signtype":"RSA","sdk_timestamp":"1764146842","session_id":"wx2616472230110302f0d4f5e0571ae80001","sign":"aca102ea52e3cd6da553fdb33f64d24f","sub_appid":"wxfa089da95020ba1a","sub_mer_id":"320371539","sub_openid":"ogdvH6h9jPp5R3f1fyLsQjdB-fAc","term_id":"226681494"}
	pp := pppp.(*Pay)
	r, err := pp.Pay(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	if err != nil {
		fmt.Println(err.Error())
	}
	if r != nil {
		fmt.Println("-----------------------------------------------")
		by, _ := json.Marshal(r)
		fmt.Println(string(by))
	}
}

func TestPay_PayQuery(t *testing.T) {
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
	pppp, err := NewPay(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	ctx := context.Background()
	req := dto.Query{
		OrderNo: "176579233208001223041",
	}

	// {"mchnt_cd":"0002900F1503036","qr_code":"","random_str":"20251126164722483439","reserved_addn_inf":"","reserved_channel_order_id":"106617641468421811000000008530","reserved_fund_state":"","reserved_fy_order_no":"","reserved_fy_settle_dt":"","reserved_fy_trace_no":"230659899430","reserved_pay_info":"{\"timeStamp\":\"1764146842\",\"package\":\"prepay_id=wx2616472230110302f0d4f5e0571ae80001\",\"paySign\":\"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==\",\"appId\":\"wxfa089da95020ba1a\",\"signType\":\"RSA\",\"nonceStr\":\"68af5752a9e0489eaf1865e451519f42\"}","reserved_transaction_id":"","result_code":"000000","result_msg":"SUCCESS","sdk_appid":"wxfa089da95020ba1a","sdk_noncestr":"68af5752a9e0489eaf1865e451519f42","sdk_package":"prepay_id=wx2616472230110302f0d4f5e0571ae80001","sdk_partnerid":"","sdk_paysign":"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==","sdk_signtype":"RSA","sdk_timestamp":"1764146842","session_id":"wx2616472230110302f0d4f5e0571ae80001","sign":"aca102ea52e3cd6da553fdb33f64d24f","sub_appid":"wxfa089da95020ba1a","sub_mer_id":"320371539","sub_openid":"ogdvH6h9jPp5R3f1fyLsQjdB-fAc","term_id":"226681494"}
	pp := pppp.(*Pay)
	r, err := pp.Query(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	if err != nil {
		fmt.Println(err.Error())
	}
	if r != nil {
		fmt.Println("-----------------------------------------------")
		by, _ := json.Marshal(r)
		fmt.Println(string(by))
	}
}
