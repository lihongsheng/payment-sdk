package enum

type Env struct {
	Sftp FileHost `json:"sftp"`
}

type FileHost struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
	User string `json:"user,omitempty"`
	Pwd  string `json:"pwd,omitempty"`
}

//
//func Sandbox() Env {
//	return Env{
//		ApiHost: "https://richOperationFront-test.fuioupay.com",
//		Sftp: FileHost{
//			Host: "ftp-1.fuioupay.com",
//			Port: 9022,
//			User: "FGJ370542",
//			Pwd:  "BnjiKm3CVI7rm1Dx",
//		},
//	}
//}
//
//func Production() Env {
//	return Env{
//		ApiHost: "https://richOperationFront.fuioupay.com",
//		Sftp: FileHost{
//			Host: "ftp-1.fuioupay.com",
//			Port: 9022,
//			User: "FGJ370542",
//			Pwd:  "BnjiKm3CVI7rm1Dx",
//		},
//	}
//}
