package main

/*
#cgo LDFLAGS: build/lib/libbluetooth.a build/lib/libbluetooth_example.a build/lib/libsimpleble.a build/lib/libsimpleble-c.a -lstdc++ -ldl -framework Foundation -framework CoreBluetooth
#cgo CFLAGS: -I../SimpleBLE/simpleble/include -I./build/simpleble/export
#include <stdlib.h>
#include "./src/bluetooth.h"
#include <simpleble_c/simpleble.h>
*/
import "C"

var myChannel chan string

//export SendGoMessage
func SendGoMessage(cString *C.char) {
	goString := C.GoString(cString)
	myChannel <- goString
}

func BluetoothScan() {
	C.bluetooth_scan()
}
