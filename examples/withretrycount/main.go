//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// WithRetryCount enables retry for the client and sets the maximum retry count.

	// Example: client retry count
	c := httpx.New(httpx.WithRetryCount(2))
	_ = c
}
