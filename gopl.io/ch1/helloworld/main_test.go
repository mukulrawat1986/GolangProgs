package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("when a name is passed", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Hello(buffer, "Chris")
		assertStrings(t, buffer.String(), "Hello, Chris")
	})

	t.Run("when no name is passed to the function", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Hello(buffer, "")
		assertStrings(t, buffer.String(), "Hello, World")
	})
}
