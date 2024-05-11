package main

import (
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	must("enable BLE stack", adapter.Enable())

	// create channel
	scanResults := make(chan bluetooth.ScanResult)
	devices := make(chan bluetooth.Device)
	println("scanning...")
	go scan(scanResults, devices)
	println("oh yeah")
	scanResult := <-scanResults
	println("all right")
	device := <-devices
	println("for sure")
	println(scanResult.LocalName())
	services, err := device.DiscoverServices(nil)
	must("discover services", err)
	loopThroughServices(services)
	println("Done!")
}

func loopThroughServices(services []bluetooth.DeviceService) {
	for i := range services {
		service := services[i]
		println(service.String())
		characteristics, err := service.DiscoverCharacteristics(nil)
		must("discover characteristics", err)
		loopThroughCharacteristics(characteristics)

		println("")
	}
}

func loopThroughCharacteristics(characteristic []bluetooth.DeviceCharacteristic) {
	for i := range characteristic {
		char := characteristic[i]
		// search for 6e400002b5a3f393e0a9e50e24dcca9e
		println(char.String())
	}
}

func scan(scanResults chan bluetooth.ScanResult, devices chan bluetooth.Device) {
	err := adapter.Scan(func(adapter *bluetooth.Adapter, scanResult bluetooth.ScanResult) {
		name := scanResult.LocalName()
		if deviceIsBanglejs(name) {
			_ = adapter.StopScan()
			scanResults <- scanResult
			println("got a Bangle.js device", scanResult.LocalName())
			device, err := adapter.Connect(scanResult.Address, bluetooth.ConnectionParams{})
			must("connect", err)
			devices <- device
		}
	})

	must("start scan", err)
}

func deviceIsBanglejs(name string) bool {
	return len(name) > 9 && name[:9] == "Bangle.js"
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
