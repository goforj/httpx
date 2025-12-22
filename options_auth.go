package httpx

import (
	"encoding/base64"

	"github.com/imroc/req/v3"
)

// Auth sets the Authorization header using a scheme and token.
// @group Auth
// Applies to both client defaults and request-time overrides.
//
// Example: custom auth scheme
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Auth("Token", "abc123"))
func Auth(scheme, token string) OptionBuilder {
	return OptionBuilder{}.Auth(scheme, token)
}

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
// Applies to both client defaults and request-time overrides.
//
// Example: bearer auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Bearer("token"))
func Bearer(token string) OptionBuilder {
	return OptionBuilder{}.Bearer(token)
}

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
// Applies to both client defaults and request-time overrides.
//
// Example: basic auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Basic("user", "pass"))
func Basic(user, pass string) OptionBuilder {
	return OptionBuilder{}.Basic(user, pass)
}

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
