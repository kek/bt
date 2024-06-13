package main

import "fmt"

func main() {
	myChannel = make(chan ScanMsg)
	go BluetoothScan()
	done := false
	for !done {
		s := <-myChannel
		switch s.(type) {
		case DeviceFound:
			fmt.Println("Found", s)
		case ScanDone:
			done = true
		}
	}
}
