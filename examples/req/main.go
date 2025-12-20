//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Req returns the underlying req client for advanced usage.

	// Example: enable req debugging
	c := httpx.New()
	c.Req().EnableDumpEachRequest()
}
