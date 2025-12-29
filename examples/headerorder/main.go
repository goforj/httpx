//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// HeaderOrder sets the header order for requests.

	// Example: set header order
	c := httpx.New(httpx.HeaderOrder("host", "user-agent", "accept"))
	_ = c
}
