package main

import "testing"

func TestHelloReturn(t *testing.T) {
	want := "Hello World"
	returnValue := SayHello()
	if want != returnValue {
		t.Errorf("want: 'Hello World' got: '%s'\n", returnValue)
	}
}
