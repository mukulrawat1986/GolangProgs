package main

import (
	"bytes"
	"testing"
)

func TestDup(t *testing.T) {

	inputArgs := "Hello\nHelloooo\nHello\nHello World\nHello World\n"

	inputBuff := bytes.NewBuffer([]byte(inputArgs))
	outputBuff := &bytes.Buffer{}

	Dup(inputBuff, outputBuff)

	got := outputBuff.String()
	want := "2\tHello\n2\tHello World\n"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
