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

	ctx := context.Background()
	c := httpx.New()
	res, err := httpx.PatchCtx[UpdateUser, User](c, ctx, "https://httpbin.org/patch", UpdateUser{Name: "Ana"})
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps User
	// #User {
	//   Name => "Ana" #string
	// }
}
