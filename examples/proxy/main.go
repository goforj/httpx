//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Proxy sets a proxy URL for the client.

	// Example: set proxy URL
	c := httpx.New(httpx.Proxy("http://localhost:8080"))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// map[string]interface {}(nil)
}
