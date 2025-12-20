//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// WithRetryFixedInterval sets a fixed retry interval for the client.

	// Example: client retry interval
	c := httpx.New(httpx.WithRetryFixedInterval(200 * time.Millisecond))
	_ = c
}
