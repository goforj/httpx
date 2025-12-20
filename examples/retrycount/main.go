//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// RetryCount enables retry for a request and sets the maximum retry count.

	// Example: request retry count
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCount(2))
}
