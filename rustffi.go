package main

/*
#cgo LDFLAGS: build/libbluetooth.a -ldl
#include <stdlib.h>
#include "./bluetooth.h"
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
