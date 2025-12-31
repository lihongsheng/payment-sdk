package fuiou

import (
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/client"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	enum2 "github.com/lihongsheng/payment-sdk/adapter/lakala/enum"
	"strings"
)

type Api struct {
	Client *client.Client
	C      config.Config
	Env    enum.Env
	Extra  enum.Extra
}

func NewApi(c config.Config) (*Api, error) {
	if c.ApiHost == "" {
		c.ApiHost = enum2.ApiHost
	} else {
		c.ApiHost = strings.TrimRight(c.ApiHost, "/")
	}
	if c.Version == "" {
		c.Version = enum2.Version
	}
	// 转账client
	newClient, err := client.NewClient(c)
	return &Api{
		Client: newClient,
		C:      c,
	}, err
}
