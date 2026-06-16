package driver

import "github.com/lihongsheng/payment-sdk/config/params"

// schema 微信支付配置的字段树定义。
var schema = &params.Schema{
	Type:     params.SchemaObject,
	Title:    "微信支付配置",
	Required: []string{"merchant", "cert"},
	Order:    []string{"merchant", "cert", "service"},
	Properties: map[string]*params.Schema{
		"merchant": {
			Type:        params.SchemaObject,
			Title:       "商户信息",
			Description: "微信支付商户基础信息",
			Required:    []string{"mch_id", "app_id", "app_secret"},
			Order:       []string{"mch_id", "app_id", "app_secret"},
			Properties: map[string]*params.Schema{
				"mch_id": {
					Type:     params.SchemaString,
					Title:    "微信商户ID",
					Pattern:  "^[0-9]{10,15}$",
					Examples: []any{"1900000109"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"app_id": {
					Type:     params.SchemaString,
					Title:    "微信AppID",
					Pattern:  "^wx[0-9a-zA-Z]{16,32}$",
					Examples: []any{"wx8888888888888888"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"app_secret": {
					Type:     params.SchemaString,
					Title:    "应用Secret",
					Pattern:  "^.{6,}$",
					Examples: []any{"32位字母数字混合字符串"},
					UI:       &params.UISchema{InputType: params.InputPassword, ValidateType: params.ValidateReg},
				},
			},
		},
		"cert": {
			Type:        params.SchemaObject,
			Title:       "证书配置",
			Description: "API 秘钥、私钥、公钥及对应证书序列号",
			Required:    []string{"api_secret", "rsa_private", "rsa_private_number"},
			Order:       []string{"api_secret", "rsa_private", "rsa_private_number", "rsa_public_number", "rsa_public"},
			Properties: map[string]*params.Schema{
				"api_secret": {
					Type:     params.SchemaString,
					Title:    "微信API V3秘钥",
					Pattern:  "^.{6,}$",
					Examples: []any{"32位字母数字混合字符串"},
					UI:       &params.UISchema{InputType: params.InputPassword, ValidateType: params.ValidateReg},
				},
				"rsa_private": {
					Type:     params.SchemaString,
					Title:    "应用私钥",
					Examples: []any{"-----BEGIN PRIVATE KEY-----\nMIIE...\n-----END PRIVATE KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPrivate, Rows: 6},
				},
				"rsa_private_number": {
					Type:     params.SchemaString,
					Title:    "应用证书序列号",
					Pattern:  "^.{6,}$",
					Examples: []any{"0123456789ABCDEF0123456789ABCDEF01234567"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"rsa_public_number": {
					Type:     params.SchemaString,
					Title:    "微信公钥证书序列号",
					Pattern:  "^.{6,}$",
					Examples: []any{"PUB_KEY_ID_XXXXX"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"rsa_public": {
					Type:     params.SchemaString,
					Title:    "微信公钥",
					Examples: []any{"-----BEGIN PUBLIC KEY-----\nMIIB...\n-----END PUBLIC KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPublic, Rows: 6},
				},
			},
		},
		"service": {
			Type:        params.SchemaObject,
			Title:       "增值服务",
			Description: "支付分等增值服务配置（按需填写）",
			Order:       []string{"score_service_id"},
			Properties: map[string]*params.Schema{
				"score_service_id": {
					Type:     params.SchemaString,
					Title:    "支付分服务ID",
					Examples: []any{"服务商在微信支付分申请的 service_id"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateString},
				},
			},
		},
	},
}