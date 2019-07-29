package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare speeeds of servers, returning url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Duration(20 * time.Millisecond))
		fastServer := makeDelayedServer(time.Duration(0 * time.Millisecond))

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return an error if a server doesn't respond within 10 sec", func(t *testing.T) {
		serverA := makeDelayedServer(time.Duration(11 * time.Second))
		serverB := makeDelayedServer(time.Duration(12 * time.Second))

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
