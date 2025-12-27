package payscore

import "encoding/json"

type CreateServiceOrderRequest struct {
	OutOrderNo          *string              `json:"out_order_no,omitempty"`
	Appid               *string              `json:"appid,omitempty"`
	ServiceId           *string              `json:"service_id,omitempty"`
	ServiceIntroduction *string              `json:"service_introduction,omitempty"`
	PostPayments        []Payment            `json:"post_payments,omitempty"`
	PostDiscounts       []ServiceOrderCoupon `json:"post_discounts,omitempty"`
	TimeRange           *TimeRange           `json:"time_range,omitempty"`
	Location            *Location            `json:"location,omitempty"`
	RiskFund            *RiskFund            `json:"risk_fund,omitempty"`
	Attach              *string              `json:"attach,omitempty"`
	NotifyUrl           *string              `json:"notify_url,omitempty"`
	NeedUserConfirm     *bool                `json:"need_user_confirm,omitempty"`
	Device              *Device              `json:"device,omitempty"`
}

type CreateServiceOrderResponse struct {
	Appid               *string              `json:"appid,omitempty"`
	Mchid               *string              `json:"mchid,omitempty"`
	OutOrderNo          *string              `json:"out_order_no,omitempty"`
	ServiceId           *string              `json:"service_id,omitempty"`
	ServiceIntroduction *string              `json:"service_introduction,omitempty"`
	State               *string              `json:"state,omitempty"`
	StateDescription    *string              `json:"state_description,omitempty"`
	PostPayments        []Payment            `json:"post_payments,omitempty"`
	PostDiscounts       []ServiceOrderCoupon `json:"post_discounts,omitempty"`
	RiskFund            *RiskFund            `json:"risk_fund,omitempty"`
	TimeRange           *TimeRange           `json:"time_range,omitempty"`
	Location            *Location            `json:"location,omitempty"`
	Attach              *string              `json:"attach,omitempty"`
	NotifyUrl           *string              `json:"notify_url,omitempty"`
	OrderId             *string              `json:"order_id,omitempty"`
	Package             *string              `json:"package,omitempty"`
}

type Payment struct {
	Name        *string `json:"name,omitempty"`
	Amount      *int64  `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
	Count       *int64  `json:"count,omitempty"`
}

type ServiceOrderCoupon struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Amount      *int64  `json:"amount,omitempty"`
	Count       *int64  `json:"count,omitempty"`
}

type TimeRange struct {
	StartTime       *string `json:"start_time,omitempty"`
	EndTime         *string `json:"end_time,omitempty"`
	StartTimeRemark *string `json:"start_time_remark,omitempty"`
	EndTimeRemark   *string `json:"end_time_remark,omitempty"`
}

type Location struct {
	StartLocation *string `json:"start_location,omitempty"`
	EndLocation   *string `json:"end_location,omitempty"`
}

type RiskFund struct {
	Name        *string `json:"name,omitempty"`
	Amount      *int64  `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
}

type Device struct {
	StartDeviceId *string `json:"start_device_id,omitempty"`
	EndDeviceId   *string `json:"end_device_id,omitempty"`
	MaterielNo    *string `json:"materiel_no,omitempty"`
}

type GetServiceOrderRequest struct {
	OutOrderNo *string `json:"out_order_no,omitempty"`
	ServiceId  *string `json:"service_id,omitempty"`
	Appid      *string `json:"appid,omitempty"`
	QueryId    *string `json:"query_id,omitempty"`
}

func (o *GetServiceOrderRequest) MarshalJSON() ([]byte, error) {
	type Alias GetServiceOrderRequest
	a := &struct {
		OutOrderNo *string `json:"out_order_no,omitempty"`
		ServiceId  *string `json:"service_id,omitempty"`
		Appid      *string `json:"appid,omitempty"`
		QueryId    *string `json:"query_id,omitempty"`
		*Alias
	}{
		// 序列化时移除非 Body 字段
		OutOrderNo: nil,
		ServiceId:  nil,
		Appid:      nil,
		QueryId:    nil,
		Alias:      (*Alias)(o),
	}
	return json.Marshal(a)
}

