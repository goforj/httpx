//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Path sets a path parameter by name.

	// Example: path parameter
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/anything/{id}", httpx.Path("id", 42))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/anything/42" #string
	// }
}
