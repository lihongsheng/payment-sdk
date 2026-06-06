package client

import (
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	"testing"
)

const (
	testRsaPrivate = `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`
	testRsaPublic = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`
)

func newTestConfig() config.Config {
	return config.Config{
		Merchant: config.Merchant{MchID: "0002900F0370542"},
		Cert: config.Cert{
			RsaPrivate: testRsaPrivate,
			RsaPublic:  testRsaPublic,
		},
		API: config.API{ApiHost: "https://richOperationFront-test.fuioupay.com"},
	}
}

func TestClient_GetResponseSignContent(t *testing.T) {
	c := newTestConfig()
	s := NewSign(c)
	ss, err := s.DecryptByKey("zT+IwAosT7l4vS9N7GM6MI/kjVg6ldHb7CvsKZT87nU8DR6bQ55kzkqt5X8QmSbRVRMaNKVHC5BAoiEupJnSnIBj5DOeDh79ykkxNrTFq98SQ=", []byte(c.Cert.RsaPrivate))
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(";;;;;;;;;;;")

	ssUtf8, _ := util.GBKToUTF8Byte(ss)
	t.Log("ssssss", string(ssUtf8))
	t.Log(len(ss))
}

func TestClient_EncryptByPublicKey(t *testing.T) {
	c := newTestConfig()
	s := NewSign(c)
	sss := "<xml><traceNo>1761902585812</traceNo><mchntCd>0002900F0370542</mchntCd><signature>LTXjUHybEfzR2tYkVjZxun2PQfWTKytrZqvq/ioa4VIc7QYG9OKC7VfBORCuMhb+rErc02pBUVlRxrqKnd3RhFHkvn4tOsHYIvsrdfhEnPfC75y1m4eEpShQ8QSvFsFaqJfX9OowttUMS86OuVnNVxiET4BjsK/Ygb0HG2o8GcQ=</signature></xml>"
	sssGbk, _ := util.Utf8ToGbk(sss)
	en, _ := s.EncryptByPublicKey([]byte(sssGbk), []byte(c.Cert.RsaPublic))
	t.Log(en)
}