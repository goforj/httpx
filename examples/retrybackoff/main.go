//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"time"
)

func main() {
	// RetryBackoff sets a capped exponential backoff retry interval for a request.

	// Example: retry backoff
	// Apply to all requests
	c := httpx.New(httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }

	// Apply to a single request
	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
