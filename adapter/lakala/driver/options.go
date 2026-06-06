package driver

import "github.com/lihongsheng/payment-sdk/config/params"

// schema 拉卡拉支付配置的字段树定义，由前端按 JSON Schema 渲染表单。
var schema = &params.Schema{
	Type:     params.SchemaObject,
	Title:    "拉卡拉支付配置",
	Required: []string{"merchant", "cert", "api"},
	Order:    []string{"merchant", "cert", "api"},
	Properties: map[string]*params.Schema{
		"merchant": {
			Type:        params.SchemaObject,
			Title:       "商户信息",
			Description: "拉卡拉商户基础信息",
			Required:    []string{"app_id", "mch_id", "term_no"},
			Order:       []string{"app_id", "mch_id", "term_no"},
			Properties: map[string]*params.Schema{
				"app_id": {
					Type:     params.SchemaString,
					Title:    "应用id",
					Examples: []any{"OUP000000xxxxx"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateString},
				},
				"mch_id": {
					Type:     params.SchemaString,
					Title:    "商户号",
					Pattern:  "^[0-9a-zA-Z]{8,15}$",
					Examples: []any{"8228xxxxxxx"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"term_no": {
					Type:     params.SchemaString,
					Title:    "终端号",
					Pattern:  "^[0-9a-fA-F]{4,20}$",
					Examples: []any{"A0099999"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
			},
		},
		"cert": {
			Type:        params.SchemaObject,
			Title:       "证书配置",
			Description: "私钥/公钥及证书序列号",
			Required:    []string{"rsa_private_number", "rsa_private_key", "rsa_public_key"},
			Order:       []string{"rsa_private_number", "rsa_private_key", "rsa_public_key"},
			Properties: map[string]*params.Schema{
				"rsa_private_number": {
					Type:     params.SchemaString,
					Title:    "应用证书序列号",
					Pattern:  "^[0-9a-fA-F]{10,64}$",
					Examples: []any{"0123456789abcdef0123456789abcdef"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"rsa_private_key": {
					Type:     params.SchemaString,
					Title:    "私钥",
					Examples: []any{"-----BEGIN PRIVATE KEY-----\nMIIE...\n-----END PRIVATE KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPrivate, Rows: 6},
				},
				"rsa_public_key": {
					Type:     params.SchemaString,
					Title:    "公钥",
					Examples: []any{"-----BEGIN PUBLIC KEY-----\nMIIB...\n-----END PUBLIC KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPublic, Rows: 6},
				},
			},
		},
		"api": {
			Type:        params.SchemaObject,
			Title:       "接口地址",
			Description: "拉卡拉接口相关配置",
			Required:    []string{"api_host"},
			Order:       []string{"api_host"},
			Properties: map[string]*params.Schema{
				"api_host": {
					Type:     params.SchemaString,
					Title:    "API默认地址",
					Default:  "https://s2.lakala.com",
					Format:   "url",
					Examples: []any{"https://s2.lakala.com"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateUrl},
				},
			},
		},
	},
}