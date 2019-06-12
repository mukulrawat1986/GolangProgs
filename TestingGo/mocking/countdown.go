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

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
