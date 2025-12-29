//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// Timeout sets a per-request timeout using context cancellation.

	// Example: timeout
	// Apply to all requests
	c := httpx.New(httpx.Timeout(2 * time.Second))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.Timeout(2*time.Second))
}
