//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// PutCtx issues a PUT request using the provided client and context.

	// Example: context-aware PUT
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	ctx := context.Background()
	c := httpx.New()
	res, err := httpx.PutCtx[UpdateUser, User](c, ctx, "https://httpbin.org/put", UpdateUser{Name: "Ana"})
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps User
	// #User {
	//   Name => "Ana" #string
	// }
}
