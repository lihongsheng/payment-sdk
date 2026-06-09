# 新增支付渠道开发指南

> 本文档详细说明如何为一个新的支付渠道开发适配器，并将其集成到 Payment SDK 中。

---

## 目录

1. [概述](#1-概述)
2. [准备步骤](#2-准备步骤)
3. [创建适配器目录结构](#3-创建适配器目录结构)
4. [实现渠道配置](#4-实现渠道配置)
5. [实现 HTTP 客户端](#5-实现-http-客户端)
6. [实现支付接口](#6-实现支付接口)
7. [实现退款接口](#7-实现退款接口)
8. [实现驱动注册](#8-实现驱动注册)
9. [注册到核心入口](#9-注册到核心入口)
10. [添加渠道枚举](#10-添加渠道枚举)
11. [添加到全局配置](#11-添加到全局配置)
12. [完整示例](#12-完整示例)
13. [最佳实践](#13-最佳实践)

---

## 1. 概述

Payment SDK 采用**驱动注册模式**（与 Go `database/sql` 相同），新增渠道的核心工作就是实现 `driver/iface/` 中定义的接口，并通过 `driver/register.go` 的注册函数将驱动注册到全局注册表。

### 新增渠道需要实现的接口

| 接口 | 文件 | 说明 |
|------|------|------|
| `PaymentDriver` | `driver/iface/pay.go` | 支付驱动，负责实例化 Pay 接口 |
| `Pay` | `driver/iface/pay.go` | 支付操作：支付、查询、关闭、回调 |
| `RefundDriver` | `driver/iface/refund.go` | 退款驱动，负责实例化 Refund 接口 |
| `Refund` | `driver/iface/refund.go` | 退款操作：退款、查询、回调 |
| `Validate` | `driver/iface/validate.go` | 可选，用于配置参数校验 |

### 新增渠道需要修改的文件

| 文件 | 操作 | 说明 |
|------|------|------|
| `enum/channel/channel.proto` | 添加新枚举值 | 增加渠道标识 |
| `config/conf.go` | 添加配置字段 + Option | 增加渠道配置结构体和 WithXxx 选项 |
| `privoder.go` | 添加空导入 | 触发新渠道驱动的 init() 注册 |
| `adapter/{channel}/` | 新增目录 | 实现全部适配器代码 |

---

## 2. 准备步骤

### 2.1 了解渠道 API

在开发前，请先梳理目标支付渠道的以下信息：

| 信息 | 说明 |
|------|------|
| 支付接口 | 统一下单（JSAPI / H5 / 扫码 / APP 等） |
| 支付查询接口 | 查询订单状态 |
| 关闭订单接口 | 关闭未支付订单 |
| 退款接口 | 发起退款 |
| 退款查询接口 | 查询退款状态 |
| 支付回调 | 异步通知格式与签名验证方式 |
| 退款回调 | 退款异步通知 |
| 签名算法 | RSA / HMAC / MD5 等 |
| 加密要求 | 请求/响应是否需要加解密 |
| 证书要求 | 是否需要商户证书、平台公钥等 |

### 2.2 确认渠道枚举值

确认在 `enum/channel/channel.proto` 中为新渠道分配一个**未使用的枚举值**：

```protobuf
enum Channel {
  Channel_UNKNOWN = 0;
  Alipay = 1;   // 支付宝
  Wechat = 2;   // 微信
  Fuiou = 3;    // 富友
  Lakala = 4;   // 拉卡拉
  // NewChannel = 5;  ← 新增渠道在这里
}
```

---

## 3. 创建适配器目录结构

在 `adapter/` 目录下以渠道名为标识创建目录：

```
adapter/{channel}/
├── driver/
│   ├── driver.go     # 实现 PaymentDriver / RefundDriver，在 init() 中注册
│   └── options.go    # 配置的 JSON Schema（可选，用于前端动态表单）
├── config/
│   └── config.go     # 渠道专属配置结构体 & Validate() 方法
├── client/
│   └── client.go     # HTTP 客户端、签名逻辑、请求/响应处理
├── payment/
│   ├── jsapi.go      # JSAPI/公众号支付
│   ├── h5.go         # H5 支付
│   ├── native.go     # 扫码支付（可选，渠道支持的话）
│   ├── app.go        # APP 支付（可选）
│   └── lite.go       # 小程序支付（可选）
├── refund/
│   └── refund.go     # 退款接口实现
└── model/
    └── model.go      # 渠道内部请求/响应结构体（可选）
```

> **命名惯例**：目录名使用全小写英文字母，不包含空格或特殊字符（如 `wxpay`、`alipay`、`fuiou`）。

---

## 4. 实现渠道配置

### 4.1 定义配置结构体

新建 `adapter/{channel}/config/config.go`，定义渠道特有的配置参数：

```go
package config

import "github.com/lihongsheng/payment-sdk/errors"

type Config struct {
    Merchant Merchant `json:"merchant"`
    Cert     Cert     `json:"cert"`
    API      API      `json:"api"`
}

type Merchant struct {
    AppID string `json:"app_id"`
    MchID string `json:"mch_id"`
}

type Cert struct {
    RsaPrivate string `json:"rsa_private_key"`
    RsaPublic  string `json:"rsa_public_key"`
}

type API struct {
    ApiHost string `json:"api_host"`
    Version string `json:"version"`
}

// Validate 校验配置参数完整性，在初始化时调用
func (c Config) Validate() error {
    if c.Merchant.AppID == "" {
        return errors.ErrorParamError("新渠道: 应用ID is empty")
    }
    if c.Cert.RsaPrivate == "" {
        return errors.ErrorParamError("新渠道: 私钥 is empty")
    }
    return nil
}
```

**设计原则：**

- 配置字段按功能分组（`Merchant` / `Cert` / `API` / `Service` 等）
- 所有字段添加 `json` tag，支持 JSON 字符串反序列化
- 实现 `Validate()` 方法校验必填参数
- 错误码使用 `errors` 包的预定义函数（`ErrorParamError` / `ErrorNotFound` 等）

### 4.2 注册全局配置选项

在 `config/conf.go` 中为新渠道添加配置字段和 `WithXxx` 选项函数：

```go
import (
    // 添加 import
    newch "github.com/lihongsheng/payment-sdk/adapter/{channel}/config"
)

type Config struct {
    // ... 原有字段

    // 新增字段
    NewChannelConfig *newch.Config `json:"{channel}_config"`
}

// 新增选项函数
func With{Channel}Config(cf *newch.Config) Option {
    return func(c *Config) {
        c.NewChannelConfig = cf
    }
}
```

---

## 5. 实现 HTTP 客户端

新建 `adapter/{channel}/client/client.go`，负责渠道 API 的 HTTP 通信、签名和加解密：

```go
package client

import (
    "context"
    "net/http"
    "github.com/lihongsheng/payment-sdk/adapter/{channel}/config"
    "github.com/lihongsheng/payment-sdk/config/proxy"
)

type Client struct {
    Config     config.Config
    HttpClient *http.Client
    // 其他内部状态（如私钥对象、token 等）
}

func NewClient(cf config.Config, proxyCfg *proxy.Proxy) (*Client, error) {
    // 1. 调用 cf.Validate() 校验配置
    // 2. 解析 RSA 私钥/公钥
    // 3. 构建 HTTP Client（支持代理）
    // 4. 返回 Client 实例
}

// Post 发送 POST 请求，自动完成签名
func (c *Client) Post(ctx context.Context, path string, req any) ([]byte, error) {
    // 1. 序列化请求参数
    // 2. 加签
    // 3. 发送 HTTP 请求
    // 4. 验签响应
    // 5. 返回响应 body
}
```

**功能要点：**

- **签名** — 实现渠道要求的签名算法（RSA-SHA256 / HMAC-SHA256 / MD5 等）
- **代理** — 支持通过 `config/proxy.Proxy` 配置 HTTP 代理
- **错误处理** — 将渠道 HTTP 错误码转为 SDK 内部错误，附加原始响应便于排查
- **重试** — 对网络级别的临时错误实现幂等重试（可选）
- **日志** — 使用 `github.com/zeromicro/go-zero/core/logc` 记录请求和响应日志

---

## 6. 实现支付接口

### 6.1 实现 Pay 接口

在 `adapter/{channel}/payment/` 下创建支付方式实现文件（如 `jsapi.go`）：

```go
package payment

import (
    "context"
    "net/http"

    "github.com/lihongsheng/payment-sdk/adapter/{channel}/client"
    "github.com/lihongsheng/payment-sdk/driver/dto"
    enum "github.com/lihongsheng/payment-sdk/enum/payment"
)

type Jsapi struct {
    client *client.Client
}

// NewJsApi 创建 JSAPI 支付实例，返回 iface.Pay 接口
func NewJsApi(cl *client.Client) (*Jsapi, error) {
    return &Jsapi{client: cl}, nil
}

// Pay 发起支付
// 渠道返回 prepay_id 后，组装前端调起支付所需的参数
func (j *Jsapi) Pay(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
    // 1. 将 dto.PayOrder 转为渠道请求结构体
    // 2. 调用 HTTP 客户端发送请求
    // 3. 将渠道响应转为 dto.PayResponse

    return &dto.PayResponse{
        OrderNo:        req.Order.OrderNo,
        TradeNo:        "channel_trade_no",
        PayAmount:      req.Order.PayAmount,
        Status:         enum.Status_Pending,
        PaymentProduct: enum.PaymentProduct_JSAPI.String(),
        Action: dto.Action{
            Action: "prepay", // "redirect" | "qrcode" | "prepay"
            Parameters: map[string]string{
                "appId":     j.client.Config.Merchant.AppID,
                "timeStamp": "...",
                "nonceStr":  "...",
                "package":   "prepay_id=...",
                "paySign":   "...",
            },
        },
    }, nil
}

// Query 查询订单
func (j *Jsapi) Query(ctx context.Context, req dto.Query) (*dto.PayDetail, error) {
    // ...
}

// Close 关闭订单
func (j *Jsapi) Close(ctx context.Context, req dto.CloseQuery) error {
    // ...
}

// Complete 完成支付（付款码等需要确认的场景）
func (j *Jsapi) Complete(ctx context.Context, req *dto.PayOrder) (*dto.PayResponse, error) {
    // ...
}

// Callback 处理支付回调
func (j *Jsapi) Callback(ctx context.Context, req *http.Request) (*dto.CallbackPayDetail, error) {
    // 1. 读取请求 body
    // 2. 验签
    // 3. 解析回调数据
    // 4. 返回 dto.CallbackPayDetail
}
```

### 6.2 PayResponse.Action 说明

`Action` 字段告诉前端如何唤起支付：

| Action 值 | 说明 | Parameters | URL |
|-----------|------|-----------|-----|
| `"prepay"` | JSAPI/小程序唤起支付 | `appId`, `timeStamp`, `nonceStr`, `package`, `paySign` | — |
| `"redirect"` | H5 跳转支付 | — | 支付链接 URL |
| `"qrcode"` | 展示二维码 | `qrcode` 或 `qrcode_url` | 二维码 URL |
| `"app"` | APP 唤起支付 | 渠道约定的 APP 调起参数 | — |

### 6.3 创建不同支付方式的实例

在 `driver/driver.go` 的 `Open()` 方法中根据 `PaymentProduct` 返回不同的支付实现：

```go
func (p Payment) Open(c config.Config) (iface.Pay, error) {
    cl, err := initConfig(c)
    if err != nil {
        return nil, err
    }
    switch c.PaymentProduct {
    case payment.PaymentProduct_JSAPI:
        return payment.NewJsApi(cl)
    case payment.PaymentProduct_H5:
        return payment.NewH5(cl)
    case payment.PaymentProduct_Qrcode:
        return payment.NewQrcode(cl)
    case payment.PaymentProduct_APP:
        return payment.NewApp(cl)
    case payment.PaymentProduct_LITE:
        return payment.NewLite(cl)
    }
    return nil, errors.ErrorParamError("{channel}: unknown payment product")
}
```

---

## 7. 实现退款接口

新建 `adapter/{channel}/refund/refund.go`：

```go
package refund

import (
    "context"
    "net/http"

    "github.com/lihongsheng/payment-sdk/adapter/{channel}/client"
    "github.com/lihongsheng/payment-sdk/driver/dto"
)

type Refund struct {
    client *client.Client
}

func NewRefund(cl *client.Client) (*Refund, error) {
    return &Refund{client: cl}, nil
}

// Refund 发起退款
func (r *Refund) Refund(ctx context.Context, req *dto.RefundRequest) (*dto.RefundDetail, error) {
    // 1. 映射退款请求参数
    // 2. 调用渠道退款接口
    // 3. 返回统一退款返回结构

    return &dto.RefundDetail{
        RefundNo:         req.RefundNo,
        TradeRefundNo:    "channel_refund_no",
        TradeNo:          req.TradeNo,
        OrderNo:          req.OrderNo,
        Status:           refundStatus.Success,
        Amount:           req.Amount,
    }, nil
}

// Query 查询退款状态
func (r *Refund) Query(ctx context.Context, req dto.RefundQuery) (*dto.RefundDetail, error) {
    // ...
}

// Callback 处理退款回调
func (r *Refund) Callback(ctx context.Context, req *http.Request) (*dto.CallbackRefundDetail, error) {
    // ...
}

// IsSupportCallback 是否支持退款回调
func (r *Refund) IsSupportCallback() bool {
    return true // 根据渠道实际情况返回
}
```

---

## 8. 实现驱动注册

新建 `adapter/{channel}/driver/driver.go`，实现 `PaymentDriver` 和 `RefundDriver` 接口，并在 `init()` 中完成注册：

### 8.1 基础驱动实现

```go
package driver

import (
    "encoding/json"

    "github.com/lihongsheng/payment-sdk/adapter/{channel}/client"
    conf "github.com/lihongsheng/payment-sdk/adapter/{channel}/config"
    payment2 "github.com/lihongsheng/payment-sdk/adapter/{channel}/payment"
    "github.com/lihongsheng/payment-sdk/adapter/{channel}/refund"

    "github.com/lihongsheng/payment-sdk/config"
    "github.com/lihongsheng/payment-sdk/driver"
    "github.com/lihongsheng/payment-sdk/driver/iface"
    "github.com/lihongsheng/payment-sdk/enum"
    "github.com/lihongsheng/payment-sdk/enum/channel"
    "github.com/lihongsheng/payment-sdk/enum/payment"
    "github.com/lihongsheng/payment-sdk/errors"
)

func init() {
    driver.PaymentRegister(channel.Channel_{Channel}, Payment{})
    driver.RefundRegister(channel.Channel_{Channel}, Refund{})
}

type Payment struct{}

func (p Payment) Open(c config.Config) (iface.Pay, error) {
    if c.PaymentProduct == payment.PaymentProduct_PaymentMethod_UNKNOWN {
        return nil, errors.ErrorParamError("payment: unknown payment product")
    }
    cl, err := initConfig(c)
    if err != nil {
        return nil, err
    }

    switch c.PaymentProduct {
    case payment.PaymentProduct_JSAPI:
        return payment2.NewJsApi(cl)
    // 根据渠道支持情况添加更多支付方式
    }
    return nil, errors.ErrorParamError("{channel}: unknown payment product")
}

func (p Payment) GetConfigOptions() *iface.ChannelOption {
    return &iface.ChannelOption{
        Channel: channel.Channel_{Channel}.String(),
        Label:   "{渠道中文名称}",
        Schema:  schema, // 见下方 8.2
    }
}

func (p Payment) GetSupportProduct() []iface.PaymentMethod {
    return []iface.PaymentMethod{
        {
            Method: payment.Payment_Wechat.String(),
            Label:  "微信支付",
            Product: []iface.PaymentProduct{
                {Product: payment.PaymentProduct_JSAPI.String(), Label: "公众号支付(JSAPI)"},
                {Product: payment.PaymentProduct_LITE.String(), Label: "小程序支付"},
            },
        },
        {
            Method: payment.Payment_Alipay.String(),
            Label:  "支付宝支付",
            Product: []iface.PaymentProduct{
                {Product: payment.PaymentProduct_JSAPI.String(), Label: "公众号支付(JSAPI)"},
                {Product: payment.PaymentProduct_LITE.String(), Label: "小程序支付"},
            },
        },
    }
}

func (p Payment) IsSupportPayment(product payment.PaymentProduct, device enum.Device) bool {
    switch product {
    case payment.PaymentProduct_JSAPI:
        switch device {
        case enum.Device_Wechat, enum.Device_Wechat_Lite, enum.Device_Alipay, enum.Device_Alipay_Lite:
            return true
        }
    }
    return false
}

func (p Payment) CallbackResponse() string {
    return "success" // 渠道要求的回调成功响应内容
}

type Refund struct{}

func (p Refund) Open(c config.Config) (iface.Refund, error) {
    cf, err := initConfig(c)
    if err != nil {
        return nil, err
    }
    return refund.NewRefund(cf)
}

func (p Refund) CallbackResponse() string {
    return "success"
}

// initConfig 通用的配置初始化逻辑
func initConfig(c config.Config) (*client.Client, error) {
    var cf conf.Config
    if c.{Channel}Config != nil {
        cf = *c.{Channel}Config
    } else if c.Config != "" {
        if err := json.Unmarshal([]byte(c.Config), &cf); err != nil {
            return nil, errors.ErrorParamError("parse config err: %v", err)
        }
    } else {
        return nil, errors.ErrorParamError("payment: config is empty")
    }

    cl, err := client.NewClient(cf, c.Proxy)
    if err != nil {
        return nil, err
    }
    return cl, nil
}
```

### 8.2 配置 Schema（建议实现）

新建 `adapter/{channel}/driver/options.go`，定义渠道配置的 JSON Schema，用于前端动态渲染配置表单：

```go
package driver

import "github.com/lihongsheng/payment-sdk/config/params"

var schema = &params.Schema{
    Type:     params.SchemaObject,
    Title:    "{渠道中文名称}配置",
    Required: []string{"merchant", "cert"},
    Order:    []string{"merchant", "cert"},
    Properties: map[string]*params.Schema{
        "merchant": {
            Type:        params.SchemaObject,
            Title:       "商户信息",
            Description: "{渠道中文名称}商户基础信息",
            Required:    []string{"app_id", "mch_id"},
            Order:       []string{"app_id", "mch_id"},
            Properties: map[string]*params.Schema{
                "app_id": {
                    Type:     params.SchemaString,
                    Title:    "应用ID",
                    Examples: []any{"app_id_example"},
                    UI:       &params.UISchema{InputType: params.InputText},
                },
                "mch_id": {
                    Type:     params.SchemaString,
                    Title:    "商户号",
                    Examples: []any{"mch_id_example"},
                    UI:       &params.UISchema{InputType: params.InputText},
                },
            },
        },
        "cert": {
            Type:        params.SchemaObject,
            Title:       "证书配置",
            Description: "密钥与证书信息",
            Required:    []string{"rsa_private_key"},
            Order:       []string{"rsa_private_key", "rsa_public_key"},
            Properties: map[string]*params.Schema{
                "rsa_private_key": {
                    Type:     params.SchemaString,
                    Title:    "应用私钥",
                    Examples: []any{"-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----"},
                    UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPrivate, Rows: 6},
                },
                "rsa_public_key": {
                    Type:     params.SchemaString,
                    Title:    "平台公钥",
                    Examples: []any{"-----BEGIN PUBLIC KEY-----\n...\n-----END PUBLIC KEY-----"},
                    UI:       &params.UISchema{InputType: params.InputTextarea, ValidateType: params.ValidateRsaPublic, Rows: 6},
                },
            },
        },
    },
}
```

---

## 9. 注册到核心入口

在 `privoder.go` 中添加新渠道驱动的空导入，触发其 `init()` 注册：

```go
package payment_sdk

import (
    _ "github.com/lihongsheng/payment-sdk/adapter/alipay/driver"
    _ "github.com/lihongsheng/payment-sdk/adapter/fuiou/driver"
    _ "github.com/lihongsheng/payment-sdk/adapter/lakala/driver"
    _ "github.com/lihongsheng/payment-sdk/adapter/wxpay/driver"
    _ "github.com/lihongsheng/payment-sdk/adapter/{channel}/driver"  // ← 新增
    // ...
)
```

---

## 10. 添加渠道枚举

在 `enum/channel/channel.proto` 中添加新渠道枚举值：

```protobuf
enum Channel {
  Channel_UNKNOWN = 0;
  Alipay = 1;
  Wechat = 2;
  Fuiou = 3;
  Lakala = 4;
  {Channel} = 5;  // ← 新渠道，添加注释说明
}
```

运行 `make enum` 生成 Go 代码：

```bash
make enum
```

---

## 11. 添加到全局配置

在 `config/conf.go` 中完成以下修改：

### 11.1 添加配置字段

```go
type Config struct {
    // ... 原有字段
    {Channel}Config *{channel}.Config `json:"{channel}_config"`
}
```

### 11.2 添加 WithXxx 选项函数

```go
func With{Channel}Config(cf *{channel}.Config) Option {
    return func(c *Config) {
        c.{Channel}Config = cf
    }
}
```

### 11.3 添加 import

```go
import (
    // ... 原有 import
    {channel} "github.com/lihongsheng/payment-sdk/adapter/{channel}/config"
)
```

---

## 12. 完整示例

### 以"新支付渠道 NewPay"为例

#### 目录结构

```
adapter/newpay/
├── driver/
│   ├── driver.go       # PaymentDriver / RefundDriver 实现 + init() 注册
│   └── options.go      # 配置 Schema
├── config/
│   └── config.go       # Config 结构体
├── client/
│   └── client.go       # HTTP 客户端
├── payment/
│   └── jsapi.go        # JSAPI 支付
└── refund/
    └── refund.go       # 退款实现
```

#### config/conf.go 修改

```go
import (
    newpay "github.com/lihongsheng/payment-sdk/adapter/newpay/config"
)

type Config struct {
    // ...
    NewPayConfig *newpay.Config `json:"newpay_config"`
}

func WithNewPayConfig(cf *newpay.Config) Option {
    return func(c *Config) {
        c.NewPayConfig = cf
    }
}
```

#### enum/channel/channel.proto 修改

```protobuf
enum Channel {
  Channel_UNKNOWN = 0;
  Alipay = 1;
  Wechat = 2;
  Fuiou = 3;
  Lakala = 4;
  NewPay = 5; // 新支付
}
```

然后运行 `make enum` 重新生成代码。

#### privoder.go 修改

```go
import (
    // ...
    _ "github.com/lihongsheng/payment-sdk/adapter/newpay/driver"
)
```

---

## 13. 最佳实践

### 13.1 配置参数

- ❌ **不要**在配置中存放业务运行时参数（如订单号、金额等）
- ✅ **应该**在配置中存放静态凭据（商户号、AppID、API 密钥等）
- ● 配置通过 `Validate()` 做充分校验，尽早暴露配置错误
- ● 敏感信息（私钥、密钥）支持通过环境变量或密钥管理服务注入

### 13.2 错误处理

- ● 使用 `errors` 包预定义的错误码函数
- ✅ 渠道返回的业务错误码映射为 SDK 统一错误（如 `ErrorNoSupport`、`ErrorParamError` 等）
- ✅ 保留原始响应（`OriginResponse` 字段）用于排查问题

### 13.3 签名与安全

- ● 在 `client` 层统一处理签名和验签
- ● 回调处理中务必先验签再执行业务逻辑
- ● 使用 HTTPS 发起 HTTP 请求
- ✅ 支持通过 `proxy.Proxy` 配置代理（用于抓包调试）

### 13.4 回调

- ✅ 支付回调成功后调用 `CallbackResponse()` 返回渠道期望的响应内容
- ❌ 不要在回调处理中做耗时操作（如数据库写入），应异步处理

### 13.5 日志

- ● 使用 `github.com/zeromicro/go-zero/core/logc` 记录关键请求和响应日志
- ● 不要在日志中记录完整私钥或密钥原文

### 13.6 测试

- ● 为每个支付产品编写单元测试（参考 `adapter/wxpay/payment/jsapi_test.go`）
- ● 测试覆盖正常流程和错误流程
- ● 支持 Mock 渠道 HTTP 接口进行测试
- ● 测试中使用测试环境配置，不使用生产凭据

### 13.7 完整的实现清单

完成新渠道接入后，请逐项确认：

- [ ] `enum/channel/channel.proto` 添加了枚举值并运行 `make enum`
- [ ] `config/conf.go` 添加了配置字段和 `WithXxx` 选项函数
- [ ] `adapter/{channel}/config/config.go` 配置结构体实现 `Validate()`
- [ ] `adapter/{channel}/client/client.go` HTTP 客户端实现签名和验签
- [ ] `adapter/{channel}/payment/jsapi.go` 至少实现一种支付方式
- [ ] `adapter/{channel}/refund/refund.go` 退款实现
- [ ] `adapter/{channel}/driver/driver.go` 驱动注册 + `init()` 函数
- [ ] `adapter/{channel}/driver/options.go` 配置 Schema（推荐实现）
- [ ] `privoder.go` 添加了空导入
- [ ] `go build ./...` 构建通过
- [ ] `go test ./...` 测试通过