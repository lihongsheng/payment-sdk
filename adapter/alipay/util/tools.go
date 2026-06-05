package util

import (
	"bytes"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"github.com/lihongsheng/payment-sdk/errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// AmountToCents
// 将金额字符串转换为分单位整数
func AmountToCents(amountStr string) (int, error) {
	if amountStr == "" {
		return 0, nil
	}
	// 1. 移除可能存在的逗号（如 1,000.00 处理为 1000.00）
	cleaned := strings.ReplaceAll(amountStr, ",", "")

	// 2. 验证输入是否为合法金额格式
	re := regexp.MustCompile(`^\d+(\.\d{1,2})?$`)
	if !re.MatchString(cleaned) {
		return 0, errors.ErrorParamError("无效的金额格式: %s", amountStr)
	}

	// 3. 处理小数点，确保保留两位小数
	var dollars, cents string
	if strings.Contains(cleaned, ".") {
		parts := strings.Split(cleaned, ".")
		dollars = parts[0]
		centsPart := parts[1]

		// 小数部分补0或截断至两位
		switch len(centsPart) {
		case 1:
			cents = centsPart + "0" // 如 88.0 → 00
		case 2:
			cents = centsPart // 如 88.09 → 09
		default:
			cents = centsPart[:2] // 超过两位则截断（如 88.123 → 12）
		}
	} else {
		// 无小数点，补两位0
		dollars = cleaned
		cents = "00"
	}

	// 4. 拼接并转换为整数（分）
	totalStr := dollars + cents
	return strconv.Atoi(totalStr)
}

func Int64ToFloatString(amount int64) string {
	return fmt.Sprintf("%.2f", float64(amount/100))
}

//	var PaymentStatusMap = map[string]payment.Status{
//		enum.PAYMENT_STATUS_WAIT_BUYER_PAY: payment.Status_Pending,
//		enum.PAYMENT_STATUS_TRADE_CLOSED:   payment.Status_Cancel,
//		enum.PAYMENT_STATUS_TRADE_SUCCESS:  payment.Status_Success,
//		enum.PAYMENT_STATUS_TRADE_FINISHED: payment.Status_Success,
//	}
func PaymentStatus(status string) payment.Status {
	switch status {
	case enum.PAYMENT_STATUS_WAIT_BUYER_PAY:
		return payment.Status_Pending
	case enum.PAYMENT_STATUS_TRADE_CLOSED:
		return payment.Status_Close
	case enum.PAYMENT_STATUS_TRADE_SUCCESS:
		return payment.Status_Success
	case enum.PAYMENT_STATUS_TRADE_FINISHED:
		return payment.Status_Success
	}
	return payment.Status_Status_UNKNOWN
}

// ParseBeijingDateTime 将 time.DateTime 格式的字符串解析为北京时区的 time.Time
func ParseBeijingDateTime(dateTimeStr string) time.Time {
	if dateTimeStr == "" {
		return time.Time{}
	}
	// 北京时间时区
	beijingLoc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// 出错时用东八区
		beijingLoc = time.FixedZone("CST", 8*3600)
	}
	// 使用 time.ParseInLocation 在指定时区解析时间
	ret, err := time.ParseInLocation(time.DateTime, dateTimeStr, beijingLoc)
	if err != nil {
		return time.Time{}
	}
	return ret
}

func GetUnitTransferStatus(status string) transfer.UnitTransferStatus {
	// 描述】转账单据状态。可能出现的状态如下： SUCCESS：转账成功； WAIT_PAY：等待支付； CLOSED：订单超时关闭； FAIL：失败（适用于"单笔转账到银行卡"）； DEALING：处理中（适用于"单笔转账到银行卡"）； REFUND：退票（适用于"单笔转账到银行卡"）； alipay.fund.trans.app.pay涉及的状态： WAIT_PAY、SUCCESS、CLOSED alipay.fund.trans.refund涉及的状态：SUCCESS alipay.fund.trans.uni.transfer涉及的状态：SUCCESS、FAIL、DEALING、REFUND
	switch status {
	case "SUCCESS":
		return transfer.UnitTransferStatus_UnitTransferStatus_SUCCESS
	case "WAIT_PAY":
		return transfer.UnitTransferStatus_UnitTransferStatus_WAIT_PAY
	case "CLOSED":
		return transfer.UnitTransferStatus_UnitTransferStatus_CLOSE
	case "FAIL":
		return transfer.UnitTransferStatus_UnitTransferStatus_FAILED
	case "DEALING":
		return transfer.UnitTransferStatus_UnitTransferStatus_PROCESSING
	case "REFUND":
		return transfer.UnitTransferStatus_UnitTransferStatus_FAILED
	}
	return transfer.UnitTransferStatus_UnitTransferStatus_UNKNOWN
}

func GetRequestBody(request *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.ErrorParamError("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
