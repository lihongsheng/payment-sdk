package driver

import (
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"sync"
)

var driversMu sync.RWMutex

// paymentDrivers
var paymentDrivers = make(map[string]iface.PaymentDriver)

// refundDrivers
var refundDrivers = make(map[string]iface.RefundDriver)

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

func PaymentRegister(channelName string, driver iface.PaymentDriver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("Payment: Register driver is nil")
	}
	if _, dup := paymentDrivers[channelName]; dup {
		panic("Payment: Register called twice for driver " + channelName)
	}
	paymentDrivers[channelName] = driver
}

func RefundRegister(channelName string, driver iface.RefundDriver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("Refund: Register driver is nil")
	}
	if _, dup := refundDrivers[channelName]; dup {
		panic("Refund: Register called twice for driver " + channelName)
	}
	refundDrivers[channelName] = driver
}

func Payment(channelName string, cf config.Config) (iface.Pay, error) {
	driversMu.RLock()
	driver, ok := paymentDrivers[channelName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("payment: unknown driver %s (forgotten import?)", channelName)
	}
	connector, err := driver.Open(cf)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

func Refund(channelName string, cf config.Config) (iface.Refund, error) {
	driversMu.RLock()
	driver, ok := refundDrivers[channelName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("refund: unknown driver %s (forgotten import?)", channelName)
	}
	connector, err := driver.Open(cf)
	if err != nil {
		return nil, err
	}
	return connector, nil
}
