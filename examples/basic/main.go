//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Basic sets HTTP basic authentication headers.

	// Example: basic auth
	// Apply to all requests
	c := httpx.New(httpx.Basic("user", "pass"))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Authorization => "Basic dXNlcjpwYXNz" #string
	//   }
	// }

	// Apply to a single request
	res, _ = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.Basic("user", "pass"))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     Authorization => "Basic dXNlcjpwYXNz" #string
	//   }
	// }
}
