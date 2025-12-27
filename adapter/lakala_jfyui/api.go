package lakala_jfyui

import (
	"github.com/singer-stack-lab/payment-sdk/adapter/lakala_jfyui/client"
	"github.com/singer-stack-lab/payment-sdk/config"
	"strings"
)

type Api struct {
	Client *client.Client
	C      config.Config
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
