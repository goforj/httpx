//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// PostCtx issues a context-aware POST request and decodes its response.

	// Example: client context-aware POST
	type CreateUser struct {
		Name string `json:"name"`
	}
	type CreateUserResponse struct {
		JSON CreateUser `json:"json"`
	}
	ctx := context.Background()
	c := httpx.New()
	res, err := c.PostCtx[CreateUserResponse](ctx, "https://httpbin.org/post", CreateUser{Name: "Ana"})
	if err != nil {
		return
	}
	fmt.Println(res.JSON.Name)
	// Ana
}
