//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Delete issues a DELETE request using the provided client.

	// Example: typed DELETE
	type DeleteResponse struct {
		OK bool `json:"ok"`
	}

	c := httpx.New()
	res := httpx.Delete[DeleteResponse](c, "https://api.example.com/users/1")
	_, _ = res.Body, res.Err
}
