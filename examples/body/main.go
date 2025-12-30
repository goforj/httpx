//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Body sets the request body and infers JSON for structs and maps.

	// Example: send JSON body with inference
	type Payload struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res, _ := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.Body(Payload{Name: "Ana"}))
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   json => #map[string]interface {} {
	//     name => "Ana" #string
	//   }
	// }
}
