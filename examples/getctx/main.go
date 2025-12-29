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
	type GetResponse struct {
		URL string `json:"url"`
	}

	ctx := context.Background()
	c := httpx.New()
	res, err := httpx.GetCtx[GetResponse](c, ctx, "https://httpbin.org/get")
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps GetResponse
	// #GetResponse {
	//   URL => "https://httpbin.org/get" #string
	// }
}
