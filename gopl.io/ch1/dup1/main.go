// dup prints the text of each line that appears more than
// once in the standar input, preceeded by its count
package main

import (
	"io"
	"os"
)

func Dup(r io.Reader, w io.Writer) {
}

func main() {
	Dup(os.Stdin, os.Stdout)
}
