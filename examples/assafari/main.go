//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// AsSafari applies the Safari browser profile (headers including User-Agent, TLS, and HTTP/2 behavior).

	// Example: use a Safari profile
	_ = httpx.New(httpx.AsSafari())
}
