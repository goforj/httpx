//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// Head issues a HEAD request and decodes its response using the established response contract.

	// Example: client HEAD request
	c := httpx.New()
	_, err := c.Head[[]byte]("https://httpbin.org/get")
	fmt.Println(err == nil)
	// true
}
