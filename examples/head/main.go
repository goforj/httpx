//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// Head issues a HEAD request using the provided client.

	// Example: HEAD request
	c := httpx.New()
	_, err := httpx.Head[string](c, "https://httpbin.org/get")
	if err != nil {
		return
	}
}
