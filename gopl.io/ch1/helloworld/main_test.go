package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("when a name is passed", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Hello(buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
