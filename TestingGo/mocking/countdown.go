package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Sleeper interface to handle our dependency on Sleep()
type Sleeper interface {
	 Sleep()
}

type DefaultSleeper struct {}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
