package fund

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	nethttp "net/http"
	neturl "net/url"
	"strconv"
	"strings"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type FundApiService services.Service

// QueryBalance 查询账户实时余额
// 商户通过此接口可以查询本商户号的账号余额情况
// 支持商户：【普通服务商】
// 接口文档：https://pay.weixin.qq.com/doc/v3/partner/4012712163
func (s *FundApiService) QueryBalance(ctx context.Context, req QueryBalanceRequest) (resp *QueryBalanceResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	if req.AccountType == "" {
		return nil, nil, fmt.Errorf("account_type must not be empty")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/merchant/fund/balance/{account_type}"
	localVarPath = strings.Replace(localVarPath, "{account_type}", neturl.PathEscape(core.ParameterToString(req.AccountType, "")), -1)

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Response from Http Response
	resp = new(QueryBalanceResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// GetFundFlowBill 申请资金账单
// 微信支付按天提供商户各账户的资金流水账单文件，商户可以通过该接口获取账单文件的下载地址
// 账单文件详细记录了账户资金操作的相关信息，包括业务单号、收支金额及记账时间等
// 支持商户：【普通商户】
// 注意：
// - 资金账单中的数据反映的是商户微信账户资金变动情况
// - 当日账单将在次日上午9点开始生成，建议商户在次日上午10点以后获取
// - 资金账单中所有涉及金额的字段均以"元"为单位
// - 以商户号维度频率限制为3QPS
func (s *FundApiService) GetFundFlowBill(ctx context.Context, req FundFlowBillRequest) (resp *FundFlowBillResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	if req.BillDate == "" {
		return nil, nil, fmt.Errorf("bill_date must not be empty")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/bill/fundflowbill"

	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("bill_date", req.BillDate)
	if req.AccountType != "" {
		localVarQueryParams.Add("account_type", string(req.AccountType))
	}
	if req.TarType != "" {
		localVarQueryParams.Add("tar_type", req.TarType)
	}

	localVarHTTPContentTypes := []string{"application/json"}
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	resp = new(FundFlowBillResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DownloadBill 下载账单文件
// 通过GetFundFlowBill获取的下载地址下载账单文件，该接口也需要认证
// downloadUrl: 从FundFlowBillResponse.DownloadUrl获取
// isGzip: 是否是GZIP压缩格式（申请账单时指定了TarType=GZIP）
// 注意：账单下载接口返回的是原始文件，不带响应签名，需要跳过验签
func (s *FundApiService) DownloadBill(ctx context.Context, downloadUrl string, isGzip bool) (resp *DownloadBillResponse, err error) {
	if downloadUrl == "" {
		return nil, fmt.Errorf("download_url must not be empty")
	}

	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarHeaderParams.Set("Accept", "*/*")
	var localVarHTTPContentTypes []string
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	result, err := s.Client.Request(ctx, localVarHTTPMethod, downloadUrl, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)

	// 即使SDK报验签错误，仍可获取响应数据（账单下载接口不返回签名头）
	if result != nil && result.Response != nil {
		defer result.Response.Body.Close()

		if result.Response.StatusCode != nethttp.StatusOK {
			bodyBytes, _ := io.ReadAll(result.Response.Body)
			return nil, fmt.Errorf("download bill failed with status %d: %s", result.Response.StatusCode, string(bodyBytes))
		}

		data, readErr := io.ReadAll(result.Response.Body)
		if readErr != nil {
			return nil, fmt.Errorf("read response body failed: %w", readErr)
		}

		if isGzip {
			gzReader, gzErr := gzip.NewReader(bytes.NewReader(data))
			if gzErr != nil {
				return nil, fmt.Errorf("create gzip reader failed: %w", gzErr)
			}
			defer gzReader.Close()
			data, readErr = io.ReadAll(gzReader)
			if readErr != nil {
				return nil, fmt.Errorf("decompress gzip failed: %w", readErr)
			}
		}

		hash := sha1.New()
		hash.Write(data)
		hashValue := hex.EncodeToString(hash.Sum(nil))

		resp = &DownloadBillResponse{
			Data:      data,
			HashValue: hashValue,
		}
		return resp, nil
	}

	if err != nil {
		return nil, fmt.Errorf("download bill request failed: %w", err)
	}

	return nil, fmt.Errorf("download bill failed: no response received")
}

// 表头字段名映射（支持多种可能的列名）
var headerMapping = map[string]string{
	"记账时间":     "accounting_time",
	"微信支付业务单号": "transaction_id",
	"资金流水单号":   "transaction_id",
	"业务单号":     "biz_no",
	"商户单号":     "biz_no",
	"业务类型":     "biz_type",
	"交易类型":     "biz_type",
	"业务名称":     "biz_name",
	"交易名称":     "biz_name",
	"收/支":      "income_expense_type",
	"收支类型":     "income_expense_type",
	"收入/支出":    "income_expense_type",
	"收支金额(元)":  "amount",
	"金额(元)":    "amount",
	"收支金额":     "amount",
	"金额":       "amount",
	"账户结余(元)":  "balance",
	"结余(元)":    "balance",
	"账户结余":     "balance",
	"结余":       "balance",
	"资金变更提交者":  "submitter",
	"提交者":      "submitter",
	"操作者":      "submitter",
	"备注":       "remark",
	"业务凭证号":    "voucher_no",
	"凭证号":      "voucher_no",
}

// ParseFundFlowBill 解析资金账单内容
// data: 账单文件内容（从DownloadBill获取）
// 返回资金流水记录列表
// 支持动态表头解析，能够应对字段顺序变化和字段名更改
func ParseFundFlowBill(data []byte) ([]FundFlowRecord, error) {
	var records []FundFlowRecord
	var columnIndex = make(map[string]int)

	// 移除 UTF-8 BOM（如果存在）
	if len(data) >= 3 && data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
		data = data[3:]
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	headerParsed := false
	inSummary := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(line, "资金账单汇总") || strings.Contains(line, "汇总数据") {
			inSummary = true
			continue
		}

		if inSummary {
			continue
		}

		fields := parseCSVLine(line)

		if !headerParsed && isHeaderLine(fields) {
			columnIndex = parseHeader(fields)
			headerParsed = true
			continue
		}

		if !headerParsed {
			continue
		}

		record := parseDataLine(fields, columnIndex)
		if record != nil {
			records = append(records, *record)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan bill content failed: %w", err)
	}

	return records, nil
}

// parseCSVLine 解析CSV行，处理`前缀
func parseCSVLine(line string) []string {
	fields := strings.Split(line, ",")
	for i := range fields {
		fields[i] = strings.TrimPrefix(fields[i], "`")
		fields[i] = strings.TrimSpace(fields[i])
	}
	return fields
}

// isHeaderLine 判断是否为表头行（至少匹配3个已知字段）
func isHeaderLine(fields []string) bool {
	if len(fields) < 3 {
		return false
	}
	matchCount := 0
	for _, field := range fields {
		if _, ok := headerMapping[field]; ok {
			matchCount++
		}
	}
	return matchCount >= 3
}

// parseHeader 解析表头，建立列名到索引的映射
func parseHeader(fields []string) map[string]int {
	index := make(map[string]int)
	for i, field := range fields {
		if mappedName, ok := headerMapping[field]; ok {
			index[mappedName] = i
		}
	}
	return index
}

// parseDataLine 根据列索引解析数据行
func parseDataLine(fields []string, columnIndex map[string]int) *FundFlowRecord {
	if len(fields) < 3 {
		return nil
	}

	record := &FundFlowRecord{}
	getValue := func(key string) string {
		if idx, ok := columnIndex[key]; ok && idx < len(fields) {
			return fields[idx]
		}
		return ""
	}

	record.AccountingTime = getValue("accounting_time")
	record.TransactionId = getValue("transaction_id")
	record.BizNo = getValue("biz_no")
	record.BizType = getValue("biz_type")
	record.BizName = getValue("biz_name")
	record.IncomeExpenseType = getValue("income_expense_type")
	record.Amount = getValue("amount")
	record.Balance = getValue("balance")
	record.Submitter = getValue("submitter")
	record.Remark = getValue("remark")
	record.VoucherNo = getValue("voucher_no")

	if record.Amount == "" {
		return nil
	}

	return record
}

// 收支类型关键词映射
var incomeKeywords = []string{"收入", "入账", "转入", "收款"}
var expenseKeywords = []string{"支出", "出账", "转出", "付款", "退款"}

// SummarizeFundFlow 汇总资金流水
func SummarizeFundFlow(records []FundFlowRecord) (*FundFlowSummary, error) {
	summary := &FundFlowSummary{
		RecordCount: len(records),
	}

	for _, record := range records {
		amountStr := strings.ReplaceAll(record.Amount, ",", "")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			continue
		}

		if isIncomeType(record.IncomeExpenseType) {
			summary.TotalIncome += amount
		} else if isExpenseType(record.IncomeExpenseType) {
			summary.TotalExpense += amount
		}
	}

	summary.NetAmount = summary.TotalIncome - summary.TotalExpense

	if len(records) > 0 {
		balanceStr := strings.ReplaceAll(records[len(records)-1].Balance, ",", "")
		lastBalance, err := strconv.ParseFloat(balanceStr, 64)
		if err == nil {
			summary.EndBalance = lastBalance
		}
	}

	return summary, nil
}

// isIncomeType 判断是否为收入类型
func isIncomeType(typeStr string) bool {
	for _, keyword := range incomeKeywords {
		if strings.Contains(typeStr, keyword) {
			return true
		}
	}
	return false
}

// isExpenseType 判断是否为支出类型
func isExpenseType(typeStr string) bool {
	for _, keyword := range expenseKeywords {
		if strings.Contains(typeStr, keyword) {
			return true
		}
	}
	return false
}

// GetAccountBalance 获取账户余额（普通商户版本，基于账单数据，获取指定日期结束时的余额）
func (s *FundApiService) GetAccountBalance(ctx context.Context, billDate string, accountType AccountType) (*AccountBalanceResponse, error) {
	billResp, _, err := s.GetFundFlowBill(ctx, FundFlowBillRequest{
		BillDate:    billDate,
		AccountType: accountType,
	})
	if err != nil {
		return nil, fmt.Errorf("get fund flow bill failed: %w", err)
	}

	downloadResp, err := s.DownloadBill(ctx, billResp.DownloadUrl, false)
	if err != nil {
		return nil, fmt.Errorf("download bill failed: %w", err)
	}

	if !strings.EqualFold(downloadResp.HashValue, billResp.HashValue) {
		return nil, fmt.Errorf("bill hash mismatch: expected %s, got %s", billResp.HashValue, downloadResp.HashValue)
	}

	records, err := ParseFundFlowBill(downloadResp.Data)
	if err != nil {
		return nil, fmt.Errorf("parse bill failed: %w", err)
	}

	summary, err := SummarizeFundFlow(records)
	if err != nil {
		return nil, fmt.Errorf("summarize fund flow failed: %w", err)
	}

	return &AccountBalanceResponse{
		BillDate:     billDate,
		AccountType:  accountType,
		EndBalance:   summary.EndBalance,
		TotalIncome:  summary.TotalIncome,
		TotalExpense: summary.TotalExpense,
		RecordCount:  summary.RecordCount,
		Records:      records,
	}, nil
}
