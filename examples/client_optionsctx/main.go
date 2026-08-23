//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// OptionsCtx issues a context-aware OPTIONS request and decodes its response.

	// Example: client context-aware OPTIONS
	ctx := context.Background()
	c := httpx.New()
	_, err := c.OptionsCtx[[]byte](ctx, "https://httpbin.org/get")
	fmt.Println(err == nil)
	// true
}
