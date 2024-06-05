package main

import "fmt"

func main() {
	myChannel = make(chan ScanMsg)
	go BluetoothScan()
	done := false
	for !done {
		s := <-myChannel
		switch s.(type) {
		case MsgInfo:
			fmt.Println("got", s, "on my channel")
		case MsgDone:
			println("DONE")
			done = true
		}
	}
}
