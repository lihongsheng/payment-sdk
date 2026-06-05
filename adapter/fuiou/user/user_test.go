package user

import (
	"context"
	"fmt"
	config2 "github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser_Create(t *testing.T) {

}

func TestUser_Query(t *testing.T) {

}

func TestUser_Delete(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config2.Config{
		MchID: "0002900F0370542",
		RsaPrivate: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
		RsaPublic: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c, nil)
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
	c := config2.Config{
		MchID: "0002900F0370542",
		RsaPrivate: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
		RsaPublic: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c, nil)
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
	c := config2.Config{
		MchID: "0002900F0370542",
		RsaPrivate: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
		RsaPublic: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c, nil)
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
		Mobile:    "1856576****",
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
	c := config2.Config{
		MchID: "0002900F0370542",
		RsaPrivate: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
		RsaPublic: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c, nil)
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
		Mobile:    "15538***220",
		OutAcntNo: "62148****909****",
		OutAcntNm: "阮**",
	})
	assert.NoError(t, err2)

}

func TestUser_UnsetBank(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Error(err)
		}
	}()
	c := config2.Config{
		MchID: "0002900F0370542",
		RsaPrivate: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
		RsaPublic: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	user, err := NewUser(c, nil)
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
		OutAcntNo: "6217857500009654440",
	})
	fmt.Println(fmt.Sprintf("%+v", r))
	fmt.Println(fmt.Sprintf("%+v", err))
	assert.NoError(t, err)
}
