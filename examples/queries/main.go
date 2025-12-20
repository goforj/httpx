//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Queries adds multiple query parameters.

	// Example: add query params
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com/search", httpx.Queries(map[string]string{
		"q":  "go",
		"ok": "1",
	}))
}
