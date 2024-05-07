package main

import (
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	must("enable BLE stack", adapter.Enable())

	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		name := device.LocalName()
		if deviceIsBangle(name) {
			println("found:", adapter, name)
		}
	})
	must("start scan", err)
}

func deviceIsBangle(name string) bool {
	return len(name) > 9 && name[:9] == "Bangle.js"
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
