package httpx

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/imroc/req/v3"
)

// Header sets a header on a request.
// @group Request Options
//
// Example: apply a header
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Header("X-Trace", "1"))
func Header(key, value string) Option {
	return func(r *req.Request) {
		r.SetHeader(key, value)
	}
}

// Headers sets multiple headers on a request.
// @group Request Options
//
// Example: apply headers
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Headers(map[string]string{
//		"X-Trace": "1",
//		"Accept":  "application/json",
//	}))
func Headers(values map[string]string) Option {
	return func(r *req.Request) {
		r.SetHeaders(values)
	}
}

// Query adds query parameters as key/value pairs.
// @group Request Options
//
// Example: add query params
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go", "ok", "1"))
func Query(kv ...string) Option {
	return func(r *req.Request) {
		if len(kv)%2 != 0 {
			panic("httpx: Query expects even number of key/value arguments")
		}
		for i := 0; i < len(kv); i += 2 {
			r.AddQueryParam(kv[i], kv[i+1])
		}
	}
}

// Queries adds multiple query parameters.
// @group Request Options
//
// Example: add query params
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/search", httpx.Queries(map[string]string{
//		"q":  "go",
//		"ok": "1",
//	}))
func Queries(values map[string]string) Option {
	return func(r *req.Request) {
		r.SetQueryParams(values)
	}
}

// Path sets a path parameter by name.
// @group Request Options
//
// Example: path parameter
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	_ = httpx.Get[User](c, "https://example.com/users/{id}", httpx.Path("id", 42))
func Path(key string, value any) Option {
	return func(r *req.Request) {
		r.SetPathParam(key, fmt.Sprint(value))
	}
}

// Paths sets multiple path parameters.
// @group Request Options
//
// Example: multiple path parameters
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	_ = httpx.Get[User](c, "https://example.com/orgs/{org}/users/{id}", httpx.Paths(map[string]any{
//		"org": "goforj",
//		"id":  42,
//	}))
func Paths(values map[string]any) Option {
	return func(r *req.Request) {
		params := make(map[string]string, len(values))
		for key, value := range values {
			params[key] = fmt.Sprint(value)
		}
		r.SetPathParams(params)
	}
}

// Body sets the request body and infers JSON for structs and maps.
// @group Request Options
//
// Example: send JSON body with inference
//
//	type Payload struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Body(Payload{Name: "Ana"}))
func Body(value any) Option {
	return func(r *req.Request) {
		setBody(r, value)
	}
}

// JSON sets the request body as JSON.
// @group Request Options
//
// Example: force JSON body
//
//	type Payload struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.JSON(Payload{Name: "Ana"}))
func JSON(value any) Option {
	return func(r *req.Request) {
		r.SetBodyJsonMarshal(value)
	}
}

// Form sets form data for the request.
// @group Request Options
//
// Example: submit a form
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Form(map[string]string{
//		"name": "Ana",
//	}))
func Form(values map[string]string) Option {
	return func(r *req.Request) {
		r.SetFormData(values)
	}
}

// Timeout sets a per-request timeout using context cancellation.
// @group Request Options
//
// Example: per-request timeout
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Timeout(2*time.Second))
func Timeout(d time.Duration) Option {
	return func(r *req.Request) {
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, d)
		r.SetContext(ctx)
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			cancel()
			return nil
		})
	}
}

// Before runs a hook before the request is sent.
// @group Request Options
//
// Example: mutate req.Request
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Before(func(r *req.Request) {
//		r.EnableDump()
//	}))
func Before(fn func(*req.Request)) Option {
	return func(r *req.Request) {
		fn(r)
	}
}

func setBody(r *req.Request, value any) {
	switch value.(type) {
	case nil:
		return
	case string, []byte, io.Reader:
		r.SetBody(value)
	default:
		r.SetBodyJsonMarshal(value)
	}
}
