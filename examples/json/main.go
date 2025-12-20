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
	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.JSON(Payload{Name: "Ana"}))
}
