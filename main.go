package main

func main() {
	myChannel = make(chan string)
	go BluetoothScan()
	var s string
	select {
	case s = <-myChannel:
		println("got", s, "on my channel")
	}
}
