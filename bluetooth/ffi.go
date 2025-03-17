package bluetooth

/*
#cgo LDFLAGS: bluetooth/build/lib/libbluetooth.a bluetooth/build/lib/libbluetooth_example.a bluetooth/build/lib/libsimpleble.a bluetooth/build/lib/libsimpleble-c.a -lstdc++ -ldl -framework Foundation -framework CoreBluetooth
#cgo CFLAGS: -I../../SimpleBLE/simpleble/include -I./build/simpleble/export
#include <stdlib.h>
#include "./src/bluetooth.h"
#include <simpleble_c/simpleble.h>
*/
import "C"
import "fmt"

var deviceScanChannel chan ScanMsg

type ScanDone struct{}
type Device struct {
	Identifier string
	Address    string
}

type ScanMsg any

func (i Device) String() string {
	switch i.Identifier {
	case "":
		return fmt.Sprintf("%-20s %s", "Unknown", i.Address)
	default:
		return fmt.Sprintf("%-20s %s", i.Identifier, i.Address)
	}
}

func CreateChannel() chan ScanMsg {
	deviceScanChannel = make(chan ScanMsg)
	return deviceScanChannel
}

//export AnnounceDeviceWasFound
func AnnounceDeviceWasFound(cIdentifier *C.char, cAddress *C.char) {
	identifier := C.GoString(cIdentifier)
	address := C.GoString(cAddress)
	deviceScanChannel <- Device{identifier, address}
}

//export AnnounceScanIsDone
func AnnounceScanIsDone() {
	deviceScanChannel <- ScanDone{}
}

func StartScan() {
	C.bluetooth_scan()
}
