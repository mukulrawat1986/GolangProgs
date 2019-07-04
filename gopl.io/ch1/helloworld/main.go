package main

import (
	"io"
	"os"
)

func Hello(w io.Writer, name string) {
}

func main() {
	Hello(os.Stdout, "")
}
