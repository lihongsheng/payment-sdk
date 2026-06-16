<p align="center">
  <h1 align="center">Payment SDK</h1>
  <p align="center">Go 语言统一支付网关 SDK — 支付宝 / 微信 / 富友 / 拉卡拉</p>
</p>

---

## 概述

Payment SDK 是一款基于 Go 语言的支付软件开发工具包，为**支付宝、微信支付、富友、拉卡拉**等多种支付渠道提供**统一调用接口**。SDK 采用与 Go `database/sql` 相似的**驱动注册模式**，各类支付渠道均实现统一接口，接入新渠道只需实现接口并注册驱动即可。

### 核心特性

- **统一接口** — 所有渠道共享 `Pay` / `Refund` / `Transfer` 等接口，切换渠道无需修改业务代码
- **驱动注册** — 与 `database/sql` 一致的驱动注册模式，新渠道即插即用
- **函数式选项** — 通过 `config.Option` 灵活配置渠道参数
- **动态配置** — 内置 JSON Schema 配置描述，支持前端动态渲染渠道配置表单
- **完整能力** — 支付、退款、转账、分账、回调全链路覆盖

---

## 快速开始

### 安装

```bash
go get github.com/lihongsheng/payment-sdk
```

### 发起支付

```go
package main

import (
    "context"
    "fmt"
    "time"

    payment_sdk "github.com/lihongsheng/payment-sdk"
    "github.com/lihongsheng/payment-sdk/config"
    "github.com/lihongsheng/payment-sdk/driver/dto"
    "github.com/lihongsheng/payment-sdk/enum/channel"
    "github.com/lihongsheng/payment-sdk/enum/payment"
    wxCf "github.com/lihongsheng/payment-sdk/adapter/wxpay/config"
)

func main() {
    // 1. 配置渠道参数（函数式选项模式）
    opts := []config.Option{
        config.WithPaymentProduct(payment.PaymentProduct_JSAPI),
        config.WithWxConfig(&wxCf.Config{
            Merchant: wxCf.Merchant{
                AppID:     "wx8888888888888888",
                MchID:     "1900000109",
                AppSecret: "your_app_secret",
            },
            Cert: wxCf.Cert{
                APISecret:        "your_api_v3_secret",
                RsaPrivate:       "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----",
                RsaPrivateNumber: "your_cert_serial_no",
            },
        }),
    }

    // 2. 获取渠道支付接口
    pay, err := payment_sdk.Payment(channel.Channel_Wechat, opts...)
    if err != nil {
        panic(err)
    }

    // 3. 发起支付
    ctx := context.Background()
    resp, err := pay.Pay(ctx, &dto.PayOrder{
        Order: dto.Order{
            OrderNo:   fmt.Sprintf("%d", time.Now().Unix()),
            PayAmount: dto.Amount{Total: 100, Currency: "CNY"}, // 1 元 = 100 分
            Subject:   "测试订单",
        },
        Payer: dto.Payer{
            OpenID: "user_openid_here",
        },
        NotifyUrl: "https://your-domain.com/callback/wechat",
    })
    if err != nil {
        panic(err)
    }

    // 4. 返回前端唤起支付所需参数
    fmt.Printf("Action: %+v\n", resp.Action)
}
```

### 发起退款

```go
refund, err := payment_sdk.Refund(channel.Channel_Wechat, opts...)
if err != nil {
    panic(err)
}

detail, err := refund.Refund(ctx, &dto.RefundRequest{
    RefundNo:    "REFUND_001",
    TradeNo:     "wechat_trade_no",
    OrderNo:     "your_order_no",
    Reason:      "用户申请退款",
    NotifyUrl:   "https://your-domain.com/callback/refund",
    Amount:      dto.Amount{Total: 100, Currency: "CNY"},
    OrderAmount: dto.Amount{Total: 100, Currency: "CNY"},
})
```

### 处理回调

```go
// 支付回调
callbackDetail, err := pay.Callback(ctx, httpReq)
fmt.Printf("OrderNo: %s, Status: %s\n", callbackDetail.OrderNo, callbackDetail.Status)

// 退款回调
refundDetail, err := refund.Callback(ctx, httpReq)
fmt.Printf("RefundNo: %s, Status: %s\n", refundDetail.RefundNo, refundDetail.Status)
```

### 查询订单

