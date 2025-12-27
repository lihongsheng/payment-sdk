package fund

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/stretchr/testify/assert"
)

func TestQueryBalance(t *testing.T) {
	cccc := config.Config{
		AppID: "wxc4cbb00d34676a39",
		MchID: "1734655963",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNFZ7+CFXv+cY6
aZA3qrObYdkq2Hr5nrziw728o6spMBlux5LYCiqsiCcVk2EHdGVG+NkL4+7qFZZn
bfFLOlJ71AWEY6AxCUDzQU0FlTLfbVvoBTfuD1cHLJ4eWOVxd69ccSDmyqowML3l
rIy5VkSxFPN9MMbbjBinCKEcq7qqp89rGJ/DHXTfDWlze3kIQ1I5pkIpOHBYCJpw
TCIMjZaD7RyRs3g6Lor0kgGKF/wvJpD1y97DOH31ae3q1mUuCAvCgHsaVJ8Xy3Fe
FHjj5k60e490ZSVIxquaLKgdLCOHzanNM0b2hBge4JFs5WiefnldbFoSCLnmZoFM
Z1k8NLgJAgMBAAECggEBALKBFjFAaiSFQD/JcEizoL8nLNH5ORiaTQgHnddakfk4
r3pX5X17Y+dmXraN4A+QBTMAGFMcIvDGt3HxJOv4qKC1S9QOrvjfeBQDC5pHYc9N
LYjHOgZQfcx1zm7Vq2fW5+DLavrW/ckSM8d8J8SNzR5hAQo/cWbZCwAgJ7N/2er8
tE6o29QIGDRfXknVPQILQ4szN09u+wAKDN2fqWMJK93LYnRFImkG95GvOyNNeFST
goeYzTSa9EDouXbpc/6yi+dWtirSSdz9JcNJCSPwNkGeN8gY9g2oqmIIyleuAJSh
InOc7FVS6GN0T96wF/MURqLikuxd8If7wYHT3QuM4k0CgYEA5uWjFdOQBowI+TSm
fNgfgbxdsg5wXYWC74CcQY1VyYfzc4KCr9WkCp4q7QKzuNl4EcfZblFQPOtlFiC4
HKP/PWrahEGHbABzR8/COvnXJlyFbZYBZcqQx6UVpEAbLinr1KeNUdxM2+49AEDC
DNiwxPLavg1JbZ3IXUvZcckt69MCgYEA42GQi2GezD+NDezE9/2bBt3Sj9CBpfqG
y2qP0NTMco3tf0m9pOXNes0TWMlcNqbfNRWwUhwbotD989FW8DROIastgPinew1w
jgu3i0RRLo43o7kXo1VgfACZtnyNu5eGeSW9p0qEb6+1hFnEJD1E1fd32Hl7LEkj
FPA3hvRRLzMCgYEAu3suhuR4B7jg+Gi13p0mSBOJCLEMsANeo9YMCnYWeEM5fEs2
jkusnSp8tGIDSE7cXaOmexrvSefg9qyukiLfdyEyOT10Fk/vSxn5HiYLeoyZkrfA
fsSR9mmnPu0GpN9JLCy4ZQW6KsIxKelrJz8zmVeQIc1sF+OL86VI9ElkTh0CgYBa
mIT/A+ZBexC4e53/MQyTr841ft4pQ6hxZAPpMOBExEfcb4UgLm+wfgU3WwNiwxos
DHg0Pv7D4IFOpBn/mteGkq2OtYQFg1VVQ1XeJ1oxMnj4wsaoTVwkDBkynG3cUIND
wcXO4eHinh+0nA+KYN8MYY1eswhgnMNdlFiLbMzsAQKBgHY4qvew/WbVJIT292vH
2Llytr4O3R/f/6FEbTu6FdHoVdrjhmBunps9/PQPZOGUkprpp9wJ0Z5a1WXOHN8H
JDfXQyKpPZJzNia5r3AIZ7YrYfj18dBDkEAIOqQmCvQeU56BkbGQFHOMUWIHG1Jt
1CUZZKuSZ48ycukR6Y5Z3ELt
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "6CF729BC6F8B1596BB1C4CD9DB6338C9D02E1694",
			PublicKeyID:             "PUB_KEY_ID_0117346559632025120500192382001802",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Gj4CGj9c5zzjWw0PQOU
oX4O+cP8vQVokLuJ4DkRUV6kdGkmcDWNSKJydWDOKGH4IrEbqsiLAm00ghU6TWVO
lgNpNZmXMMGWQPH60SycHUW89J9A+8JOkwo4/J/ubf/iPCfaTvEKCWaNxcO9AI0g
sMcAWgRtuvSYLqAUZZ9hycqZYGfQQyvN4FgAN3NbC6x7t9SAfY21uaClMuZTqeTs
KHLzAJWnTJ/5dBlF9oaz78wwihDk7zEvPHH6y218kGKwWuDNARCZk13vdQQuP6V8
fjIYkPGo2hCHp61wNaKwwEJHPRm7Lw+zxKS4ywJeUJCbZd5N2CxNGQu5M/AEIvL0
iwIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "0a3cd5200286392cdb400fa6e0aa09c8",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := FundApiService{Client: api.Client}
	ctx := context.Background()

	// 测试查询基本户余额
	resp, apiR, err := c.QueryBalance(ctx, QueryBalanceRequest{
		AccountType: AccountType_BASIC,
	})
	assert.NoError(t, err)
	t.Log(fmt.Sprintf("Response: %+v", resp))
	if apiR != nil {
		t.Log(fmt.Sprintf("StatusCode: %d", apiR.Response.StatusCode))
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
}

func TestQueryBalanceOperation(t *testing.T) {
	cccc := config.Config{
		AppID: "wxc4cbb00d34676a39",
		MchID: "1734655963",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNFZ7+CFXv+cY6
aZA3qrObYdkq2Hr5nrziw728o6spMBlux5LYCiqsiCcVk2EHdGVG+NkL4+7qFZZn
bfFLOlJ71AWEY6AxCUDzQU0FlTLfbVvoBTfuD1cHLJ4eWOVxd69ccSDmyqowML3l
rIy5VkSxFPN9MMbbjBinCKEcq7qqp89rGJ/DHXTfDWlze3kIQ1I5pkIpOHBYCJpw
TCIMjZaD7RyRs3g6Lor0kgGKF/wvJpD1y97DOH31ae3q1mUuCAvCgHsaVJ8Xy3Fe
FHjj5k60e490ZSVIxquaLKgdLCOHzanNM0b2hBge4JFs5WiefnldbFoSCLnmZoFM
Z1k8NLgJAgMBAAECggEBALKBFjFAaiSFQD/JcEizoL8nLNH5ORiaTQgHnddakfk4
r3pX5X17Y+dmXraN4A+QBTMAGFMcIvDGt3HxJOv4qKC1S9QOrvjfeBQDC5pHYc9N
LYjHOgZQfcx1zm7Vq2fW5+DLavrW/ckSM8d8J8SNzR5hAQo/cWbZCwAgJ7N/2er8
tE6o29QIGDRfXknVPQILQ4szN09u+wAKDN2fqWMJK93LYnRFImkG95GvOyNNeFST
goeYzTSa9EDouXbpc/6yi+dWtirSSdz9JcNJCSPwNkGeN8gY9g2oqmIIyleuAJSh
InOc7FVS6GN0T96wF/MURqLikuxd8If7wYHT3QuM4k0CgYEA5uWjFdOQBowI+TSm
fNgfgbxdsg5wXYWC74CcQY1VyYfzc4KCr9WkCp4q7QKzuNl4EcfZblFQPOtlFiC4
HKP/PWrahEGHbABzR8/COvnXJlyFbZYBZcqQx6UVpEAbLinr1KeNUdxM2+49AEDC
DNiwxPLavg1JbZ3IXUvZcckt69MCgYEA42GQi2GezD+NDezE9/2bBt3Sj9CBpfqG
y2qP0NTMco3tf0m9pOXNes0TWMlcNqbfNRWwUhwbotD989FW8DROIastgPinew1w
jgu3i0RRLo43o7kXo1VgfACZtnyNu5eGeSW9p0qEb6+1hFnEJD1E1fd32Hl7LEkj
FPA3hvRRLzMCgYEAu3suhuR4B7jg+Gi13p0mSBOJCLEMsANeo9YMCnYWeEM5fEs2
jkusnSp8tGIDSE7cXaOmexrvSefg9qyukiLfdyEyOT10Fk/vSxn5HiYLeoyZkrfA
fsSR9mmnPu0GpN9JLCy4ZQW6KsIxKelrJz8zmVeQIc1sF+OL86VI9ElkTh0CgYBa
mIT/A+ZBexC4e53/MQyTr841ft4pQ6hxZAPpMOBExEfcb4UgLm+wfgU3WwNiwxos
DHg0Pv7D4IFOpBn/mteGkq2OtYQFg1VVQ1XeJ1oxMnj4wsaoTVwkDBkynG3cUIND
wcXO4eHinh+0nA+KYN8MYY1eswhgnMNdlFiLbMzsAQKBgHY4qvew/WbVJIT292vH
2Llytr4O3R/f/6FEbTu6FdHoVdrjhmBunps9/PQPZOGUkprpp9wJ0Z5a1WXOHN8H
JDfXQyKpPZJzNia5r3AIZ7YrYfj18dBDkEAIOqQmCvQeU56BkbGQFHOMUWIHG1Jt
1CUZZKuSZ48ycukR6Y5Z3ELt
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "6CF729BC6F8B1596BB1C4CD9DB6338C9D02E1694",
			PublicKeyID:             "PUB_KEY_ID_0117346559632025120500192382001802",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Gj4CGj9c5zzjWw0PQOU
oX4O+cP8vQVokLuJ4DkRUV6kdGkmcDWNSKJydWDOKGH4IrEbqsiLAm00ghU6TWVO
lgNpNZmXMMGWQPH60SycHUW89J9A+8JOkwo4/J/ubf/iPCfaTvEKCWaNxcO9AI0g
sMcAWgRtuvSYLqAUZZ9hycqZYGfQQyvN4FgAN3NbC6x7t9SAfY21uaClMuZTqeTs
KHLzAJWnTJ/5dBlF9oaz78wwihDk7zEvPHH6y218kGKwWuDNARCZk13vdQQuP6V8
fjIYkPGo2hCHp61wNaKwwEJHPRm7Lw+zxKS4ywJeUJCbZd5N2CxNGQu5M/AEIvL0
iwIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "0a3cd5200286392cdb400fa6e0aa09c8",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := FundApiService{Client: api.Client}
	ctx := context.Background()

	// 测试查询运营账户余额
	resp, apiR, err := c.QueryBalance(ctx, QueryBalanceRequest{
		AccountType: AccountType_OPERATION,
	})
	assert.NoError(t, err)
	t.Log(fmt.Sprintf("Response: %+v", resp))
	if apiR != nil {
		t.Log(fmt.Sprintf("StatusCode: %d", apiR.Response.StatusCode))
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
}

func TestQueryBalanceFees(t *testing.T) {
	cccc := config.Config{
		AppID: "wxc4cbb00d34676a39",
		MchID: "1734655963",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNFZ7+CFXv+cY6
aZA3qrObYdkq2Hr5nrziw728o6spMBlux5LYCiqsiCcVk2EHdGVG+NkL4+7qFZZn
bfFLOlJ71AWEY6AxCUDzQU0FlTLfbVvoBTfuD1cHLJ4eWOVxd69ccSDmyqowML3l
rIy5VkSxFPN9MMbbjBinCKEcq7qqp89rGJ/DHXTfDWlze3kIQ1I5pkIpOHBYCJpw
TCIMjZaD7RyRs3g6Lor0kgGKF/wvJpD1y97DOH31ae3q1mUuCAvCgHsaVJ8Xy3Fe
FHjj5k60e490ZSVIxquaLKgdLCOHzanNM0b2hBge4JFs5WiefnldbFoSCLnmZoFM
Z1k8NLgJAgMBAAECggEBALKBFjFAaiSFQD/JcEizoL8nLNH5ORiaTQgHnddakfk4
r3pX5X17Y+dmXraN4A+QBTMAGFMcIvDGt3HxJOv4qKC1S9QOrvjfeBQDC5pHYc9N
LYjHOgZQfcx1zm7Vq2fW5+DLavrW/ckSM8d8J8SNzR5hAQo/cWbZCwAgJ7N/2er8
tE6o29QIGDRfXknVPQILQ4szN09u+wAKDN2fqWMJK93LYnRFImkG95GvOyNNeFST
goeYzTSa9EDouXbpc/6yi+dWtirSSdz9JcNJCSPwNkGeN8gY9g2oqmIIyleuAJSh
InOc7FVS6GN0T96wF/MURqLikuxd8If7wYHT3QuM4k0CgYEA5uWjFdOQBowI+TSm
fNgfgbxdsg5wXYWC74CcQY1VyYfzc4KCr9WkCp4q7QKzuNl4EcfZblFQPOtlFiC4
HKP/PWrahEGHbABzR8/COvnXJlyFbZYBZcqQx6UVpEAbLinr1KeNUdxM2+49AEDC
DNiwxPLavg1JbZ3IXUvZcckt69MCgYEA42GQi2GezD+NDezE9/2bBt3Sj9CBpfqG
y2qP0NTMco3tf0m9pOXNes0TWMlcNqbfNRWwUhwbotD989FW8DROIastgPinew1w
jgu3i0RRLo43o7kXo1VgfACZtnyNu5eGeSW9p0qEb6+1hFnEJD1E1fd32Hl7LEkj
FPA3hvRRLzMCgYEAu3suhuR4B7jg+Gi13p0mSBOJCLEMsANeo9YMCnYWeEM5fEs2
jkusnSp8tGIDSE7cXaOmexrvSefg9qyukiLfdyEyOT10Fk/vSxn5HiYLeoyZkrfA
fsSR9mmnPu0GpN9JLCy4ZQW6KsIxKelrJz8zmVeQIc1sF+OL86VI9ElkTh0CgYBa
mIT/A+ZBexC4e53/MQyTr841ft4pQ6hxZAPpMOBExEfcb4UgLm+wfgU3WwNiwxos
DHg0Pv7D4IFOpBn/mteGkq2OtYQFg1VVQ1XeJ1oxMnj4wsaoTVwkDBkynG3cUIND
wcXO4eHinh+0nA+KYN8MYY1eswhgnMNdlFiLbMzsAQKBgHY4qvew/WbVJIT292vH
2Llytr4O3R/f/6FEbTu6FdHoVdrjhmBunps9/PQPZOGUkprpp9wJ0Z5a1WXOHN8H
JDfXQyKpPZJzNia5r3AIZ7YrYfj18dBDkEAIOqQmCvQeU56BkbGQFHOMUWIHG1Jt
1CUZZKuSZ48ycukR6Y5Z3ELt
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "6CF729BC6F8B1596BB1C4CD9DB6338C9D02E1694",
			PublicKeyID:             "PUB_KEY_ID_0117346559632025120500192382001802",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Gj4CGj9c5zzjWw0PQOU
oX4O+cP8vQVokLuJ4DkRUV6kdGkmcDWNSKJydWDOKGH4IrEbqsiLAm00ghU6TWVO
lgNpNZmXMMGWQPH60SycHUW89J9A+8JOkwo4/J/ubf/iPCfaTvEKCWaNxcO9AI0g
sMcAWgRtuvSYLqAUZZ9hycqZYGfQQyvN4FgAN3NbC6x7t9SAfY21uaClMuZTqeTs
KHLzAJWnTJ/5dBlF9oaz78wwihDk7zEvPHH6y218kGKwWuDNARCZk13vdQQuP6V8
fjIYkPGo2hCHp61wNaKwwEJHPRm7Lw+zxKS4ywJeUJCbZd5N2CxNGQu5M/AEIvL0
iwIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "0a3cd5200286392cdb400fa6e0aa09c8",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := FundApiService{Client: api.Client}
	ctx := context.Background()

	// 测试查询手续费账户余额
	resp, apiR, err := c.QueryBalance(ctx, QueryBalanceRequest{
		AccountType: AccountType_FEES,
	})
	assert.NoError(t, err)
	t.Log(fmt.Sprintf("Response: %+v", resp))
	if apiR != nil {
		t.Log(fmt.Sprintf("StatusCode: %d", apiR.Response.StatusCode))
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
}

// TestGetFundFlowBill 测试申请资金账单（普通商户可用）
func TestGetFundFlowBill(t *testing.T) {
	cccc := config.Config{
		AppID: "wxc4cbb00d34676a39",
		MchID: "1734655963",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNFZ7+CFXv+cY6
aZA3qrObYdkq2Hr5nrziw728o6spMBlux5LYCiqsiCcVk2EHdGVG+NkL4+7qFZZn
bfFLOlJ71AWEY6AxCUDzQU0FlTLfbVvoBTfuD1cHLJ4eWOVxd69ccSDmyqowML3l
rIy5VkSxFPN9MMbbjBinCKEcq7qqp89rGJ/DHXTfDWlze3kIQ1I5pkIpOHBYCJpw
TCIMjZaD7RyRs3g6Lor0kgGKF/wvJpD1y97DOH31ae3q1mUuCAvCgHsaVJ8Xy3Fe
FHjj5k60e490ZSVIxquaLKgdLCOHzanNM0b2hBge4JFs5WiefnldbFoSCLnmZoFM
Z1k8NLgJAgMBAAECggEBALKBFjFAaiSFQD/JcEizoL8nLNH5ORiaTQgHnddakfk4
r3pX5X17Y+dmXraN4A+QBTMAGFMcIvDGt3HxJOv4qKC1S9QOrvjfeBQDC5pHYc9N
LYjHOgZQfcx1zm7Vq2fW5+DLavrW/ckSM8d8J8SNzR5hAQo/cWbZCwAgJ7N/2er8
tE6o29QIGDRfXknVPQILQ4szN09u+wAKDN2fqWMJK93LYnRFImkG95GvOyNNeFST
goeYzTSa9EDouXbpc/6yi+dWtirSSdz9JcNJCSPwNkGeN8gY9g2oqmIIyleuAJSh
InOc7FVS6GN0T96wF/MURqLikuxd8If7wYHT3QuM4k0CgYEA5uWjFdOQBowI+TSm
fNgfgbxdsg5wXYWC74CcQY1VyYfzc4KCr9WkCp4q7QKzuNl4EcfZblFQPOtlFiC4
HKP/PWrahEGHbABzR8/COvnXJlyFbZYBZcqQx6UVpEAbLinr1KeNUdxM2+49AEDC
DNiwxPLavg1JbZ3IXUvZcckt69MCgYEA42GQi2GezD+NDezE9/2bBt3Sj9CBpfqG
y2qP0NTMco3tf0m9pOXNes0TWMlcNqbfNRWwUhwbotD989FW8DROIastgPinew1w
jgu3i0RRLo43o7kXo1VgfACZtnyNu5eGeSW9p0qEb6+1hFnEJD1E1fd32Hl7LEkj
FPA3hvRRLzMCgYEAu3suhuR4B7jg+Gi13p0mSBOJCLEMsANeo9YMCnYWeEM5fEs2
jkusnSp8tGIDSE7cXaOmexrvSefg9qyukiLfdyEyOT10Fk/vSxn5HiYLeoyZkrfA
fsSR9mmnPu0GpN9JLCy4ZQW6KsIxKelrJz8zmVeQIc1sF+OL86VI9ElkTh0CgYBa
mIT/A+ZBexC4e53/MQyTr841ft4pQ6hxZAPpMOBExEfcb4UgLm+wfgU3WwNiwxos
DHg0Pv7D4IFOpBn/mteGkq2OtYQFg1VVQ1XeJ1oxMnj4wsaoTVwkDBkynG3cUIND
wcXO4eHinh+0nA+KYN8MYY1eswhgnMNdlFiLbMzsAQKBgHY4qvew/WbVJIT292vH
2Llytr4O3R/f/6FEbTu6FdHoVdrjhmBunps9/PQPZOGUkprpp9wJ0Z5a1WXOHN8H
JDfXQyKpPZJzNia5r3AIZ7YrYfj18dBDkEAIOqQmCvQeU56BkbGQFHOMUWIHG1Jt
1CUZZKuSZ48ycukR6Y5Z3ELt
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "6CF729BC6F8B1596BB1C4CD9DB6338C9D02E1694",
			PublicKeyID:             "PUB_KEY_ID_0117346559632025120500192382001802",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Gj4CGj9c5zzjWw0PQOU
oX4O+cP8vQVokLuJ4DkRUV6kdGkmcDWNSKJydWDOKGH4IrEbqsiLAm00ghU6TWVO
lgNpNZmXMMGWQPH60SycHUW89J9A+8JOkwo4/J/ubf/iPCfaTvEKCWaNxcO9AI0g
sMcAWgRtuvSYLqAUZZ9hycqZYGfQQyvN4FgAN3NbC6x7t9SAfY21uaClMuZTqeTs
KHLzAJWnTJ/5dBlF9oaz78wwihDk7zEvPHH6y218kGKwWuDNARCZk13vdQQuP6V8
fjIYkPGo2hCHp61wNaKwwEJHPRm7Lw+zxKS4ywJeUJCbZd5N2CxNGQu5M/AEIvL0
iwIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "0a3cd5200286392cdb400fa6e0aa09c8",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := FundApiService{Client: api.Client}
	ctx := context.Background()

	resp, apiR, err := c.GetFundFlowBill(ctx, FundFlowBillRequest{
		BillDate:    "2025-12-14",
		AccountType: AccountType_BASIC,
	})
	if err != nil {
		t.Logf("Error: %v", err)
	}
	t.Log(fmt.Sprintf("Response: %+v", resp))
	if apiR != nil {
		t.Log(fmt.Sprintf("StatusCode: %d", apiR.Response.StatusCode))
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
}

// TestGetFundFlowBillWithGzip 测试申请资金账单（带GZIP压缩）
func TestGetFundFlowBillWithGzip(t *testing.T) {
	cccc := config.Config{
		AppID: "wxc4cbb00d34676a39",
		MchID: "1734655963",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNFZ7+CFXv+cY6
aZA3qrObYdkq2Hr5nrziw728o6spMBlux5LYCiqsiCcVk2EHdGVG+NkL4+7qFZZn
bfFLOlJ71AWEY6AxCUDzQU0FlTLfbVvoBTfuD1cHLJ4eWOVxd69ccSDmyqowML3l
rIy5VkSxFPN9MMbbjBinCKEcq7qqp89rGJ/DHXTfDWlze3kIQ1I5pkIpOHBYCJpw
TCIMjZaD7RyRs3g6Lor0kgGKF/wvJpD1y97DOH31ae3q1mUuCAvCgHsaVJ8Xy3Fe
FHjj5k60e490ZSVIxquaLKgdLCOHzanNM0b2hBge4JFs5WiefnldbFoSCLnmZoFM
Z1k8NLgJAgMBAAECggEBALKBFjFAaiSFQD/JcEizoL8nLNH5ORiaTQgHnddakfk4
r3pX5X17Y+dmXraN4A+QBTMAGFMcIvDGt3HxJOv4qKC1S9QOrvjfeBQDC5pHYc9N
LYjHOgZQfcx1zm7Vq2fW5+DLavrW/ckSM8d8J8SNzR5hAQo/cWbZCwAgJ7N/2er8
tE6o29QIGDRfXknVPQILQ4szN09u+wAKDN2fqWMJK93LYnRFImkG95GvOyNNeFST
goeYzTSa9EDouXbpc/6yi+dWtirSSdz9JcNJCSPwNkGeN8gY9g2oqmIIyleuAJSh
InOc7FVS6GN0T96wF/MURqLikuxd8If7wYHT3QuM4k0CgYEA5uWjFdOQBowI+TSm
fNgfgbxdsg5wXYWC74CcQY1VyYfzc4KCr9WkCp4q7QKzuNl4EcfZblFQPOtlFiC4
HKP/PWrahEGHbABzR8/COvnXJlyFbZYBZcqQx6UVpEAbLinr1KeNUdxM2+49AEDC
DNiwxPLavg1JbZ3IXUvZcckt69MCgYEA42GQi2GezD+NDezE9/2bBt3Sj9CBpfqG
y2qP0NTMco3tf0m9pOXNes0TWMlcNqbfNRWwUhwbotD989FW8DROIastgPinew1w
jgu3i0RRLo43o7kXo1VgfACZtnyNu5eGeSW9p0qEb6+1hFnEJD1E1fd32Hl7LEkj
FPA3hvRRLzMCgYEAu3suhuR4B7jg+Gi13p0mSBOJCLEMsANeo9YMCnYWeEM5fEs2
jkusnSp8tGIDSE7cXaOmexrvSefg9qyukiLfdyEyOT10Fk/vSxn5HiYLeoyZkrfA
fsSR9mmnPu0GpN9JLCy4ZQW6KsIxKelrJz8zmVeQIc1sF+OL86VI9ElkTh0CgYBa
mIT/A+ZBexC4e53/MQyTr841ft4pQ6hxZAPpMOBExEfcb4UgLm+wfgU3WwNiwxos
DHg0Pv7D4IFOpBn/mteGkq2OtYQFg1VVQ1XeJ1oxMnj4wsaoTVwkDBkynG3cUIND
wcXO4eHinh+0nA+KYN8MYY1eswhgnMNdlFiLbMzsAQKBgHY4qvew/WbVJIT292vH
2Llytr4O3R/f/6FEbTu6FdHoVdrjhmBunps9/PQPZOGUkprpp9wJ0Z5a1WXOHN8H
JDfXQyKpPZJzNia5r3AIZ7YrYfj18dBDkEAIOqQmCvQeU56BkbGQFHOMUWIHG1Jt
1CUZZKuSZ48ycukR6Y5Z3ELt
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "6CF729BC6F8B1596BB1C4CD9DB6338C9D02E1694",
			PublicKeyID:             "PUB_KEY_ID_0117346559632025120500192382001802",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Gj4CGj9c5zzjWw0PQOU
oX4O+cP8vQVokLuJ4DkRUV6kdGkmcDWNSKJydWDOKGH4IrEbqsiLAm00ghU6TWVO
lgNpNZmXMMGWQPH60SycHUW89J9A+8JOkwo4/J/ubf/iPCfaTvEKCWaNxcO9AI0g
sMcAWgRtuvSYLqAUZZ9hycqZYGfQQyvN4FgAN3NbC6x7t9SAfY21uaClMuZTqeTs
KHLzAJWnTJ/5dBlF9oaz78wwihDk7zEvPHH6y218kGKwWuDNARCZk13vdQQuP6V8
fjIYkPGo2hCHp61wNaKwwEJHPRm7Lw+zxKS4ywJeUJCbZd5N2CxNGQu5M/AEIvL0
iwIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "0a3cd5200286392cdb400fa6e0aa09c8",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := FundApiService{Client: api.Client}
	ctx := context.Background()

	resp, apiR, err := c.GetFundFlowBill(ctx, FundFlowBillRequest{
		BillDate:    "2025-12-14",
		AccountType: AccountType_BASIC,
		TarType:     string(TarType_GZIP),
	})
	if err != nil {
		t.Logf("Error: %v", err)
	}
	t.Log(fmt.Sprintf("Response: %+v", resp))
	if apiR != nil {
		t.Log(fmt.Sprintf("StatusCode: %d", apiR.Response.StatusCode))
		bodyBytes, _ := io.ReadAll(apiR.Response.Body)
		t.Log(string(bodyBytes))
	}
}

// TestDownloadBill 测试下载账单文件
func TestDownloadBill(t *testing.T) {
	cccc := config.Config{
		AppID: "wxc4cbb00d34676a39",
		MchID: "1734655963",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNFZ7+CFXv+cY6
aZA3qrObYdkq2Hr5nrziw728o6spMBlux5LYCiqsiCcVk2EHdGVG+NkL4+7qFZZn
bfFLOlJ71AWEY6AxCUDzQU0FlTLfbVvoBTfuD1cHLJ4eWOVxd69ccSDmyqowML3l
rIy5VkSxFPN9MMbbjBinCKEcq7qqp89rGJ/DHXTfDWlze3kIQ1I5pkIpOHBYCJpw
TCIMjZaD7RyRs3g6Lor0kgGKF/wvJpD1y97DOH31ae3q1mUuCAvCgHsaVJ8Xy3Fe
FHjj5k60e490ZSVIxquaLKgdLCOHzanNM0b2hBge4JFs5WiefnldbFoSCLnmZoFM
Z1k8NLgJAgMBAAECggEBALKBFjFAaiSFQD/JcEizoL8nLNH5ORiaTQgHnddakfk4
r3pX5X17Y+dmXraN4A+QBTMAGFMcIvDGt3HxJOv4qKC1S9QOrvjfeBQDC5pHYc9N
LYjHOgZQfcx1zm7Vq2fW5+DLavrW/ckSM8d8J8SNzR5hAQo/cWbZCwAgJ7N/2er8
tE6o29QIGDRfXknVPQILQ4szN09u+wAKDN2fqWMJK93LYnRFImkG95GvOyNNeFST
goeYzTSa9EDouXbpc/6yi+dWtirSSdz9JcNJCSPwNkGeN8gY9g2oqmIIyleuAJSh
InOc7FVS6GN0T96wF/MURqLikuxd8If7wYHT3QuM4k0CgYEA5uWjFdOQBowI+TSm
fNgfgbxdsg5wXYWC74CcQY1VyYfzc4KCr9WkCp4q7QKzuNl4EcfZblFQPOtlFiC4
HKP/PWrahEGHbABzR8/COvnXJlyFbZYBZcqQx6UVpEAbLinr1KeNUdxM2+49AEDC
DNiwxPLavg1JbZ3IXUvZcckt69MCgYEA42GQi2GezD+NDezE9/2bBt3Sj9CBpfqG
y2qP0NTMco3tf0m9pOXNes0TWMlcNqbfNRWwUhwbotD989FW8DROIastgPinew1w
jgu3i0RRLo43o7kXo1VgfACZtnyNu5eGeSW9p0qEb6+1hFnEJD1E1fd32Hl7LEkj
FPA3hvRRLzMCgYEAu3suhuR4B7jg+Gi13p0mSBOJCLEMsANeo9YMCnYWeEM5fEs2
jkusnSp8tGIDSE7cXaOmexrvSefg9qyukiLfdyEyOT10Fk/vSxn5HiYLeoyZkrfA
fsSR9mmnPu0GpN9JLCy4ZQW6KsIxKelrJz8zmVeQIc1sF+OL86VI9ElkTh0CgYBa
mIT/A+ZBexC4e53/MQyTr841ft4pQ6hxZAPpMOBExEfcb4UgLm+wfgU3WwNiwxos
DHg0Pv7D4IFOpBn/mteGkq2OtYQFg1VVQ1XeJ1oxMnj4wsaoTVwkDBkynG3cUIND
wcXO4eHinh+0nA+KYN8MYY1eswhgnMNdlFiLbMzsAQKBgHY4qvew/WbVJIT292vH
2Llytr4O3R/f/6FEbTu6FdHoVdrjhmBunps9/PQPZOGUkprpp9wJ0Z5a1WXOHN8H
JDfXQyKpPZJzNia5r3AIZ7YrYfj18dBDkEAIOqQmCvQeU56BkbGQFHOMUWIHG1Jt
1CUZZKuSZ48ycukR6Y5Z3ELt
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "6CF729BC6F8B1596BB1C4CD9DB6338C9D02E1694",
			PublicKeyID:             "PUB_KEY_ID_0117346559632025120500192382001802",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Gj4CGj9c5zzjWw0PQOU
oX4O+cP8vQVokLuJ4DkRUV6kdGkmcDWNSKJydWDOKGH4IrEbqsiLAm00ghU6TWVO
lgNpNZmXMMGWQPH60SycHUW89J9A+8JOkwo4/J/ubf/iPCfaTvEKCWaNxcO9AI0g
sMcAWgRtuvSYLqAUZZ9hycqZYGfQQyvN4FgAN3NbC6x7t9SAfY21uaClMuZTqeTs
KHLzAJWnTJ/5dBlF9oaz78wwihDk7zEvPHH6y218kGKwWuDNARCZk13vdQQuP6V8
fjIYkPGo2hCHp61wNaKwwEJHPRm7Lw+zxKS4ywJeUJCbZd5N2CxNGQu5M/AEIvL0
iwIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "0a3cd5200286392cdb400fa6e0aa09c8",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := FundApiService{Client: api.Client}
	ctx := context.Background()

	billDate := "2025-12-14"
	billResp, _, err := c.GetFundFlowBill(ctx, FundFlowBillRequest{
		BillDate:    billDate,
		AccountType: AccountType_BASIC,
	})
	if err != nil {
		t.Logf("GetFundFlowBill Error: %v", err)
		t.Skip("跳过测试：无法获取账单下载地址")
		return
	}
	t.Logf("BillResponse: %+v", billResp)
	t.Logf("DownloadUrl: %s", billResp.DownloadUrl)

	downloadResp, err := c.DownloadBill(ctx, billResp.DownloadUrl, false)
	if err != nil {
		t.Fatalf("DownloadBill Error: %v", err)
	}

	t.Logf("Downloaded %d bytes", len(downloadResp.Data))
	t.Logf("SHA1: %s", downloadResp.HashValue)
	t.Logf("Expected SHA1: %s", billResp.HashValue)

	if downloadResp.HashValue != billResp.HashValue {
		t.Errorf("SHA1 mismatch: expected %s, got %s", billResp.HashValue, downloadResp.HashValue)
	}

	// 保存账单文件到本地
	saveDir := "./bills"
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		t.Fatalf("Create directory failed: %v", err)
	}

	fileName := fmt.Sprintf("fundflow_%s_%s.csv", AccountType_BASIC, billDate)
	filePath := filepath.Join(saveDir, fileName)

	if err := os.WriteFile(filePath, downloadResp.Data, 0644); err != nil {
		t.Fatalf("Write file failed: %v", err)
	}
	t.Logf("账单已保存到: %s", filePath)

	content := string(downloadResp.Data)
	if len(content) > 1000 {
		content = content[:1000] + "..."
	}
	t.Logf("Bill Content:\n%s", content)
}

// TestGetAccountBalance 测试获取账户余额（普通商户版本）
func TestGetAccountBalance(t *testing.T) {
	cccc := config.Config{
		AppID: "wxc4cbb00d34676a39",
		MchID: "1734655963",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNFZ7+CFXv+cY6
aZA3qrObYdkq2Hr5nrziw728o6spMBlux5LYCiqsiCcVk2EHdGVG+NkL4+7qFZZn
bfFLOlJ71AWEY6AxCUDzQU0FlTLfbVvoBTfuD1cHLJ4eWOVxd69ccSDmyqowML3l
rIy5VkSxFPN9MMbbjBinCKEcq7qqp89rGJ/DHXTfDWlze3kIQ1I5pkIpOHBYCJpw
TCIMjZaD7RyRs3g6Lor0kgGKF/wvJpD1y97DOH31ae3q1mUuCAvCgHsaVJ8Xy3Fe
FHjj5k60e490ZSVIxquaLKgdLCOHzanNM0b2hBge4JFs5WiefnldbFoSCLnmZoFM
Z1k8NLgJAgMBAAECggEBALKBFjFAaiSFQD/JcEizoL8nLNH5ORiaTQgHnddakfk4
r3pX5X17Y+dmXraN4A+QBTMAGFMcIvDGt3HxJOv4qKC1S9QOrvjfeBQDC5pHYc9N
LYjHOgZQfcx1zm7Vq2fW5+DLavrW/ckSM8d8J8SNzR5hAQo/cWbZCwAgJ7N/2er8
tE6o29QIGDRfXknVPQILQ4szN09u+wAKDN2fqWMJK93LYnRFImkG95GvOyNNeFST
goeYzTSa9EDouXbpc/6yi+dWtirSSdz9JcNJCSPwNkGeN8gY9g2oqmIIyleuAJSh
InOc7FVS6GN0T96wF/MURqLikuxd8If7wYHT3QuM4k0CgYEA5uWjFdOQBowI+TSm
fNgfgbxdsg5wXYWC74CcQY1VyYfzc4KCr9WkCp4q7QKzuNl4EcfZblFQPOtlFiC4
HKP/PWrahEGHbABzR8/COvnXJlyFbZYBZcqQx6UVpEAbLinr1KeNUdxM2+49AEDC
DNiwxPLavg1JbZ3IXUvZcckt69MCgYEA42GQi2GezD+NDezE9/2bBt3Sj9CBpfqG
y2qP0NTMco3tf0m9pOXNes0TWMlcNqbfNRWwUhwbotD989FW8DROIastgPinew1w
jgu3i0RRLo43o7kXo1VgfACZtnyNu5eGeSW9p0qEb6+1hFnEJD1E1fd32Hl7LEkj
FPA3hvRRLzMCgYEAu3suhuR4B7jg+Gi13p0mSBOJCLEMsANeo9YMCnYWeEM5fEs2
jkusnSp8tGIDSE7cXaOmexrvSefg9qyukiLfdyEyOT10Fk/vSxn5HiYLeoyZkrfA
fsSR9mmnPu0GpN9JLCy4ZQW6KsIxKelrJz8zmVeQIc1sF+OL86VI9ElkTh0CgYBa
mIT/A+ZBexC4e53/MQyTr841ft4pQ6hxZAPpMOBExEfcb4UgLm+wfgU3WwNiwxos
DHg0Pv7D4IFOpBn/mteGkq2OtYQFg1VVQ1XeJ1oxMnj4wsaoTVwkDBkynG3cUIND
wcXO4eHinh+0nA+KYN8MYY1eswhgnMNdlFiLbMzsAQKBgHY4qvew/WbVJIT292vH
2Llytr4O3R/f/6FEbTu6FdHoVdrjhmBunps9/PQPZOGUkprpp9wJ0Z5a1WXOHN8H
JDfXQyKpPZJzNia5r3AIZ7YrYfj18dBDkEAIOqQmCvQeU56BkbGQFHOMUWIHG1Jt
1CUZZKuSZ48ycukR6Y5Z3ELt
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "6CF729BC6F8B1596BB1C4CD9DB6338C9D02E1694",
			PublicKeyID:             "PUB_KEY_ID_0117346559632025120500192382001802",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Gj4CGj9c5zzjWw0PQOU
oX4O+cP8vQVokLuJ4DkRUV6kdGkmcDWNSKJydWDOKGH4IrEbqsiLAm00ghU6TWVO
lgNpNZmXMMGWQPH60SycHUW89J9A+8JOkwo4/J/ubf/iPCfaTvEKCWaNxcO9AI0g
sMcAWgRtuvSYLqAUZZ9hycqZYGfQQyvN4FgAN3NbC6x7t9SAfY21uaClMuZTqeTs
KHLzAJWnTJ/5dBlF9oaz78wwihDk7zEvPHH6y218kGKwWuDNARCZk13vdQQuP6V8
fjIYkPGo2hCHp61wNaKwwEJHPRm7Lw+zxKS4ywJeUJCbZd5N2CxNGQu5M/AEIvL0
iwIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "0a3cd5200286392cdb400fa6e0aa09c8",
		Proxy:  config.Proxy{},
	}

	api, err := wxpay.InitClient(cccc)
	assert.NoError(t, err)
	c := FundApiService{Client: api.Client}
	ctx := context.Background()

	billDate := "2025-12-14"
	resp, err := c.GetAccountBalance(ctx, billDate, AccountType_BASIC)
	if err != nil {
		t.Logf("GetAccountBalance Error: %v", err)
		t.Skip("跳过测试：无法获取账户余额")
		return
	}

	t.Logf("账单日期: %s", resp.BillDate)
	t.Logf("账户类型: %s", resp.AccountType)
	t.Logf("期末余额: %.2f 元", resp.EndBalance)
	t.Logf("当日总收入: %.2f 元", resp.TotalIncome)
	t.Logf("当日总支出: %.2f 元", resp.TotalExpense)
	t.Logf("交易笔数: %d", resp.RecordCount)

	for i, record := range resp.Records {
		if i >= 5 {
			t.Logf("... 还有 %d 条记录", len(resp.Records)-5)
			break
		}
		t.Logf("记录%d: 时间=%s, 类型=%s, 金额=%s, 余额=%s, 备注=%s",
			i+1, record.AccountingTime, record.IncomeExpenseType,
			record.Amount, record.Balance, record.Remark)
	}
}
