//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// MultipartBoundary overrides the default multipart boundary generator.

	// Example: customize multipart boundary
	c := httpx.New(httpx.MultipartBoundary(func() string { return "boundary" }))
	_ = c
}
