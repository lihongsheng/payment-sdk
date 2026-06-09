# payment-sdk

Go 语言统一支付网关 SDK — 支付宝 / 微信 / 富友 / 拉卡拉

## 文档

- 📖 [项目文档](./docs/README.md) — 概述、快速开始、架构设计、配置说明
- 🛠️ [新增支付渠道开发指南](./docs/develop.md) — 详细的渠道接入教程与完整代码示例

## 快速开始

```go
import (
    payment_sdk "github.com/lihongsheng/payment-sdk"
    "github.com/lihongsheng/payment-sdk/config"
    "github.com/lihongsheng/payment-sdk/enum/channel"
    "github.com/lihongsheng/payment-sdk/enum/payment"
)

opts := []config.Option{
    config.WithPaymentProduct(payment.PaymentProduct_JSAPI),
    config.WithWxConfig(&wxConfig.Config{ /* ... */ }),
}

pay, err := payment_sdk.Payment(channel.Channel_Wechat, opts...)
resp, err := pay.Pay(ctx, &dto.PayOrder{ /* ... */ })
```

详见 [文档](./docs/README.md)。

## License

See [LICENSE](./LICENSE) for details.
