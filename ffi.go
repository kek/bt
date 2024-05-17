package main

/*
#cgo LDFLAGS: src/build/lib/libbluetooth.a src/build/lib/libsimpleble.a src/build/lib/libsimpleble-c.a -lstdc++ -ldl -framework Foundation -framework CoreBluetooth
#cgo CFLAGS: -I../SimpleBLE/simpleble/include -I./src/build/simpleble/export
#include <stdlib.h>
#include "./bluetooth.h"
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
	C.something()
}
