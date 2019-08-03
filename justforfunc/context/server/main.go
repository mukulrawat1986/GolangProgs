package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	// the context in handler
	ctx := r.Context()

	log.Println("Handler started")
	defer log.Println("Handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello, World")

	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
