//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Header sets a header on a request.

	// Example: apply a header
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Header("X-Trace", "1"))
}
