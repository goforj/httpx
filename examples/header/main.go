//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Header sets a header on a request or client.

	// Example: apply a header
	// Apply to all requests
	c := httpx.New(httpx.Header("X-Trace", "1"))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.Header("X-Trace", "1"))
}
