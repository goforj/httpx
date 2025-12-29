//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Head issues a HEAD request using the provided client.

	// Example: HEAD request
	c := httpx.New()
	res := httpx.Head[string](c, "https://example.com")
	_ = res
}
