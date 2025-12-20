//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Dump enables req's request-level dump output.

	// Example: dump a single request
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Dump())
}
