package _context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}

func TestHandler(t *testing.T) {

	data := "Hello, World"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{
			response:  "Hello, World",
			cancelled: false,
			t: t,
		}
		svr := Server(store)


		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q want %q", response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{
			response:  data,
			cancelled: false,
			t: t,
		}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)

		request = request.WithContext(ctx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}
