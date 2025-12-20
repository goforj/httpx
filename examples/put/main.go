//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Put issues a PUT request using the provided client.

	// Example: typed PUT
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res := httpx.Put[UpdateUser, User](c, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
	_, _ = res.Body, res.Err
}
