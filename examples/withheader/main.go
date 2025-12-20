//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// WithHeader sets a default header for all requests.

	// Example: client header
	c := httpx.New(httpx.WithHeader("X-Trace", "1"))
	_ = c
}
