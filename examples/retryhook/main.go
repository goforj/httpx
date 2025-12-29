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
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.RetryHook(func(_ *req.Response, _ error) {}))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
