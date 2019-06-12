package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// create a buffer that we inject to the Countdown function
	// This buffer is where the writing happens in our test, so we can
	// capture what data is being generated
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
