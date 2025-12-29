//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// AsFirefox applies the Firefox browser profile (headers including User-Agent, TLS, and HTTP/2 behavior).

	// Example: use a Firefox profile
	c := httpx.New(httpx.AsFirefox())
	_ = c
}
