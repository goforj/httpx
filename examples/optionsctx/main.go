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
	ctx := context.Background()
	c := httpx.New()
	_, err := httpx.OptionsCtx[string](c, ctx, "https://httpbin.org/get")
	if err != nil {
		return
	}
}
