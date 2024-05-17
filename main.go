package main

func main() {
	myChannel = make(chan string)
	go BluetoothScan()
	s := <-myChannel
	println("got", s, "on my channel")
}
