//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Bearer sets the Authorization header with a bearer token.

	// Example: bearer auth
	// Apply to all requests
	c := httpx.New(httpx.Bearer("token"))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.Bearer("token"))
}
