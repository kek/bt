package main

/*
#cgo LDFLAGS: build/lib/libbluetooth.a build/lib/libbluetooth_example.a build/lib/libsimpleble.a build/lib/libsimpleble-c.a -lstdc++ -ldl -framework Foundation -framework CoreBluetooth
#cgo CFLAGS: -I../SimpleBLE/simpleble/include -I./build/simpleble/export
#include <stdlib.h>
#include "./src/bluetooth.h"
#include <simpleble_c/simpleble.h>
*/
import "C"

var myChannel chan ScanMsg

type ScanDone struct{}
type DeviceFound struct {
	Identifier string
	Address    string
}

type ScanMsg interface{}

func (i DeviceFound) String() string {
	switch i.Identifier {
	case "":
		return "Unknown" + "/" + i.Address
	default:
		return i.Identifier + "/" + i.Address
	}
}

//export GoDeviceFound
func GoDeviceFound(cIdentifier *C.char, cAddress *C.char) {
	identifier := C.GoString(cIdentifier)
	address := C.GoString(cAddress)
	myChannel <- DeviceFound{identifier, address}
}

//export SendGoDone
func SendGoDone() {
	myChannel <- ScanDone{}
}

func BluetoothScan() {
	C.bluetooth_scan()
}
