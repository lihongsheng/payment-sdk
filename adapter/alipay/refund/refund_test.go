package refund

import (
	"context"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRefund(t *testing.T) {

	c := config.Config{
		AppID:  "2021005199606982",
		MchID:  "",
		APIKey: "",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCVHxmnZCHt57/RGbQQCfagILP37StXMV9H9xmPNL5UDusEfvzZZaUJ/xrneiLgGkmrVYSlvvu4VqNWUmufHCFiOuSZawlXud2hhTLYOjMme31mgXfwA/LD2zL273Njl/WQcnSjDZ+/i8/VieESSd+VfB2rem0EMSg2IskL8FRtjOKUFH7UxhLltBwieKYRJhrAC3d1iK/AYxq8580kkhnPRkEq97nQDhmObxPytB2jdbrg5VcuJul98wSrh/rmK6NmKwYxoEHXkXtazrPOMcWepcjvPc6LGls/pkjPHTsNuQ2zok2eouXxT4pXsKaXat//lC4DcN7ZYLMgXW6U1GghAgMBAAECggEAT0WcJp8VZloXXzoRvaD+STt5UGf07UIiy6fxH27Dav4PM3GqSos9Q+YoUDteRNHUrc8lV9MCD4mHBQGFkqyqloSGN4RuEAOPbSjY0ZeGz4NAM3jJ4H7I1BM3sYuzRJNoiGQ/2toIPRQ+5N6KFmXNBLNZNmo7f8n9/WFajQ0yZJV7r1dUzyylUYoi3MXJcZdZ9GPtYnDgm1mYxTFiUZH8Vs7SIARDV2PXsTPGrPoCuvWGLjGJ70TmD9lf7qbT2XQ9ECA837o+mB/DRCXPkW1YZibEo/dSn1dgN45va3OwLlKmHP9nmas49fACajENXEcH31OSLMKn+/8ibf4jg8WxQQKBgQDGdzFrwc1zcTNeZwpqSiqPHkveikA0jPKhJXQEQ7L4LqvlXrbNZOztMKegtoiXcXylFs6ktVB6DoXJRO7ko9lj0wXku68ykPg9lTKRBUPwxAqylwvXr27kEraYbc01UNMjYdWoRmW2r0xWYuzKCgt0OaBh8I996mJtwBlqHr1TOQKBgQDAWemqJaWMk5T7b0jynfAMjG639kgCuFP4ieibsCEnks82wABh9CQCUftjJuKJra0Y90Jo9LDcsiEhJzj2g2Ux/ElfGMJ6UAnL4nNZRsuIKIYQ4SYqkaywRUpXrqSfppzT+gIMIRIkaTFWyNl32+YcPgYvjLpEQ43dOw0Bg+60KQKBgA/ZoRbMCY54sfOOIyTN+4wmjUSeofYQc6gvB473oSl1AMF2yP6qWsmuoxSQv2Po6l187s/0sxKsaj7ixHl6JHh2d/gGjT1qmojAi6NNWlm2LmyI56c4GZEZdX/C9SqW4+XlgFzKEoF+iogzqlmQZ83GdGqd/be+qlG9j8oL3m7xAoGBAJWqIC7cpghQmL5e46KrkykgZ9HJ1lQPen7LR2tESzfQely+Xk3ZOd94TKLkkaXQtfvUAB9MLQU7sJ1sOF1O9YFELQ7KZB/pgQXcCCZx+FHUPiSKGzlKKdTpxSOeZsc2S5aAU/F+FfCwzMwa8WLafhyBgiyjUhdlfM+jb8Y7EpOpAoGAIFM8TGpNBPL/RK8uJ1QBXTn6n1b6SDGpj3HUZyEaa20BORRF7+mHvLkeXhLH09XaSmyI3+igK67/Exj07DkDJsaTA/xlYvzqkpCds2wZHs+/rmC5KC4tM46wxAkOK1VtjAvhh+4FnHOUGjhTOS6hAba5wP3sgkHfTwfRYDWDkQw=
-----END RSA PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsyA5/k4SyfoHhZCPIr/VjBG1vRhOwPLoOddf5WCH6168h/el+BGMiNOdczbP11HmrfpJWsXKDR7W/T/RMwX032RY2ah7HTJG1Auh5RXu4SoKmZYq84nvT0i0N8hp222E1fECts1NaWSfQP8dZcbFTd5HEWwbBmdFR/ExWl2u3wZB0OFHgIz4rfVV3qQnqFLzPqYIV23QfkHS0cVrlZU+GjSO7BVLRozOPD+0O1N9cuq8gjqSYDbR83VTAG27hfJel3prg5FTDAeLSi1eyUIrNRB3LZjaj5X90gyRHY4nNOx7Jh01YTuZR/FSHVHVdGdkw3h131FqmJpB2xAdc/qnVwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "",
		Version: "",
	}
	pay, err := NewRefund(c)
	assert.NoError(t, err)
	ctx := context.Background()
	req := &dto.RefundRequest{
		OrderNo:  "176068626732100521009",
		TradeNo:  "",
		Reason:   "测试退款",
		RefundNo: "176068626732100521009",
		Amount: dto.Amount{
			Total:    2,
			Currency: "CNY",
		},
	}
	r, err := pay.Refund(ctx, req)
	assert.NoError(t, err)
	// {"alipay_trade_refund_response":{"code":"10000","msg":"Success","buyer_logon_id":"158******07","fund_change":"N","gmt_refund_pay":"2025-10-16 11:14:30","out_trade_no":"Test11110003","refund_detail_item_list":[{"amount":"0.01","fund_channel":"COUPON"}],"refund_fee":"0.01","send_back_fee":"0.01","trade_no":"2025101622001429831436657633","buyer_open_id":"083luGU2ef8yeC95sroZ9WGyojnYtj3lBoBA6lY_2PlBrU4"},"sign":"eLBVZGQKREKg8IECb5XDC4iPkdID6G6L9myt9v63YcmSS7TNytHpbmzRxGkPdTPq+YaJyQS1OE3XdwTYh/nPiYHEtdfpIIbUV7XwH8znqBaYD95X+B0tLkeuaXIv8dTwh0OY1byMzJqmf5UJLROP9B+ewHOOFV10IMUdNAE7exXUHT+x6qanx4bNCs14QCiGcWo7gXqn8x67j6CPxtxsamb9idSaMgVcJkyETG4BH6+YfKePYUY9A6teWuDMii3jihpmF0L5emXZT9OoJjGuq1R1nn176dj3COffLyW+IMVHCVbu3a82XyzMLL+1h2xxxzflP1waf0yWPl+fd5T/6Q=="}
	t.Log(r)
}

func TestRefund_Query(t *testing.T) {

	c := config.Config{
		AppID:  "2021005199606982",
		MchID:  "",
		APIKey: "",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCVHxmnZCHt57/RGbQQCfagILP37StXMV9H9xmPNL5UDusEfvzZZaUJ/xrneiLgGkmrVYSlvvu4VqNWUmufHCFiOuSZawlXud2hhTLYOjMme31mgXfwA/LD2zL273Njl/WQcnSjDZ+/i8/VieESSd+VfB2rem0EMSg2IskL8FRtjOKUFH7UxhLltBwieKYRJhrAC3d1iK/AYxq8580kkhnPRkEq97nQDhmObxPytB2jdbrg5VcuJul98wSrh/rmK6NmKwYxoEHXkXtazrPOMcWepcjvPc6LGls/pkjPHTsNuQ2zok2eouXxT4pXsKaXat//lC4DcN7ZYLMgXW6U1GghAgMBAAECggEAT0WcJp8VZloXXzoRvaD+STt5UGf07UIiy6fxH27Dav4PM3GqSos9Q+YoUDteRNHUrc8lV9MCD4mHBQGFkqyqloSGN4RuEAOPbSjY0ZeGz4NAM3jJ4H7I1BM3sYuzRJNoiGQ/2toIPRQ+5N6KFmXNBLNZNmo7f8n9/WFajQ0yZJV7r1dUzyylUYoi3MXJcZdZ9GPtYnDgm1mYxTFiUZH8Vs7SIARDV2PXsTPGrPoCuvWGLjGJ70TmD9lf7qbT2XQ9ECA837o+mB/DRCXPkW1YZibEo/dSn1dgN45va3OwLlKmHP9nmas49fACajENXEcH31OSLMKn+/8ibf4jg8WxQQKBgQDGdzFrwc1zcTNeZwpqSiqPHkveikA0jPKhJXQEQ7L4LqvlXrbNZOztMKegtoiXcXylFs6ktVB6DoXJRO7ko9lj0wXku68ykPg9lTKRBUPwxAqylwvXr27kEraYbc01UNMjYdWoRmW2r0xWYuzKCgt0OaBh8I996mJtwBlqHr1TOQKBgQDAWemqJaWMk5T7b0jynfAMjG639kgCuFP4ieibsCEnks82wABh9CQCUftjJuKJra0Y90Jo9LDcsiEhJzj2g2Ux/ElfGMJ6UAnL4nNZRsuIKIYQ4SYqkaywRUpXrqSfppzT+gIMIRIkaTFWyNl32+YcPgYvjLpEQ43dOw0Bg+60KQKBgA/ZoRbMCY54sfOOIyTN+4wmjUSeofYQc6gvB473oSl1AMF2yP6qWsmuoxSQv2Po6l187s/0sxKsaj7ixHl6JHh2d/gGjT1qmojAi6NNWlm2LmyI56c4GZEZdX/C9SqW4+XlgFzKEoF+iogzqlmQZ83GdGqd/be+qlG9j8oL3m7xAoGBAJWqIC7cpghQmL5e46KrkykgZ9HJ1lQPen7LR2tESzfQely+Xk3ZOd94TKLkkaXQtfvUAB9MLQU7sJ1sOF1O9YFELQ7KZB/pgQXcCCZx+FHUPiSKGzlKKdTpxSOeZsc2S5aAU/F+FfCwzMwa8WLafhyBgiyjUhdlfM+jb8Y7EpOpAoGAIFM8TGpNBPL/RK8uJ1QBXTn6n1b6SDGpj3HUZyEaa20BORRF7+mHvLkeXhLH09XaSmyI3+igK67/Exj07DkDJsaTA/xlYvzqkpCds2wZHs+/rmC5KC4tM46wxAkOK1VtjAvhh+4FnHOUGjhTOS6hAba5wP3sgkHfTwfRYDWDkQw=
-----END RSA PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsyA5/k4SyfoHhZCPIr/VjBG1vRhOwPLoOddf5WCH6168h/el+BGMiNOdczbP11HmrfpJWsXKDR7W/T/RMwX032RY2ah7HTJG1Auh5RXu4SoKmZYq84nvT0i0N8hp222E1fECts1NaWSfQP8dZcbFTd5HEWwbBmdFR/ExWl2u3wZB0OFHgIz4rfVV3qQnqFLzPqYIV23QfkHS0cVrlZU+GjSO7BVLRozOPD+0O1N9cuq8gjqSYDbR83VTAG27hfJel3prg5FTDAeLSi1eyUIrNRB3LZjaj5X90gyRHY4nNOx7Jh01YTuZR/FSHVHVdGdkw3h131FqmJpB2xAdc/qnVwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "",
		Version: "",
	}
	pay, err := NewRefund(c)
	assert.NoError(t, err)
	ctx := context.Background()
	req := dto.RefundQuery{
		OrderNo:  "176068626732100521009",
		TradeNo:  "",
		RefundNo: "176068626732100521009",
	}
	r, err := pay.Query(ctx, req)
	assert.NoError(t, err)
	// {"alipay_trade_fastpay_refund_query_response":{"code":"10000","msg":"Success","out_request_no":"TestRefund11110003","out_trade_no":"Test11110003","refund_amount":"0.01","refund_detail_item_list":[{"amount":"0.01","fund_channel":"COUPON"}],"refund_status":"REFUND_SUCCESS","send_back_fee":"0.01","total_amount":"0.01","trade_no":"2025101622001429831436657633"},"sign":"kbqtzbpop6wvdYdhsX/Jd+CAbcpMsbiX6Ymrz/bzeP0BsoUwomhpPdazFTYSZVNV+jIIrx08LKyzBkdIe24fDUNY/PGjrj8Sr837cFH3/TN1Krp/yxIpNgePg3/INgd8UaTjrsyZkA2YwlbE+MqRnTKnrRjyzH5rtWwVvhf205gfTTd7ojuz0rsNF1KEDfSLx7bGV7et0UnjHt7zqNYv6akU8kBZQVEEtv/l+lKeAD0Tt+saaMYn43x20en5+VHs4/W861Ul8ZoFFqgZCdTkh0DagXP4L596dEUex4AnbTzcc775iB1ReeQd9Pj+PQ1URXUCuFY+vi1pbTZ5jkFHdw=="}
	t.Log(r)
}
