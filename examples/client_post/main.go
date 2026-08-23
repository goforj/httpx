//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// Post issues a POST request and decodes its response.

	// Example: typed POST
	type CreateUser struct {
		Name string `json:"name"`
	}
	type CreateUserResponse struct {
		JSON CreateUser `json:"json"`
	}

	c := httpx.New()
	res, err := c.Post[CreateUserResponse]("https://httpbin.org/post", CreateUser{Name: "Ana"})
	if err != nil {
		return
	}
	fmt.Println(res.JSON.Name)
	// Ana
}
