//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// HTTP2ConnectionFlow sets the HTTP/2 connection flow control window increment.

	// Example: customize HTTP/2 connection flow
	c := httpx.New(httpx.HTTP2ConnectionFlow(1_048_576))
	_ = c
}
