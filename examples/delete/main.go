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
	res, err := httpx.Delete[DeleteResponse](c, "https://httpbin.org/delete")
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps DeleteResponse
	// #DeleteResponse {
	//   OK => true #bool
	// }
}
