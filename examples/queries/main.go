//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Queries sets query parameters from a map.

	// Example: add query params
	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.Queries(map[string]string{
		"q":  "search",
		"ok": "1",
	}))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   args => #map[string]interface {} {
	//     ok => "1" #string
	//     q => "search" #string
	//   }
	// }
}
