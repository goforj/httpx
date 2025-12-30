//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// RetryCondition sets the retry condition for a request.
	// Overrides the default behavior (retry only when a request error occurs).

	// Example: retry on 503
	// Apply to all requests
	c := httpx.New(httpx.RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/status/503")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// map[string]interface {}(nil)

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/status/503", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// map[string]interface {}(nil)
}
