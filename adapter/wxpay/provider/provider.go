package provider

import (
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client"
	"github.com/lihongsheng/payment-sdk/adapter/wxpay/client/complaints"
	"github.com/lihongsheng/payment-sdk/config"
)

func NewComplaint(c config.Config) (*complaints.ComplaintApiService, error) {
	api, err := client.InitClient(c)
	if err != nil {
		return nil, err
	}
	client := complaints.ComplaintApiService{Client: api.Client}
	return &client, nil
}
