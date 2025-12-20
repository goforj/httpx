//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// Timeout sets a per-request timeout using context cancellation.

	// Example: per-request timeout
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Timeout(2*time.Second))
}
