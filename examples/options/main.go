//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Options issues an OPTIONS request using the provided client.

	// Example: OPTIONS request
	c := httpx.New()
	_, err := httpx.Options[string](c, "https://httpbin.org/get")
	if err != nil {
		return
	}
}
