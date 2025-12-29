//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Auth sets the Authorization header using a scheme and token.

	// Example: custom auth scheme
	// Apply to all requests
	c := httpx.New(httpx.Auth("Token", "abc123"))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.Auth("Token", "abc123"))
}
