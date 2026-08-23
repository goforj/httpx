//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// Patch issues a PATCH request and decodes its response.

	// Example: typed PATCH
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type UpdateUserResponse struct {
		JSON UpdateUser `json:"json"`
	}

	c := httpx.New()
	res, err := c.Patch[UpdateUserResponse]("https://httpbin.org/patch", UpdateUser{Name: "Ana"})
	if err != nil {
		return
	}
	fmt.Println(res.JSON.Name)
	// Ana
}
