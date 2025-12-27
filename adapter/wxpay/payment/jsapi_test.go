package payment

import (
	"context"
	"fmt"
	"github.com/singer-stack-lab/payment-sdk/adapter/wxpay/until"
	"time"

	"github.com/singer-stack-lab/payment-sdk/config"
	"github.com/singer-stack-lab/payment-sdk/driver/dto"
	enum "github.com/singer-stack-lab/payment-sdk/enum/payment"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestJsapi_Pay(t *testing.T) {
	cccc := config.Config{
		AppID: "wx6c663032961e5e4b",
		MchID: "1730424164",
		Cert: config.Cert{
			CertPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC9h7iMvdgwoQgZ
jspOMMNptXdR6eyPLtN0DA1AU/laYTSKvA6POYDzvv6XA8fOQblGYpM6rV8ebsGN
O2rCCTllZ8S79S0JD/WoYfGqCU+A/B5qgFM+IxfH1o1yk7s4kdeKm5e8mXR22fyu
n60+Aoxsc2rG1BepxbVB44UD+zJyHd+qPnpfcpe+3llPgz9UBWCKRI7bWg/ai33+
6DnOtsW8d96zk1zZlhD0tvYoda21Pn6D2kC+PctrbRTORm6Ms8B9Exm519T7Qv4+
gTYut5LKODPWA8/yQK2O0usohL0Oa12J5YYyXmFgM/6/xb5eeRBkHKytje7ErjeN
h/DKnKNZAgMBAAECggEBAKchx9xUjRBU8I+ZG01YMvpEy7OwVKru4Ai1l/niE0Ff
9rVgoHYsf0pyYo9JBikcTAWFZ8+WkwqDIKsqejohaDvEBfi5e71CFZ8mv7TyAOn9
adUA1Gc3CwFucc1X+QEpmqjgDC6EI++qyugyZtMH7Ey2erJa1YMglXZE7EdWfGWj
7aVDwQJrpBhFMGg9nIRuclZnzqGVuXubJg8PWy7rnDkIi/Z/QM52jZZW50X1W6mL
sRNy+eloxg6MgXx+KjAypplQMVWqu21gyudktHAXFSCf3P+NdYb0N1RL/JouUgwz
suGGxRQVi7WOzrP9EJgilFhW07uv6X3hd+KzFJSDDwECgYEA464B3Ucg/j3Oc1KX
IN184VQYzzJlPxu2Nx8MwhzekSDU++/E4jvBd65BDgHc+77OwrOO4HNrCGD+iV/S
idqf4bIlGrQlp5Qew+tBb4nhzi+zEsYtwCnJozDACk9PssqP5Hj+pFDvhjy3sMon
cPNSm2+jOihcAqEIB7/4jsOFG/MCgYEA1RrrXP8sSmP2WFINRHFhssl5jAV/5MwQ
8jRjE6CFX9IwVHPPXjSwwOh16tugM2KPFYuHIqWsR3kyp73Vh4Ce97AJu2G1gO2t
wl5g0nVwiRr+ejmbJGNasrp62QZggMrHEKTqriosqOj3m6A7n9KvMKd0bsC2uA5W
kZeBPx+m0oMCgYEAw9GvDM/WUpR58bnA/aVBeNNJmzru1X5SE8qCwJjv28ZvKFgp
76IRXYvjq9ZyZ5rOXartYaIjFkvF4AUoISSFiiobu4HhOOYuJ7c4ymO+cAWacLU+
OB44rECLitJ364BIjep6qHxr5fpmyoizr3O3QrSboLOBn0k8jN3RO4hx/X0CgYAA
7T4KyH1L0YV3utud6ZRQL7oclsWInC6SrxGjOzZ5RTO6mkpTkY0XOauRmuTmdE5E
/LdYujm2kdtbiWLNVQzb7OMN8o3UgrQXvUtUfvg/UGO86lU3Yks5rb/tA68VwEv/
UYhHu504GtNA1QCNYGAsqP3DoYjp4f4UYgFI4f1auwKBgCScG5m93LGgJQlE4XI1
hGtWCs7qMcdvz8a7X84NXuCxBuw4WZlikeLTQENZzny1u9GKWrczEcKFzMiF7u+e
7kfnZzOO1Dph+Ut/8Y0XwVJZDfsZca+dOfDDgJIjZMKqGqJ+67VpVTyguXu7NW7a
tdnfEMZWEZvzWlXYE6j4DVXF
-----END PRIVATE KEY-----`,
			CertificateSerialNumber: "1A6DEA49B006CC47A477C387244985669DB0EA0B",
			PublicKeyID:             "PUB_KEY_ID_0117304241642025110600182326001207",
			PublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9CUWaMY8Fd5HfSVNKJtQ
FjwZ57sruwz4cRK2HCvHUaNsWMdy5S2zfGIwD8Z9YfJESa3AmSV42jb1tFu3tb+W
ew6JMLRHolknZuEokDO0zH49fhnVP/6x2FItGy4/lIBaoPg49LHXxCZxVqshykV7
m3/bd9MGYgoMHSQKKJhYzNCbL2Gp8L7ChtVctReG6eGjcOOk+a4hztWA7e1Qig8z
LsGGnI+rcowMHn9ZTLwYa9taBNotVwa8qqOTCs9s22/6ZCDJHJELMrNrvjRNomjl
C0N/dHAbJuftpoXWSxAxJdI0r2+Vy1VqDglJ9KxCiPf9B87eirOnZgn6YgKJ8a35
iQIDAQAB
-----END PUBLIC KEY-----`,
		},
		APIKey: "hafkjfhkhjewhkjjhfkw135213213131",
		Proxy:  config.Proxy{},
	}
	api, err := NewJsApi(cccc)
	if api == nil {
		fmt.Println(err.Error())
		return
	}
	// 45D3967D4B9491CEBDA222FE8DC2600920E42722
	// 45D3967D4B9491CEBDA222FE8DC2600920E42722
	assert.NoError(t, err)
	ctx := context.Background()
	req := &dto.PayOrder{
		Order: dto.Order{
			OrderNo: "123456789",
			PayAmount: dto.Amount{
				Currency: "CNY",
				Total:    1,
			},
			Subject: "test pay",
		},
		NotifyUrl:      "https://api-cabinet.test.jianxindianzi.com/callback/v1/tenant/payment/5/175870989522000055545",
		PassbackParams: "tpdp_id=5",
		Payer: dto.Payer{
			OpenID: "o7uyb2G1X2nRyhPCwGSt6TFd3sAs",
		},
		RedirectUrl: "",
		TimeExpire:  time.Now().Unix() + 90,
	}
	resp, err := api.Pay(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestStatus(t *testing.T) {
	assert.Equal(t, enum.Status_Status_UNKNOWN, until.PaymentStatus["UNKNOWN"])
}
