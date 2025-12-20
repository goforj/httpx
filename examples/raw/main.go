//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Raw returns the underlying req client for chaining raw requests.

	// Example: drop down to req
	c := httpx.New()
	resp, err := c.Raw().R().Get("https://httpbin.org/uuid")
	_, _ = resp, err
}
