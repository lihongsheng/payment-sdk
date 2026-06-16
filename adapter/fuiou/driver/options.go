package driver

import "github.com/lihongsheng/payment-sdk/config/params"

// schema 富友支付配置的字段树定义。
var schema = &params.Schema{
	Type:     params.SchemaObject,
	Title:    "富友支付配置",
	Required: []string{"merchant", "cert", "api"},
	Order:    []string{"merchant", "cert", "api", "wechat", "alipay"},
	Properties: map[string]*params.Schema{
		"merchant": {
			Type:        params.SchemaObject,
			Title:       "商户信息",
			Description: "富友商户基础信息",
			Required:    []string{"mch_id", "api_secret", "order_prefix"},
			Order:       []string{"mch_id", "api_secret", "order_prefix"},
			Properties: map[string]*params.Schema{
				"mch_id": {
					Type:     params.SchemaString,
					Title:    "商户号",
					Pattern:  "^[0-9a-zA-Z]{8,64}$",
					Examples: []any{"0002900F0xxxxxxx"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"api_secret": {
					Type:     params.SchemaString,
					Title:    "API秘钥",
					Pattern:  "^[0-9a-zA-Z]{32,64}$",
					Examples: []any{"32位英文数字混合的接口秘钥"},
					UI:       &params.UISchema{InputType: params.InputPassword, ValidateType: params.ValidateReg},
				},
				"order_prefix": {
					Type:     params.SchemaString,
					Title:    "订单前缀",
					Pattern:  "^[0-9a-fA-F]{4,20}$",
					Examples: []any{"FY"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
			},
		},
		"cert": {
			Type:        params.SchemaObject,
			Title:       "证书配置",
			Description: "富友主账户私钥与公钥（用于转账类签名）",
			Required:    []string{"rsa_private_key", "rsa_public_key"},
			Order:       []string{"rsa_private_key", "rsa_public_key"},
			Properties: map[string]*params.Schema{
				"rsa_private_key": {
					Type:     params.SchemaString,
					Title:    "私钥(转账)",
					Examples: []any{"-----BEGIN PRIVATE KEY-----\nMIIE...\n-----END PRIVATE KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPrivate, Rows: 6},
				},
				"rsa_public_key": {
					Type:     params.SchemaString,
					Title:    "公钥(转账)",
					Examples: []any{"-----BEGIN PUBLIC KEY-----\nMIIB...\n-----END PUBLIC KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPublic, Rows: 6},
				},
			},
		},
		"api": {
			Type:        params.SchemaObject,
			Title:       "接口配置",
			Description: "富友接口地址及协议版本",
			Required:    []string{"api_host", "version"},
			Order:       []string{"api_host", "version"},
			Properties: map[string]*params.Schema{
				"api_host": {
					Type:     params.SchemaString,
					Title:    "API默认地址",
					Default:  "https://aipay-cloud.fuioupay.com",
					Format:   "url",
					Examples: []any{"https://aipay-cloud.fuioupay.com"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateUrl},
				},
				"version": {
					Type:     params.SchemaString,
					Title:    "版本号",
					Default:  "1.0",
					Examples: []any{"1.0"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateString},
				},
			},
		},
		"wechat": {
			Type:        params.SchemaObject,
			Title:       "微信子渠道",
			Description: "用于获取微信用户 openid，按需填写",
			Order:       []string{"app_id", "app_secret"},
			Properties: map[string]*params.Schema{
				"app_id": {
					Type:     params.SchemaString,
					Title:    "微信应用id",
					Examples: []any{"wx8888888888888888"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateString},
				},
				"app_secret": {
					Type:     params.SchemaString,
					Title:    "微信应用秘钥",
					Examples: []any{"32位字母数字混合字符串"},
					UI:       &params.UISchema{InputType: params.InputPassword, ValidateType: params.ValidateString},
				},
			},
		},
		"alipay": {
			Type:        params.SchemaObject,
			Title:       "支付宝子渠道",
			Description: "富友支付宝相关配置，按需填写",
			Order:       []string{"app_id", "rsa_private_key", "rsa_root_crt"},
			Properties: map[string]*params.Schema{
				"app_id": {
					Type:     params.SchemaString,
					Title:    "支付宝应用id",
					Pattern:  "^[0-9a-zA-Z]{8,64}$",
					Examples: []any{"2021000000000000"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"rsa_private_key": {
					Type:     params.SchemaString,
					Title:    "支付宝应用私钥",
					Examples: []any{"-----BEGIN PRIVATE KEY-----\nMIIE...\n-----END PRIVATE KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPrivate, Rows: 6},
				},
				"rsa_root_crt": {
					Type:     params.SchemaString,
					Title:    "根证书-crt",
					Examples: []any{"-----BEGIN CERTIFICATE-----\nMIIE...\n-----END CERTIFICATE-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaCert, Rows: 6},
				},
			},
		},
	},
}