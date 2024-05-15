package main

/*
#cgo LDFLAGS: ./rust/target/release/libbluetooth.a -ldl
#include <stdlib.h>
#include "./bluetooth.h"
*/
import "C"

//export addTogether
func addTogether(x C.int, y C.int) C.int { return C.int(x + y) }

func main() {
	cS := C.make_string()
	println("C string:", cS)
	gS := C.GoString(cS)
	println("go string: >>>", gS, "<<<")

	//C.bluetooth_upload(C.CString("console.log('hello world!')"))

	// create channel
	//devices := make(chan bluetooth.Device)
	//go scan(scanResults, devices)
	//scanResult := <-scanResults
}

/*
func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
*/
