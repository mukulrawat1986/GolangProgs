package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// create a child of main context and add cancellation to it
	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(time.Second, cancel)
	sleepAndTalk(ctx, 5*time.Second, "Hello, World")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
