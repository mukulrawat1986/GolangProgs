package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	fmt.Print(writer, "3")
}

func main() {
	Countdown(os.Stdout)
}
