//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Options issues an OPTIONS request using the provided client.

	// Example: OPTIONS request
	c := httpx.New()
	res := httpx.Options[string](c, "https://example.com")
	_ = res
}
