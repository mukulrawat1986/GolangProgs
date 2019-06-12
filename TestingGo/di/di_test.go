package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// create a buffer of bytes.Buffer and inject it to our Greet function
	buffer := &bytes.Buffer{}
	Greet(buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
