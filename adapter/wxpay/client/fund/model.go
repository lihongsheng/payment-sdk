package fund

// QueryBalanceRequest 查询余额请求（服务商专用）
type QueryBalanceRequest struct {
	// AccountType 账户类型
	// 可选取值：BASIC(基本账户)、OPERATION(运营账户)、FEES(手续费账户)
	AccountType AccountType `json:"account_type,omitempty"`
}

// QueryBalanceResponse 查询余额响应（服务商专用）
type QueryBalanceResponse struct {
	// AvailableAmount 可用余额（单位：分），此余额可做提现操作
	AvailableAmount *int64 `json:"available_amount,omitempty"`
	// PendingAmount 不可用余额（单位：分）
	PendingAmount *int64 `json:"pending_amount,omitempty"`
}

// FundFlowBillRequest 申请资金账单请求（普通商户）
type FundFlowBillRequest struct {
	// BillDate 账单日期，格式yyyy-MM-DD，仅支持三个月内的账单下载申请
	BillDate string `json:"bill_date"`
	// AccountType 资金账户类型，不填默认是BASIC
	// 可选取值：BASIC(基本账户)、OPERATION(运营账户)、FEES(手续费账户)
	AccountType AccountType `json:"account_type,omitempty"`
	// TarType 压缩类型，不填则以不压缩的方式返回账单文件流
	// 可选取值：GZIP(下载账单时返回.gzip格式的压缩文件流)
	TarType string `json:"tar_type,omitempty"`
}

// FundFlowBillResponse 申请资金账单响应
type FundFlowBillResponse struct {
	// HashType 哈希类型，固定为SHA1
	HashType string `json:"hash_type"`
	// HashValue 账单文件的SHA1摘要值，用于商户侧校验文件的一致性
	HashValue string `json:"hash_value"`
	// DownloadUrl 供下一步请求账单文件的下载地址，该地址5min内有效
	DownloadUrl string `json:"download_url"`
}

// DownloadBillRequest 下载账单请求
type DownloadBillRequest struct {
	// DownloadUrl 账单下载地址（从FundFlowBillResponse获取）
	DownloadUrl string `json:"download_url"`
}

// DownloadBillResponse 下载账单响应
type DownloadBillResponse struct {
	// Data 账单文件内容（原始字节）
	Data []byte `json:"-"`
	// HashValue 实际计算的SHA1值，用于校验
	HashValue string `json:"-"`
}

// FundFlowRecord 资金流水记录
type FundFlowRecord struct {
	// AccountingTime 记账时间
	AccountingTime string `json:"accounting_time"`
	// TransactionId 微信支付业务单号
	TransactionId string `json:"transaction_id"`
	// BizNo 业务单号
	BizNo string `json:"biz_no"`
	// BizType 业务类型
	BizType string `json:"biz_type"`
	// BizName 业务名称
	BizName string `json:"biz_name"`
	// IncomeExpenseType 收/支类型
	IncomeExpenseType string `json:"income_expense_type"`
	// Amount 收支金额（元）
	Amount string `json:"amount"`
	// Balance 账户结余（元）
	Balance string `json:"balance"`
	// Submitter 资金变更提交者
	Submitter string `json:"submitter"`
	// Remark 备注
	Remark string `json:"remark"`
	// VoucherNo 业务凭证号
	VoucherNo string `json:"voucher_no"`
}

// FundFlowSummary 资金流水汇总
type FundFlowSummary struct {
	// TotalIncome 总收入（元）
	TotalIncome float64 `json:"total_income"`
	// TotalExpense 总支出（元）
	TotalExpense float64 `json:"total_expense"`
	// NetAmount 净额（元）= 总收入 - 总支出
	NetAmount float64 `json:"net_amount"`
	// RecordCount 记录数
	RecordCount int `json:"record_count"`
	// EndBalance 期末余额（元）- 账单最后一条记录的余额
	EndBalance float64 `json:"end_balance"`
}

// AccountBalanceResponse 账户余额响应（普通商户版本，基于账单数据）
type AccountBalanceResponse struct {
	// BillDate 账单日期
	BillDate string `json:"bill_date"`
	// AccountType 账户类型
	AccountType AccountType `json:"account_type"`
	// EndBalance 期末余额（元）
	EndBalance float64 `json:"end_balance"`
	// TotalIncome 当日总收入（元）
	TotalIncome float64 `json:"total_income"`
	// TotalExpense 当日总支出（元）
	TotalExpense float64 `json:"total_expense"`
	// RecordCount 当日交易笔数
	RecordCount int `json:"record_count"`
	// Records 资金流水明细
	Records []FundFlowRecord `json:"records,omitempty"`
}
