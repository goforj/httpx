//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// PutCtx issues a context-aware PUT request and decodes its response.

	// Example: client context-aware PUT
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type UpdateUserResponse struct {
		JSON UpdateUser `json:"json"`
	}
	ctx := context.Background()
	c := httpx.New()
	res, err := c.PutCtx[UpdateUserResponse](ctx, "https://httpbin.org/put", UpdateUser{Name: "Ana"})
	if err != nil {
		return
	}
	fmt.Println(res.JSON.Name)
	// Ana
}
