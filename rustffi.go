package main

/*
#cgo LDFLAGS: ./build/libbluetooth.a build/libsimpleble.a build/libcxxbridge1.a build/libsimpleble_bindings.a build/libbluetooth.a build/libsimpleble-c.a build/liblink-cplusplus.a -L./build -ldl
#cgo CFLAGS: -I./include -x c++ -std=c++17 -stdlib=libc++
#cgo CPPFLAGS: -I./include -std=c++17 -stdlib=libc++
#include <stdlib.h>
#include "./bluetooth.h"
#include "./include/simpleble/SimpleBLE.h"
#include <functional>
#include <optional>
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
