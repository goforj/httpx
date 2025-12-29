package httpx

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/imroc/req/v3"
)

// Header sets a header on a request or client.
// @group Request Options
//
// Applies to both client defaults and request-time headers.
// Example: apply a header
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Header("X-Trace", "1"))
func Header(key, value string) OptionBuilder {
	return OptionBuilder{}.Header(key, value)
}

func (b OptionBuilder) Header(key, value string) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonHeader(key, value)
		},
		func(r *req.Request) {
			r.SetHeader(key, value)
		},
	))
}

// Headers sets multiple headers on a request or client.
// @group Request Options
//
// Applies to both client defaults and request-time headers.
// Example: apply headers
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Headers(map[string]string{
//		"X-Trace": "1",
//		"Accept":  "application/json",
//	}))
func Headers(values map[string]string) OptionBuilder {
	return OptionBuilder{}.Headers(values)
}

func (b OptionBuilder) Headers(values map[string]string) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetCommonHeaders(values)
		},
		func(r *req.Request) {
			r.SetHeaders(values)
		},
	))
}

// UserAgent sets the User-Agent header on a request or client.
// @group Request Options
//
// Applies to both client defaults and request-time headers.
// Example: set a User-Agent
//
//	c := httpx.New(httpx.UserAgent("my-app/1.0"))
//	_ = httpx.Get[string](c, "https://example.com")
func UserAgent(value string) OptionBuilder {
	return OptionBuilder{}.UserAgent(value)
}

func (b OptionBuilder) UserAgent(value string) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetUserAgent(value)
		},
		func(r *req.Request) {
			r.SetHeader("User-Agent", value)
		},
	))
}

// Query adds query parameters as key/value pairs.
// @group Request Options
//
// Applies to individual requests only.
// Example: add query params
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go", "ok", "1"))
func Query(kv ...string) OptionBuilder {
	return OptionBuilder{}.Query(kv...)
}

func (b OptionBuilder) Query(kv ...string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		if len(kv)%2 != 0 {
			panic("httpx: Query expects even number of key/value arguments")
		}
		for i := 0; i < len(kv); i += 2 {
			r.AddQueryParam(kv[i], kv[i+1])
		}
	}))
}

// Queries adds multiple query parameters.
// @group Request Options
//
// Applies to individual requests only.
// Example: add query params
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/search", httpx.Queries(map[string]string{
//		"q":  "go",
//		"ok": "1",
//	}))
func Queries(values map[string]string) OptionBuilder {
	return OptionBuilder{}.Queries(values)
}

func (b OptionBuilder) Queries(values map[string]string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetQueryParams(values)
	}))
}

// Path sets a path parameter by name.
// @group Request Options
//
// Applies to individual requests only.
// Example: path parameter
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	_ = httpx.Get[User](c, "https://example.com/users/{id}", httpx.Path("id", 42))
func Path(key string, value any) OptionBuilder {
	return OptionBuilder{}.Path(key, value)
}

func (b OptionBuilder) Path(key string, value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetPathParam(key, fmt.Sprint(value))
	}))
}

// Paths sets multiple path parameters.
// @group Request Options
//
// Applies to individual requests only.
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
func Paths(values map[string]any) OptionBuilder {
	return OptionBuilder{}.Paths(values)
}

func (b OptionBuilder) Paths(values map[string]any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		params := make(map[string]string, len(values))
		for key, value := range values {
			params[key] = fmt.Sprint(value)
		}
		r.SetPathParams(params)
	}))
}

// Body sets the request body and infers JSON for structs and maps.
// @group Request Options
//
// Applies to individual requests only.
// Example: send JSON body with inference
//
//	type Payload struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Body(Payload{Name: "Ana"}))
func Body(value any) OptionBuilder {
	return OptionBuilder{}.Body(value)
}

func (b OptionBuilder) Body(value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		setBody(r, value)
	}))
}

// JSON sets the request body as JSON.
// @group Request Options
//
// Applies to individual requests only.
// Example: force JSON body
//
//	type Payload struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.JSON(Payload{Name: "Ana"}))
func JSON(value any) OptionBuilder {
	return OptionBuilder{}.JSON(value)
}

func (b OptionBuilder) JSON(value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetBodyJsonMarshal(value)
	}))
}

// Form sets form data for the request.
// @group Request Options
//
// Applies to individual requests only.
// Example: submit a form
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Form(map[string]string{
//		"name": "Ana",
//	}))
func Form(values map[string]string) OptionBuilder {
	return OptionBuilder{}.Form(values)
}

func (b OptionBuilder) Form(values map[string]string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetFormData(values)
	}))
}

// Timeout sets a per-request timeout using context cancellation.
// @group Request Options
//
// Applies to both client defaults (via WithTimeout) and individual requests.
// Example: per-request timeout
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Timeout(2*time.Second))
func Timeout(d time.Duration) OptionBuilder {
	return OptionBuilder{}.Timeout(d)
}

func (b OptionBuilder) Timeout(d time.Duration) OptionBuilder {
	return b.add(bothOption(
		func(c *Client) {
			c.req.SetTimeout(d)
		},
		func(r *req.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, d)
			r.SetContext(ctx)
			r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
				cancel()
				return nil
			})
		},
	))
}

// Before runs a hook before the request is sent.
// @group Request Options
//
// Applies to individual requests only.
// Example: mutate req.Request
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Before(func(r *req.Request) {
//		r.EnableDump()
//	}))
func Before(fn func(*req.Request)) OptionBuilder {
	return OptionBuilder{}.Before(fn)
}

func (b OptionBuilder) Before(fn func(*req.Request)) OptionBuilder {
	if fn == nil {
		return b
	}
	return b.add(requestOnly(func(r *req.Request) {
		fn(r)
	}))
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
