package main

/*
#cgo LDFLAGS: ./rust/target/release/libbluetooth.a -ldl
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

func getStringFromC() string {
	cString := C.make_string()
	goString := C.GoString(cString)
	return goString
}

func sendMeAMessage() {
	C.send_me_a_message()
}
