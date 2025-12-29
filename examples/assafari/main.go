//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// AsSafari applies the Safari browser profile (headers, TLS, and HTTP/2 behavior).

	// Example: use a Safari profile
	c := httpx.New(httpx.AsSafari())
	_ = c
}
