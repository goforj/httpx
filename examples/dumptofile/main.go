//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// DumpToFile enables req's request-level dump output to a file path.

	// Example: dump to a file
	c := httpx.New()
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.DumpToFile("httpx.dump"))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
