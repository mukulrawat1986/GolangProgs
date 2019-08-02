package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	// create a child of main context and add cancellation to it
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel() // to release resources if sleepAndTalk completes before timeout
	sleepAndTalk(ctx, 5*time.Second, "Hello, World")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
