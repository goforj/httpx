//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// DeleteCtx issues a context-aware DELETE request and decodes its response.

	// Example: client context-aware DELETE
	type DeleteResponse struct {
		URL string `json:"url"`
	}
	ctx := context.Background()
	c := httpx.New()
	res, err := c.DeleteCtx[DeleteResponse](ctx, "https://httpbin.org/delete")
	if err != nil {
		return
	}
	fmt.Println(res.URL)
	// https://httpbin.org/delete
}
