//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Paths sets multiple path parameters from a map.

	// Example: path parameters
	c := httpx.New()
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/anything/{org}/users/{id}", httpx.Paths(map[string]any{
		"org": "goforj",
		"id":  7,
	}))
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/anything/goforj/users/7" #string
	// }
}
