//go:build ignore
// +build ignore

package main

import (
	"context"
	"github.com/goforj/httpx"
)

func main() {
	// PatchCtx issues a PATCH request using the provided client and context.

	// Example: context-aware PATCH
	type UpdateUser struct {
		Name string `json:"name"`
	}
	type UpdateUserResponse struct {
		JSON UpdateUser `json:"json"`
	}

	ctx := context.Background()
	c := httpx.New()
	res, err := httpx.PatchCtx[UpdateUser, UpdateUserResponse](c, ctx, "https://httpbin.org/patch", UpdateUser{Name: "Ana"})
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
