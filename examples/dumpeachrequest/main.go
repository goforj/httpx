//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// DumpEachRequest enables request-level dumps for each request on the client.

	// Example: dump each request as it is sent
	c := httpx.New(httpx.DumpEachRequest())
	_ = c
}
