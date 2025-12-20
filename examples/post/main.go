//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Post issues a POST request using the provided client.

	// Example: typed POST
	type CreateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res := httpx.Post[CreateUser, User](c, "https://api.example.com/users", CreateUser{Name: "Ana"})
	_, _ = res.Body, res.Err
}
