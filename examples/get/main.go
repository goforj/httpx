//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/httpx"
)

func main() {
	// Get issues a GET request using the provided client.

	// Example: fetch GitHub pull requests (typed)
	type PullRequest struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
	}

	c := httpx.New(httpx.Header("Accept", "application/vnd.github+json"))
	res := httpx.Get[[]PullRequest](c, "https://api.github.com/repos/goforj/httpx/pulls")
	if res.Err != nil {
		return
	}
	godump.Dump(res.Body)

	// Example: bind to a string body
	c2 := httpx.New()
	res2 := httpx.Get[string](c2, "https://httpbin.org/uuid")
	_, _ = res2.Body, res2.Err // Body is string
}
