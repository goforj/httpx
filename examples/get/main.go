//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	type Request struct {
		Uuid string `json:"uuid"`
	}

	// client
	c := httpx.New()

	// request, binds result to Request
	res, _ := httpx.Get[Request](c, "https://httpbin.org/uuid")

	httpx.Dump(res)
	// #main.Request {
	//  +Uuid => "6101eccc-8f59-444f-9ccc-9c39a85d5da5" #string
	// }
}
