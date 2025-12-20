//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Bearer sets the Authorization header with a bearer token.

	// Example: bearer auth
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.Bearer("token"))
}
