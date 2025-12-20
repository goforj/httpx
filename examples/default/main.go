//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Default returns the shared default client.

	// Example: use default client for quick calls
	c := httpx.Default()
	_ = c
}
