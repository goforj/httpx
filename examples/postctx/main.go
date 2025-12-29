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

	ctx := context.Background()
	c := httpx.New()
	res, err := httpx.PostCtx[CreateUser, User](c, ctx, "https://httpbin.org/post", CreateUser{Name: "Ana"})
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps User
	// #User {
	//   Name => "Ana" #string
	// }
}
