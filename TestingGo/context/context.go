package _context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get context from request
		ctx := r.Context()

		// create a channel to send data in
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)

		case <-ctx.Done():
			// request is cancelled
			store.Cancel()
		}

	}
}
