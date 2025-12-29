//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// RetryCount enables retry for a request and sets the maximum retry count.

	// Example: retry count
	// Apply to all requests
	c := httpx.New(httpx.RetryCount(2))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryCount(2))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
