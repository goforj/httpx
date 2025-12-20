package httpx

import "github.com/imroc/req/v3"

// Result contains the decoded response body and response metadata.
//
// Example: access body and response
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res := httpx.Get[User](c, "https://example.com/users/1")
//	if res.Err != nil {
//		return
//	}
//	_ = res.Body
//	_ = res.Response
type Result[T any] struct {
	Body     T
	Response *req.Response
	Err      error
}

// Option configures a request before it is sent.
//
// Example: custom header
//
//	c := httpx.New()
//	res := httpx.Get[string](c, "https://example.com", httpx.Header("X-Trace", "1"))
//	_ = res
type Option func(*req.Request)

// ClientOption configures a client instance.
//
// Example: configure client
//
//	c := httpx.New(httpx.WithTimeout(3 * time.Second))
//	_ = c
type ClientOption func(*Client)

// ErrorMapper customizes error creation for non-2xx responses.
//
// Example: map error responses
//
//	c := httpx.New(httpx.WithErrorMapper(func(resp *req.Response) error {
//		return fmt.Errorf("status %d", resp.StatusCode)
//	}))
//	_ = c
type ErrorMapper func(*req.Response) error
