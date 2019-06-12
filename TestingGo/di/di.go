package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)

}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	Greet(os.Stdout, "Chris")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}
