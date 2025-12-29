//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// DeleteCtx issues a DELETE request using the provided client and context.

	// Example: context-aware DELETE
	type DeleteResponse struct {
		OK bool `json:"ok"`
	}

	c := httpx.New()
	ctx := context.Background()
	res := httpx.DeleteCtx[DeleteResponse](c, ctx, "https://api.example.com/users/1")
	_, _ = res.Body, res.Err // Body is DeleteResponse
}
