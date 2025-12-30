//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Query adds query parameters as key/value pairs.

	// Example: add query params
	c := httpx.New()
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.Query("q", "search"))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   args => #map[string]interface {} {
	//     q => "search" #string
	//   }
	// }
}
