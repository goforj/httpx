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
	_ = c
}
