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
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <- ctx.Done():
				s.t.Log("spy store got cancelled")
				return

				default:
					time.Sleep(10 * time.Millisecond)
					result +=  string(c)
			}
		}

		data <- result
	}()

	select {
	case <- ctx.Done():
		return "", ctx.Err()

		case res := <- data:
			return res, nil
	}
}


func TestHandler(t *testing.T) {

	data := "Hello, World"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{
			response:  "Hello, World",
			t: t,
		}
		svr := Server(store)


		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q want %q", response.Body.String(), data)
		}

	})

	//t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
	//	store := &SpyStore{
	//		response:  data,
	//		cancelled: false,
	//		t: t,
	//	}
	//	svr := Server(store)
	//
	//	request := httptest.NewRequest(http.MethodGet, "/", nil)
	//
	//	ctx, cancel := context.WithCancel(request.Context())
	//	time.AfterFunc(5 * time.Millisecond, cancel)
	//
	//	request = request.WithContext(ctx)
	//	response := httptest.NewRecorder()
	//
	//	svr.ServeHTTP(response, request)
	//
	//	store.assertWasCancelled()
	//})
}
