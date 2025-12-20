//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// WithTimeout sets the default timeout for the client.

	// Example: client timeout
	c := httpx.New(httpx.WithTimeout(3 * time.Second))
	_ = c
}
