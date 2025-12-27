package payscore

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
	nethttp "net/http"
	neturl "net/url"
	"strings"
)

type ScoreApiService services.Service

func (a *ScoreApiService) Prepay(ctx context.Context, req CreateServiceOrderRequest) (resp *CreateServiceOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/payscore/serviceorder"

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CreateServiceOrderResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ScoreApiService) GetServiceOrder(ctx context.Context, req GetServiceOrderRequest) (resp *ServiceOrderEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutOrderNo == nil && req.QueryId == nil {
		return nil, nil, fmt.Errorf("field `OutOrderNo` or QueryId is required and must be specified in QueryOrderByIdRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/payscore/serviceorder"

	// Make sure All Required Params are properly set
	if req.Appid == nil {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in QueryOrderByIdRequest")
	}
	if req.ServiceId == nil {
		return nil, nil, fmt.Errorf("field `ServiceId` is required and must be specified in QueryOrderByIdRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.OutOrderNo != nil && *req.OutOrderNo != "" {
		localVarQueryParams.Add("out_order_no", core.ParameterToString(*req.OutOrderNo, ""))
	}
	if req.QueryId != nil && *req.QueryId != "" {
		localVarQueryParams.Add("query_id", core.ParameterToString(*req.QueryId, ""))
	}
	localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))
	localVarQueryParams.Add("service_id", core.ParameterToString(*req.ServiceId, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract payments.Transaction from Http Response
	resp = new(ServiceOrderEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ScoreApiService) CancelServiceOrder(ctx context.Context, req CancelServiceOrderRequest) (resp *CancelServiceOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	if req.OutOrderNo == nil {
		return nil, nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in CancelServiceOrderRequest")
	}
	if req.Reason == nil {
		return nil, nil, fmt.Errorf("field `Reason` is required and must be specified in CancelServiceOrderRequest")
	}
	// Make sure All Required Params are properly set
	if req.Appid == nil {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in QueryOrderByIdRequest")
	}
	if req.ServiceId == nil {
		return nil, nil, fmt.Errorf("field `ServiceId` is required and must be specified in QueryOrderByIdRequest")
	}
	localVarPath := consts.WechatPayAPIServer + "/v3/payscore/serviceorder/{out_order_no}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"out_order_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutOrderNo, "")), -1)
	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CancelServiceOrderResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *ScoreApiService) Complete(ctx context.Context, req CompleteServiceOrderRequest) (resp *CompleteServiceOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	if req.OutOrderNo == nil {
		return nil, nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in CancelServiceOrderRequest")
	}
	// Make sure All Required Params are properly set
	if req.Appid == nil {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in QueryOrderByIdRequest")
	}
	if req.ServiceId == nil {
		return nil, nil, fmt.Errorf("field `ServiceId` is required and must be specified in QueryOrderByIdRequest")
	}
	localVarPath := consts.WechatPayAPIServer + "/v3/payscore/serviceorder/{out_order_no}/complete"
	localVarPath = strings.Replace(localVarPath, "{"+"out_order_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutOrderNo, "")), -1)
	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CompleteServiceOrderResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
