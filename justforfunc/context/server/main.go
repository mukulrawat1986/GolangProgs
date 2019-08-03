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
	log.Println("Handler started")
	defer log.Println("Handler ended")
	// make it really slow
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Hello, World")
}
