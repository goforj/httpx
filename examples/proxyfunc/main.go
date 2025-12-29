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
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   url => "https://httpbin.org/get" #string
	// }
}
