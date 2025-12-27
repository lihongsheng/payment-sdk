package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/client"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/model"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/errors"
)

type User struct {
	Client *client.Client
	conf   config.Config
}

func NewUser(conf config.Config) (*User, error) {
	newClient, err := client.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &User{
		Client: newClient,
		conf:   conf,
	}, nil
}

func (u *User) AuthToken(ctx context.Context, req *model.UserAuthRequest) (*model.UserTokenResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqParam := map[string]string{
		"grant_type": string(req.GrantType),
		"code":       req.Code,
	}
	if req.RefreshToken != "" {
		reqParam["refresh_token"] = req.RefreshToken
	}
	commonParams := u.Client.GetCommonRequestParams()
	commonParams[enum.COMMON_PARAM_METHOD_NAME] = enum.ALIPAY_SYSTEM_OAUTH_TOKEN
	resp, err := u.Client.DoPost(commonParams, reqParam, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var response model.AuthTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	fmt.Println(response.ErrorResponse)
	if response.ErrorResponse != nil {
		return nil, errors.ErrorSystemError(response.ErrorResponse.SubMsg+";"+string(body), nil)
	}
	respTrue := false
	if response.AlipayOpenAuthTokenAppResponse.OpenId != "" {
		respTrue = true
	}
	if !respTrue {
		return nil, errors.ErrorSystemError("not return OpenId;"+string(body), nil)
	}
	return &response.AlipayOpenAuthTokenAppResponse, nil
}
