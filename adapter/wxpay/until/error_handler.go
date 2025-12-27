package until

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	errors2 "github.com/singer-stack-lab/payment-sdk/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/zeromicro/go-zero/core/logc"
)

// 微信支付错误处理
// {
//  "code" : "SIGN_ERROR",
//  "detail" : {
//    "detail" : {
//      "issue" : "sign not match"
//    },
//    "field" : "signature",
//    "location" : "authorization",
//    "sign_information" : {
//      "method" : "POST",
//      "sign_message_length" : 349,
//      "truncated_sign_message" : "POST\n/v3/pay/transactions/jsapi\n1756951532\nIXlHDnWguGbz7uMOL7xdPibFoV342jP3\n{\"amount\n",
//      "url" : "/v3/pay/transactions/jsapi"
//    }
//  },
//  "message" : "签名错误"
//}

const (
	// 订单号已使用
	OUT_TRADE_NO_USED     = "OUT_TRADE_NO_USED"
	FREQUENCY_LIMITED     = "REQUENCY_LIMITED"
	PARAM_ERROR           = "PARAM_ERROR"
	INVALID_REQUEST       = "INVALID_REQUEST"
	SIGN_ERROR            = "SIGN_ERROR"
	APPID_MCHID_NOT_MATCH = "APPID_MCHID_NOT_MATCH"
	MCH_NOT_EXISTS        = "MCH_NOT_EXISTS"
	SYSTEM_ERROR          = "SYSTEM_ERROR"
)

func ErrorHandler(ctx context.Context, result *core.APIResult, err error, message string) error {
	if err != nil {
		var apiErr *core.APIError
		if errors.As(err, &apiErr) {
			if message == "" {
				message = apiErr.Message
			}
			return handlerErr(apiErr, message)
		}
		return errors2.ErrorSystemError(message, nil).WithCause(err)
	}
	if result == nil {
		return errors2.ErrorSystemError("not return result;"+message, nil)
	}
	if result.Response.StatusCode > 300 {
		var errResp = core.APIError{}
		body, err := io.ReadAll(result.Response.Body)
		if err != nil {
			logc.Debugw(ctx, "wxpay-is-error", logc.Field("Response", result))
			return errors2.ErrorSystemError("wxpay is error;err:%s", err.Error()).WithCause(err)
		}
		_ = json.Unmarshal(body, &errResp)
		if message == "" {
			message = errResp.Message
		}
		return handlerErr(&errResp, message)
	}
	return nil
}

func handlerErr(err *core.APIError, message string) error {
	switch err.Code {
	case OUT_TRADE_NO_USED:
		return errors2.ErrorDuplicateRequest("order is used;"+message, nil).WithCause(err)
	case FREQUENCY_LIMITED:
		return errors2.ErrorLimited("frequency is limited;"+message, nil).WithCause(err)
	case PARAM_ERROR:
		return errors2.ErrorParamError("param is error;"+message, nil).WithCause(err)
	case INVALID_REQUEST:
		return errors2.ErrorInvalidRequest("invalid request;"+message, nil).WithCause(err)
	case SIGN_ERROR:
		return errors2.ErrorSignError("sign is error;"+message, nil).WithCause(err)
	case APPID_MCHID_NOT_MATCH:
		return errors2.ErrorAppidMchidNotMatch("appid or mchid is not match;"+message, nil).WithCause(err)
	case MCH_NOT_EXISTS:
		return errors2.ErrorMchNotExists("mch is not exists;"+message, nil).WithCause(err)
	case SYSTEM_ERROR:
		return errors2.ErrorRetrySystemError("system error;please retry;"+message, nil).WithCause(err)
	}

	return errors2.ErrorSystemError(message, nil).WithCause(errors.New(fmt.Sprintf("code:%s;message:%s", err.Code, err.Message)))
}
