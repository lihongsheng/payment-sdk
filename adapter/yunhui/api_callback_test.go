package yunhui

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApiCallback_ParseRefund(t *testing.T) {
	c := config.Config{
		AppID: "Y1765007866-5",
		//MchID: "Y17661083936167484",
		MchID:  "Y1765007866",
		APIKey: "oNV6w3sdxHyKDqOEJRkWaoWnfvW0qvxiA8PgznIeaEGXJlUqq1Yyfr4v1j5LmC0gf9FGYeRkg3Yc2BYtVeeSRds7ChPx19qmXpsGAPcHWGuLMUqK4wmbACqDQmUZDUIN",
		Cert: config.Cert{
			CertificateSerialNumber: "",
			CertPrivateKey:          ``,
			PublicKey:               ``,
		},
		Proxy:   config.Proxy{},
		ApiHost: "https://payment.yiyunhuipay.com",
		Version: "",
		Extra:   `{"term_no":"63751"}`,
	}
	pppp, err := NewApiCallback(c)
	assert.NoError(t, err)
	assert.NotNil(t, pppp)
	fmt.Println("-----------------------------------------------")
	req := &http.Request{}
	body := `payOrderId=P2001852669056032770&sign=38046DDDD42E9A05AD59019203F92D90&channelOrderNo=4200002973202512196128651013&reqTime=1766128466156&refundOrderId=R2001913957153021953&createdAt=1766128455805&payAmount=10&appId=Y1765007866-5&mchRefundNo=175879416443300038879&successTime=1766128466000&currency=cny&state=2&mchNo=Y1765007866&refundAmount=1`
	req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(body)))
	ctx := context.Background()
	r, err := pppp.ParseRefund(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, r)
	if err != nil {
		fmt.Println(err.Error())
	}
	if r != nil {
		fmt.Println("-----------------------------------------------")
		by, _ := json.Marshal(r)
		fmt.Println(string(by))
	}
}
