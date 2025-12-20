//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Paths sets multiple path parameters.

	// Example: multiple path parameters
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	_ = httpx.Get[User](c, "https://example.com/orgs/{org}/users/{id}", httpx.Paths(map[string]any{
		"org": "goforj",
		"id":  42,
	}))
}
