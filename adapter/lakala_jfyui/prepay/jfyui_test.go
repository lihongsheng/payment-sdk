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

func Test_NewJFYUI(t *testing.T) {
	c := config.Config{
		AppID:  "1995090100510236673",
		MchID:  "82270304816033L",
		APIKey: "f00dac5077ea11e754e14c9541bc0170",
		Cert: config.Cert{
			CertificateSerialNumber: "12323213123123213213",
		},
		Proxy: config.Proxy{},
		// ApiHost: "https://test.wsmsd.cn/sit",
		ApiHost: "https://s2.lakala.com/",
		Version: "",
		Extra:   `{"term_no":"P1449896"}`,
	}
	pppp, err := NewJFYUI(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
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
			OpenID: "owDb32P8qIDRcPrjkTRmA5pXUmp8",
			AppID:  "wxc233b52c912ad8f7",
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
	rrr, err := pppp.Pay(ctx, req)
	assert.NoError(t, err)
	if err != nil {
		t.Log(err.Error())
	}
	if rrr != nil {
		by, _ := json.Marshal(rrr)
		t.Log(string(by))
	}
}

func Test_NewJFYUI_QUERY(t *testing.T) {
	c := config.Config{
		AppID:  "1995090100510236673",
		MchID:  "82270304816033L",
		APIKey: "f00dac5077ea11e754e14c9541bc0170",
		Cert: config.Cert{
			CertificateSerialNumber: "12323213123123213213",
		},
		Proxy: config.Proxy{},
		// ApiHost: "https://test.wsmsd.cn/sit",
		ApiHost: "https://s2.lakala.com/",
		Version: "",
		Extra:   `{"term_no":"P1449896"}`,
	}
	pppp, err := NewJFYUI(c, enum.PaymentProduct_JSAPI, enum.Payment_Wxpay)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	ctx := context.Background()
	req := dto.Query{
		TradeNo: "25120211012001101011544312821",
	}
	rrr, err := pppp.Query(ctx, req)
	assert.NoError(t, err)
	if rrr != nil {
		by, _ := json.Marshal(rrr)
		t.Log(string(by))
	}
}
