package refund

import (
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/until"
	enum "github.com/lihongsheng/payment-sdk/enum/payment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsapi_Pay(t *testing.T) {

}

func TestStatus(t *testing.T) {
	assert.Equal(t, enum.Status_Status_UNKNOWN, until.PaymentStatus["UNKNOWN"])
}
