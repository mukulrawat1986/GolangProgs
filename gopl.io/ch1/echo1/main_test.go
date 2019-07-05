package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	args := []string{"filename", "hello", "world"}
	buff := &bytes.Buffer{}

	Echo(args, buff)

	got := buff.String()
	want := "hello world\n"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
