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
		AppID:   "",
		MchID:   "0008710FA369999",
		APIKey:  "71a01b40caa811f09a88a57d8f5260d8",
		Cert:    config.Cert{},
		Proxy:   config.Proxy{},
		ApiHost: "https://aipay-cloud.fuioupay.com",
		Version: "1.0",
		Extra:   `{"wx_app_id": "wxc233b52c912ad8f7", "order_prefix": "19856", "wx_app_secret": "1660bd2be85e638319381da7867ce294"}`,
	}
	pppp, err := NewRefund(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	ctx := context.Background()
	req := &dto.RefundRequest{
		OrderNo:  "176423592652800749889",
		RefundNo: "17622222",
		Amount:   dto.Amount{Total: 2},
	}
	r, err := pppp.Refund(ctx, req)
	assert.NoError(t, err)
	fmt.Println(err)
	if r != nil {
		fmt.Println("---------------------------------")
		b, _ := json.Marshal(r)
		fmt.Println(string(b))
	}
	return
}

func TestRefund_Query(t *testing.T) {
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
	pppp, err := NewRefund(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	ctx := context.Background()
	req := dto.RefundQuery{
		RefundNo: "17622222",
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

func TestRefund_Parse(t *testing.T) {
	bbb := `{"ins_cd":"08M0061071","mchnt_cd":"0008710FA369999","mchnt_order_no":"19856176423592652800749889","order_type":"WECHAT","random_str":"20251127191303162439","refund_id":"50101505292025112744978914295","refund_order_no":"1985617622222","reserved_fund_state":"","reserved_fy_settle_dt":"20251127","reserved_fy_trace_no":"230711415467","reserved_refund_amt":"2","result_code":"000000","result_msg":"SUCCESS","sign":"d937952375e35d8687f7841b8fc813fb","term_id":"6286","trans_stat":"SUCCESS","transaction_id":"4200002918202511270538002877"}`
	resp := &RefundQueryResponse{}
	err := json.Unmarshal([]byte(bbb), resp)
	assert.NoError(t, err)

	fmt.Println(resp.TransStat, resp.GetRefundStatus())
}
