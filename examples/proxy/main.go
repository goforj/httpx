//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Proxy sets a proxy URL for the client.

	// Example: set proxy URL
	c := httpx.New(httpx.Proxy("http://localhost:8080"))
	_ = c
}
