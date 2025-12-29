//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// GetCtx issues a GET request using the provided client and context.

	// Example: context-aware GET
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	ctx := context.Background()
	res := httpx.GetCtx[User](c, ctx, "https://api.example.com/users/1")
	_, _ = res.Body, res.Err // Body is User
}
