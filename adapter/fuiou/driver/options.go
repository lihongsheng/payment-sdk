package driver

import "github.com/lihongsheng/payment-sdk/config/params"

var options = []params.Option{
	{
		Label:        "商户号",
		Name:         "mch_id",
		Type:         params.String,
		ValidateReg:  "^[0-9a-zA-Z]{8,64}$", //
		ValidateType: "Reg",
		InputType:    "text",
		Default:      "",
		Values:       nil,

		Require: true,
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

		Require: true,
	},
	{
		Label:        "订单前缀",
		Name:         "order_prefix",
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
		Default:      "https://aipay-cloud.fuioupay.com", // 默认值设为拉卡拉官方地址
		Values:       nil,

		Require: true,
	}, {
		Label:        "版本号",
		Name:         "version",
		Type:         params.String,
		ValidateReg:  "", // 校验HTTP/HTTPS地址格式
		ValidateType: params.ValidateString,
		InputType:    params.InputText,
		Default:      "1.0", // 默认值设为拉卡拉官方地址
		Values:       nil,

		Require: true,
	},
	{
		Label:        "私钥(转账)",
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
		Label:        "公钥(转账)",
		Name:         "rsa_public_key",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaPublic,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "微信应用id",
		Name:         "wechat_app_id",
		Type:         params.String,
		ValidateReg:  "", // 校验HTTP/HTTPS地址格式
		ValidateType: params.ValidateString,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,
		Require:      false,
	},
	{
		Label:        "微信应用秘钥",
		Name:         "wechat_app_secret",
		Type:         params.String,
		ValidateReg:  "", // 校验HTTP/HTTPS地址格式
		ValidateType: params.ValidateString,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,
		Require:      false,
	},
	{
		Label:        "支付宝应用id",
		Name:         "alipay_app_id",
		Type:         params.String,
		ValidateReg:  "^[0-9a-zA-Z]{8,64}$", //
		ValidateType: params.ValidateReg,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,
		Require:      false,
	},
	{
		Label:        "支付宝应用私钥",
		Name:         "alipay_rsa_private_key",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaPrivate,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "根证书-crt",
		Name:         "alipay_rsa_root_crt",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaCert,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,
		Require:      false,
	},
}
