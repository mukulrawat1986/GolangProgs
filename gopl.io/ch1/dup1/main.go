// dup prints the text of each line that appears more than
// once in the standar input, preceeded by its count
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Dup(r io.Reader, w io.Writer) {
	counts := make(map[string]int)

	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(w, "%d\t%s\n", n, line)
		}
	}
}

func main() {
	Dup(os.Stdin, os.Stdout)
}
