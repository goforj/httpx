//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Get issues a GET request using the provided client.

	// Example: bind to a struct
	type GetResponse struct {
		URL string `json:"url"`
	}

	c := httpx.New()
	res, _ := httpx.Get[GetResponse](c, "https://httpbin.org/get")
	httpx.Dump(res)
	// #GetResponse {
	//   URL => "https://httpbin.org/get" #string
	// }

	// Example: bind to a string body
	resString, _ := httpx.Get[string](c, "https://httpbin.org/uuid")
	println(resString) // dumps string
	// {
	//   "uuid": "becbda6d-9950-4966-ae23-0369617ba065"
	// }
}
