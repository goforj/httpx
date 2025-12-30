//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// RetryCount enables retry for a request and sets the maximum retry count.
	// Default behavior from req: retries are disabled (count = 0). When enabled,
	// retries happen only on request errors unless RetryCondition is set, and the
	// default interval is a fixed 100ms between attempts.

	// Example: retry count
	// Apply to all requests
	c := httpx.New(httpx.RetryCount(2))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }

	// Apply to a single request
	res, _ = httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.RetryCount(2))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
