package refund

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPay_Refund(t *testing.T) {
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
	pppp, err := NewRefund(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	ctx := context.Background()
	req := &dto.RefundRequest{
		RefundNo:  "175879416443300038879",
		TradeNo:   "",
		OrderNo:   "1766113843",
		Reason:    "reason",
		NotifyUrl: "https://api-cabinet.test.jianxindianzi.com/public/v1/callback/refund/5/175879416443300038879",
		Amount: dto.Amount{
			Total:    1,
			Currency: "CNY",
		},
		Goods: nil,
		OrderAmount: dto.Amount{
			Total:    1,
			Currency: "CNY",
		},
	}
	// {"mchnt_cd":"0002900F1503036","qr_code":"","random_str":"20251126164722483439","reserved_addn_inf":"","reserved_channel_order_id":"106617641468421811000000008530","reserved_fund_state":"","reserved_fy_order_no":"","reserved_fy_settle_dt":"","reserved_fy_trace_no":"230659899430","reserved_pay_info":"{\"timeStamp\":\"1764146842\",\"package\":\"prepay_id=wx2616472230110302f0d4f5e0571ae80001\",\"paySign\":\"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==\",\"appId\":\"wxfa089da95020ba1a\",\"signType\":\"RSA\",\"nonceStr\":\"68af5752a9e0489eaf1865e451519f42\"}","reserved_transaction_id":"","result_code":"000000","result_msg":"SUCCESS","sdk_appid":"wxfa089da95020ba1a","sdk_noncestr":"68af5752a9e0489eaf1865e451519f42","sdk_package":"prepay_id=wx2616472230110302f0d4f5e0571ae80001","sdk_partnerid":"","sdk_paysign":"WupERdk6N+Wg89k8fQMq6gNqPGpAh4jwd+9+C5dHOyvBc2o+9EkoT27ws+GwRoSw8ycCF2Tj+qLHl6Ps6d+NcbaH3AnUXrUTGpaD24QZ2wi6+dDoWiMK0U/EO1npUSQtZ+RX1JO2RobwweDSdySDHOQdDnaFb1Q+uGOy/Sx84G1EWnAyQre158FSLI/whVMifoHh4t9A94F83lIYtVuO6EjFFLMIENrGtLACADwRzs0Hd5s+de4ou76b4Uo1+2kG7OE7nwkmxIgSwPOcaiuROTdeWBByHAB3Q5J49RMHf+VoeUqQUnnccvk65s7ArV/yHborMLO2sGLnVLIAV49U3A==","sdk_signtype":"RSA","sdk_timestamp":"1764146842","session_id":"wx2616472230110302f0d4f5e0571ae80001","sign":"aca102ea52e3cd6da553fdb33f64d24f","sub_appid":"wxfa089da95020ba1a","sub_mer_id":"320371539","sub_openid":"ogdvH6h9jPp5R3f1fyLsQjdB-fAc","term_id":"226681494"}
	// 1766113843
	r, err := pppp.Refund(ctx, req)
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
	pppp, err := NewRefund(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	ctx := context.Background()
	req := dto.RefundQuery{
		RefundNo: "175879416443300038879",
		TradeNo:  "",
		OrderNo:  "",
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
