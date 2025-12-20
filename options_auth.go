package httpx

import "github.com/imroc/req/v3"

// Auth sets the Authorization header using a scheme and token.
// @group Auth
//
// Example: custom auth scheme
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Auth("Token", "abc123"))
func Auth(scheme, token string) Option {
	return func(r *req.Request) {
		r.SetHeader("Authorization", scheme+" "+token)
	}
}

// Bearer sets the Authorization header with a bearer token.
// @group Auth
//
// Example: bearer auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Bearer("token"))
func Bearer(token string) Option {
	return func(r *req.Request) {
		r.SetBearerAuthToken(token)
	}
}

// Basic sets HTTP basic authentication headers.
// @group Auth
//
// Example: basic auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Basic("user", "pass"))
func Basic(user, pass string) Option {
	return func(r *req.Request) {
		r.SetBasicAuth(user, pass)
	}
}
