// echo1 prints its command line arguments
package main

import (
	"io"
	"os"
)

func Echo(args []string, w io.Writer) {

}

func main() {
	Echo(os.Args, os.Stdout)
}
