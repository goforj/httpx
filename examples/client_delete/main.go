//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// Delete issues a DELETE request and decodes its response.

	// Example: typed DELETE
	type DeleteResponse struct {
		URL string `json:"url"`
	}

	c := httpx.New()
	res, err := c.Delete[DeleteResponse]("https://httpbin.org/delete")
	if err != nil {
		return
	}
	fmt.Println(res.URL)
	// https://httpbin.org/delete
}
