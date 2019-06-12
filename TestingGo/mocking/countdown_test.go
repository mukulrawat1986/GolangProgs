package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	write = "write"
	sleep = "sleep"
)

// CountdownOperationSpy is a spy to count the number of times Sleep is called, and the order
// of operations
type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("print 3 to go", func(t *testing.T) {
		// create a buffer that we inject to the Countdown function
		// This buffer is where the writing happens in our test, so we can
		// capture what data is being generated
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted called %v got calls %v", want, spySleepPrinter.Calls)
		}
	})
}
