//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// WithMiddleware adds request middleware to the client.

	// Example: add request middleware
	c := httpx.New(httpx.WithMiddleware(func(_ *req.Client, r *req.Request) error {
		r.SetHeader("X-Trace", "1")
		return nil
	}))
	_ = c
}
