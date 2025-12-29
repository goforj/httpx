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

	ctx := context.Background()
	c := httpx.New()
	res, err := httpx.DeleteCtx[DeleteResponse](c, ctx, "https://httpbin.org/delete")
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps DeleteResponse
	// #DeleteResponse {
	//   OK => true #bool
	// }
}
