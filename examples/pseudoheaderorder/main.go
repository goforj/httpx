//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// PseudoHeaderOrder sets the HTTP/2 pseudo header order for requests.

	// Example: set pseudo header order
	c := httpx.New(httpx.PseudoHeaderOrder(":method", ":authority", ":scheme", ":path"))
	_ = c
}
