//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// EnableDump enables req's request-level dump output.

	// Example: dump a single request
	c := httpx.New()
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.EnableDump())
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
