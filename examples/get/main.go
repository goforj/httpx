//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/httpx"
)

func main() {
	// Get issues a GET request using the provided client.

	// Example: fetch GitHub pull requests
	type PullRequest struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
	}

	c := httpx.New(httpx.Opts().Header("Accept", "application/vnd.github+json"))
	res := httpx.Get[[]PullRequest](c, "https://api.github.com/repos/goforj/httpx/pulls")
	if res.Err != nil {
		return
	}
	godump.Dump(res.Body)
}
