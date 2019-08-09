package server

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) string {
	switch player {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	default:
		return ""
	}
}
