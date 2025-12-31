package fuiou

import (
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/client"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/config"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"strings"
)

type Api struct {
	Client *client.Client
	C      config.Config
	Env    enum.Env
	Extra  enum.Extra
}

func NewApi(c config.Config) (*Api, error) {
	c.ApiHost = strings.TrimRight(c.ApiHost, "/")
	// 转账client
	newClient, err := client.NewClient(c)
	return &Api{
		Client: newClient,
		C:      c,
	}, err
}
