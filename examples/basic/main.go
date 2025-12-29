//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Basic sets HTTP basic authentication headers.

	// Example: basic auth
	// Apply to all requests
	c := httpx.New(httpx.Basic("user", "pass"))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.Basic("user", "pass"))
}
