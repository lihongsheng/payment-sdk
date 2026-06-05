package driver

import (
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	errors2 "github.com/lihongsheng/payment-sdk/errors"
	"sync"
)

var driversMu sync.RWMutex

// paymentDrivers
var paymentDrivers = make(map[channel.Channel]iface.PaymentDriver)

// refundDrivers
var refundDrivers = make(map[channel.Channel]iface.RefundDriver)

//var paymentSupportProduct = make(map[channel.Channel]map[payment.PaymentProduct]enum.Device)
// unitTransferDrivers
// var unitTransferDrivers = make(map[string]iface.UnitTransfer)

// UnitTransferRegister
// 单人单次转账
//func UnitTransferRegister(name channel.Channel, driver iface.PaymentDriver) {
//	driversMu.Lock()
//	defer driversMu.Unlock()
//	if driver == nil {
//		panic("Payment: Register driver is nil")
//	}
//	if _, dup := paymentDrivers[name]; dup {
//		panic("Payment: Register called twice for driver " + name.String())
//	}
//	paymentDrivers[name] = driver
//}

func PaymentRegister(channelName channel.Channel, driver iface.PaymentDriver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("Payment: Register driver is nil")
	}
	if _, dup := paymentDrivers[channelName]; dup {
		panic("Payment: Register called twice for driver " + channelName.String())
	}
	paymentDrivers[channelName] = driver
}

func GetPaymentDriver(channelName channel.Channel) (iface.PaymentDriver, error) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver, dup := paymentDrivers[channelName]; dup {
		return driver, nil
	}
	return nil, errors2.ErrorNotFound("not find dirve by [%s]", channelName.String())
}

func RefundRegister(channelName channel.Channel, driver iface.RefundDriver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("Refund: Register driver is nil")
	}
	if _, dup := refundDrivers[channelName]; dup {
		panic("Refund: Register called twice for driver " + channelName.String())
	}
	refundDrivers[channelName] = driver
}

func GetRefundDriver(channelName channel.Channel) (iface.RefundDriver, error) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver, dup := refundDrivers[channelName]; dup {
		return driver, nil
	}
	return nil, errors2.ErrorNotFound("not find dirve by [%s]", channelName.String())
}

func Payment(channelName channel.Channel, cf config.Config) (iface.Pay, error) {
	driversMu.RLock()
	driver, ok := paymentDrivers[channelName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("payment: unknown driver %s (forgotten import?)", channelName.String())
	}
	connector, err := driver.Open(cf)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

func Refund(channelName channel.Channel, cf config.Config) (iface.Refund, error) {
	driversMu.RLock()
	driver, ok := refundDrivers[channelName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("refund: unknown driver %s (forgotten import?)", channelName.String())
	}
	connector, err := driver.Open(cf)
	if err != nil {
		return nil, err
	}
	return connector, nil
}
