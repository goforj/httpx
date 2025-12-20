//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// RetryBackoff sets a capped exponential backoff retry interval for a request.

	// Example: request retry backoff
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
}
