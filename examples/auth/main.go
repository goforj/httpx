//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Auth sets the Authorization header using a scheme and token.

	// Example: custom auth scheme
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Auth("Token", "abc123"))
}
