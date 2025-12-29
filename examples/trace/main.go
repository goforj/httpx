//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Trace enables req's request-level trace output.

	// Example: trace a single request
	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.Trace())
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
