package payment

//
//func TestH5Payment(t *testing.T) {
//	c := config.Config{
//		AppID:  "2021005199606982",
//		MchID:  "",
//		APIKey: "",
//		Cert: config.Cert{
//			CertPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
//MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCVHxmnZCHt57/RGbQQCfagILP37StXMV9H9xmPNL5UDusEfvzZZaUJ/xrneiLgGkmrVYSlvvu4VqNWUmufHCFiOuSZawlXud2hhTLYOjMme31mgXfwA/LD2zL273Njl/WQcnSjDZ+/i8/VieESSd+VfB2rem0EMSg2IskL8FRtjOKUFH7UxhLltBwieKYRJhrAC3d1iK/AYxq8580kkhnPRkEq97nQDhmObxPytB2jdbrg5VcuJul98wSrh/rmK6NmKwYxoEHXkXtazrPOMcWepcjvPc6LGls/pkjPHTsNuQ2zok2eouXxT4pXsKaXat//lC4DcN7ZYLMgXW6U1GghAgMBAAECggEAT0WcJp8VZloXXzoRvaD+STt5UGf07UIiy6fxH27Dav4PM3GqSos9Q+YoUDteRNHUrc8lV9MCD4mHBQGFkqyqloSGN4RuEAOPbSjY0ZeGz4NAM3jJ4H7I1BM3sYuzRJNoiGQ/2toIPRQ+5N6KFmXNBLNZNmo7f8n9/WFajQ0yZJV7r1dUzyylUYoi3MXJcZdZ9GPtYnDgm1mYxTFiUZH8Vs7SIARDV2PXsTPGrPoCuvWGLjGJ70TmD9lf7qbT2XQ9ECA837o+mB/DRCXPkW1YZibEo/dSn1dgN45va3OwLlKmHP9nmas49fACajENXEcH31OSLMKn+/8ibf4jg8WxQQKBgQDGdzFrwc1zcTNeZwpqSiqPHkveikA0jPKhJXQEQ7L4LqvlXrbNZOztMKegtoiXcXylFs6ktVB6DoXJRO7ko9lj0wXku68ykPg9lTKRBUPwxAqylwvXr27kEraYbc01UNMjYdWoRmW2r0xWYuzKCgt0OaBh8I996mJtwBlqHr1TOQKBgQDAWemqJaWMk5T7b0jynfAMjG639kgCuFP4ieibsCEnks82wABh9CQCUftjJuKJra0Y90Jo9LDcsiEhJzj2g2Ux/ElfGMJ6UAnL4nNZRsuIKIYQ4SYqkaywRUpXrqSfppzT+gIMIRIkaTFWyNl32+YcPgYvjLpEQ43dOw0Bg+60KQKBgA/ZoRbMCY54sfOOIyTN+4wmjUSeofYQc6gvB473oSl1AMF2yP6qWsmuoxSQv2Po6l187s/0sxKsaj7ixHl6JHh2d/gGjT1qmojAi6NNWlm2LmyI56c4GZEZdX/C9SqW4+XlgFzKEoF+iogzqlmQZ83GdGqd/be+qlG9j8oL3m7xAoGBAJWqIC7cpghQmL5e46KrkykgZ9HJ1lQPen7LR2tESzfQely+Xk3ZOd94TKLkkaXQtfvUAB9MLQU7sJ1sOF1O9YFELQ7KZB/pgQXcCCZx+FHUPiSKGzlKKdTpxSOeZsc2S5aAU/F+FfCwzMwa8WLafhyBgiyjUhdlfM+jb8Y7EpOpAoGAIFM8TGpNBPL/RK8uJ1QBXTn6n1b6SDGpj3HUZyEaa20BORRF7+mHvLkeXhLH09XaSmyI3+igK67/Exj07DkDJsaTA/xlYvzqkpCds2wZHs+/rmC5KC4tM46wxAkOK1VtjAvhh+4FnHOUGjhTOS6hAba5wP3sgkHfTwfRYDWDkQw=
//-----END RSA PRIVATE KEY-----`,
//			PublicKey: `-----BEGIN PUBLIC KEY-----
//MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsyA5/k4SyfoHhZCPIr/VjBG1vRhOwPLoOddf5WCH6168h/el+BGMiNOdczbP11HmrfpJWsXKDR7W/T/RMwX032RY2ah7HTJG1Auh5RXu4SoKmZYq84nvT0i0N8hp222E1fECts1NaWSfQP8dZcbFTd5HEWwbBmdFR/ExWl2u3wZB0OFHgIz4rfVV3qQnqFLzPqYIV23QfkHS0cVrlZU+GjSO7BVLRozOPD+0O1N9cuq8gjqSYDbR83VTAG27hfJel3prg5FTDAeLSi1eyUIrNRB3LZjaj5X90gyRHY4nNOx7Jh01YTuZR/FSHVHVdGdkw3h131FqmJpB2xAdc/qnVwIDAQAB
//-----END PUBLIC KEY-----`,
//		},
//		Proxy:   config.Proxy{},
//		ApiHost: "",
//		Version: "",
//	}
//	pay, err := NewH5(c)
//	assert.NoError(t, err)
//	ctx := context.Background()
//	p, err := pay.Pay(ctx, &dto.PayOrder{
//		Order: dto.Order{
//			OrderNo: "Test11110001",
//			Amount: dto.Amount{
//				Total: 1,
//			},
//			PayAmount: dto.Amount{
//				Total: 1,
//			},
//			Goods:    nil,
//			Subject:  "Test",
//			Desc:     "",
//			CreateAt: time.Now().Add(time.Hour * 2),
//		},
//		Payer: dto.Payer{
//			OpenID: "6cdb9e12101f4ec69525fc2d149cAD83",
//		},
//		RedirectUrl:    "",
//		TimeExpire:     0,
//		NotifyUrl:      "https://api-cabinet.test.jianxindianzi.com//public/v1/callback/payment/20/Test11110001",
//		PassBackParams: "",
//		SettleInfo:     nil,
//		SceneInfo:      nil,
//		RiskFund:       nil,
//	})
//	assert.NoError(t, err)
//	t.Log(p)
//}
