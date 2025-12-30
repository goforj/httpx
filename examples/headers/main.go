//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Headers sets multiple headers on a request or client.

	// Example: apply headers
	// Apply to all requests
	c := httpx.New(httpx.Headers(map[string]string{
		"X-Trace": "1",
		"Accept":  "application/json",
	}))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Accept => "application/json" #string
	//     X-Trace => "1" #string
	//   }
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.Headers(map[string]string{
		"X-Trace": "1",
		"Accept":  "application/json",
	}))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Accept => "application/json" #string
	//     X-Trace => "1" #string
	//   }
	// }
}
