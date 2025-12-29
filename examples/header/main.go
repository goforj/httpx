//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Header sets a header on a request or client.

	// Example: apply a header
	// Apply to all requests
	c := httpx.New(httpx.Header("X-Trace", "1"))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     X-Trace => "1" #string
	//   }
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.Header("X-Trace", "1"))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     X-Trace => "1" #string
	//   }
	// }
}
