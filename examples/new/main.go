//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// New creates a client with opinionated defaults and optional overrides.

	// Example: custom base URL and timeout
	c := httpx.New(
		httpx.WithBaseURL("https://api.example.com"),
		httpx.WithTimeout(5*time.Second),
	)
	_ = c
}
