//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// RetryFixedInterval sets a fixed retry interval for a request.

	// Example: retry interval
	// Apply to all requests
	c := httpx.New(httpx.RetryFixedInterval(200 * time.Millisecond))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
}
