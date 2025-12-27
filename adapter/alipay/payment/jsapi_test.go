package payment

import (
	"context"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJsapi_Pay(t *testing.T) {
	c := config.Config{
		AppID: "2021005199606982",
		MchID: "",
		// 02941eef3187dddf3d3b83462e1dfcf6
		APIKey: "",
		Cert: config.Cert{
			CertificateSerialNumber: "5875781f2c5b439ed3999d3c21aa5562",
			CertPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC3fcEnCC2/djBnEYFDp41ZWuM4T49Pfr6p0gHXtfHGrO9CqegJhFPd4HT4BmVWEGgsHBMSiTce4b4VsrOQG0j7Lr2It0ziPpBwJUyQti5zZUHZoNl6FDVJIUo4p1DOxoqmYqj4iK/AgZVu0YiVNnbuKGWRoLt6B2++O05lX63qgAKxitqeRvGca73vbeiavUEhVaDVYbUBZOvLOuH/FDgLAHETOjXm59+Q7OEdT/cV+tltNyn/zTCUbql3I8MZQQkFcjOylREm+Jg7nGrK0o/9Bjw413LuQwe4Y0uXO4r7mUwV6QTS6L3L+8QzcBzpOUj1oC+iNM5Xe7b6/f8RssntAgMBAAECggEBAJg0xXM2MpSeWno0qBBRvUuAu/iq7krX7Rp7LLIdp8fRDcszm1nBnbvwA7b1vzuL8o2cYCnBDqscQIxJJvAD4h8R34C35BvMIA4WspNbu95XAS+gYKYGe0kFR0cFa1+Ki9qG/njjET7Tfvxk7TIw8qiNc+s/9CO+ZW/cTCSFIIPicYGwpIw1s15o8j/W9OMsxt7+RCiPdj0Kgu6SUAvwZ5hTCKFPO1sJJYU6DJJDCLSdKXrcez+fQKsWsimbHV1G3Jk4PFzaGruk9c+SUgUnPnRRAH/kxSa6bijNJjK/5apIXL0SIOfqztWV0ciG6evPu60L8nADb2dfKAzKudxtfnUCgYEA53Clnz3wlBStGqaOsR0FxngzS84Pe2uCguwcB1BYiVfY7OaN/a4YHm2MdaNfZI1sfJAbHiJJw2cbcb51iNOlWI6KMLqpXJzz2GLhhf8BlZuVQwtY/ogRTFOCrejs+3LarZR1ct8b56193wCh90Q4ZysCLxZLdvV0Ey56vMJpc8MCgYEAyvaE+iRdsdYj2g/OaBfKRxB3BQWlg/aO8169g7f7/yCoDkMZGZaLI/TiyEp6ja9J+OYt74aIwrAOlhXZOMIpkxuPa+opzRpYExPCAG1FUV5xtXYomTLHGZpY/VURu0HVx0GMvritzgDllTol7U+hHaF/ziGZiz637o+AvJD2YI8CgYBU18XPd6xvDJlc0Lw0j3gjhsL/Qh4I16OzjQzFXZ3nU23xgp+A5MZwuSYpped9fB7OFyHpzRYPbQfxjamWrEfMDAz6yiL2EY+OvskADziseKmiN1G7lXQJ7/9S87WUhElIIslfgAzBxKcFDj1R4Q9wOqMFUf3+MQMJWRujgP2ZcwKBgCaiVF+1KmyeoYZxuc2Qsb1jQfSfxYjgFwf2gcaf8AT5d2P3m8CGKog2pjCQFNIBiugpEJdmmYPNtYkWMZofQ4CwH1XgZgKXwCopeaVRJ3+8pZJwInw+8S1LdyMJ3W0ys/nQ3XS0fMkY3JrSTcPlg3q7cjOPq8WYs7RidpPuWCe7AoGBAM7j6/hIktBmsCsjZgXQOvH+I4RpXVg6finIRDEjjPPlpujCQZuIUL8/HXznyoT/yuqTFFOLTLOSrgASwNp175WlgCpZ+RNktsczMP/PfZTqzQWyYf/gzRH9OjdD8E+LTYtH89gfiO9XtdxYxE0ke8Bz7CQLFbFfH8vrTelkJq2n
-----END RSA PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsyA5/k4SyfoHhZCPIr/VjBG1vRhOwPLoOddf5WCH6168h/el+BGMiNOdczbP11HmrfpJWsXKDR7W/T/RMwX032RY2ah7HTJG1Auh5RXu4SoKmZYq84nvT0i0N8hp222E1fECts1NaWSfQP8dZcbFTd5HEWwbBmdFR/ExWl2u3wZB0OFHgIz4rfVV3qQnqFLzPqYIV23QfkHS0cVrlZU+GjSO7BVLRozOPD+0O1N9cuq8gjqSYDbR83VTAG27hfJel3prg5FTDAeLSi1eyUIrNRB3LZjaj5X90gyRHY4nNOx7Jh01YTuZR/FSHVHVdGdkw3h131FqmJpB2xAdc/qnVwIDAQAB
-----END PUBLIC KEY-----`,
			// 687b59193f3f462dd5336e5abf83c5d8_
			PublicKeyID: "02941eef3187dddf3d3b83462e1dfcf6",
		},
		Proxy:   config.Proxy{},
		ApiHost: "",
		Version: "",
	}
	pay, err := NewJsApi(c)
	assert.NoError(t, err)
	ctx := context.Background()
	p, err := pay.Pay(ctx, &dto.PayOrder{
		Order: dto.Order{
			OrderNo: "Test11110003",
			Amount: dto.Amount{
				Total: 1,
			},
			PayAmount: dto.Amount{
				Total: 1,
			},
			Goods:    nil,
			Subject:  "Test",
			Desc:     "",
			CreateAt: time.Now().Add(24 * time.Hour),
		},
		Payer: dto.Payer{
			OpenID: "083luGU2ef8yeC95sroZ9WGyojnYtj3lBoBA6lY_2PlBrU4",
		},
		RedirectUrl:    "",
		TimeExpire:     0,
		NotifyUrl:      "https://api-cabinet.test.jianxindianzi.com//public/v1/callback/payment/20/Test11110001",
		PassbackParams: "",
		SettleInfo:     nil,
		SceneInfo:      nil,
		RiskFund:       nil,
	})
	if err != nil {
		t.Log(err.Error())
	}
	assert.NoError(t, err)
	t.Log(p)
}

func TestJsapi_Query(t *testing.T) {
	c := config.Config{
		AppID:  "2021005199606982",
		MchID:  "",
		APIKey: "",
		Cert: config.Cert{
			CertificateSerialNumber: "5875781f2c5b439ed3999d3c21aa5562",
			CertPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC3fcEnCC2/djBnEYFDp41ZWuM4T49Pfr6p0gHXtfHGrO9CqegJhFPd4HT4BmVWEGgsHBMSiTce4b4VsrOQG0j7Lr2It0ziPpBwJUyQti5zZUHZoNl6FDVJIUo4p1DOxoqmYqj4iK/AgZVu0YiVNnbuKGWRoLt6B2++O05lX63qgAKxitqeRvGca73vbeiavUEhVaDVYbUBZOvLOuH/FDgLAHETOjXm59+Q7OEdT/cV+tltNyn/zTCUbql3I8MZQQkFcjOylREm+Jg7nGrK0o/9Bjw413LuQwe4Y0uXO4r7mUwV6QTS6L3L+8QzcBzpOUj1oC+iNM5Xe7b6/f8RssntAgMBAAECggEBAJg0xXM2MpSeWno0qBBRvUuAu/iq7krX7Rp7LLIdp8fRDcszm1nBnbvwA7b1vzuL8o2cYCnBDqscQIxJJvAD4h8R34C35BvMIA4WspNbu95XAS+gYKYGe0kFR0cFa1+Ki9qG/njjET7Tfvxk7TIw8qiNc+s/9CO+ZW/cTCSFIIPicYGwpIw1s15o8j/W9OMsxt7+RCiPdj0Kgu6SUAvwZ5hTCKFPO1sJJYU6DJJDCLSdKXrcez+fQKsWsimbHV1G3Jk4PFzaGruk9c+SUgUnPnRRAH/kxSa6bijNJjK/5apIXL0SIOfqztWV0ciG6evPu60L8nADb2dfKAzKudxtfnUCgYEA53Clnz3wlBStGqaOsR0FxngzS84Pe2uCguwcB1BYiVfY7OaN/a4YHm2MdaNfZI1sfJAbHiJJw2cbcb51iNOlWI6KMLqpXJzz2GLhhf8BlZuVQwtY/ogRTFOCrejs+3LarZR1ct8b56193wCh90Q4ZysCLxZLdvV0Ey56vMJpc8MCgYEAyvaE+iRdsdYj2g/OaBfKRxB3BQWlg/aO8169g7f7/yCoDkMZGZaLI/TiyEp6ja9J+OYt74aIwrAOlhXZOMIpkxuPa+opzRpYExPCAG1FUV5xtXYomTLHGZpY/VURu0HVx0GMvritzgDllTol7U+hHaF/ziGZiz637o+AvJD2YI8CgYBU18XPd6xvDJlc0Lw0j3gjhsL/Qh4I16OzjQzFXZ3nU23xgp+A5MZwuSYpped9fB7OFyHpzRYPbQfxjamWrEfMDAz6yiL2EY+OvskADziseKmiN1G7lXQJ7/9S87WUhElIIslfgAzBxKcFDj1R4Q9wOqMFUf3+MQMJWRujgP2ZcwKBgCaiVF+1KmyeoYZxuc2Qsb1jQfSfxYjgFwf2gcaf8AT5d2P3m8CGKog2pjCQFNIBiugpEJdmmYPNtYkWMZofQ4CwH1XgZgKXwCopeaVRJ3+8pZJwInw+8S1LdyMJ3W0ys/nQ3XS0fMkY3JrSTcPlg3q7cjOPq8WYs7RidpPuWCe7AoGBAM7j6/hIktBmsCsjZgXQOvH+I4RpXVg6finIRDEjjPPlpujCQZuIUL8/HXznyoT/yuqTFFOLTLOSrgASwNp175WlgCpZ+RNktsczMP/PfZTqzQWyYf/gzRH9OjdD8E+LTYtH89gfiO9XtdxYxE0ke8Bz7CQLFbFfH8vrTelkJq2n
-----END RSA PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsyA5/k4SyfoHhZCPIr/VjBG1vRhOwPLoOddf5WCH6168h/el+BGMiNOdczbP11HmrfpJWsXKDR7W/T/RMwX032RY2ah7HTJG1Auh5RXu4SoKmZYq84nvT0i0N8hp222E1fECts1NaWSfQP8dZcbFTd5HEWwbBmdFR/ExWl2u3wZB0OFHgIz4rfVV3qQnqFLzPqYIV23QfkHS0cVrlZU+GjSO7BVLRozOPD+0O1N9cuq8gjqSYDbR83VTAG27hfJel3prg5FTDAeLSi1eyUIrNRB3LZjaj5X90gyRHY4nNOx7Jh01YTuZR/FSHVHVdGdkw3h131FqmJpB2xAdc/qnVwIDAQAB
-----END PUBLIC KEY-----`,
			// 687b59193f3f462dd5336e5abf83c5d8_
			PublicKeyID: "02941eef3187dddf3d3b83462e1dfcf6",
		},
		Proxy:   config.Proxy{},
		ApiHost: "",
		Version: "",
	}
	pay, err := NewJsApi(c)
	assert.NoError(t, err)
	ctx := context.Background()
	p, err := pay.Query(ctx, dto.Query{
		OrderNo: "",
		TradeNo: "2025101622001429831436657633",
	})
	assert.NoError(t, err)
	t.Log(p)
}

func TestJsapi_Query2(t *testing.T) {
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
	pay, err := NewJsApi(c)
	assert.NoError(t, err)
	ctx := context.Background()
	p, err := pay.Query(ctx, dto.Query{
		OrderNo: "Test11110003",
		TradeNo: "",
	})
	assert.NoError(t, err)
	t.Log(p)
}
