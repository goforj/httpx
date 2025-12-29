//go:build ignore
// +build ignore

package main

import "github.com/goforj/httpx"

func main() {
	// Patch issues a PATCH request using the provided client.

	// Example: typed PATCH
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type UpdateUserResponse struct {
		JSON UpdateUser `json:"json"`
	}

	c := httpx.New()
	res, err := httpx.Patch[UpdateUser, UpdateUserResponse](c, "https://httpbin.org/patch", UpdateUser{Name: "Ana"})
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
