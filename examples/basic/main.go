//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Basic sets HTTP basic authentication headers.

	// Example: basic auth
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Basic("user", "pass"))
}
