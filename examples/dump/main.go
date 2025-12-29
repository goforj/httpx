//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Dump prints values using the bundled godump formatter.

	// Example: dump a response
	res, err := httpx.Get[map[string]any](httpx.Default(), "https://httpbin.org/uuid")
	_ = err
	httpx.Dump(res)
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
