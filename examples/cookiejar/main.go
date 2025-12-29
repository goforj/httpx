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
	u, _ := url.Parse("https://example.com")
	jar.SetCookies(u, []*http.Cookie{
		{Name: "session", Value: "abc123"},
	})
	c := httpx.New(httpx.CookieJar(jar))
	_ = c
}
