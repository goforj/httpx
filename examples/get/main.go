//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Get issues a GET request using the provided client.

	// Example: basic GET
	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get")
	if err != nil {
		return
	}
	httpx.Dump(res)
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }

	// Example: bind to a string body
	resText, err := httpx.Get[string](c, "https://httpbin.org/get")
	if err != nil {
		return
	}
	_ = resText // resText is string
}
