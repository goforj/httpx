//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"github.com/imroc/req/v3"
	"time"
)

func main() {
	// RetryInterval sets a custom retry interval function for a request.
	// Overrides the req default interval (fixed 100ms).

	// Example: custom retry interval
	// Apply to all requests
	c := httpx.New(httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * 100 * time.Millisecond
	}))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }

	// Apply to a single request
	res, _ = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * 100 * time.Millisecond
	}))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
