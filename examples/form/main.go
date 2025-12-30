//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Form sets form data for the request.

	// Example: submit a form
	c := httpx.New()
	res, _ := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.Form(map[string]string{
		"name": "alice",
	}))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   form => #map[string]interface {} {
	//     name => "alice" #string
	//   }
	// }
}
