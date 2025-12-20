//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Query adds a single query parameter.

	// Example: add query param
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go"))
}
