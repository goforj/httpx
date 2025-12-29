//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// AsChrome applies the Chrome browser profile (headers, TLS, and HTTP/2 behavior).

	// Example: use a Chrome profile
	c := httpx.New(httpx.AsChrome())
	_ = c
}
