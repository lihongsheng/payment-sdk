package alipay

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/util"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/dto"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"io"
	"net/http"
	"net/url"
)

// curl -X POST 'NOTIFY_URL' \
//--header 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
//--data-urlencode 'charset=UTF-8' \
//--data-urlencode 'biz_content={
//	"out_biz_no":"201806300001",
//	"product_code":"STD_RED_PACKET",
//	"biz_scene":"PERSONAL_PAY",
//	"origin_interface":"alipay.fund.trans.app.pay",
//	"pay_fund_order_id":"20190801110070001506380000251556",
//	"order_id":"20190624110075000006530000014566",
//	"status":"SUCCESS",
//	"sub_status":"SUCCESS",
//	"receiver_open_id":"074a1CcTG1LelxKe4xQC0zgNdId0nxi95b5lsNpazWYoCo5",
//	"receiver_user_id":"2088872048673300",
//	"action_type":"FINISH",
//	"trans_amount":"32.00",
//	"pay_date":"2013-01-01 08:08:08",
//	"refund_date":"2013-01-02 08:08:08",
//	"entrust_order_id":"202007162000000000461",
//	"settle_serial_no":"2023052993044491260542090100400",
//	"sub_order_status":"FAIL",
//	"sub_order_error_code":"MID_ACCOUNT_CARD_INFO_ERROR",
//	"sub_order_fail_reason":"收款方银行卡信息有误"
//}' \
//--data-urlencode 'utc_timestamp=${now}' \
//--data-urlencode 'sign=${sign}' \
//--data-urlencode 'app_id=${appid}' \
//--data-urlencode 'version=1.1' \
//--data-urlencode 'sign_type=RSA2' \
//--data-urlencode 'notify_id=${notify_id}' \
//--data-urlencode 'msg_method=alipay.fund.trans.order.changed'

type EventCallback struct {
	C config.Config
}

func NewEventCallback(config2 config.Config) *EventCallback {
	return &EventCallback{
		C: config2,
	}
}

func (e *EventCallback) Parse(ctx context.Context, req *http.Request) (eventType enum.EventType, resp any, err error) {
	bodyBytes, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return enum.EventTypeUnKnown, nil, err
	}

	sign, signValue, err := GenerateSignString(values)
	if err != nil {
		return enum.EventTypeUnKnown, nil, err
	}
	verifg, err := RsaVerify(sign, signValue, e.C.Cert.PublicKey)
	if err != nil {
		return enum.EventTypeUnKnown, nil, err
	}
	if !verifg {
		return enum.EventTypeUnKnown, nil, errors2.ErrorSignError("签名验证失败："+string(bodyBytes), nil)
	}

	respFrom := e.buildCallbackParams(values)
	switch respFrom.MsgMethod {
	case string(enum.EventTypeOrderChanged):
		respOrigin := &model.AlipayTransferOrderAmountChangeEvent{}
		err = json.Unmarshal([]byte(respFrom.BizContent), respOrigin)
		eventType = enum.EventTypeOrderChanged
		resp = e.AlipayTransferOrderAmountChangeEventToDtoUnitResponse(respOrigin, string(bodyBytes))
		return
	}
	return enum.EventTypeUnKnown, nil, errors2.ErrorNoSupport("not support event type[%s]", respFrom.MsgMethod)
}

func (e *EventCallback) AlipayTransferOrderAmountChangeEventToDtoUnitResponse(req *model.AlipayTransferOrderAmountChangeEvent, origin string) *dto.UintTransferDetailResponse {
	amount, _ := util.AmountToCents(req.TransAmount)
	resp := &dto.UintTransferDetailResponse{
		TransferNo:      req.OutBizNo,
		ThirdTransferNo: req.OrderId,
		Status:          util.GetUnitTransferStatus(req.Status),
		TransferAmount: dto.Amount{
			Total:    int64(amount),
			Currency: "CNY",
		},
		Remark:     req.SubOrderErrorCode,
		FailReason: req.SubOrderFailReason,
		User: dto.User{
			UnionID: req.ReceiverUserId,
			OpenID:  req.ReceiverOpenId,
		},
		OriginResponse: origin,
	}
	return resp
}

func (e *EventCallback) buildCallbackParams(values url.Values) *model.AlipayCallbackEventBody {
	return &model.AlipayCallbackEventBody{
		AppId:         values.Get("app_id"),
		BizContent:    values.Get("biz_content"),
		NotifyId:      values.Get("notify_id"),
		Sign:          values.Get("sign"),
		SignType:      values.Get("sign_type"),
		UtcTimestamp:  values.Get("utc_timestamp"),
		Version:       values.Get("version"),
		MsgMethod:     values.Get("msg_method"),
		MerchantAppId: values.Get("merchant_app_id"),
	}
}
