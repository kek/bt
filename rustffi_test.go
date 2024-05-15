package main

import (
	"testing"
)

func TestReturnString(t *testing.T) {
	want := "Whatevery"
	got := getStringFromC()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
func TestReceiveMessage(t *testing.T) {
	myChannel = make(chan string)
	want := "Hey there"
	go sendMeAMessage()
	got := <-myChannel
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
