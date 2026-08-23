//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// GetCtx issues a context-aware GET request and decodes its response.

	// Example: client context-aware GET
	type GetResponse struct {
		URL string `json:"url"`
	}
	ctx := context.Background()
	c := httpx.New()
	res, err := c.GetCtx[GetResponse](ctx, "https://httpbin.org/get")
	fmt.Println(err == nil, res.URL)
	// true https://httpbin.org/get
}
