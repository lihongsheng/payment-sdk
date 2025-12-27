package client

import (
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	"github.com/lihongsheng/payment-sdk/config"
	"testing"
)

func TestClient_GetResponseSignContent(t *testing.T) {
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
	s := NewSign(c)
	ss, err := s.DecryptByKey("zT+IwAosT7l4vS9N7GM6MI/kjVg6ldHb7CvsKZT87nU8DR6bQ55kzkqt5X8QmSbRVRMaNKVHC5BAoiEupJnSnIBj5DOeDh79ykkxNrTFq98SQ=", []byte(c.Cert.CertPrivateKey))
	//ss, err := s.DecryptByKey("QXozoE8Rx6Qb6AkVcW90ZxxA7duSJWB+eB9cIFE02/KKZZWByB5aEEvBUVNiE/+VUbcMUNivCafWkkqCp5QSYOHGMLZmtW4etbe3GIJEf3MvKUCIUlnqwMpdUKsylAE5mu5dnuDgy7iJqTbXHtPctZyBqBfTYk1eEiMPzavlZhJiNIDXsLdZu1M2cRf69FWkZnYBu+331yxUQOdCda1dL1asFV6w82z11D29eGliJEwejIDJFVrOXBok5mZRPStrdccSKYHrd2OgVmLmEenPj8vnEP6f5dgJUnD4NuFWMN7gYvQQOXgMraobCjz9oboHjck946Lg3+O4oEj+cef6xpDXeIYbAXRCfQUPn7JoMEkX6d6MWq8zG6mCztT2RgSAvG0UFhjvlEH4W3FFebcw/YZHFiRkN8ZzJjbEy0dY0VJQh3XZAD5jS0OUG9ndO+wKPc8pg3qqog1uuoTgmo55DvqsOl8NghSfPZOAj9ItY0jwmJWBfHX9IeWgXNs1upkDax632FttLOS/kIEAp3foe/poYiroEVXgYrJE8duZGdYznm1CKWin3KPdenshIH3USpaOtEoSBumcX7jobDR2Cma1ra4V9u+LSU9zV5nqPACwPlsmIN8CVppCF9wn0Ae+Ykq+YozRuzouqfOylpGWWYhZqB9+a9XP6HXyvmVvTCw=", []byte(c.Cert.CertPrivateKey))
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(";;;;;;;;;;;")

	ssUtf8, _ := util.GBKToUTF8Byte(ss)
	t.Log("ssssss", string(ssUtf8))
	t.Log(len(ss))
}

func TestClient_EncryptByPublicKey(t *testing.T) {
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
	s := NewSign(c)
	sss := "<xml><traceNo>1761902585812</traceNo><mchntCd>0002900F0370542</mchntCd><signature>LTXjUHybEfzR2tYkVjZxun2PQfWTKytrZqvq/ioa4VIc7QYG9OKC7VfBORCuMhb+rErc02pBUVlRxrqKnd3RhFHkvn4tOsHYIvsrdfhEnPfC75y1m4eEpShQ8QSvFsFaqJfX9OowttUMS86OuVnNVxiET4BjsK/Ygb0HG2o8GcQ=</signature><cleanType>02</cleanType><interBankNo>103222000905</interBankNo><outAcntNm>测试公司567</outAcntNm><mobile>13377566869</mobile><outAcntNo>623668271000025678</outAcntNo><allocateScale>10000</allocateScale><certTp>1</certTp><certNo>123456789</certNo><protocolType>03</protocolType><mchntCdUserId/><checkType>2</checkType><miniAppReturnPath/><channel>01</channel><organizationType>1</organizationType><bcpNo/><busiLicValidateStart>20131114</busiLicValidateStart><busiLicValidateEnd>20331113</busiLicValidateEnd><busiLicAddr>中国（上海）自由贸易试验区</busiLicAddr><busiLicPic>/20240607/busiLicPic.png</busiLicPic><legalName>测试法人</legalName><legalMobile>13663704488</legalMobile><legalCertTp>2</legalCertTp><legalCertNo>34042dfgdfgdf</legalCertNo><legalValidateStart>20050119</legalValidateStart><legalValidateEnd>20250119</legalValidateEnd><legalImagF>/20240607/legalImag.png</legalImagF><legalImagB>/20240607/legalImag.png</legalImagB><contactName>测试联系人</contactName><contactEmail>lh3334sd@fuioupay.com</contactEmail><contactCertNo>34042dfgdfgdf</contactCertNo><outAcntNoType/><extendInfo><mcc>5933</mcc><registerProvince>310</registerProvince><registerCity>2900</registerCity><registerDistrict>290C</registerDistrict><registeredCapital>30000000</registeredCapital><businessScope>电信业务</businessScope><contactValidDateStart>20050119</contactValidDateStart><contactValidDateEnd>20250119</contactValidDateEnd><contactPortraitFilePath>/20240607/legalImag.png</contactPortraitFilePath><contactBadgeFilePath>/20240607/legalImag.png</contactBadgeFilePath><cfcaPath>/20240607/legalImag.png</cfcaPath><beneficiaryInfoList><beneficiaryInfoList><name>测试人</name><certTp>0</certTp><certNo>34042dfgdfgdf</certNo><validDateStart>20050119</validDateStart><validDateEnd>20250119</validDateEnd><mobile>13666504488</mobile><address>中国（上海）自由贸易试验区</address><portraitFilePath>/20240607/legalImag.png</portraitFilePath><badgeFilePath>/20240607/legalImag.png</badgeFilePath><beneficiaryType>0</beneficiaryType></beneficiaryInfoList></beneficiaryInfoList><shareholderInfoList><shareholderInfoList><name>测试人</name><certTp>0</certTp><certNo>34042dfgdfgdf</certNo><validDateStart>20050119</validDateStart><validDateEnd>20250119</validDateEnd><portraitFilePath>/20240607/legalImag.png</portraitFilePath><badgeFilePath>/20240607/legalImag.png</badgeFilePath></shareholderInfoList></shareholderInfoList></extendInfo></xml>"
	sssGbk, _ := util.Utf8ToGbk(sss)
	// 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
	en, _ := s.EncryptByPublicKey([]byte(sssGbk), []byte(c.Cert.PublicKey))
	t.Log(en)
}
