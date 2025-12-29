//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// OptionsCtx issues an OPTIONS request using the provided client and context.

	// Example: context-aware OPTIONS
	c := httpx.New()
	ctx := context.Background()
	res := httpx.OptionsCtx[string](c, ctx, "https://example.com")
	_ = res
}
