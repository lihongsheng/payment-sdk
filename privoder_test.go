package payment_sdk

import (
	"context"
	"fmt"
	wxCf "github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	enum2 "github.com/lihongsheng/payment-sdk/enum"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPayment_WxpayJSAPI(t *testing.T) {
	opts := []config.Option{
		config.WithPaymentProduct(payment.PaymentProduct_JSAPI),
		config.WithWxConfig(&wxCf.Config{
			AppID:     "wx6666",
			MchID:     "17*****",
			APISecret: "***",
			AppSecret: "***",
			RsaPrivate: `-----BEGIN PRIVATE KEY-----
***
-----END PRIVATE KEY-----`,
			RsaPrivateNumber: "***",
			ScoreServiceID:   "",
		}),
	}

	pay, err := Payment(channel.Channel_Wechat, opts...)
	assert.NoError(t, err)
	ctx := context.Background()
	req := &dto.PayOrder{
		Order: dto.Order{
			OrderNo: fmt.Sprintf("%d", time.Now().Unix()),
			PayAmount: dto.Amount{
				Currency: "CNY",
				Total:    1,
			},
			Subject: "test pay",
		},
		NotifyUrl:      "https://***/callback/v1/tenant/payment/5/175870989522000055545",
		PassBackParams: "tpdp_id=5",
		Payer: dto.Payer{
			OpenID: "o7uyb2G1X2nRyhPCwGSt6TFd3sAs",
		},
		RedirectUrl: "",
		TimeExpire:  time.Now().Unix() + 90,
	}
	resp, err := pay.Pay(ctx, req)
	assert.NoError(t, err)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(resp)
}

func TestPayment_WxpayH5API(t *testing.T) {
	opts := []config.Option{
		config.WithPaymentProduct(payment.PaymentProduct_H5),
		config.WithWxConfig(&wxCf.Config{
			AppID:     "wx6c663032961e5e4b",
			MchID:     "1730424164",
			APISecret: "***",
			AppSecret: "***",
			RsaPrivate: `-----BEGIN PRIVATE KEY-----
***
-----END PRIVATE KEY-----`,
			RsaPrivateNumber: "***",
			ScoreServiceID:   "",
		}),
	}

	pay, err := Payment(channel.Channel_Wechat, opts...)
	assert.NoError(t, err)
	ctx := context.Background()
	req := &dto.PayOrder{
		Order: dto.Order{
			OrderNo: fmt.Sprintf("%d", time.Now().Unix()),
			PayAmount: dto.Amount{
				Currency: "CNY",
				Total:    1,
			},
			Subject: "test pay",
		},
		NotifyUrl:      "https://***/callback/v1/tenant/payment/5/175870989522000055545",
		PassBackParams: "tpdp_id=5",
		Payer: dto.Payer{
			OpenID: "o7uyb2G1X2nRyhPCwGSt6TFd3sAs",
		},
		RedirectUrl: "",
		TimeExpire:  time.Now().Unix() + 90,
		SceneInfo: &dto.SceneInfo{
			ClientIp:        "127.0.0.1",
			Device:          enum2.Device_H5,
			ApplicationInfo: dto.ApplicationInfo{},
		},
	}
	resp, err := pay.Pay(ctx, req)
	assert.NoError(t, err)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(resp)
}
