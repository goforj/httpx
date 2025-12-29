package httpx

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/imroc/req/v3"
)

// Header sets a header on a request or client.
// @group Request Composition
//
// Applies to both client defaults and request-time headers.
// Example: apply a header
//
//	// Apply to all requests
//	c := httpx.New(httpx.Header("X-Trace", "1"))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   headers => #map[string]interface {} {
//	//     X-Trace => "1" #string
//	//   }
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.Header("X-Trace", "1"))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   headers => #map[string]interface {} {
//	//     X-Trace => "1" #string
//	//   }
//	// }
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

func headers(values map[string]string) OptionBuilder {
	return OptionBuilder{}.add(bothOption(
		func(c *Client) {
			c.req.SetCommonHeaders(values)
		},
		func(r *req.Request) {
			r.SetHeaders(values)
		},
	))
}

// UserAgent sets the User-Agent header on a request or client.
// @group Request Composition
//
// Applies to both client defaults and request-time headers.
// Example: set a User-Agent
//
//	// Apply to all requests
//	c := httpx.New(httpx.UserAgent("my-app/1.0"))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/headers")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   headers => #map[string]interface {} {
//	//     User-Agent => "my-app/1.0" #string
//	//   }
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/headers", httpx.UserAgent("my-app/1.0"))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   headers => #map[string]interface {} {
//	//     User-Agent => "my-app/1.0" #string
//	//   }
//	// }
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
// @group Request Composition
//
// Applies to individual requests only.
// Example: add query params
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.Query("q", "search"))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   args => #map[string]interface {} {
//	//     q => "search" #string
//	//   }
//	// }
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

// Path sets a path parameter by name.
// @group Request Composition
//
// Applies to individual requests only.
// Example: path parameter
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/anything/{id}", httpx.Path("id", 42))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   url => "https://httpbin.org/anything/42" #string
//	// }
func Path(key string, value any) OptionBuilder {
	return OptionBuilder{}.Path(key, value)
}

func (b OptionBuilder) Path(key string, value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetPathParam(key, fmt.Sprint(value))
	}))
}

// Body sets the request body and infers JSON for structs and maps.
// @group Request Composition
//
// Applies to individual requests only.
// Example: send JSON body with inference
//
//	type Payload struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.Body(Payload{Name: "Ana"}))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   json => #map[string]interface {} {
//	//     name => "Ana" #string
//	//   }
//	// }
func Body(value any) OptionBuilder {
	return OptionBuilder{}.Body(value)
}

func (b OptionBuilder) Body(value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		setBody(r, value)
	}))
}

// JSON sets the request body as JSON.
// @group Request Composition
//
// Applies to individual requests only.
// Example: force JSON body
//
//	type Payload struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.JSON(Payload{Name: "Ana"}))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   json => #map[string]interface {} {
//	//     name => "Ana" #string
//	//   }
//	// }
func JSON(value any) OptionBuilder {
	return OptionBuilder{}.JSON(value)
}

func (b OptionBuilder) JSON(value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetBodyJsonMarshal(value)
	}))
}

// Form sets form data for the request.
// @group Request Composition
//
// Applies to individual requests only.
// Example: submit a form
//
//	c := httpx.New()
//	res, err := httpx.Post[any, map[string]any](c, "https://httpbin.org/post", nil, httpx.Form(map[string]string{
//		"name": "alice",
//	}))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   form => #map[string]interface {} {
//	//     name => "alice" #string
//	//   }
//	// }
func Form(values map[string]string) OptionBuilder {
	return OptionBuilder{}.Form(values)
}

func (b OptionBuilder) Form(values map[string]string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetFormData(values)
	}))
}

// Timeout sets a per-request timeout using context cancellation.
// @group Request Control
//
// Applies to both client defaults (via WithTimeout) and individual requests.
// Example: timeout
//
//	// Apply to all requests
//	c := httpx.New(httpx.Timeout(2 * time.Second))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/delay/2")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   url => "https://httpbin.org/delay/2" #string
//	// }
//
//	// Apply to a single request
//	res, err = httpx.Get[map[string]any](c, "https://httpbin.org/delay/2", httpx.Timeout(2*time.Second))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   url => "https://httpbin.org/delay/2" #string
//	// }
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
// @group Request Control
//
// Applies to individual requests only.
// Example: mutate req.Request
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get", httpx.Before(func(r *req.Request) {
//		r.EnableDump()
//	}))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   url => "https://httpbin.org/get" #string
//	// }
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
