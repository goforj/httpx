//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// JSON sets the request body as JSON.

	// Example: force JSON body
	type Payload struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res, _ := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.JSON(Payload{Name: "Ana"}))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   json => #map[string]interface {} {
	//     name => "Ana" #string
	//   }
	// }
}
