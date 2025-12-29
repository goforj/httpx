//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// RetryCondition sets the retry condition for a request.

	// Example: retry on 503
	// Apply to all requests
	c := httpx.New(httpx.RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}))
	httpx.Get[string](c, "https://example.com")

	// Apply to a single request
	httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}))
}
