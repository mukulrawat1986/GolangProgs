package server

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := r.URL.Path[len("/players/"):]

	if score := p.Store.GetPlayerScore(player); score == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, score)
	}
}
