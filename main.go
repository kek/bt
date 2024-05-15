package main

func main() {
	gS := getStringFromC()
	println("go string: >>>", gS, "<<<")

	//C.bluetooth_upload(C.CString("console.log('hello world!')"))

	// create channel
	//devices := make(chan bluetooth.Device)
	//go scan(scanResults, devices)
	//scanResult := <-scanResults
}

/*
func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
*/
