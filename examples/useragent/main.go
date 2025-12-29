//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// UserAgent sets the User-Agent header on a request or client.

	// Example: set a User-Agent
	c := httpx.New(httpx.UserAgent("my-app/1.0"))
	_ = httpx.Get[string](c, "https://example.com")
}
