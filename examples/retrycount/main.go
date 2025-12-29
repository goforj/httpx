//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// RetryCount enables retry for a request and sets the maximum retry count.

	// Example: retry count
	// Apply to all requests
	c := httpx.New(httpx.RetryCount(2))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.RetryCount(2))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
