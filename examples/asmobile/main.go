//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// AsMobile applies a mobile Chrome-like profile (headers, TLS, and HTTP/2 behavior).

	// Example: use a mobile profile
	c := httpx.New(httpx.AsMobile())
	_ = c
}
