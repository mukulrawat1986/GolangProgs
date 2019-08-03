package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/mukulrawat1986/GolangProgs/justforfunc/context/log"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	// the context in handler
	ctx := r.Context()

	log.Println(ctx, "Handler started")
	defer log.Println(ctx, "Handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello, World")

	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
