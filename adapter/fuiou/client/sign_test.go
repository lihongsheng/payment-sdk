package client

import (
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/util"
	"testing"
)

func TestClient_GetResponseSignContent(t *testing.T) {
	c := config.Config{
		MchID: "0002900F0370542",
		RsaPrivate: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
		RsaPublic: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	s := NewSign(c)
	ss, err := s.DecryptByKey("zT+IwAosT7l4vS9N7GM6MI/kjVg6ldHb7CvsKZT87nU8DR6bQ55kzkqt5X8QmSbRVRMaNKVHC5BAoiEupJnSnIBj5DOeDh79ykkxNrTFq98SQ=", []byte(c.RsaPrivate))
	//ss, err := s.DecryptByKey("QXozoE8Rx6Qb6AkVcW90ZxxA7duSJWB+eB9cIFE02/KKZZWByB5aEEvBUVNiE/+VUbcMUNivCafWkkqCp5QSYOHGMLZmtW4etbe3GIJEf3MvKUCIUlnqwMpdUKsylAE5mu5dnuDgy7iJqTbXHtPctZyBqBfTYk1eEiMPzavlZhJiNIDXsLdZu1M2cRf69FWkZnYBu+331yxUQOdCda1dL1asFV6w82z11D29eGliJEwejIDJFVrOXBok5mZRPStrdccSKYHrd2OgVmLmEenPj8vnEP6f5dgJUnD4NuFWMN7gYvQQOXgMraobCjz9oboHjck946Lg3+O4oEj+cef6xpDXeIYbAXRCfQUPn7JoMEkX6d6MWq8zG6mCztT2RgSAvG0UFhjvlEH4W3FFebcw/YZHFiRkN8ZzJjbEy0dY0VJQh3XZAD5jS0OUG9ndO+wKPc8pg3qqog1uuoTgmo55DvqsOl8NghSfPZOAj9ItY0jwmJWBfHX9IeWgXNs1upkDax632FttLOS/kIEAp3foe/poYiroEVXgYrJE8duZGdYznm1CKWin3KPdenshIH3USpaOtEoSBumcX7jobDR2Cma1ra4V9u+LSU9zV5nqPACwPlsmIN8CVppCF9wn0Ae+Ykq+YozRuzouqfOylpGWWYhZqB9+a9XP6HXyvmVvTCw=", []byte(c.RsaPrivate))
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
		MchID: "0002900F0370542",
		RsaPrivate: `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAJMr8NnRV7ve7Y5FEBium/TsU0fK5NvzvFpsYxPAQhBXY+EN0Bi2JEg790C1njk9Q3U36u2JBDHAiDIomlgh6wWkJsFn7dghV/fCWSX1VVJ+dRINZy1432fRaJ8GqspvMneBpeLjBe94IwlWKpN+AOR+BNX8QL/uHmfCPlVQXos9AgMBAAECgYAzqbMs434m50UBMmFKKNF6kxNRGnpodBFktLO7FTybu/HF6TFp21a1PMe5IYhfk5AAsBZ6OCUOygWFhhdYZN+5W+dweF3kp1rLE4y5CjwqNlk/g22TAndf9znh/ltHFLvITToqu/eh/34tE1gyNxRbsi1olw/1wv8ZRjM3vtM9QQJBANvNwFq+CJHUyFzkXQB7+ycQFnY8wDq8Uw2Hv9ZMjgIntH7FSlJtdu5mAYPPo6f74slO5tFUMNP7EVppqsjYaNkCQQCraD6iKHo+OIlvvYIKiMXatJGD7N1GNhq5CrhUNPWLHwv/Ih2D3JJdF8IUZOPIJfUxTfM2fZYI+EVdsv6s4RcFAkAGjNYbnighOGcUJZYD6q3sVxVkRqEv3ubWs2HrH/Lna4l8caKqXCq8JfwLkod8/QugFiLYwBqIZqX4vMdjHtfZAkBsAl9dbWZCaPvpxp/4JWGPxDLhz9NLV/KU4bVvkoObq++yUHwKyGYOdVcd5MlIKOsNq5Hzp0Vw14lWVuF2bMxFAkBuNrZksvUULNIaWDKd4rQ6GVzUxXuIZW0ZE6atHYDiXPB4jVAjKRtLxZAV1qH9cr1zNJlcg+RbGYUdF9t4A9n5
-----END PRIVATE KEY-----`,
		RsaPublic: `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGg8ORm9u3s8o60tcgD7ZsL/0WR5vbR0krgmOkpm3yt1hgBoYABoUhiLTTt6V6N4RJg9YK2ku9J3hTNPmiGbfaUobcYs/9RLM1TCw9xyJESoQSRXKvKdIFh/pdAaZdMBsZX+Ltpk5H+PdGBIj/lvxMhJ5LELKCHR6bJ/v62dtApwIDAQAB
-----END PUBLIC KEY-----`,
		ApiHost: "https://richOperationFront-test.fuioupay.com",
	}
	s := NewSign(c)
	sss := "<xml><traceNo>1761902585812</traceNo><mchntCd>0002900F0370542</mchntCd><signature>LTXjUHybEfzR2tYkVjZxun2PQfWTKytrZqvq/ioa4VIc7QYG9OKC7VfBORCuMhb+rErc02pBUVlRxrqKnd3RhFHkvn4tOsHYIvsrdfhEnPfC75y1m4eEpShQ8QSvFsFaqJfX9OowttUMS86OuVnNVxiET4BjsK/Ygb0HG2o8GcQ=</signature><cleanType>02</cleanType><interBankNo>103222000905</interBankNo><outAcntNm>测试公司567</outAcntNm><mobile>13377566869</mobile><outAcntNo>623668271000025678</outAcntNo><allocateScale>10000</allocateScale><certTp>1</certTp><certNo>123456789</certNo><protocolType>03</protocolType><mchntCdUserId/><checkType>2</checkType><miniAppReturnPath/><channel>01</channel><organizationType>1</organizationType><bcpNo/><busiLicValidateStart>20131114</busiLicValidateStart><busiLicValidateEnd>20331113</busiLicValidateEnd><busiLicAddr>中国（上海）自由贸易试验区</busiLicAddr><busiLicPic>/20240607/busiLicPic.png</busiLicPic><legalName>测试法人</legalName><legalMobile>13663704488</legalMobile><legalCertTp>2</legalCertTp><legalCertNo>34042dfgdfgdf</legalCertNo><legalValidateStart>20050119</legalValidateStart><legalValidateEnd>20250119</legalValidateEnd><legalImagF>/20240607/legalImag.png</legalImagF><legalImagB>/20240607/legalImag.png</legalImagB><contactName>测试联系人</contactName><contactEmail>lh3334sd@fuioupay.com</contactEmail><contactCertNo>34042dfgdfgdf</contactCertNo><outAcntNoType/><extendInfo><mcc>5933</mcc><registerProvince>310</registerProvince><registerCity>2900</registerCity><registerDistrict>290C</registerDistrict><registeredCapital>30000000</registeredCapital><businessScope>电信业务</businessScope><contactValidDateStart>20050119</contactValidDateStart><contactValidDateEnd>20250119</contactValidDateEnd><contactPortraitFilePath>/20240607/legalImag.png</contactPortraitFilePath><contactBadgeFilePath>/20240607/legalImag.png</contactBadgeFilePath><cfcaPath>/20240607/legalImag.png</cfcaPath><beneficiaryInfoList><beneficiaryInfoList><name>测试人</name><certTp>0</certTp><certNo>34042dfgdfgdf</certNo><validDateStart>20050119</validDateStart><validDateEnd>20250119</validDateEnd><mobile>13666504488</mobile><address>中国（上海）自由贸易试验区</address><portraitFilePath>/20240607/legalImag.png</portraitFilePath><badgeFilePath>/20240607/legalImag.png</badgeFilePath><beneficiaryType>0</beneficiaryType></beneficiaryInfoList></beneficiaryInfoList><shareholderInfoList><shareholderInfoList><name>测试人</name><certTp>0</certTp><certNo>34042dfgdfgdf</certNo><validDateStart>20050119</validDateStart><validDateEnd>20250119</validDateEnd><portraitFilePath>/20240607/legalImag.png</portraitFilePath><badgeFilePath>/20240607/legalImag.png</badgeFilePath></shareholderInfoList></shareholderInfoList></extendInfo></xml>"
	sssGbk, _ := util.Utf8ToGbk(sss)
	en, _ := s.EncryptByPublicKey([]byte(sssGbk), []byte(c.RsaPublic))
	t.Log(en)
}
