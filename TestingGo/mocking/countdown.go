package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord = "Go!"
	countdownStart = 3
)

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w,finalWord)
}

func main() {
	Countdown(os.Stdout)
}