type ServiceOrderEntity struct {
	OutOrderNo          *string              `json:"out_order_no,omitempty"`
	ServiceId           *string              `json:"service_id,omitempty"`
	Appid               *string              `json:"appid,omitempty"`
	Mchid               *string              `json:"mchid,omitempty"`
	ServiceIntroduction *string              `json:"service_introduction,omitempty"`
	State               *string              `json:"state,omitempty"`
	StateDescription    *string              `json:"state_description,omitempty"`
	PostPayments        *Payment             `json:"post_payments,omitempty"`
	PostDiscounts       []ServiceOrderCoupon `json:"post_discounts,omitempty"`
	RiskFund            *RiskFund            `json:"risk_fund,omitempty"`
	TotalAmount         *int64               `json:"total_amount,omitempty"`
	NeedCollection      *bool                `json:"need_collection,omitempty"`
	Collection          *Collection          `json:"collection,omitempty"`
	TimeRange           *TimeRange           `json:"time_range,omitempty"`
	Location            *Location            `json:"location,omitempty"`
	Attach              *string              `json:"attach,omitempty"`
	NotifyUrl           *string              `json:"notify_url,omitempty"`
	Openid              *string              `json:"openid,omitempty"`
	OrderId             *string              `json:"order_id,omitempty"`
}

type Collection struct {
	State        *string  `json:"state,omitempty"`
	TotalAmount  *int64   `json:"total_amount,omitempty"`
	PayingAmount *int64   `json:"paying_amount,omitempty"`
	PaidAmount   *int64   `json:"paid_amount,omitempty"`
	Details      []Detail `json:"details,omitempty"`
}

type Detail struct {
	Seq           *int64  `json:"seq,omitempty"`
	Amount        *int64  `json:"amount,omitempty"`
	PaidType      *string `json:"paid_type,omitempty"`
	PaidTime      *string `json:"paid_time,omitempty"`
	TransactionId *string `json:"transaction_id,omitempty"`
}

type CancelServiceOrderRequest struct {
	OutOrderNo *string `json:"out_order_no,omitempty"`
	Appid      *string `json:"appid,omitempty"`
	ServiceId  *string `json:"service_id,omitempty"`
	Reason     *string `json:"reason,omitempty"`
}

func (o *CancelServiceOrderRequest) MarshalJSON() ([]byte, error) {
	type Alias CancelServiceOrderRequest
	a := &struct {
		OutOrderNo *string `json:"out_order_no,omitempty"`
		*Alias
	}{
		// 序列化时移除非 Body 字段
		OutOrderNo: nil,
		Alias:      (*Alias)(o),
	}
	return json.Marshal(a)
}

type CancelServiceOrderResponse struct {
	Appid      *string `json:"appid,omitempty"`
	Mchid      *string `json:"mchid,omitempty"`
	OutOrderNo *string `json:"out_order_no,omitempty"`
	ServiceId  *string `json:"service_id,omitempty"`
	OrderId    *string `json:"order_id,omitempty"`
}

type CompleteServiceOrderRequest struct {
	OutOrderNo    *string              `json:"out_order_no,omitempty"`
	Appid         *string              `json:"appid,omitempty"`
	ServiceId     *string              `json:"service_id,omitempty"`
	PostPayments  []Payment            `json:"post_payments,omitempty"`
	PostDiscounts []ServiceOrderCoupon `json:"post_discounts,omitempty"`
	TotalAmount   *int64               `json:"total_amount,omitempty"`
	TimeRange     *TimeRange           `json:"time_range,omitempty"`
	Location      *Location            `json:"location,omitempty"`
	ProfitSharing *bool                `json:"profit_sharing,omitempty"`
	GoodsTag      *string              `json:"goods_tag,omitempty"`
	Device        *Device              `json:"device,omitempty"`
}

func (o *CompleteServiceOrderRequest) MarshalJSON() ([]byte, error) {
	type Alias CompleteServiceOrderRequest
	a := &struct {
		OutOrderNo *string `json:"out_order_no,omitempty"`
		*Alias
	}{
		// 序列化时移除非 Body 字段
		OutOrderNo: nil,
		Alias:      (*Alias)(o),
	}
	return json.Marshal(a)
}

type CompleteServiceOrderResponse struct {
	Appid               *string              `json:"appid,omitempty"`
	Mchid               *string              `json:"mchid,omitempty"`
	OutOrderNo          *string              `json:"out_order_no,omitempty"`
	ServiceId           *string              `json:"service_id,omitempty"`
	ServiceIntroduction *string              `json:"service_introduction,omitempty"`
	State               *string              `json:"state,omitempty"`
	StateDescription    *string              `json:"state_description,omitempty"`
	TotalAmount         *int64               `json:"total_amount,omitempty"`
	PostPayments        []Payment            `json:"post_payments,omitempty"`
	PostDiscounts       []ServiceOrderCoupon `json:"post_discounts,omitempty"`
	RiskFund            *RiskFund            `json:"risk_fund,omitempty"`
	TimeRange           *TimeRange           `json:"time_range,omitempty"`
	Location            *Location            `json:"location,omitempty"`
	OrderId             *string              `json:"order_id,omitempty"`
	NeedCollection      *bool                `json:"need_collection,omitempty"`
}
