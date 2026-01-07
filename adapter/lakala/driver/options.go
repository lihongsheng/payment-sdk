package driver

import "github.com/lihongsheng/payment-sdk/config/params"

var options = &params.Option{
	Label:        "拉卡拉配置项",
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
			ValidateReg:  "", // 无特定正则，仅非空校验
			ValidateType: "String",
			InputType:    "text",
			Default:      "",
			Values:       nil,
			Properties:   nil,
			Require:      true,
		},
		{
			Label:        "商户号",
			Name:         "mch_id",
			Type:         params.String,
			ValidateReg:  "^[0-9a-zA-Z]{8,15}$", // 拉卡拉商户号通常为8-15位数字
			ValidateType: "Reg",
			InputType:    "text",
			Default:      "",
			Values:       nil,
			Properties:   nil,
			Require:      true,
		},
		{
			Label:        "API秘钥",
			Name:         "api_secret",
			Type:         params.String,
			ValidateReg:  "^[0-9a-zA-Z]{32,64}$", // 秘钥长度32-64位
			ValidateType: "Reg",
			InputType:    "password", // 敏感字段用密码框
			Default:      "",
			Values:       nil,
			Properties:   nil,
			Require:      true,
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
					Label:        "私钥证书序列号",
					Name:         "private_number",
					Type:         params.String,
					ValidateReg:  "^[0-9a-fA-F]{10,64}$", // 序列号为32-64位十六进制
					ValidateType: params.ValidateReg,
					InputType:    params.InputText,
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
			},
		},
		{
			Label:        "终端号",
			Name:         "term_no",
			Type:         params.String,
			ValidateReg:  "^[0-9a-fA-F]{4,20}$", // 终端号通常4-10位数字
			ValidateType: params.ValidateReg,
			InputType:    params.InputText,
			Default:      "",
			Values:       nil,
			Properties:   nil,
			Require:      true,
		},
		{
			Label:        "API默认地址",
			Name:         "api_host",
			Type:         params.String,
			ValidateReg:  "", // 校验HTTP/HTTPS地址格式
			ValidateType: params.ValidateUrl,
			InputType:    params.InputText,
			Default:      "https://s2.lakala.com", // 默认值设为拉卡拉官方地址
			Values:       nil,
			Properties:   nil,
			Require:      true,
		},
	},
}
