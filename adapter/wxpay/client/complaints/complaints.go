package complaints

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type ComplaintApiService services.Service

// Search 投诉单列表
func (s *ComplaintApiService) Search(ctx context.Context, req SearchComplaintRequest) (resp *SearchComplaintResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaints-v2"
	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.Limit > 0 {
		localVarQueryParams.Add("limit", core.ParameterToString(req.Limit, ""))
	}
	if req.Offset > 0 {
		localVarQueryParams.Add("offset", core.ParameterToString(req.Offset, ""))
	}
	if req.BeginDate != "" && req.EndDate != "" {
		localVarQueryParams.Add("begin_date", req.BeginDate)
		localVarQueryParams.Add("end_date", req.EndDate)
	}
	if req.ComplaintedMchid != "" {
		localVarQueryParams.Add("complainted_mchid", req.ComplaintedMchid)
	}
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(SearchComplaintResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryDetail 获取投诉单详情
func (s *ComplaintApiService) QueryDetail(ctx context.Context, complaintId string) (resp *ComplaintDetailResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaints-v2/" + complaintId
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(ComplaintDetailResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// Refund 退款审批
func (s *ComplaintApiService) Refund(ctx context.Context, complaintId string, req RefundHandlerRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaints-v2/" + complaintId + "/update-refund-progress"
	// Setup Body Params
	localVarPostBody = req
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	return result, err
}

// Replay 回复
func (s *ComplaintApiService) Replay(ctx context.Context, complaintId string, req ReplayComplaintRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaints-v2/" + complaintId + "/response"
	// Setup Body Params
	localVarPostBody = req
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	return result, err
}

// Complete 完结订单
func (s *ComplaintApiService) Complete(ctx context.Context, complaintId string, req CompleteRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaints-v2/" + complaintId + "/complete"
	// Setup Body Params
	localVarPostBody = req
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	return result, err
}

// AddNoticeUrl 完结订单
func (s *ComplaintApiService) AddNoticeUrl(ctx context.Context, req AddComplaintNoticeUrl) (resp *QueryCallbackUrlResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaint-notifications"
	// Setup Body Params
	localVarPostBody = req
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(QueryCallbackUrlResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// UpdateNoticeUrl 完结订单
func (s *ComplaintApiService) UpdateNoticeUrl(ctx context.Context, req UpdateComplaintNoticeUrl) (resp *QueryCallbackUrlResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPut
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaint-notifications"
	// Setup Body Params
	localVarPostBody = req
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(QueryCallbackUrlResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryNoticeUrl 完结订单
func (s *ComplaintApiService) QueryNoticeUrl(ctx context.Context) (resp *QueryCallbackUrlResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaint-notifications"
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(QueryCallbackUrlResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DeleteNoticeUrl 完结订单
func (s *ComplaintApiService) DeleteNoticeUrl(ctx context.Context) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaint-notifications"
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	return result, err

}

// QueryNegotiationHistory 查询投诉单协商历史
func (s *ComplaintApiService) QueryNegotiationHistory(ctx context.Context, complaintId string, req QueryNegotiationHistoryRequest) (resp *QueryNegotiationHistoryResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/merchant-service/complaints-v2/" + complaintId + "/negotiation-historys"
	localVarQueryParams = neturl.Values{}
	if req.Limit > 0 {
		localVarQueryParams.Add("limit", core.ParameterToString(req.Limit, ""))
	}
	if req.Offset > 0 {
		localVarQueryParams.Add("offset", core.ParameterToString(req.Offset, ""))
	}
	localVarHTTPContentTypes := []string{"application/json"}
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	resp = new(QueryNegotiationHistoryResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
