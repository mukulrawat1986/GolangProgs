package main

import (
	"log"
	"net/http"

	"github.com/mukulrawat1986/GolangProgs/TestingGo/serverex/server"
)

func main() {

	server := &server.PlayerServer{}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
