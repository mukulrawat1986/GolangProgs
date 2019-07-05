package main

import (
	"fmt"
	"io"
	"os"
)

func Echo(args []string, w io.Writer) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Fprintf(w, "%s\n", s)
}

func main() {
	Echo(os.Args, os.Stdout)
}
