//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// HeadCtx issues a HEAD request using the provided client and context.

	// Example: context-aware HEAD
	ctx := context.Background()
	c := httpx.New()
	_, err := httpx.HeadCtx[string](c, ctx, "https://httpbin.org/get")
	if err != nil {
		return
	}
}
