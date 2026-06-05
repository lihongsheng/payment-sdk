package client

import (
	"github.com/lihongsheng/payment-sdk/adapter/alipay/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetResponseSignContent(t *testing.T) {
	c, err := NewClient(config.Config{}, nil)
	assert.NoError(t, err)

	s, err := c.GetResponseSignContent("{\"alipay_trade_create_response\":{\"code\":\"10000\",\"msg\":\"Success\",\"out_trade_no\":\"20150423001001\",\"trade_no\":\"2015042321001004720200028594\"},\"sign\":\"ERITJKEIJKJHKKKKKKKHJEREEEEEEEEEEE\"}", "alipay.trade.create")
	t.Log(s)
	//if err != nil {
	//	t.Log(err.Error())
	//}
	//t.Log(err)
	assert.NoError(t, err)
	t.Log(s)
}
