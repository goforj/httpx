//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// RetryBackoff sets a capped exponential backoff retry interval for a request.

	// Example: retry backoff
	// Apply to all requests
	c := httpx.New(httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
}