```go
// 支付查询
payDetail, err := pay.Query(ctx, dto.Query{
    OrderNo: "your_order_no",
    TradeNo: "wechat_trade_no",
})

// 退款查询
refundDetail, err := refund.Query(ctx, dto.RefundQuery{
    RefundNo: "REFUND_001",
    OrderNo:  "your_order_no",
})
```

---

## 架构设计

### 驱动注册模式

SDK 采用与 Go `database/sql` 相似的驱动注册模式：

```
┌─────────────────────────────────────────────────────┐
│                    用户代码                          │
│  payment_sdk.Payment(channel.Channel_Wechat, opts…) │
└───────────────────────┬─────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────┐
│              privoder.go (入口层)                    │
│  Payment() → driver.Payment(channel, config)        │
│  Refund()  → driver.Refund(channel, config)         │
└───────────────────────┬─────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────┐
│           driver/register.go (注册中心)              │
│  paymentDrivers map[Channel]PaymentDriver           │
│  refundDrivers  map[Channel]RefundDriver            │
│  PaymentRegister() / RefundRegister()               │
└───────────────────────┬─────────────────────────────┘
                        │
          ┌─────────────┼─────────────┐
          ▼             ▼             ▼
   ┌──────────┐  ┌──────────┐  ┌──────────┐
   │ 微信驱动 │  │ 支付宝驱动 │  │ 富友驱动 │  …
   │  wxpay   │  │  alipay  │  │  fuiou   │
   └──────────┘  └──────────┘  └──────────┘
```

**工作流程：**

1. **驱动注册** — 每个支付渠道通过 `init()` 函数调用 `driver.PaymentRegister()` / `driver.RefundRegister()` 完成自身注册
2. **导入触发** — `privoder.go` 中的空导入 `_ "adapter/xxx/driver"` 触发 `init()` 函数
3. **工厂函数** — `Payment()` 和 `Refund()` 根据渠道名称从注册表查找驱动，调用 `Open()` 返回具体实现

### 核心接口

#### PaymentDriver（支付驱动接口）

```go
type PaymentDriver interface {
    Open(c config.Config) (Pay, error)                       // 实例化支付接口
    GetConfigOptions() *ChannelOption                        // 获取渠道配置描述（JSON Schema）
    GetSupportProduct() []PaymentMethod                      // 获取支持的支付方式
    IsSupportPayment(payment.PaymentProduct, Device) bool    // 判断是否支持某种支付产品+设备组合
    CallbackResponse() string                                // 回调成功后返回给渠道的响应内容
}
```

#### Pay（支付接口）

```go
type Pay interface {
    Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)            // 发起支付
    Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error)                // 查询订单
    Close(ctx context.Context, req dto.CloseQuery) error                             // 关闭订单
    Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error)       // 完成支付（付款码场景）
    Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) // 支付回调
}
```

#### RefundDriver / Refund（退款接口）

```go
type RefundDriver interface {
    Open(c config.Config) (Refund, error)
    CallbackResponse() string
}

type Refund interface {
    Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error)        // 发起退款
    Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error)            // 退款查询
    Callback(ctx context.Context, req *http.Request) (*dto.CallbackRefundDetail, error)   // 退款回调
    IsSupportCallback() bool                                                              // 是否支持退款回调
}
```

---

## 项目结构

