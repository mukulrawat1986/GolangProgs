package log

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type key int

const requestIDKey = key(42)

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		fmt.Println("Could not find request id in context")
		return
	}

	log.Printf("[%d] %s\n", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
