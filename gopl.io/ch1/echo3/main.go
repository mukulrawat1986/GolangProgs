package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo(args []string, w io.Writer) {
	fmt.Fprintf(w, "%s\n", strings.Join(args[1:], " "))
}

func main() {
	Echo(os.Args, os.Stdout)
}
