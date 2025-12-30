//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"net/http"
)

func main() {
	// ProxyFunc sets a proxy function for the client.

	// Example: set proxy function
	c := httpx.New(httpx.ProxyFunc(http.ProxyFromEnvironment))
	res, _ := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   uuid => "<uuid>" #string
	// }
}
