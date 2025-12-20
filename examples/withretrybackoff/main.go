//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// WithRetryBackoff sets a capped exponential backoff retry interval for the client.

	// Example: client retry backoff
	c := httpx.New(httpx.WithRetryBackoff(100*time.Millisecond, 2*time.Second))
	_ = c
}
