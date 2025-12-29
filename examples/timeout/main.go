//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// Timeout sets a per-request timeout using context cancellation.

	// Example: timeout
	// Apply to all requests
	c := httpx.New(httpx.Timeout(2 * time.Second))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/delay/2")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/delay/2" #string
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/delay/2", httpx.Timeout(2*time.Second))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/delay/2" #string
	// }
}
