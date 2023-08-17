package clienthook_test

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cehnhy/go-snippets/internal/http/clienthook"
)

func ExampleNewTransport() {
	hook := func(req *http.Request) {
		ctx := req.Context()
		fmt.Println(ctx.Value("key"))
	}
	transport := clienthook.NewTransport(nil, hook)
	client := http.Client{
		Transport: transport,
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Output:
	// value
}
