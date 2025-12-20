//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Query adds query parameters as key/value pairs.

	// Example: add query params
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go", "ok", "1"))
}
