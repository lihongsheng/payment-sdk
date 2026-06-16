package until

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	sdkerr "github.com/lihongsheng/payment-sdk/errors"
	"github.com/lihongsheng/payment-sdk/log"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
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
			return handlerErr(ctx, apiErr, message)
		}
		log.Error(ctx, "wxpay request failed",
			log.F(log.FieldKeyChannel, "wxpay"),
			log.F(log.FieldKeyError, err.Error()),
		)
		return sdkerr.ErrorSystemError(message, nil).WithCause(err)
	}
	if result == nil {
		log.Error(ctx, "wxpay response empty",
			log.F(log.FieldKeyChannel, "wxpay"),
		)
		return sdkerr.ErrorSystemError("not return result;"+message, nil)
	}
	if result.Response.StatusCode > 300 {
		var errResp = core.APIError{}
		body, err := io.ReadAll(result.Response.Body)
		if err != nil {
			log.Error(ctx, "wxpay read error response failed",
				log.F(log.FieldKeyChannel, "wxpay"),
				log.F(log.FieldKeyError, err.Error()),
			)
			return sdkerr.ErrorSystemError("wxpay is error;err:%s", err.Error()).WithCause(err)
		}
		_ = json.Unmarshal(body, &errResp)
		if message == "" {
			message = errResp.Message
		}
		return handlerErr(ctx, &errResp, message)
	}
	return nil
}

func handlerErr(ctx context.Context, err *core.APIError, message string) error {
	log.Error(ctx, "wxpay api error",
		log.F(log.FieldKeyChannel, "wxpay"),
		log.F("err_code", err.Code),
		log.F("err_message", message),
	)
	if strings.Contains(message, "此商家的收款功能已被限制") {
		return sdkerr.ErrorPaymentLimited("code:%s;message:%s", err.Code, message)
	}
	if strings.Contains(message, "invalid ip") {
		return sdkerr.ErrorIpLimited("IP被限制;code:%s;message:%s", err.Code, message)
	}
	switch err.Code {
	case OUT_TRADE_NO_USED:
		return sdkerr.ErrorDuplicateRequest("order is used; code:%s;message:%s", err.Code, message).WithCause(err)
	case FREQUENCY_LIMITED:
		return sdkerr.ErrorLimited("frequency is limited; code:%s;message:%s", err.Code, message).WithCause(err)
	case PARAM_ERROR:
		return sdkerr.ErrorParamError("param is error; code:%s;message:%s", err.Code, message).WithCause(err)
	case INVALID_REQUEST:
		return sdkerr.ErrorInvalidRequest("invalid request; code:%s;message:%s", err.Code, message).WithCause(err)
	case SIGN_ERROR:
		return sdkerr.ErrorSignError("sign is error; code:%s;message:%s", err.Code, message).WithCause(err)
	case APPID_MCHID_NOT_MATCH:
		return sdkerr.ErrorAppidMchidNotMatch("appid or mchid is not match; code:%s;message:%s", err.Code, message).WithCause(err)
	case MCH_NOT_EXISTS:
		return sdkerr.ErrorMchNotExists("mch is not exists; code:%s;message:%s", err.Code, message).WithCause(err)
	case SYSTEM_ERROR:
		return sdkerr.ErrorRetrySystemError("system error;please retry; code:%s;message:%s", err.Code, message).WithCause(err)
	}

	return sdkerr.ErrorSystemError("code:%s;message:%s", err.Code, message).WithCause(errors.New(fmt.Sprintf("code:%s;message:%s", err.Code, err.Message)))
}
