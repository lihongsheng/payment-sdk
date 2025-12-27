package user

import (
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_AuthToken(t *testing.T) {
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
	// 6e19cbba50b848169d431e44ec32GX83
	// app_id=2021005199606982&charset=UTF-8&code=6e19cbba50b848169d431e44ec32GX83&format=json&grant_type=authorization_code&method=alipay.system.oauth.token&refresh_token=&sign_type=RSA2&timestamp=2025-10-15 16:12:57&version=1.0
	// app_id=2021005199606982&amp;charset=UTF-8&amp;code=6e19cbba50b848169d431e44ec32GX83&amp;format=json&amp;grant_type=authorization_code&amp;method=alipay.system.oauth.token&amp;sign_type=RSA2&amp;timestamp=2025-10-15 16:12:57&amp;version=1.0
	u, err := NewUser(c)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	//  {"alipay_system_oauth_token_response":{"access_token":"authusrB3379f70b688f4a63a3d5a5fdbea72X83","auth_start":"2025-10-15 13:19:21","expires_in":1296000,"re_expires_in":2592000,"refresh_token":"authusrB79e8d26898f3401b900b666db1aeeX83","open_id":"083luGU2ef8yeC95sroZ9WGyojnYtj3lBoBA6lY_2PlBrU4"},"sign":"GfIz3Q2Rwk2t72xXohXun9bDUrUxUmLr6ATzbaPTGCtBxzRHcN3JAkWJtdtuD/GFxt8lvYVagrFWb50F6n7CYy2dE6hu/7m19md+y6kV60BeORvKCypKC61EegKB4GjeA6uLKdSnjrZI8vaohrHsjw9SRV8lnQVKi/p50Y8XF35nHGIW4t11AvYMQy6+NyscP9bSDodIFXWxN9XT5vnOtd5av4ij8ZR6MFI9v5aWakAOZNWkM4P+GBsZ9ePCwiGD8Se5puvAjljUi5O3oriRaald0e8gPaz1t3BNwWteY6N91B+4EAlzy2oOmpME7ey6U0EDbXB2auQxbF+LKNbQkQ=="}%!(EXTRA <nil>) metadata = map[] cause = <nil>
	//  {"error_response":{"msg":"Invalid Arguments","code":"40002","sub_msg":"授权码code无效","sub_code":"isv.code-invalid"},"     sign":"RxKKJutoX1EMMIUXgv0YW8SVI8PUN8mS+m6irmULx/xe5rS4YfkWoTGBNAhdw4Ezd/JcVUT13ai878qjwJMDAtb85Kab5JH5MtpQ2f6tbCPn40q3jJMDUCYUFlJgm8YtBbfhau4PrXqA4CxwytQVUzbKgC89bUVakMUejeeUMTNGqpBfS2lcFhzmisimuBbfWrRwrNZddsQYwnolPEbmuHSWrLexkEYGGsM2YStF82qmRD70EpIvvgwLqW3aQMbxA95FX3Hzr/hrwui8yFKXsk9kHKP2fYyQPtgIMeIAIxlYu/fJu7b+kJQWd+d2yzecqkYAVnVIxLQ7kWc3hnKxkQ=="}
	r, err2 := u.AuthToken(ctx, &model.UserAuthRequest{
		GrantType: enum.AUTH_TYPE_AUTHORIZATION_CODE,
		Code:      "3cadac90d8444d8aa0842896a5caME99",
	})
	t.Log(r, err2)
	t.Log(r.OpenId)
	t.Log(r.GetExpiresIn())
	if err2 != nil {
		t.Log(err2.Error())
	}
	assert.NoError(t, err2)
}

func TestUser_AuthToken2(t *testing.T) {
	body1 := []byte(`{"alipay_system_oauth_token_response":{"access_token":"authusrB3379f70b688f4a63a3d5a5fdbea72X83","auth_start":"2025-10-15 13:19:21","expires_in":1296000,"re_expires_in":2592000,"refresh_token":"authusrB79e8d26898f3401b900b666db1aeeX83","open_id":"083luGU2ef8yeC95sroZ9WGyojnYtj3lBoBA6lY_2PlBrU4"},"sign":"GfIz3Q2Rwk2t72xXohXun9bDUrUxUmLr6ATzbaPTGCtBxzRHcN3JAkWJtdtuD/GFxt8lvYVagrFWb50F6n7CYy2dE6hu/7m19md+y6kV60BeORvKCypKC61EegKB4GjeA6uLKdSnjrZI8vaohrHsjw9SRV8lnQVKi/p50Y8XF35nHGIW4t11AvYMQy6+NyscP9bSDodIFXWxN9XT5vnOtd5av4ij8ZR6MFI9v5aWakAOZNWkM4P+GBsZ9ePCwiGD8Se5puvAjljUi5O3oriRaald0e8gPaz1t3BNwWteY6N91B+4EAlzy2oOmpME7ey6U0EDbXB2auQxbF+LKNbQkQ=="}`)
	body2 := []byte(`{"alipay_system_oauth_token_response":{"access_token":"authusrB3379f70b688f4a63a3d5a5fdbea72X83","auth_start":"2025-10-15 13:19:21","expires_in":"1296000","re_expires_in":"2592000","refresh_token":"authusrB79e8d26898f3401b900b666db1aeeX83","open_id":"083luGU2ef8yeC95sroZ9WGyojnYtj3lBoBA6lY_2PlBrU4"},"sign":"GfIz3Q2Rwk2t72xXohXun9bDUrUxUmLr6ATzbaPTGCtBxzRHcN3JAkWJtdtuD/GFxt8lvYVagrFWb50F6n7CYy2dE6hu/7m19md+y6kV60BeORvKCypKC61EegKB4GjeA6uLKdSnjrZI8vaohrHsjw9SRV8lnQVKi/p50Y8XF35nHGIW4t11AvYMQy6+NyscP9bSDodIFXWxN9XT5vnOtd5av4ij8ZR6MFI9v5aWakAOZNWkM4P+GBsZ9ePCwiGD8Se5puvAjljUi5O3oriRaald0e8gPaz1t3BNwWteY6N91B+4EAlzy2oOmpME7ey6U0EDbXB2auQxbF+LKNbQkQ=="}`)
	var model model.AuthTokenResponse
	err := json.Unmarshal(body1, &model)
	if err != nil {
		t.Error(err)
	}
	t.Log(model.AlipayOpenAuthTokenAppResponse.OpenId)
	t.Log(model.AlipayOpenAuthTokenAppResponse.GetExpiresIn())
	err = json.Unmarshal(body2, &model)
	if err != nil {
		t.Error(err)
	}
	t.Log(model.AlipayOpenAuthTokenAppResponse.GetReExpiresIn())

}
