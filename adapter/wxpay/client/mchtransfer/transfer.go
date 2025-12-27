package mchtransfer

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

type TransferApiService services.Service

func (s *TransferApiService) TransferBill(ctx context.Context, req TransferBillRequest) (resp *TransferBillResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/fund-app/mch-transfer/transfer-bills"
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
	resp = new(TransferBillResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (s *TransferApiService) Cancel(ctx context.Context, req TransferBillCancelRequest) (resp *TransferBillCancelResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	if req.OutBillNo == "" {
		return nil, result, fmt.Errorf("out_bill_no must not empty")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/{out_bill_no}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"out_bill_no"+"}", neturl.PathEscape(core.ParameterToString(req.OutBillNo, "")), -1)
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
	resp = new(TransferBillCancelResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (s *TransferApiService) GetTransferBillByOutBillNo(ctx context.Context, req GetTransferBilByOutBillNORequest) (resp *TransferBillDetailResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/{out_bill_no}"
	localVarPath = strings.Replace(localVarPath, "{"+"out_bill_no"+"}", neturl.PathEscape(core.ParameterToString(req.OutBillNo, "")), -1)
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
	resp = new(TransferBillDetailResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (s *TransferApiService) GetTransferBillByNo(ctx context.Context, req GetTransferBilByNORequest) (resp *TransferBillDetailResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/fund-app/mch-transfer/transfer-bills/transfer-bill-no/{transfer_bill_no}"
	localVarPath = strings.Replace(localVarPath, "{"+"transfer_bill_no"+"}", neturl.PathEscape(core.ParameterToString(req.TransferBillNo, "")), -1)
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
	resp = new(TransferBillDetailResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
