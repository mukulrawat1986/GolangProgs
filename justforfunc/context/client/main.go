package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// get the context
	ctx := context.Background()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// add context to our request
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	io.Copy(os.Stdout, res.Body)
}
