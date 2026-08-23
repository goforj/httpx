//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// PatchCtx issues a context-aware PATCH request and decodes its response.

	// Example: client context-aware PATCH
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type UpdateUserResponse struct {
		JSON UpdateUser `json:"json"`
	}
	ctx := context.Background()
	c := httpx.New()
	res, err := c.PatchCtx[UpdateUserResponse](ctx, "https://httpbin.org/patch", UpdateUser{Name: "Ana"})
	if err != nil {
		return
	}
	fmt.Println(res.JSON.Name)
	// Ana
}
