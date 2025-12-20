//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// WithBaseURL sets a base URL on the client.

	// Example: client base URL
	c := httpx.New(httpx.WithBaseURL("https://api.example.com"))
	_ = c
}
