//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"time"
)

func main() {
	// RetryInterval sets a custom retry interval function for a request.

	// Example: custom retry interval
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * 100 * time.Millisecond
	}))
}
