package model

// JFCreateOrderRequest 创建订单请求
type JFCreateOrderRequest struct {
	MerchID           string                `json:"merchId"`
	TradeAmount       string                `json:"tradeAmount"`
	Remark            string                `json:"remark"`
	OrderTemplateData []JFOrderTemplateItem `json:"orderTemplateData"`
}

// JFOrderTemplateItem 订单模板项
type JFOrderTemplateItem struct {
	Key                 int64                  `json:"key"`
	Type                string                 `json:"type"`
	Index               int                    `json:"index"`
	Label               string                 `json:"label"`
	Value               string                 `json:"value"`
	Origin              string                 `json:"origin"`
	Options             map[string]interface{} `json:"options"`
	DisplayName         string                 `json:"displayName"`
	FormItemFlag        bool                   `json:"formItemFlag"`
	SettingsTitle       string                 `json:"settingsTitle"`
	MarginLeftRight     int                    `json:"marginLeftRight"`
	MarginTopBottom     int                    `json:"marginTopBottom"`
	CashierTemplateName string                 `json:"cashierTemplateName"`
	State               bool                   `json:"state"`
}

// JFCreateOrderResponse 创建订单响应
type JFCreateOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data struct {
		PayURL string `json:"payUrl"`
	} `json:"data,omitempty"`
}

// JFQueryOrderRequest 查询订单请求
type JFQueryOrderRequest struct {
	ReqTime string         `json:"reqTime"`
	Version string         `json:"version"`
	ReqData JFQueryReqData `json:"reqData"`
}

// JFQueryReqData 查询请求数据
type JFQueryReqData struct {
	ChannelID  string `json:"channelId"`
	PayOrderNo string `json:"payOrderNo"`
	MerchantNo string `json:"merchantNo"`
}

// JFQueryOrderResponse 查询订单响应
type JFQueryOrderResponse struct {
	Code     string `json:"code"` // 可能是字符串"000000"或整数0
	Msg      string `json:"msg,omitempty"`
	RespData struct {
		OrderStatus string `json:"orderStatus"` // 2=已支付
		PayOrderNo  string `json:"payOrderNo"`
		TotalAmount string `json:"totalAmount"`
	} `json:"respData,omitempty"`
}
