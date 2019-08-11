package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mukulrawat1986/GolangProgs/TestingGo/serverex/server"
)

// Integration test
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	svr := &server.PlayerServer{
		Store: store,
	}
	player := "Pepper"

	svr.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	svr.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))

	response := httptest.NewRecorder()
	svr.ServeHTTP(response, server.NewGetScoreRequest(player))

	server.AssertStatus(t, response.Code, http.StatusOK)

	server.AssertResponseBody(t, response.Body.String(), "3")
}
