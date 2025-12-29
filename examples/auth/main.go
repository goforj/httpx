//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Auth sets the Authorization header using a scheme and token.

	// Example: custom auth scheme
	// Apply to all requests
	c := httpx.New(httpx.Auth("Token", "abc123"))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Authorization => "Token abc123" #string
	//   }
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.Auth("Token", "abc123"))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Authorization => "Token abc123" #string
	//   }
	// }
}
