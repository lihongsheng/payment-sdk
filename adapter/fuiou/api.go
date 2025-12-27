package fuiou

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/client"
	"github.com/lihongsheng/payment-sdk/adapter/fuiou/enum"
	"github.com/lihongsheng/payment-sdk/config"
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
	extra := enum.Extra{}
	if c.Extra != "" {
		err := json.Unmarshal([]byte(c.Extra), &extra)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("not Unmarshal extra:[%s]", err.Error()))
		}
	}
	return &Api{
		Client: newClient,
		C:      c,
		Extra:  extra,
	}, err
}
