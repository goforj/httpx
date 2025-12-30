//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// BaseURL sets a base URL on the client.

	// Example: client base URL
	c := httpx.New(httpx.BaseURL("https://httpbin.org"))
	res, _ := httpx.Get[map[string]any](c, "/uuid")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
