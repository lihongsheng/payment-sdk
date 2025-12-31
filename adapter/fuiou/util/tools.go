package util

import (
	"bytes"
	"fmt"
	"github.com/lihongsheng/payment-sdk/enum/transfer"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// URLEncodeGBK 对应 Java 的 URLEncoder.encode(str, "GBK")
// 将字符串按 GBK 编码后进行 URL 编码
func URLEncodeGBK(str string) (string, error) {
	// 1. 将 UTF-8 字符串转换为 GBK 字节流
	gbkBytes, err := ioutil.ReadAll(
		transform.NewReader(
			strings.NewReader(str),
			simplifiedchinese.GBK.NewEncoder(), // UTF-8 → GBK
		),
	)
	if err != nil {
		return "", fmt.Errorf("GBK 编码失败: %v", err)
	}

	// 2. 对 GBK 字节流进行 URL 编码（类似 Java 的 URLEncoder）
	return url.QueryEscape(string(gbkBytes)), nil
}

// URLDecodeGBK 对应 Java 的 URLDecoder.decode(str, "GBK")
// 将 URL 编码的字符串解码后按 GBK 转换为 UTF-8
func URLDecodeGBK(encodedStr string) (string, error) {
	// 1. 先进行 URL 解码（得到 GBK 编码的字节流）
	decodedBytes, err := url.QueryUnescape(encodedStr)
	if err != nil {
		return "", fmt.Errorf("URL 解码失败: %v", err)
	}
	gbkBytes, err := ioutil.ReadAll(
		transform.NewReader(
			strings.NewReader(decodedBytes),
			simplifiedchinese.GBK.NewDecoder(),
		),
	)
	if err != nil {
		return "", fmt.Errorf("GBK 解码失败: %v", err)
	}

	return string(gbkBytes), nil
}

func Utf8ToGbk(utf8Str string) ([]byte, error) {
	gbkBytes, err := ioutil.ReadAll(
		transform.NewReader(
			strings.NewReader(utf8Str),
			simplifiedchinese.GBK.NewEncoder(), // UTF-8 → GBK
		),
	)
	if err != nil {
		return nil, fmt.Errorf("GBK 编码失败: %v", err)
	}
	return gbkBytes, nil
}

// GbkEncode 将字符串按GBK编码转换为字节数组（模拟Java的String.getBytes("GBK")）
func GbkEncode(s string) ([]byte, error) {
	encoder := simplifiedchinese.GBK.NewEncoder()
	result, _, err := transform.Bytes(encoder, []byte(s))
	return result, err
}

// GBKToUTF8 将 GBK 编码的字节流转换为 UTF-8 字符串
func GBKToUTF8(gbkStr string) (string, error) {
	// 创建 GBK 解码器（输入 GBK 字节流，输出 UTF-8 字节流）
	decoder := simplifiedchinese.GBK.NewDecoder()
	// 通过 transform.Reader 转换编码
	utf8Bytes, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(gbkStr), decoder))
	if err != nil {
		return "", fmt.Errorf("GBK 转 UTF-8 失败: %v", err)
	}
	return string(utf8Bytes), nil
}

func GBKToUTF8Byte(gbk []byte) (string, error) {
	// 创建 GBK 解码器（输入 GBK 字节流，输出 UTF-8 字节流）
	decoder := simplifiedchinese.GBK.NewDecoder()
	// 通过 transform.Reader 转换编码
	b := &bytes.Buffer{}
	b.Write(gbk)
	utf8Bytes, err := ioutil.ReadAll(transform.NewReader(b, decoder))
	if err != nil {
		return "", fmt.Errorf("GBK 转 UTF-8 失败: %v", err)
	}
	return string(utf8Bytes), nil
}

func GetTransferStatus(s string) (transfer.DetailStatus, string) {
	// 枚举值：
	//01 处理中
	//05 成功
	//06 失败
	//07 已撤销
	//08 有回退
	switch s {
	case "01":
		return transfer.DetailStatus_DetailStatus_PROCESSING, ""
	case "05":
		return transfer.DetailStatus_DetailStatus_SUCCESS, ""
	case "06":
		return transfer.DetailStatus_DetailStatus_FAILED, ""
	case "07":
		return transfer.DetailStatus_DetailStatus_FAILED, "撤销转账"
	case "08":
		return transfer.DetailStatus_DetailStatus_FAILED, "回退"
	}
	return transfer.DetailStatus_DetailStatus_UNKNOWN, "未知状态:" + s
}

func GetRequestBody(request *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}
