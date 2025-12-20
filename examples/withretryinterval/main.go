//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"time"
)

func main() {
	// WithRetryInterval sets a custom retry interval function for the client.

	// Example: client retry interval
	c := httpx.New(httpx.WithRetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * 100 * time.Millisecond
	}))
	_ = c
}
