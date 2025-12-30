//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// UserAgent sets the User-Agent header on a request or client.

	// Example: set a User-Agent
	// Apply to all requests
	c := httpx.New(httpx.UserAgent("my-app/1.0"))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     User-Agent => "my-app/1.0" #string
	//   }
	// }

	// Apply to a single request
	res, _ = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.UserAgent("my-app/1.0"))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   headers => #map[string]interface {} {
	//     User-Agent => "my-app/1.0" #string
	//   }
	// }
}
