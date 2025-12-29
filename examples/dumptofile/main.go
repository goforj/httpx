//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// DumpToFile enables req's request-level dump output to a file path.

	// Example: dump to a file
	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.DumpToFile("httpx.dump"))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
