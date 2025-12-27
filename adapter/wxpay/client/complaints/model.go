package complaints

import "time"

// 客诉详情
type ComplaintDetailResponse struct {
	ComplaintId        string    `json:"complaint_id"`
	ComplaintTime      time.Time `json:"complaint_time"`
	ComplaintDetail    string    `json:"complaint_detail"`
	ComplaintState     string    `json:"complaint_state"`
	PayerPhone         string    `json:"payer_phone"`
	PayerOpenid        string    `json:"payer_openid"`
	ComplaintOrderInfo []struct {
		TransactionId string `json:"transaction_id"`
		OutTradeNo    string `json:"out_trade_no"`
		Amount        int    `json:"amount"`
	} `json:"complaint_order_info"`
	ComplaintFullRefunded bool `json:"complaint_full_refunded"`
	IncomingUserResponse  bool `json:"incoming_user_response"`
	UserComplaintTimes    int  `json:"user_complaint_times"`
	ComplaintMediaList    []struct {
		MediaType string   `json:"media_type"`
		MediaUrl  []string `json:"media_url"`
	} `json:"complaint_media_list"`
	ProblemDescription string   `json:"problem_description"`
	ProblemType        string   `json:"problem_type"`
	ApplyRefundAmount  int      `json:"apply_refund_amount"`
	UserTagList        []string `json:"user_tag_list"`
	ServiceOrderInfo   []struct {
		OrderId    string `json:"order_id"`
		OutOrderNo string `json:"out_order_no"`
		State      string `json:"state"`
	} `json:"service_order_info"`
	AdditionalInfo struct {
		Type           string `json:"type"`
		SharePowerInfo struct {
			ReturnTime        time.Time `json:"return_time"`
			ReturnAddressInfo struct {
				ReturnAddress string `json:"return_address"`
				Longitude     string `json:"longitude"`
				Latitude      string `json:"latitude"`
			} `json:"return_address_info"`
			IsReturnedToSameMachine bool `json:"is_returned_to_same_machine"`
		} `json:"share_power_info"`
	} `json:"additional_info"`
	InPlatformService    bool `json:"in_platform_service"`
	NeedImmediateService bool `json:"need_immediate_service"`
	IsAgentMode          bool `json:"is_agent_mode"`
}

// 回复
type ReplayComplaintRequest struct {
	ComplaintedMchid    string               `json:"complainted_mchid,omitempty"`
	ResponseContent     string               `json:"response_content,omitempty"`
	ResponseImages      []string             `json:"response_images,omitempty"`
	JumpUrl             string               `json:"jump_url,omitempty"`
	JumpUrlText         string               `json:"jump_url_text,omitempty"`
	MiniProgramJumpInfo *MiniProgramJumpInfo `json:"mini_program_jump_info,omitempty"`
}

type MiniProgramJumpInfo struct {
	Appid string `json:"appid,omitempty"`
	Path  string `json:"path,omitempty"`
	Text  string `json:"text,omitempty"`
}

// 退款审批请求
type RefundHandlerRequest struct {
	Action          string   `json:"action,omitempty"`
	LaunchRefundDay int      `json:"launch_refund_day,omitempty"`
	RejectReason    string   `json:"reject_reason,omitempty"`
	RejectMediaList []string `json:"reject_media_list,omitempty"`
	Remark          string   `json:"remark,omitempty"`
}

type CompleteRequest struct {
	ComplaintedMchid string `json:"complainted_mchid"`
}

type SearchComplaintRequest struct {
	Limit            int    `json:"limit,omitempty"`
	Offset           int    `json:"offset,omitempty"`
	BeginDate        string `json:"begin_date,omitempty"`
	EndDate          string `json:"end_date,omitempty"`
	ComplaintedMchid string `json:"complainted_mchid,omitempty"`
}

