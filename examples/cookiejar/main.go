//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"net/http/cookiejar"
)

func main() {
	// CookieJar sets the cookie jar for the client.

	// Example: set cookie jar
	jar, _ := cookiejar.New(nil)
	c := httpx.New(httpx.CookieJar(jar))
	_ = c
}
