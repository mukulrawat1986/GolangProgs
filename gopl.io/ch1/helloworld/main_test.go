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

	t.Run("when no name is passed to the function", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Hello(buffer, "")

		got := buffer.String()
		want := "Hello, World"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
