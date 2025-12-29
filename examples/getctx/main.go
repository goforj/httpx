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

	ctx := context.Background()
	c := httpx.New()
	res, err := httpx.GetCtx[User](c, ctx, "https://httpbin.org/get")
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps User
	// #User {
	//   Name => "Ana" #string
	// }
}
