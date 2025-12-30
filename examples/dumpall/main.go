//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// DumpAll enables req's client-level dump output for all requests.

	// Example: dump every request and response
	c := httpx.New(httpx.DumpAll())
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