```
payment-sdk/
├── privoder.go              # 主入口，导出 Payment() / Refund() 等函数
├── privoder_test.go         # 使用示例与测试
├── config/
│   ├── conf.go              # Config 结构体 & 函数式选项（WithXxx）
│   ├── params/
│   │   ├── options.go       # JSON Schema 类型定义（Schema / UISchema）
│   │   └── SchemaNode.vue   # 前端动态表单渲染组件
│   └── proxy/
│       └── proxy.go         # 代理配置
├── driver/
│   ├── register.go          # 驱动注册中心（PaymentRegister / RefundRegister）
│   ├── iface/
│   │   ├── pay.go           # Pay / PaymentDriver 接口
│   │   ├── refund.go        # Refund / RefundDriver 接口
│   │   ├── unit_transfer.go # 单次转账接口
│   │   └── validate.go      # Validate 接口
│   └── dto/
│       ├── pay.go           # 支付请求/响应 DTO
│       ├── refund.go        # 退款请求/响应 DTO
│       ├── callback.go      # 回调详情 DTO
│       ├── unit_transfer.go # 单次转账 DTO
│       ├── batch_transfer.go# 批量转账 DTO
│       └── order_share.go   # 分账 DTO
├── adapter/                 # 各渠道适配器实现
│   ├── wxpay/               #   微信支付
│   │   ├── driver/          #     驱动注册 + 配置 Schema
│   │   ├── config/          #     渠道配置结构体 & 校验
│   │   ├── client/          #     HTTP 客户端 & 签名
│   │   ├── payment/         #     支付实现（JSAPI / H5 / Qrcode / APP / LITE）
│   │   └── refund/          #     退款实现
│   ├── alipay/              #   支付宝
│   ├── fuiou/               #   富友（聚合支付）
│   └── lakala/              #   拉卡拉（聚合支付）
├── enum/                    # 枚举类型（Proto 定义，make enum 生成）
│   ├── channel/             #   Channel — 支付渠道
│   ├── payment/             #   Payment / PaymentProduct / Status
│   ├── refund/              #   RefundChannel / RefundStatus
│   ├── transfer/            #   TransferStatus
│   ├── order_share/         #   ShareStatus
│   ├── action/              #   Action 类型
│   ├── event/               #   事件类型
│   └── const.go             #   Device / System 常量
├── errors/                  # 错误码（Proto 定义，make error 生成）
├── log/                     # 日志组件
├── tools/                   # 工具函数
└── third_party/             # Proto 第三方依赖
```

---

## 支持的渠道与支付方式

| 渠道 | 标识 | JSAPI | 小程序 | H5 | 二维码 | APP | 退款 | 转账 | 分账 |
|------|------|:-----:|:------:|:--:|:------:|:---:|:----:|:----:|:----:|
| 微信支付 | `Channel_Wechat` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 支付宝 | `Channel_Alipay` | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 富友 | `Channel_Fuiou` | ✅ | ✅ | — | — | — | ✅ | ✅ | ✅ |
| 拉卡拉 | `Channel_Lakala` | ✅ | ✅ | — | — | — | ✅ | — | — |

> 注：富友和拉卡拉为聚合支付渠道，一个渠道同时支持微信和支付宝两种支付方式。

---

## 配置说明

### 函数式选项

SDK 使用函数式选项模式配置参数，所有选项定义在 `config/conf.go`：

| 选项函数 | 说明 |
|---------|------|
| `WithPayment(payment.Payment)` | 设置支付方式（微信/支付宝） |
| `WithPaymentProduct(payment.PaymentProduct)` | 设置支付产品（JSAPI / H5 / Qrcode / APP / LITE） |
| `WithWxConfig(*wxpay.Config)` | 微信渠道配置 |
| `WithAlipayConfig(*alipay.Config)` | 支付宝渠道配置 |
| `WithFuiouConfig(*fuiou.Config)` | 富友渠道配置 |
| `WithLakalaConfig(*lakala.Config)` | 拉卡拉渠道配置 |
| `WithConfig(string)` | JSON 字符串配置（通用，可作为各渠道 Config 的替代） |
| `WithProxy(*proxy.Proxy)` | HTTP 代理配置 |

### 渠道配置结构体

#### 微信支付

```go
type Config struct {
    Merchant Merchant  // AppID, MchID, AppSecret
    Cert     Cert      // APISecret, RsaPrivate, RsaPrivateNumber, RsaPublic, RsaPublicNumber
    Service  Service   // ScoreServiceID
}
```

#### 支付宝

```go
type Config struct {
    Merchant Merchant  // AppID, AppAuthToken
    Cert     Cert      // RsaPrivate, RsaPublic, RsaAppCertSN, RsaRootCertSN, RsaRootCrt, RsaAppCrt
}
```

#### 富友

```go
type Config struct {
    Merchant Merchant  // MchID, APISecret, OrderPrefix
    Cert     Cert      // RsaPrivate, RsaPublic
    API      API       // ApiHost, Version
    Wechat   Wechat    // AppID, AppSecret（获取微信 OpenID）
    Alipay   Alipay    // AppID, RsaPrivate, RsaRootCrt
}
```

#### 拉卡拉

```go
type Config struct {
    Merchant Merchant  // AppID, MchID, TermNO
    Cert     Cert      // RsaPrivate, RsaPrivateNumber, RsaPublic
    API      API       // ApiHost
}
```

