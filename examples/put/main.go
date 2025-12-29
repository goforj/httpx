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
	type UpdateUserResponse struct {
		JSON UpdateUser `json:"json"`
	}

	c := httpx.New()
	res, err := httpx.Put[UpdateUser, UpdateUserResponse](c, "https://httpbin.org/put", UpdateUser{Name: "Ana"})
	if err != nil {
		return
	}
	httpx.Dump(res) // dumps UpdateUserResponse
	// #UpdateUserResponse {
	//   JSON => #UpdateUser {
	//     Name => "Ana" #string
	//   }
	// }
}
