package user

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/model"
	"github.com/lihongsheng/payment-sdk/errors"
)

const (
	openAccount    = "/V2/openAccount.fuiou"
	queryAccount   = "/V2/queryAllocateAccount.fuiou"
	deleteAccount  = "/invalidAllocateAccount.fuiou"
	activeAccount  = "/activeAllocateAccount.fuiou"
	updateMobile   = "/modifyProtocol.fuiou"
	updateBankCard = "/modifyAccountInCard.fuiou"
)

type User struct {
	*fuiou.Api
}

func NewUser(conf config.Config) (*User, error) {
	api, err := fuiou.NewApi(conf)
	if err != nil {
		return nil, err
	}
	return &User{
		Api: api,
	}, nil
}

func (u *User) Create(ctx context.Context, req *model.UserCreateRequest) (*model.UserCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostEncryptFrom(ctx, openAccount, req, nil)
	if err != nil {
		return nil, err
	}
	user := &model.UserCreateResponse{}
	err = xml.Unmarshal([]byte(resp.Message), user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return user, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}

func (u *User) Query(ctx context.Context, req *model.UserQueryRequest) (*model.UserQueryResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostEncryptFrom(ctx, queryAccount, req, nil)
	if err != nil {
		return nil, err
	}
	user := &model.UserQueryResponse{}
	err = xml.Unmarshal([]byte(resp.Message), user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return user, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}

func (u *User) Delete(ctx context.Context, req *model.UserDeleteRequest) (*model.UserDeleteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostReqFrom(ctx, deleteAccount, req, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var user = &model.UserDeleteResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return user, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}

func (u *User) RetryActive(ctx context.Context, req *model.UserActiveRequest) (*model.UserCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostReqFrom(ctx, activeAccount, req, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var user = &model.UserCreateResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return user, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}

// ChangeMobile
// 更改手机号
func (u *User) ChangeMobile(ctx context.Context, req *model.UserUpdateMobileRequest) (*model.UserUpdateResponse, error) {
	tReq := model.UserUpdateMobileRequestToMobileAndTransfer(req)
	if err := tReq.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostReqFrom(ctx, updateMobile, tReq, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	body := resp.Body()
	var user = &model.UserUpdateResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}

	if user.RespCode != "0000" {
		return user, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}

// ChangeAllocateScale
// 更改分账比例
func (u *User) ChangeAllocateScale(ctx context.Context, req *model.UserUpdateAllocateScaleRequest) (*model.UserUpdateResponse, error) {
	tReq := model.UserUpdateAllocateScaleRequestToMobileAndTransfer(req)
	if err := tReq.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostReqFrom(ctx, updateMobile, tReq, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var user = &model.UserUpdateResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return nil, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}

// BindBank
// 绑定银行卡
func (u *User) BindBank(ctx context.Context, req *model.UserBindBankAccountRequest) (*model.UserBankAccountResponse, error) {
	tReq := model.UserBindBankAccountRequestToUserBankAccountRequest(req)
	if err := tReq.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostReqFrom(ctx, updateBankCard, tReq, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	fmt.Println(string(body))
	var user = &model.UserBankAccountResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return nil, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}

// UnsetBank
// 解绑银行卡
func (u *User) UnsetBank(ctx context.Context, req *model.UserUnsetBankAccountRequest) (*model.UserBankAccountResponse, error) {
	tReq := model.UserUnsetBankAccountRequestToUserBankAccountRequest(req)
	if err := tReq.Validate(); err != nil {
		return nil, err
	}
	resp, err := u.Client.PostReqFrom(ctx, updateBankCard, tReq, nil)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	var user = &model.UserBankAccountResponse{}
	err = xml.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	if user.RespCode != "0000" {
		return nil, errors.ErrorSystemError(user.RespCode+":"+user.RespDesc, nil)
	}
	return user, err
}
