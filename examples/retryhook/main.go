//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
)

func main() {
	// RetryHook registers a retry hook for a request.
	// Runs before each retry attempt; no hooks are configured by default.

	// Example: hook on retry
	// Apply to all requests
	c := httpx.New(httpx.RetryHook(func(_ *req.Response, _ error) {}))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryHook(func(_ *req.Response, _ error) {}))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
