package driver

import "github.com/lihongsheng/payment-sdk/config/params"

var options = []params.Option{
	{
		Label:        "应用id",
		Name:         "app_id",
		Type:         params.String,
		ValidateReg:  "^[0-9a-zA-Z]{8,64}$", //
		ValidateType: params.ValidateReg,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,
		Require:      true,
	},

	{
		Label:        "应用私钥",
		Name:         "rsa_private_key",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaPrivate,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "支付宝证书-crt",
		Name:         "rsa_app_crt",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaCert,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,
		Require:      false,
	},
	{
		Label:        "根证书-crt",
		Name:         "rsa_root_crt",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaCert,
		InputType:    params.InputTextarea, // 长文本用多行输入框
		Default:      "",
		Values:       nil,
		Require:      false,
	},
	{
		Label:        "授权码",
		Name:         "app_auth_token",
		Type:         params.String,
		ValidateReg:  "^.{12,}$", // 秘钥长度32-64位
		ValidateType: params.ValidateReg,
		InputType:    "password", // 敏感字段用密码框
		Default:      "",
		Values:       nil,
		Require:      false,
	},
}
