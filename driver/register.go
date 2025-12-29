package driver

import (
	"fmt"
	"github.com/lihongsheng/payment-sdk/config"
	"github.com/lihongsheng/payment-sdk/driver/iface"
	"github.com/lihongsheng/payment-sdk/enum/channel"
	"sync"
)

var driversMu sync.RWMutex

//go:linkname paymentDrivers
var paymentDrivers = make(map[channel.Channel]iface.PaymentDriver)

//go:linkname refundDrivers
var refundDrivers = make(map[channel.Channel]iface.RefundDriver)

// go:linkname unitTransferDrivers
//var unitTransferDrivers = make(map[string]UnitTransfer)

func PaymentRegister(name channel.Channel, driver iface.PaymentDriver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("Payment: Register driver is nil")
	}
	if _, dup := paymentDrivers[name]; dup {
		panic("Payment: Register called twice for driver " + name.String())
	}
	paymentDrivers[name] = driver
}

func RefundRegister(name channel.Channel, driver iface.RefundDriver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("Refund: Register driver is nil")
	}
	if _, dup := refundDrivers[name]; dup {
		panic("Refund: Register called twice for driver " + name.String())
	}
	refundDrivers[name] = driver
}

func Payment(driverName channel.Channel, cf config.Config) (iface.Pay, error) {
	driversMu.RLock()
	driver, ok := paymentDrivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("payment: unknown driver %s (forgotten import?)", driverName.String())
	}
	connector, err := driver.Open(cf)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

func Refund(driverName channel.Channel, cf config.Config) (iface.Refund, error) {
	driversMu.RLock()
	driver, ok := refundDrivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("refund: unknown driver %s (forgotten import?)", driverName.String())
	}
	connector, err := driver.Open(cf)
	if err != nil {
		return nil, err
	}
	return connector, nil
}
