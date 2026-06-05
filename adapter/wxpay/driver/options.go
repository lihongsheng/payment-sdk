package driver

import "github.com/lihongsheng/payment-sdk/config/params"

var options = []params.Option{
	{
		Label:        "微信商户ID",
		Name:         "mch_id",
		Type:         params.String,
		ValidateReg:  "^[0-9]{10,15}$",
		ValidateType: "Reg",
		InputType:    "text",
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "微信AppID",
		Name:         "app_id",
		Type:         params.String,
		ValidateReg:  "^wx[0-9a-zA-Z]{16,32}$",
		ValidateType: params.ValidateReg,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "应用Secret",
		Name:         "app_secret",
		Type:         params.String,
		ValidateReg:  "^.{6,}$", //
		ValidateType: params.ValidateReg,
		InputType:    params.InputPassword,
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "微信API V3秘钥",
		Name:         "api_secret",
		Type:         params.String,
		ValidateReg:  "^.{6,}$", //
		ValidateType: params.ValidateReg,
		InputType:    params.InputPassword,
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "应用私钥",
		Name:         "rsa_private",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaPrivate,
		InputType:    params.InputTextarea,
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "应用证书序列号",
		Name:         "rsa_private_number",
		Type:         params.String,
		ValidateReg:  "^.{6,}$", //
		ValidateType: params.ValidateReg,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,
		Require:      true,
	},
	{
		Label:        "微信公钥证书序列号",
		Name:         "rsa_public_number",
		Type:         params.String,
		ValidateReg:  "^.{6,}$", //
		ValidateType: params.ValidateReg,
		InputType:    params.InputText,
		Default:      "",
		Values:       nil,
		Require:      false,
	},
	{
		Label:        "微信公钥",
		Name:         "rsa_public",
		Type:         params.String,
		ValidateReg:  "",
		ValidateType: params.ValidateRsaPublic,
		InputType:    params.InputTextarea,
		Default:      "",
		Values:       nil,
		Require:      false,
	},
}
