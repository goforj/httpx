//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Post issues a POST request using the provided client.

	// Example: typed POST
	type CreateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res, err := httpx.Post[CreateUser, User](c, "https://httpbin.org/post", CreateUser{Name: "Ana"})
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps User
	// #User {
	//   Name => "Ana" #string
	// }
}
