
# CLAUDE.md

该文件可为 Claude Code（claude.ai/code）在本代码仓库中处理代码时提供操作指引。

## 项目概述

这是一款 Go 语言支付软件开发工具包（pay-sdk），为支付宝、微信支付、富友、拉卡拉等多种支付渠道提供统一调用接口。该工具包采用基于驱动的架构设计，各类支付渠道均实现统一通用接口。
## Commands

### Build
```bash
go build ./...
```

### Run Tests
```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./adapter/alipay/payment/...

# Run a single test
go test -run TestPayment_WxpayJSAPI ./...

# Run tests with verbose output
go test -v ./...
```

### 代码生成
```bash
# 生成proto枚举文件
make enum

# 从proto文件生成错误文件
make error

# 格式化代码并准备提交至 Git
make git

# 初始化开发工具
make init
```

### 其他
```bash
go mod tidy    # 整理依赖项
go fmt ./...   # 格式化所有 Go 语言文件
```

## 架构

### 驱动注册模式

该软件开发工具包采用与 Go 语言database/sql相似的驱动注册模式：

1. **驱动注册** `（driver/register.go）`：每个支付通道都通过init()函数完成自身注册
2. **导入** `（privoder.go 文件）`：导入适配器包会触发其init()函数完成驱动注册
3. **工厂函数**: `Payment()` and `Refund()` in `privoder.go` 返回各渠道专属实现方案

### 关键文件

- `privoder.go` - 主要入口. 导出函数如 `Payment()`, `Refund()`, `GetPaymentDriver()`, `GetRefundDriver()`
- `driver/register.go` - 集中式驱动程序注册表具备 `PaymentRegister()` / `RefundRegister()`
- `driver/iface/` - 定义核心接口: `PaymentDriver`, `Pay`, `RefundDriver`, `Refund`
- `config/conf.go` - 配置结构体与函数式选项模式

### 适配器结构

`adapter/` 目录下的每一个支付通道均遵循此规则:
```
adapter/{channel}/
├── driver/driver.go    # 实现 PaymentDriver/RefundDriver, registers in init()
├── config/config.go    # 通道专属配置
├── client/             # HTTP 客户端与签名逻辑
├── payment/            # 支付方式实现方案 (JSAPI, H5, Qrcode, APP, etc.)
└── refund/             # 退款实现
```

**支持的渠道:**
- `alipay/` - Alipay 支付宝
- `wxpay/` - WeChat 微信
- `fuiou/` - Fuiou 富友
- `lakala/` - Lakala 拉卡拉

### 使用模式

```go
import (
    payment_sdk "github.com/lihongsheng/payment-sdk"
    "github.com/lihongsheng/payment-sdk/config"
    "github.com/lihongsheng/payment-sdk/driver/dto"
    "github.com/lihongsheng/payment-sdk/enum/channel"
    "github.com/lihongsheng/payment-sdk/enum/payment"
)

// 1. 使用功能选项进行配置
opts := []config.Option{
    config.WithPaymentProduct(payment.PaymentProduct_JSAPI),
    config.WithWxConfig(&wxConfig.Config{...}),
}

// 2. 获取渠道支付接口
pay, err := payment_sdk.Payment(channel.Channel_Wechat, opts...)

// 3. 使用这些接口
resp, err := pay.Pay(ctx, &dto.PayOrder{...})
detail, err := pay.Query(ctx, dto.Query{...})
callbackDetail, err := pay.Callback(ctx, req)
```

### 枚举类

枚举类型在enum/目录下以proto消息格式定义，通过make enum命令生成。核心枚举类型如下：
- `channel.Channel` - 支付渠道 (Alipay, Wechat, Fuiou, Lakala)
- `payment.Payment` - 支付方式
- `payment.PaymentProduct` - 特定产品 (JSAPI, H5, Qrcode, APP, LITE)
- `enum.Device` - 支持的设备类型 (H5, PC, APP, Alipay, Wechat, etc.)

### 新增支付渠道

1. 创建 `adapter/{new_channel}/` 目录接口
2. 实现 `iface.PaymentDriver` 和 `iface.RefundDriver` 在 `driver/driver.go`
3. 调用 `driver.PaymentRegister()` 和 `driver.RefundRegister()` 在 `init()`
4. 添加空导入至 `privoder.go`
5. 添加渠道专属配置至 `config/conf.go`
