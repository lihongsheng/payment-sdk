package model

import (
	"encoding/json"
	"github.com/lihongsheng/payment-sdk/adapter/alipay/enum"
	"github.com/lihongsheng/payment-sdk/errors"
	"time"
)

type UserAuthRequest struct {
	GrantType    enum.AuthType `json:"grant_type,omitempty"`
	Code         string        `json:"code,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
}

func (u UserAuthRequest) Validate() error {
	if u.GrantType == enum.AUTH_TYPE_AUTHORIZATION_CODE && u.Code == "" {
		return errors.ErrorParamError("grant_type is AUTHORIZATION_CODE, app auth token must not empty")
	}
	return nil
}

type UserTokenResponse struct {
	Code         string      `json:"code"`
	Msg          string      `json:"msg"`
	SubCode      string      `json:"sub_code"`
	SubMsg       string      `json:"sub_msg"`
	AccessToken  string      `json:"access_token"`
	AuthStart    string      `json:"auth_start"`
	ExpiresIn    json.Number `json:"expires_in"`
	ReExpiresIn  json.Number `json:"re_expires_in"`
	RefreshToken string      `json:"refresh_token"`
	OpenId       string      `json:"open_id"`
}

func (u *UserTokenResponse) GetExpiresIn() time.Time {
	t, _ := time.Parse(time.DateTime, u.AuthStart)
	if u.ExpiresIn.String() != "" {
		expiresIn, _ := u.ExpiresIn.Int64()
		return t.Add(time.Duration(expiresIn) * time.Second)
	}
	return time.Time{}
}

func (u *UserTokenResponse) GetReExpiresIn() time.Time {
	t, _ := time.Parse(time.DateTime, u.AuthStart)
	if u.ReExpiresIn.String() != "" {
		expiresIn, _ := u.ReExpiresIn.Int64()
		return t.Add(time.Duration(expiresIn) * time.Second)
	}
	return time.Time{}
}

type AuthTokenResponse struct {
	AlipayOpenAuthTokenAppResponse UserTokenResponse `json:"alipay_system_oauth_token_response"`
	ErrorResponse                  *ErrorResponse    `json:"error_response"`
	Sign                           string            `json:"sign"`
}

//type T struct {
//	AlipaySystemOauthTokenResponse struct {
//		AccessToken  string `json:"access_token"`
//		AuthStart    string `json:"auth_start"`
//		ExpiresIn    int    `json:"expires_in"`
//		ReExpiresIn  int    `json:"re_expires_in"`
//		RefreshToken string `json:"refresh_token"`
//		OpenId       string `json:"open_id"`
//	} `json:"alipay_system_oauth_token_response"`
//	Sign string `json:"sign"`
//}
