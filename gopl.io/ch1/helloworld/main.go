package main

import (
	"fmt"
	"io"
	"os"
)

func Hello(w io.Writer, name string) {
	var res string
	if name == "" {
		res = "Hello, World"
	} else {
		res = "Hello, " + name
	}
	fmt.Fprint(w, res)
}

func main() {
	Hello(os.Stdout, "")
}
