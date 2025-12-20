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
	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Body(Payload{Name: "Ana"}))
}
