//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Headers sets multiple headers on a request or client.

	// Example: apply headers
	// Apply to all requests
	c := httpx.New(httpx.Headers(map[string]string{
		"X-Trace": "1",
		"Accept":  "application/json",
	}))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.Headers(map[string]string{
		"X-Trace": "1",
		"Accept":  "application/json",
	}))
}
