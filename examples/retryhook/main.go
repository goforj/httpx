//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// RetryHook registers a retry hook for a request.

	// Example: hook on retry
	// Apply to all requests
	c := httpx.New(httpx.RetryHook(func(_ *req.Response, _ error) {}))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
}
