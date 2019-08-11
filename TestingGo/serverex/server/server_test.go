package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// create a stub of the PlayerStore
type StubPlayerStore struct {
	score map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayer(t *testing.T) {

	store := StubPlayerStore{
		score: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: nil,
	}

	server := &PlayerServer{
		Store: &store,
	}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {

	t.Run("it returns accepted on POST", func(t *testing.T) {
		store := StubPlayerStore{
			score: map[string]int{},
			winCalls: nil,
		}

		server := &PlayerServer{
			Store: &store,
		}
		request := httptest.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records win when Post", func(t *testing.T) {
		store := StubPlayerStore{
			score: map[string]int{},
			winCalls: nil,
		}

		server := &PlayerServer{
			Store: &store,
		}

		player := "Pepper"

		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner, got %q want %q", store.winCalls[0], player)
		}
	})

}
