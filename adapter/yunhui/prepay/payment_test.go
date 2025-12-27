package prepay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPay_Pay(t *testing.T) {
	c := config.Config{
		AppID: "Y1765007866-2",
		//MchID: "Y17661083936167484",
		MchID:  "Y1765007866",
		APIKey: "3akyc6qDCFq2yIxNKmZ02waGB6oxOhzH8oqpTZLNHKyRNOwuagot3FwyXqvrXySWmgHriPT5jRbvHC8sREjiGWCfGjdUjCMygwOv8rDgHaQzHiEXoUtjG8pW2EelhNh6",
		Cert: config.Cert{
			CertificateSerialNumber: "",
			CertPrivateKey:          ``,
			PublicKey:               ``,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://payment.yiyunhuipay.com",
		Version: "",
		Extra:   `{"term_no":"63751"}`,
	}
	pppp, err := NewPayment(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	ctx := context.Background()
	req := &dto.PayOrder{
		Order: dto.Order{
			OrderNo: fmt.Sprintf("%d", time.Now().Unix()+1),
			//OrderNo: "1766112308",
			PayAmount: dto.Amount{
				Total:    1,
				Currency: enum.Currency_CNY.String(),
			},
			Subject: "test",
		},
		Payer: dto.Payer{
			OpenID: "olymevsNYgJ9UxG1XRtQrlecKvsE",
			AppID:  "wx7159f5cc93f4ea9f",
		},
		RedirectUrl:    "https://proxy.test.jianxindianzi.com/h5/order",
		TimeExpire:     time.Now().Add(time.Hour).Unix(),
		NotifyUrl:      "https://api.test.com/order",
		PassbackParams: "",
		SettleInfo:     nil,
		SceneInfo: &dto.SceneInfo{
			ClientIp: "127.0.0.1",
		},
	}

	// {"mchnt_cd":"0002900F1503036","qr_code":"","random_str":"20251126164722483439","reserved_addn_inf":"","reserved_channel_order_id":"106617641468421811000000008530","reserved_fund_state":"","reserved_fy_order_no":"","reserved_fy_settle_dt":"","reserved_fy_trace_no":"230659899430","reserved_pay_info":"{\"timeStamp\":\"1764146842\",\"package\":\"prepay_id=wx2616472230110302f0d4f5e0571ae80001\",\"paySign\":\"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==\",\"appId\":\"wxfa089da95020ba1a\",\"signType\":\"RSA\",\"nonceStr\":\"68af5752a9e0489eaf1865e451519f42\"}","reserved_transaction_id":"","result_code":"000000","result_msg":"SUCCESS","sdk_appid":"wxfa089da95020ba1a","sdk_noncestr":"68af5752a9e0489eaf1865e451519f42","sdk_package":"prepay_id=wx2616472230110302f0d4f5e0571ae80001","sdk_partnerid":"","sdk_paysign":"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==","sdk_signtype":"RSA","sdk_timestamp":"1764146842","session_id":"wx2616472230110302f0d4f5e0571ae80001","sign":"aca102ea52e3cd6da553fdb33f64d24f","sub_appid":"wxfa089da95020ba1a","sub_mer_id":"320371539","sub_openid":"ogdvH6h9jPp5R3f1fyLsQjdB-fAc","term_id":"226681494"}
	// 1766113843
	r, err := pppp.Pay(ctx, req)
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

func TestPay_Query(t *testing.T) {
	c := config.Config{
		AppID: "Y1765007866-5",
		//MchID: "Y17661083936167484",
		MchID:  "Y1765007866",
		APIKey: "oNV6w3sdxHyKDqOEJRkWaoWnfvW0qvxiA8PgznIeaEGXJlUqq1Yyfr4v1j5LmC0gf9FGYeRkg3Yc2BYtVeeSRds7ChPx19qmXpsGAPcHWGuLMUqK4wmbACqDQmUZDUIN",
		Cert: config.Cert{
			CertificateSerialNumber: "",
			CertPrivateKey:          ``,
			PublicKey:               ``,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://payment.yiyunhuipay.com",
		Version: "",
		Extra:   `{"term_no":"63751"}`,
	}
	pppp, err := NewPayment(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	ctx := context.Background()
	req := dto.Query{
		OrderNo: "1766113843",
	}

	// {"mchnt_cd":"0002900F1503036","qr_code":"","random_str":"20251126164722483439","reserved_addn_inf":"","reserved_channel_order_id":"106617641468421811000000008530","reserved_fund_state":"","reserved_fy_order_no":"","reserved_fy_settle_dt":"","reserved_fy_trace_no":"230659899430","reserved_pay_info":"{\"timeStamp\":\"1764146842\",\"package\":\"prepay_id=wx2616472230110302f0d4f5e0571ae80001\",\"paySign\":\"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==\",\"appId\":\"wxfa089da95020ba1a\",\"signType\":\"RSA\",\"nonceStr\":\"68af5752a9e0489eaf1865e451519f42\"}","reserved_transaction_id":"","result_code":"000000","result_msg":"SUCCESS","sdk_appid":"wxfa089da95020ba1a","sdk_noncestr":"68af5752a9e0489eaf1865e451519f42","sdk_package":"prepay_id=wx2616472230110302f0d4f5e0571ae80001","sdk_partnerid":"","sdk_paysign":"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==","sdk_signtype":"RSA","sdk_timestamp":"1764146842","session_id":"wx2616472230110302f0d4f5e0571ae80001","sign":"aca102ea52e3cd6da553fdb33f64d24f","sub_appid":"wxfa089da95020ba1a","sub_mer_id":"320371539","sub_openid":"ogdvH6h9jPp5R3f1fyLsQjdB-fAc","term_id":"226681494"}
	// 1766113843
	r, err := pppp.Query(ctx, req)
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
