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
	scanResult := <-scanResults
	device := <-devices
	println(scanResult.LocalName())
	services, err := device.DiscoverServices(nil)
	must("discover services", err)
	for i := range services {
		service := services[i]
		println(service.String())
		characteristic, err := service.DiscoverCharacteristics(nil)
		must("discover characteristics", err)
		for i := range characteristic {
			char := characteristic[i]
			println(char.String())
		}
		// For the simplest control, all you need to do is connect to the Espruino bluetooth device
		// and the characteristic with ID 6e400002b5a3f393e0a9e50e24dcca9e. You can then write
		// repeatedly to it to send commands to Espruino.
		println("")
	}
	println("Done!")
}

func scan(scanResults chan bluetooth.ScanResult, devices chan bluetooth.Device) {
	err := adapter.Scan(func(adapter *bluetooth.Adapter, scanResult bluetooth.ScanResult) {
		name := scanResult.LocalName()
		if deviceIsBanglejs(name) {
			connectionParams := bluetooth.ConnectionParams{
				ConnectionTimeout: 10,
				MinInterval:       1,
				MaxInterval:       10,
				Timeout:           30,
			}
			adapter.StopScan()
			device, err := adapter.Connect(scanResult.Address, connectionParams)
			must("connect", err)
			scanResults <- scanResult
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