### 动态配置（JSON Schema）

每个渠道驱动通过 `GetConfigOptions()` 返回 **JSON Schema 描述**，前端可根据 Schema 动态渲染配置表单。SDK 内置了 Vue 组件 `config/params/SchemaNode.vue` 用于递归渲染 Schema 节点。

Schema 支持以下特性：

- **字段类型** — `string` / `integer` / `number` / `boolean` / `array` / `object`
- **校验规则** — `required`、`pattern`（正则）、`format`
- **UI 控件** — 通过 `x-ui.input_type` 指定（`text` / `password` / `textarea` / `select` / `radio` / `switch` 等）
- **校验类型** — 通过 `x-ui.validate_type` 指定后端校验（`RsaPublic` / `RsaPrivate` / `RsaCert` / `Email` / `Phone` / `Url` / `Reg` 等）
- **字段排序** — 通过 `x-order` 控制字段展示顺序（`properties` 是 map，无序）
- **示例值** — 通过 `examples` 数组提供占位符提示

---

## 枚举类型

枚举类型在 `enum/` 目录下以 Proto 消息格式定义，通过 `make enum` 命令生成 Go 代码。

### 支付渠道（Channel）

| 枚举值 | 数值 | 说明 |
|-------|:----:|------|
| `Channel_UNKNOWN` | 0 | 未知 |
| `Alipay` | 1 | 支付宝 |
| `Wechat` | 2 | 微信 |
| `Fuiou` | 3 | 富友 |
| `Lakala` | 4 | 拉卡拉 |

### 支付产品（PaymentProduct）

| 枚举值 | 数值 | 说明 |
|-------|:----:|------|
| `PaymentMethod_UNKNOWN` | 0 | 未知 |
| `H5` | 1 | 手机 H5 支付 |
| `APP` | 2 | APP 支付 |
| `JSAPI` | 3 | 公众号/小程序支付 |
| `Face` | 4 | 人脸支付 |
| `Qrcode` | 5 | 扫码支付（商户生成二维码） |
| `Scan` | 6 | 付款码支付（用户出示付款码） |
| `LITE` | 7 | 小程序支付 |
| `AFTER_PAY` | 8 | 先享后付 |
| `DEPOSIT_PAY` | 9 | 免押金付款 |
| `PC` | 10 | PC 端支付 |

### 设备类型（Device）

| 常量 | 说明 |
|------|------|
| `Device_PC` | PC 网页 |
| `Device_H5` | 手机网页 |
| `Device_Wechat` | 微信内置浏览器 |
| `Device_Alipay` | 支付宝内置浏览器 |
| `Device_Wechat_Lite` | 微信小程序 |
| `Device_Alipay_Lite` | 支付宝小程序 |
| `Device_APP` | APP |

### 订单状态（Status）

| 枚举值 | 数值 | 说明 |
|-------|:----:|------|
| `Created` | 1 | 初始化 |
| `Pending` | 2 | 预支付，等待用户完成支付 |
| `Success` | 3 | 支付成功 |
| `Failed` | 4 | 支付失败 |
| `Cancel` | 5 | 用户取消支付 |
| `Close` | 6 | 超时关闭 |
| `Refund` | 7 | 已退款 |
| `WaitConfirm` | 8 | 待确认 |
| `WaitPay` | 9 | 待支付 |
| `TimeOut` | 10 | 支付超时 |
| `TempFailed` | 11 | 临时失败（可重试） |

---

## 开发命令

```bash
# 构建
go build ./...

# 运行测试
go test ./...
go test -v ./...                              # 详细输出
go test -run TestPayment_WxpayJSAPI ./...     # 运行单个测试
go test ./adapter/alipay/payment/...          # 运行指定包测试

# 代码生成
make init          # 初始化开发工具（protoc 等插件）
make enum          # 从 proto 文件生成枚举代码
make error         # 从 proto 文件生成错误码代码

# 代码格式化 & 提交
make git           # 格式化代码并 git add
go mod tidy        # 整理依赖
go fmt ./...       # 格式化所有 Go 文件
```

---

## 扩展阅读

- [新增支付渠道开发指南](./develop.md) — 详细的渠道接入教程与完整代码示例

## License

See [LICENSE](../LICENSE) for details.
