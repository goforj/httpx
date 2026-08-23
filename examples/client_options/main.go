//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/httpx/v2"
)

func main() {
	// Options issues an OPTIONS request and decodes its response.

	// Example: client OPTIONS request
	c := httpx.New()
	_, err := c.Options[[]byte]("https://httpbin.org/get")
	fmt.Println(err == nil)
	// true
}
