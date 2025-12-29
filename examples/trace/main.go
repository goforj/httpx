//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Trace enables req's request-level trace output.

	// Example: trace a single request
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Trace())
}
