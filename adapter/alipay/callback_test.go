package alipay

import (
	"bytes"
	"context"
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"io"
	"net/http"
	"testing"
)

func TestCallbackPaymentParse(t *testing.T) {
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
	req := &http.Request{}
	paramStr := `gmt_create=2025-10-17+15%3A31%3A07&charset=UTF-8&seller_email=yuioplkjh202502%40163.com&subject=%E5%AE%A2%E6%9C%8D%3A+400%E2%80%93088%E2%80%935739&sign=dbNudH3ydq90yTvrI8HmXwj%2FkAS0Bda6xVvIiKTQLjjKenoYl8AZw%2F4%2FBu37D%2FCy66aX5zaWQwtrKNKeaHpMMACwBQLQngZLVy8puy74FvbfUeMAt5cYpQAI9n%2Bl2zsFanefNDY%2BCiEzQc8dwog8QttTmWICtgnxewDCgWY5%2B11cmobRUL5%2FqiWtclggm2egQ%2F1yxcAQFsLnOvpSE54uW8Fu7bLaKARwCagl3rDHn%2BKV592okyJK%2BcA%2FUDS%2B4g9qKrq%2FStoJniVwdwWaPbHJ66He60yQTjlQaTm8ml820iWXa7GggFHr7YXFfYSr2qaobChcCCl%2FwYlo9L25Oo2Q3g%3D%3D&buyer_open_id=099luGU2ef8yeC95sroZ9WGyuuGqL7qD-VLY8B1mtPqj4w9&invoice_amount=0.00&notify_id=2025101701222153119086991451124623&fund_bill_list=%5B%7B%22amount%22%3A%220.02%22%2C%22fundChannel%22%3A%22COUPON%22%7D%5D&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=0.02&buyer_pay_amount=0.02&app_id=2021005199606982&sign_type=RSA2&seller_id=2088070256235011&gmt_payment=2025-10-17+15%3A31%3A18&notify_time=2025-10-17+15%3A56%3A32&merchant_app_id=2021005199606982&passback_params=tpdp_id%3D17&version=1.0&out_trade_no=176068626732100521009&total_amount=0.02&trade_no=2025101722001486991458220176&auth_app_id=2021005199606982&buyer_logon_id=999***%40sina.com&point_amount=0.00`
	ctx := context.Background()
	req.Body = io.NopCloser(bytes.NewBuffer([]byte(paramStr)))
	r, err := CallbackPaymentParse(ctx, c, req)
	t.Log(r)
	t.Log(r.EventAction)
	t.Log(r.OrderNo)
	if err != nil {
		fmt.Printf("解析失败: %v\n", err)
		return
	}
}

func TestCallbackPaymentParse2(t *testing.T) {
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
	req := &http.Request{}
	paramStr := `out_biz_no=1221121221&gmt_create=2025-10-17+15%3A31%3A07&charset=UTF-8&seller_email=yuioplkjh202502%40163.com&subject=%E5%AE%A2%E6%9C%8D%3A+400%E2%80%93088%E2%80%935739&sign=dbNudH3ydq90yTvrI8HmXwj%2FkAS0Bda6xVvIiKTQLjjKenoYl8AZw%2F4%2FBu37D%2FCy66aX5zaWQwtrKNKeaHpMMACwBQLQngZLVy8puy74FvbfUeMAt5cYpQAI9n%2Bl2zsFanefNDY%2BCiEzQc8dwog8QttTmWICtgnxewDCgWY5%2B11cmobRUL5%2FqiWtclggm2egQ%2F1yxcAQFsLnOvpSE54uW8Fu7bLaKARwCagl3rDHn%2BKV592okyJK%2BcA%2FUDS%2B4g9qKrq%2FStoJniVwdwWaPbHJ66He60yQTjlQaTm8ml820iWXa7GggFHr7YXFfYSr2qaobChcCCl%2FwYlo9L25Oo2Q3g%3D%3D&buyer_open_id=099luGU2ef8yeC95sroZ9WGyuuGqL7qD-VLY8B1mtPqj4w9&invoice_amount=0.00&notify_id=2025101701222153119086991451124623&fund_bill_list=%5B%7B%22amount%22%3A%220.02%22%2C%22fundChannel%22%3A%22COUPON%22%7D%5D&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=0.02&buyer_pay_amount=0.02&app_id=2021005199606982&sign_type=RSA2&seller_id=2088070256235011&gmt_payment=2025-10-17+15%3A31%3A18&notify_time=2025-10-17+15%3A56%3A32&merchant_app_id=2021005199606982&passback_params=tpdp_id%3D17&version=1.0&out_trade_no=176068626732100521009&total_amount=0.02&trade_no=2025101722001486991458220176&auth_app_id=2021005199606982&buyer_logon_id=999***%40sina.com&point_amount=0.00`
	ctx := context.Background()
	req.Body = io.NopCloser(bytes.NewBuffer([]byte(paramStr)))
	r, err := CallbackPaymentParse(ctx, c, req)
	t.Log(r)
	t.Log(r.EventAction)
	if err != nil {
		fmt.Printf("解析失败: %v\n", err)
		return
	}
}
