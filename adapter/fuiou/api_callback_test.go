package fuiou

//
//func TestCallbackPaymentParse(t *testing.T) {
//	bbb := `{"curr_type":"CNY","full_sign":"c30191f7984fa21082186caae27f17ff","mchnt_cd":"0008710FA369999","mchnt_order_no":"19856176423592652800749889","order_amt":"2","order_type":"JSAPI","random_str":"SR5U8TH2TZRW1Z8LUV81X785ZZYBQKEZ","reserved_addn_inf":"tpdp_id=33","reserved_bank_type":"CITIC_DEBIT","reserved_buyer_logon_id":"","reserved_channel_order_id":"19856176423592652800749889","reserved_coupon_fee":"0","reserved_fund_bill_list":"","reserved_fund_state":"","reserved_fy_settle_dt":"20251127","reserved_fy_trace_no":"230705503286","reserved_is_credit":"0","reserved_promotion_detail":"","reserved_settlement_amt":"2","result_code":"000000","result_msg":"SUCCESS","settle_order_amt":"2","sign":"ab5b652fb2243917639c65252c8872ff","term_id":"","transaction_id":"4200002918202511270538002877","txn_fin_ts":"20251127173212","user_id":"owDb32P8qIDRcPrjkTRmA5pXUmp8"}`
//
//	c := config.Config{
//		AppID:   "",
//		MchID:   "0008710FA369999",
//		APIKey:  "71a01b40caa811f09a88a57d8f5260d8",
//		Cert:    config.Cert{},
//		Proxy:   config.Proxy{},
//		ApiHost: "https://aipay-cloud.fuioupay.com",
//		Version: "1.0",
//		Extra:   `{"wx_app_id": "wxc233b52c912ad8f7", "order_prefix": "19856", "wx_app_secret": "1660bd2be85e638319381da7867ce294"}`,
//	}
//	ctx := context.Background()
//	r := &http.Request{}
//	r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(bbb)))
//	client, err := NewAPICallback(c)
//	assert.NoError(t, err)
//	rr, err := client.CallbackPaymentParse(ctx, r)
//
//	assert.NoError(t, err)
//	fmt.Println(err)
//	if rr != nil {
//		fmt.Println("---------------------------------")
//		b, _ := json.Marshal(rr)
//		fmt.Println(string(b))
//	}
//}
