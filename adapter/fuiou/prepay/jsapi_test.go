package prepay

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPay_Pay(t *testing.T) {

}

func TestPay_Query(t *testing.T) {

	return
}

func TestPay_Close(t *testing.T) {

	return
}

func TestPay_Query2(t *testing.T) {

	return
}

func TestPay_QueryBody(t *testing.T) {
	body := `{"addn_inf":"tpdp_id=33","buyer_id":"owDb32P8qIDRcPrjkTRmA5pXUmp8","mchnt_cd":"0008710FA369999","mchnt_order_no":"19856176423592652800749889","order_amt":"2","order_type":"JSAPI","random_str":"20251127185516791790","reserved_bank_type":"CITIC_DEBIT","reserved_buyer_logon_id":"","reserved_channel_order_id":"19856176423592652800749889","reserved_coupon_fee":"","reserved_fund_bill_list":"","reserved_fund_state":"","reserved_fy_settle_dt":"20251127","reserved_fy_trace_no":"230705503286","reserved_is_credit":"0","reserved_promotion_detail":"","reserved_relate_order_no":"","reserved_relate_trace_no":"","reserved_settlement_amt":"2","reserved_txn_fin_ts":"20251127173212","reserved_unfreeze_time":"","result_code":"000000","result_msg":"SUCCESS","sign":"d2d694499037277731fe1352e21c9958","term_id":"6328","trans_stat":"SUCCESS","transaction_id":"4200002918202511270538002877"}`
	resp := &OrderDetail{}
	err := json.Unmarshal([]byte(body), resp)
	assert.NoError(t, err)
}
