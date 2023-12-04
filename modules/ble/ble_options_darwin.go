package ble

import (
	"github.com/Asif-Iqbal-Gazigatt"
)

var defaultBLEClientOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.CentralManager),
}

/*

var defaultBLEServerOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.PeripheralManager),
}

*/
