package user

import (
	"context"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/fuiou/enum"
	"github.com/singer-stack-lab/payment-sdk/adapter/fuiou/model"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser_Create(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config.Config{
		AppID:  "appid",
		MchID:  "0002900F0370542",
		APIKey: "api_key",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	t.Log(ctx, user)
	r, err := user.Create(ctx, &model.UserCreateRequest{
		TraceNo:       "2",
		MchntCd:       c.MchID,
		CleanType:     enum.CleanType_Personal,
		OutAcntNm:     "米昂",
		CertTp:        enum.CertTp_IDCard,
		CertNo:        "110101199001011234",
		Mobile:        "13800000001",
		ContactName:   "米昂",
		ContactCertNo: "110101199001011234",
		// 银行卡号
		OutAcntNo:     "62221199001011234",
		ProtocolType:  "01",
		MchntCdUserId: "1",
	})
	t.Log(fmt.Sprintf("%+v", r))
	fmt.Println(fmt.Sprintf("%+v", err))
	assert.NoError(t, err)
	//t.Log(r)
}

func TestUser_Query(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config.Config{
		AppID:  "appid",
		MchID:  "0006530F8893215",
		APIKey: "",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAL1XipPCUPVYHpVGDz1bLYsjdPv5Kd6teuU6y3Fk+82LNas1wQ1oxRgel8X2yNulKyG3e7m281hDOir5b+bQmSZNSYkK3tRFtSDLoAxRa6hdWTBT6USy0VeOwgLqySJNHX7VEThaKsB5OK9Q/HhSQsMcFiky499JX03plGqulgL3AgMBAAECgYBLxOZtrssbA0Jp1flvQjd9gJjl5JW+yIlvhhG3tQOXA2hctUwmA5Erz8JItDM4wmX6PiTC8tV6spxqahs/sKY4HuIP896BEMyeQPMAJaeoPkcIE21CvRNQMix7tBD1xT8PQv42lqXuG2ekcvDx6TQFQR8/sG7PbP7hM0tXXrIxQQJBAO4XHyIZXOT6dwEZTj0gPnZMj/Lb61RmnOykBWuIsBZ1Cu0Qxrrs3Fi61NGMtdvn7MpOmC06+PzdyjX0ZVe4vfECQQDLla2TRwJductB0RfqeqPOpzP8L28YHZAyNPgzel8PT9Wsy3arhDRGEzjxR2DoCSFuahCWAUyjZehWyi7nEAdnAkA/olZj2ruFR5v+4zCSDcgj/hqLIlQnXrWaWcxJDWNi3S1qZw12cFAENwsiZqVEfRxAhfkOPbDGhNDC0uszIuFBAkEAstLQ2pL/ExDF5xQhK8d552WbFiMcNFSTemZotd8BbNO1XsiBfnrr57muxNDr4CVVBkWIOBsAFG7JPKLvu+qqdQJBAJBPcmXHRlqsIPkRCA4rWHrPGdhnQEKAoiYA1ctnCWtJe+IhFyhFEDYzQV+AYQIylKq3mWzA8QNCcijqGUp8U84=
-----END RSA PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC4nlZ0Uk4/hPlmGUtrqdMCzrZ9b3Tt/zSaoH1Bc/9WN5N1C1eVDARtWYKAVbw1eLE9hCrT0GLzOKxK2YEVSfdUmhRo49gOz19jvAqpXCeTsylaxeXbaF+3ylziQ+XBtg4a8f9rLp1kmmNmumqBgENv2dJhvId+dpZjuGU1ZO/MGwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://richfront.fuioupay.com",
	}
	user, err := NewUser(c)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	t.Log(ctx, user)
	r, err := user.Query(ctx, &model.UserQueryRequest{
		TraceNo:   fmt.Sprintf("%d", time.Now().Unix()),
		MchntCd:   c.MchID,
		AccountIn: "889321520251107501766222017",
		Mobile:    "15538626220",
		OutAcntNo: "6214835269091516",
		OutAcntNm: "阮乐乐",
	})
	t.Log(fmt.Sprintf("%+v", r))
	fmt.Println(fmt.Sprintf("%+v", err))
	assert.NoError(t, err)
}

func TestUser_Delete(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config.Config{
		AppID:  "appid",
		MchID:  "0002900F0370542",
		APIKey: "api_key",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	t.Log(ctx, user)
	r, err := user.Delete(ctx, &model.UserDeleteRequest{
		TraceNo:   "2",
		MchntCd:   c.MchID,
		AccountIn: "037054220230224000049871826",
	})
	t.Log(fmt.Sprintf("%+v", r))
	fmt.Println(fmt.Sprintf("%+v", err))
	assert.NoError(t, err)
}

func TestUser_RetryActive(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config.Config{
		AppID:  "appid",
		MchID:  "0002900F0370542",
		APIKey: "api_key",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c)
	if err != nil {
		t.Log(err.Error())
		t.Error(err)
	}
	ctx := context.Background()
	r, err := user.RetryActive(ctx, &model.UserActiveRequest{
		TraceNo:   fmt.Sprintf("%d", time.Now().UnixMilli()),
		MchntCd:   c.MchID,
		AccountIn: "037054220230224000049871826",
		CheckType: enum.CheckType_Mobile,
	})
	t.Log(fmt.Sprintf("%+v", r))
	fmt.Println(fmt.Sprintf("%+v", err))
	assert.NoError(t, err)
}

func TestUser_ChangeMobile(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config.Config{
		AppID:  "appid",
		MchID:  "0002900F0370542",
		APIKey: "api_key",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c)
	if err != nil {
		t.Log(err.Error())
		t.Error(err)
	}
	ctx := context.Background()
	r, err := user.ChangeMobile(ctx, &model.UserUpdateMobileRequest{
		TraceNo:   fmt.Sprintf("%d", time.Now().UnixMilli()),
		MchntCd:   c.MchID,
		AccountIn: "037054220230224000049871826",
		CheckType: enum.CheckType_Mobile,
		Mobile:    "18565761521",
		//ProtocolType: "01",
	})
	fmt.Println(fmt.Sprintf("%+v", r))
	fmt.Println(fmt.Sprintf("%+v", err))
	assert.NoError(t, err)
}

func TestUser_BindBank(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config.Config{
		AppID:  "appid",
		MchID:  "0002900F0370542",
		APIKey: "api_key",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c)
	if err != nil {
		t.Log(err.Error())
		t.Error(err)
	}
	ctx := context.Background()
	_, err2 := user.BindBank(ctx, &model.UserBindBankAccountRequest{
		TraceNo:   fmt.Sprintf("%d", time.Now().UnixMilli()),
		MchntCd:   c.MchID,
		AccountIn: "",
		CheckType: enum.CheckType_Mobile,
		Mobile:    "15538626220",
		OutAcntNo: "6214835269091516",
		OutAcntNm: "阮乐乐",
		//ProtocolType: "01",
	})
	assert.NoError(t, err2)
	//fmt.Println(fmt.Sprintf("%+v", r))
	//if err2 != nil {
	//	t.Log(err2.Error())
	//	t.Error(err2)
	//}
	//t.Log(fmt.Sprintf("%+v", r))
}

func TestUser_UnsetBank(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config.Config{
		AppID:  "appid",
		MchID:  "0002900F0370542",
		APIKey: "api_key",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c)
	if err != nil {
		t.Log(err.Error())
		t.Error(err)
	}
	ctx := context.Background()
	r, err := user.UnsetBank(ctx, &model.UserUnsetBankAccountRequest{
		TraceNo:   fmt.Sprintf("%d", time.Now().UnixMilli()),
		MchntCd:   c.MchID,
		AccountIn: "037054220230224000049871826",
		CheckType: enum.CheckType_Mobile,
		Mobile:    "18565761521",
		//OutAcntNo: "6217857500009654440",
		OutAcntNo: "6217857500009654440",

		//ProtocolType: "01",
	})
	fmt.Println(fmt.Sprintf("%+v", r))
	fmt.Println(fmt.Sprintf("%+v", err))
	assert.NoError(t, err)
}
