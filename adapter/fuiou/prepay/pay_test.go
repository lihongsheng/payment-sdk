package prepay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
	"github.com/singer-stack-lab/payment-sdk/tools"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPay_Pay(t *testing.T) {
	c := config.Config{
		AppID:   "",
		MchID:   "0002900F1503036",
		APIKey:  "f00dac5077ea11e754e14c9541bc0170",
		Cert:    config.Cert{},
		Proxy:   config.Proxy{},
		ApiHost: "https://aipaytest.fuioupay.com/",
		Version: "",
		Extra:   `{"order_prefix":"1066"}`,
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
			OpenID: "ogdvH6h9jPp5R3f1fyLsQjdB-fAc",
			AppID:  "wxfa089da95020ba1a",
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

	r, err := pppp.Pay(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	if r != nil {
		fmt.Println("-----------------------------------------------")
		by, _ := json.Marshal(r)
		fmt.Println(string(by))
	}
}

func TestPay_Query(t *testing.T) {
	c := config.Config{
		AppID:   "",
		MchID:   "0002900F1503036",
		APIKey:  "f00dac5077ea11e754e14c9541bc0170",
		Cert:    config.Cert{},
		Proxy:   config.Proxy{},
		ApiHost: "https://aipaytest.fuioupay.com/",
		Version: "1.0",
	}
	pppp, err := NewPay(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println(pppp)
	ctx := context.Background()
	req := dto.Query{
		OrderNo: "1066" + tools.GetID(),
	}
	r, err := pppp.Query(ctx, req)
	assert.NotEmpty(t, err)
	assert.Nil(t, r)
	return
}

func TestPay_Close(t *testing.T) {
	c := config.Config{
		AppID:   "",
		MchID:   "0002900F1503036",
		APIKey:  "f00dac5077ea11e754e14c9541bc0170",
		Cert:    config.Cert{},
		Proxy:   config.Proxy{},
		ApiHost: "https://aipaytest.fuioupay.com/",
		Version: "1.0",
	}
	pppp, err := NewPay(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println(pppp)
	ctx := context.Background()
	req := dto.CloseQuery{
		OrderNo: "1066" + tools.GetID(),
	}
	err = pppp.Close(ctx, req)
	assert.NotEmpty(t, err)
	return
}

func TestPay_Query2(t *testing.T) {
	c := config.Config{
		AppID:   "",
		MchID:   "0008710FA369999",
		APIKey:  "71a01b40caa811f09a88a57d8f5260d8",
		Cert:    config.Cert{},
		Proxy:   config.Proxy{},
		ApiHost: "https://aipay-cloud.fuioupay.com",
		Version: "1.0",
		Extra:   `{"wx_app_id": "wxc233b52c912ad8f7", "order_prefix": "19856", "wx_app_secret": "1660bd2be85e638319381da7867ce294"}`,
	}
	pppp, err := NewPay(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	ctx := context.Background()
	req := dto.Query{
		OrderNo: "176423592652800749889",
	}
	r, err := pppp.Query(ctx, req)
	assert.NoError(t, err)
	fmt.Println(err)
	if r != nil {
		fmt.Println("---------------------------------")
		b, _ := json.Marshal(r)
		fmt.Println(string(b))
	}
	return
}

func TestPay_QueryBody(t *testing.T) {
	body := `{"addn_inf":"tpdp_id=33","buyer_id":"owDb32P8qIDRcPrjkTRmA5pXUmp8","mchnt_cd":"0008710FA369999","mchnt_order_no":"19856176423592652800749889","order_amt":"2","order_type":"JSAPI","random_str":"20251127185516791790","reserved_bank_type":"CITIC_DEBIT","reserved_buyer_logon_id":"","reserved_channel_order_id":"19856176423592652800749889","reserved_coupon_fee":"","reserved_fund_bill_list":"","reserved_fund_state":"","reserved_fy_settle_dt":"20251127","reserved_fy_trace_no":"230705503286","reserved_is_credit":"0","reserved_promotion_detail":"","reserved_relate_order_no":"","reserved_relate_trace_no":"","reserved_settlement_amt":"2","reserved_txn_fin_ts":"20251127173212","reserved_unfreeze_time":"","result_code":"000000","result_msg":"SUCCESS","sign":"d2d694499037277731fe1352e21c9958","term_id":"6328","trans_stat":"SUCCESS","transaction_id":"4200002918202511270538002877"}`
	resp := &OrderDetail{}
	err := json.Unmarshal([]byte(body), resp)
	assert.NoError(t, err)
}
