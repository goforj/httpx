//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/httpx"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	// CookieJar sets the cookie jar for the client.

	// Example: set cookie jar and seed cookies
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse("https://httpbin.org")
	jar.SetCookies(u, []*http.Cookie{
		{Name: "session", Value: "abc123"},
	})
	c := httpx.New(httpx.CookieJar(jar))
	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/cookies")
	_ = err
	httpx.Dump(res) // dumps map[string]any
	// #map[string]interface {} {
	//   cookies => #map[string]interface {} {
	//     session => "abc123" #string
	//   }
	// }
}
