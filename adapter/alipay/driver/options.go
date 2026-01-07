package driver

import "github.com/lihongsheng/payment-sdk/config/params"

var options = &params.Option{
	Label:        "支付宝配置项",
	Name:         "",       // 根节点无name
	Type:         "Object", // 根节点类型为Object
	ValidateReg:  "",
	ValidateType: "",
	InputType:    "",
	Default:      "",
	Values:       nil,
	// 嵌套拉卡拉Config的所有字段配置
	Properties: []params.Option{
		{
			Label:        "应用id",
			Name:         "app_id",
			Type:         params.String,
			ValidateReg:  "^[0-9a-zA-Z]{8,64}$", // 拉卡拉商户号通常为8-15位数字
			ValidateType: params.ValidateReg,
			InputType:    params.InputText,
			Default:      "",
			Values:       nil,
			Properties:   nil,
			Require:      true,
		},
		{
			Label:        "授权码（app_auth_token）",
			Name:         "app_auth_token",
			Type:         params.String,
			ValidateReg:  "^.{12,}$", // 秘钥长度32-64位
			ValidateType: params.ValidateReg,
			InputType:    "password", // 敏感字段用密码框
			Default:      "",
			Values:       nil,
			Properties:   nil,
			Require:      false,
		},
		// 证书信息（嵌套Object类型）
		{
			Label:        "证书信息",
			Name:         "cert",
			Type:         "Object",
			ValidateReg:  "",
			ValidateType: "",
			InputType:    "",
			Default:      "",
			Values:       nil,
			Require:      true,
			Properties: []params.Option{
				{
					Label:        "私钥（RSA格式）",
					Name:         "private_key",
					Type:         params.String,
					ValidateReg:  "",
					ValidateType: params.ValidateRsaPrivate,
					InputType:    params.InputTextarea, // 长文本用多行输入框
					Default:      "",
					Values:       nil,
					Properties:   nil,
					Require:      true,
				},
				{
					Label:        "公钥（RSA格式）",
					Name:         "public_key",
					Type:         params.String,
					ValidateReg:  "",
					ValidateType: params.ValidateRsaPublic,
					InputType:    params.InputTextarea, // 长文本用多行输入框
					Default:      "",
					Values:       nil,
					Properties:   nil,
					Require:      true,
				},
				{
					Label:        "aliPublicCrt文件内容(转账需要)",
					Name:         "app_crt",
					Type:         params.String,
					ValidateReg:  "",
					ValidateType: params.ValidateRsaCert,
					InputType:    params.InputTextarea, // 长文本用多行输入框
					Default:      "",
					Values:       nil,
					Properties:   nil,
					Require:      false,
				},
				{
					Label:        "aliRootCrt文件内容(转账需要)",
					Name:         "root_crt",
					Type:         params.String,
					ValidateReg:  "",
					ValidateType: params.ValidateRsaCert,
					InputType:    params.InputTextarea, // 长文本用多行输入框
					Default:      "",
					Values:       nil,
					Properties:   nil,
					Require:      false,
				},
			},
		},
	},
}
