//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// Get issues a GET request and decodes its response.

	// Example: bind to a struct
	type GetResponse struct {
		URL string `json:"url"`
	}

	c := httpx.New()
	res, err := c.Get[GetResponse]("https://httpbin.org/get")
	fmt.Println(err == nil, res.URL)
	// true https://httpbin.org/get
}
