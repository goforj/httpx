//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Headers sets multiple headers on a request or client.

	// Example: apply headers
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Headers(map[string]string{
		"X-Trace": "1",
		"Accept":  "application/json",
	}))
}
