package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request  := httptest.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	} )

	t.Run("return Floyd's score", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
