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
	c := httpx.New()
	_ = httpx.Get[string](c, "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
}
