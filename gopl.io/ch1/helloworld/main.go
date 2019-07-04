package main

import (
	"fmt"
	"io"
	"os"
)

func Hello(w io.Writer, name string) {
	res := "Hello, " + name
	fmt.Fprint(w, res)
}

func main() {
	Hello(os.Stdout, "")
}
