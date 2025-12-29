//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Bearer sets the Authorization header with a bearer token.

	// Example: bearer auth
	// Apply to all requests
	c := httpx.New(httpx.Bearer("token"))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Authorization => "Bearer token" #string
	//   }
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.Bearer("token"))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Authorization => "Bearer token" #string
	//   }
	// }
}
