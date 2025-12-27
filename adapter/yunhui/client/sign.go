package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/tools"
	"net/url"
	"sort"
	"strings"
)

type Sign struct {
	conf config.Config
}

func NewSign(conf config.Config) (*Sign, error) {
	return &Sign{
		conf: conf,
	}, nil
}

func (s *Sign) Gen(body interface{}) (map[string]any, error) {
	var params = map[string]any{}
	by, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(by, &params)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(params))
	params["signType"] = "MD5"
	params["version"] = "1.0"
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""

	for _, k := range keys {
		strVal := ""
		val, ok := params[k] // 先判断 key 是否存在，避免 panic
		if !ok {
			continue // 无该参数则跳过
		}

		switch v := val.(type) {
		case string:
			// 字符串非空则赋值
			if v != "" {
				strVal = v
			}
		case uint8:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case uint16:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case uint32:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case uint64:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case int:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case int8:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case int16:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case int32:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case int64:
			if v > 0 {
				strVal = fmt.Sprintf("%d", v)
			}
		case float64:
			if v > 0 {
				strVal = fmt.Sprintf("%d", int(v))
			}
		case float32:
			if v > 0 {
				strVal = fmt.Sprintf("%d", int(v))
			}
		default:
			// 其他类型（如 bool、struct 等）JSON 序列化
			js, err := json.Marshal(val)
			if err != nil {
				return nil, fmt.Errorf("参数 %s 序列化失败: %w", k, err)
			}
			strVal = string(js)
			// 序列化后可能是空值（如 null），需过滤
			if strVal == "null" || strVal == "" {
				strVal = ""
			}
		}

		if strVal != "" {
			signStr += k + "=" + strVal + "&"
		}
	}
	signStr = signStr + "key=" + s.conf.APIKey
	sign := tools.Md5(signStr)
	sign = strings.ToUpper(sign)
	fmt.Println(signStr, sign)
	params["sign"] = sign

	return params, nil
}

func (s *Sign) Verify(val url.Values) error {
	//var params = map[string]string{}
	keys := make([]string, 0, 20)
	for k, _ := range val {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, k := range keys {
		strVal := val.Get(k)
		if strVal != "" {
			signStr += k + "=" + strVal + "&"
		}
	}
	signStr = signStr + "key=" + s.conf.APIKey
	sign := tools.Md5(signStr)
	sign = strings.ToUpper(sign)
	if sign != val.Get("sign") {
		return errors.New("签名验证未通过")
	}
	return nil
}
