//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Patch issues a PATCH request using the provided client.

	// Example: typed PATCH
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res := httpx.Patch[UpdateUser, User](c, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
	_, _ = res.Body, res.Err // Body is User
}
