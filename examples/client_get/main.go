//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx/v2"

func main() {
	// Get issues a GET request and decodes its response.

	// Example: bind to a struct
	type GetResponse struct {
		URL string `json:"url"`
	}

	c := httpx.New()
	res, _ := c.Get[GetResponse]("https://httpbin.org/get")
	httpx.Dump(res) // dumps GetResponse
}
