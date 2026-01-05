package enum

type Device string

const (
	Device_PC     Device = "PC"
	Device_H5     Device = "H5"
	Device_Wechat Device = "Wechat"
	Device_Alipay Device = "Alipay"
)

type System string

const (
	IOS     System = "IOS"
	Android System = "Android"
	Linux   System = "Linux"
	MacOS   System = "MacOS"
	Windows System = "Windows"
)
