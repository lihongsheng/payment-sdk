package driver

import "github.com/lihongsheng/payment-sdk/config/params"

// schema 支付宝支付配置的字段树定义。
var schema = &params.Schema{
	Type:     params.SchemaObject,
	Title:    "支付宝支付配置",
	Required: []string{"merchant", "cert"},
	Order:    []string{"merchant", "cert"},
	Properties: map[string]*params.Schema{
		"merchant": {
			Type:        params.SchemaObject,
			Title:       "商户信息",
			Description: "支付宝商户基础信息",
			Required:    []string{"app_id"},
			Order:       []string{"app_id", "app_auth_token"},
			Properties: map[string]*params.Schema{
				"app_id": {
					Type:     params.SchemaString,
					Title:    "应用id",
					Pattern:  "^[0-9a-zA-Z]{8,64}$",
					Examples: []any{"2021000000000000"},
					UI:       &params.UISchema{InputType: params.InputText, ValidateType: params.ValidateReg},
				},
				"app_auth_token": {
					Type:     params.SchemaString,
					Title:    "授权码",
					Pattern:  "^.{12,}$",
					Examples: []any{"服务商授权场景才需要填写"},
					UI:       &params.UISchema{InputType: params.InputPassword, ValidateType: params.ValidateReg},
				},
			},
		},
		"cert": {
			Type:        params.SchemaObject,
			Title:       "证书配置",
			Description: "应用私钥及证书",
			Required:    []string{"rsa_private_key", "rsa_root_crt"},
			Order:       []string{"rsa_private_key", "rsa_app_crt", "rsa_root_crt"},
			Properties: map[string]*params.Schema{
				"rsa_private_key": {
					Type:     params.SchemaString,
					Title:    "应用私钥",
					Examples: []any{"-----BEGIN PRIVATE KEY-----\nMIIE...\n-----END PRIVATE KEY-----"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPrivate, Rows: 6},
				},
				"rsa_app_crt": {
					Type:     params.SchemaString,
					Title:    "支付宝证书-crt",
					Examples: []any{"应用 crt 证书全文"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaCert, Rows: 6},
				},
				"rsa_root_crt": {
					Type:     params.SchemaString,
					Title:    "根证书-crt",
					Examples: []any{"支付宝根证书全文"},
					UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaCert, Rows: 6},
				},
			},
		},
	},
}
