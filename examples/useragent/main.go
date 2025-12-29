//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// UserAgent sets the User-Agent header on a request or client.

	// Example: set a User-Agent
	// Apply to all requests
	c := httpx.New(httpx.UserAgent("my-app/1.0"))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.UserAgent("my-app/1.0"))
}
