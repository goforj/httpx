//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Put issues a PUT request using the provided client.

	// Example: typed PUT
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type User struct {
		Name string `json:"name"`
	}

	c := httpx.New()
	res, err := httpx.Put[UpdateUser, User](c, "https://httpbin.org/put", UpdateUser{Name: "Ana"})
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps User
	// #User {
	//   Name => "Ana" #string
	// }
}
