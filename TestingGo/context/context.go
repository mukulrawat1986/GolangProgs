package _context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
