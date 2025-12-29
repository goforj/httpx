//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// TraceAll enables req's client-level trace output for all requests.

	// Example: trace all requests
	c := httpx.New(httpx.TraceAll())
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
