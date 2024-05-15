package main

/*
#cgo LDFLAGS: ./rust/target/release/libbluetooth.a -ldl
#include "./bluetooth.h"
*/
import "C"
import "reflect"

func main() {
	println("Yes...")
	c_s := C.make_string()
	println(c_s)
	g_s := C.GoString(c_s)

	println(reflect.TypeOf(g_s))
	println(g_s)
	// Skriver ut HeyNoneSome.
	// Om man istället skulle ha returnerat en sträng med längd kanske...
	println("OK...")

	//C.bluetooth_upload(C.CString("console.log('hello world!')"))

	// create channel
	//devices := make(chan bluetooth.Device)
	//go scan(scanResults, devices)
	//scanResult := <-scanResults
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
