//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// EnableDump enables req's request-level dump output.

	// Example: dump a single request
	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.EnableDump())
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
