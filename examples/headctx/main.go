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
	c := httpx.New()
	ctx := context.Background()
	res := httpx.HeadCtx[string](c, ctx, "https://example.com")
	_ = res
}
