package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	// convert the timeHandler function to a HandlerFunc type
	// th := http.HandlerFunc(timeHandler)

	// add it to the ServeMux
	mux.HandleFunc("/time", timeHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
