package fund

// AccountType 账户类型
type AccountType string

func (e AccountType) Ptr() *AccountType {
	return &e
}

const (
	// AccountType_BASIC 基本账户
	AccountType_BASIC AccountType = "BASIC"
	// AccountType_OPERATION 运营账户
	AccountType_OPERATION AccountType = "OPERATION"
	// AccountType_FEES 手续费账户
	AccountType_FEES AccountType = "FEES"
)

// TarType 压缩类型
type TarType string

const (
	// TarType_GZIP 下载账单时返回.gzip格式的压缩文件流
	TarType_GZIP TarType = "GZIP"
)

