//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// HeadCtx issues a context-aware HEAD request and decodes its response using the established response contract.

	// Example: client context-aware HEAD
	ctx := context.Background()
	c := httpx.New()
	_, err := c.HeadCtx[[]byte](ctx, "https://httpbin.org/get")
	fmt.Println(err == nil)
	// true
}
