//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Path sets a path parameter by name.

	// Example: path parameter
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	_ = httpx.Get[User](c, "https://example.com/users/{id}", httpx.Path("id", 42))
}
