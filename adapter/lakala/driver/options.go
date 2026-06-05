package driver

import "github.com/lihongsheng/payment-sdk/config/params"

var options = []params.Option{
	{
		Label:        "应用id",
		Name:         "app_id",
		Type:         params.String,
		ValidateReg:  "", // 无特定正则，仅非空校验
		ValidateType: "String",
		InputType:    "text",
		Default:      "",
		Values:       nil,
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
		Require:      true,
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

		Require: true,
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

		Require: true,
	},
	{
		Label:        "应用证书序列号",
		Name:         "rsa_private_number",
		Type:         params.String,
		ValidateReg:  "^[0-9a-fA-F]{10,64}$", // 序列号为32-64位十六进制
		ValidateType: params.ValidateReg,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,

		Require: true,
	},
	{
		Label:        "私钥",
		Name:         "rsa_private_key",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaPrivate,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,

		Require: true,
	},
	{
		Label:        "公钥",
		Name:         "rsa_public_key",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaPublic,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,
		Require:      true,
	},
}
