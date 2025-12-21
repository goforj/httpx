package httpx

import (
	"encoding/base64"

	"github.com/imroc/req/v3"
)

// Auth sets the Authorization header using a scheme and token.
// @group Auth
func (b OptionBuilder) Auth(scheme, token string) OptionBuilder {
	value := scheme + " " + token
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonHeader("Authorization", value)
		},
		func(r *req.Request) {
			r.SetHeader("Authorization", value)
		},
	))
}

// Bearer sets the Authorization header with a bearer token.
// @group Auth
func (b OptionBuilder) Bearer(token string) OptionBuilder {
	value := "Bearer " + token
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonHeader("Authorization", value)
		},
		func(r *req.Request) {
			r.SetBearerAuthToken(token)
		},
	))
}

// Basic sets HTTP basic authentication headers.
// @group Auth
func (b OptionBuilder) Basic(user, pass string) OptionBuilder {
	value := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass))
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonHeader("Authorization", value)
		},
		func(r *req.Request) {
			r.SetBasicAuth(user, pass)
		},
	))
}
