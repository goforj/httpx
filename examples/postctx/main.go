//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// PostCtx issues a POST request using the provided client and context.

	// Example: context-aware POST
	type CreateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	ctx := context.Background()
	res := httpx.PostCtx[CreateUser, User](c, ctx, "https://api.example.com/users", CreateUser{Name: "Ana"})
	_, _ = res.Body, res.Err
}
