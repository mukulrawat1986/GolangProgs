package main

import (
	"bytes"
	"testing"
)

// SpySleeper to use in our tests as a spy
type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

func TestCountdown(t *testing.T) {
	// create a buffer that we inject to the Countdown function
	// This buffer is where the writing happens in our test, so we can
	// capture what data is being generated
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.calls)
	}
}
