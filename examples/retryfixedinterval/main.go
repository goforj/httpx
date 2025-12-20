//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// RetryFixedInterval sets a fixed retry interval for a request.

	// Example: request retry interval
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
}
