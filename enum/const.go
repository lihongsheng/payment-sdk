package enum

type Device string

const (
	// PC网页
	Device_PC Device = "PC"
	// 手机网页
	Device_H5 Device = "H5"
	// 微信扫码，微信内置浏览器
	Device_Wechat Device = "Wechat"
	// 支付宝扫码, 支付宝内置浏览器
	Device_Alipay Device = "Alipay"
	// 微信小程序
	Device_Wechat_Lite Device = "WechatLite"
	// 支付宝小程序
	Device_Alipay_Lite Device = "AlipayLite"
)

var DeviceDesc = map[Device]string{
	Device_PC:          "PC网页",
	Device_H5:          "手机网页",
	Device_Wechat:      "微信扫码",
	Device_Alipay:      "支付宝扫码",
	Device_Wechat_Lite: "微信小程序",
	Device_Alipay_Lite: "支付宝小程序",
}

type System string

const (
	IOS     System = "IOS"
	Android System = "Android"
	Linux   System = "Linux"
	MacOS   System = "MacOS"
	Windows System = "Windows"
)
