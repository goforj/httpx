//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// PatchCtx issues a PATCH request using the provided client and context.

	// Example: context-aware PATCH
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	ctx := context.Background()
	res := httpx.PatchCtx[UpdateUser, User](c, ctx, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
	_, _ = res.Body, res.Err // Body is User
}