type SearchComplaintResponse struct {
	Data []struct {
		ComplaintId        string    `json:"complaint_id"`
		ComplaintTime      time.Time `json:"complaint_time"`
		ComplaintDetail    string    `json:"complaint_detail"`
		ComplaintState     string    `json:"complaint_state"`
		PayerPhone         string    `json:"payer_phone"`
		ComplaintOrderInfo []struct {
			TransactionId string `json:"transaction_id"`
			OutTradeNo    string `json:"out_trade_no"`
			Amount        int    `json:"amount"`
		} `json:"complaint_order_info"`
		ComplaintFullRefunded bool `json:"complaint_full_refunded"`
		IncomingUserResponse  bool `json:"incoming_user_response"`
		UserComplaintTimes    int  `json:"user_complaint_times"`
		ComplaintMediaList    []struct {
			MediaType string   `json:"media_type"`
			MediaUrl  []string `json:"media_url"`
		} `json:"complaint_media_list"`
		ProblemDescription string   `json:"problem_description"`
		ProblemType        string   `json:"problem_type"`
		ApplyRefundAmount  int      `json:"apply_refund_amount"`
		UserTagList        []string `json:"user_tag_list"`
		ServiceOrderInfo   []struct {
			OrderId    string `json:"order_id"`
			OutOrderNo string `json:"out_order_no"`
			State      string `json:"state"`
		} `json:"service_order_info"`
		AdditionalInfo struct {
			Type           string `json:"type"`
			SharePowerInfo struct {
				ReturnTime        time.Time `json:"return_time"`
				ReturnAddressInfo struct {
					ReturnAddress string `json:"return_address"`
					Longitude     string `json:"longitude"`
					Latitude      string `json:"latitude"`
				} `json:"return_address_info"`
				IsReturnedToSameMachine bool `json:"is_returned_to_same_machine"`
			} `json:"share_power_info"`
		} `json:"additional_info"`
		InPlatformService    bool `json:"in_platform_service"`
		NeedImmediateService bool `json:"need_immediate_service"`
		IsAgentMode          bool `json:"is_agent_mode"`
	} `json:"data"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}

type AddComplaintNoticeUrl struct {
	Url string `json:"url"`
}

// 查询回调的返回结果
type QueryCallbackUrlResponse struct {
	Mchid string `json:"mchid"`
	Url   string `json:"url"`
}

type UpdateComplaintNoticeUrl struct {
	Url string `json:"url"`
}

// 查询投诉单协商历史请求
type QueryNegotiationHistoryRequest struct {
	Limit  int `json:"limit,omitempty"`  // 分页大小，范围[1,300]
	Offset int `json:"offset,omitempty"` // 分页开始位置，从0开始计数
}

// 查询投诉单协商历史响应
type QueryNegotiationHistoryResponse struct {
	Data       []NegotiationHistoryItem `json:"data,omitempty"`        // 投诉协商历史
	Limit      int                      `json:"limit"`                 // 分页大小
	Offset     int                      `json:"offset"`                // 分页开始位置
	TotalCount int                      `json:"total_count,omitempty"` // 投诉协商历史总条数，当offset=0时返回
}

// 协商历史项
type NegotiationHistoryItem struct {
	LogId                                    string              `json:"log_id,omitempty"`                                        // 日志ID
	Operator                                 string              `json:"operator,omitempty"`                                      // 操作人
	OperateTime                              time.Time           `json:"operate_time,omitempty"`                                  // 操作时间
	OperateType                              string              `json:"operate_type,omitempty"`                                  // 操作类型
	OperateDetails                           string              `json:"operate_details,omitempty"`                               // 操作详情
	ImageList                                []string            `json:"image_list,omitempty"`                                    // 图片列表
	ComplaintMediaList                       *ComplaintMediaList `json:"complaint_media_list,omitempty"`                          // 投诉媒体列表
	UserAppyPlatformServiceReason            string              `json:"user_appy_platform_service_reason,omitempty"`             // 用户申请平台介入原因
	UserAppyPlatformServiceReasonDescription string              `json:"user_appy_platform_service_reason_description,omitempty"` // 用户申请平台介入原因描述
}

// 投诉媒体列表
type ComplaintMediaList struct {
	MediaType string   `json:"media_type,omitempty"` // 媒体类型
	MediaUrl  []string `json:"media_url,omitempty"`  // 媒体URL数组
}
