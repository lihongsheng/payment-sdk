package provider

import (
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay"
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay/client/complaints"
	"github.com/singer-stack-lab/payment-sdk/config"
)

func NewComplaint(c config.Config) (*complaints.ComplaintApiService, error) {
	api, err := wxpay.InitClient(c)
	if err != nil {
		return nil, err
	}
	client := complaints.ComplaintApiService{Client: api.Client}
	return &client, nil
}
