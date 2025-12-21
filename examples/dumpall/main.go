//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// DumpAll enables req's client-level dump output for all requests.

	// Example: dump every request and response
	c := httpx.New(httpx.DumpAll())
	_ = c
}
