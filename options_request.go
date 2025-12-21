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

// Query adds query parameters as key/value pairs.
// @group Request Options
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
func (b OptionBuilder) Queries(values map[string]string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetQueryParams(values)
	}))
}

// Path sets a path parameter by name.
// @group Request Options
func (b OptionBuilder) Path(key string, value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetPathParam(key, fmt.Sprint(value))
	}))
}

// Paths sets multiple path parameters.
// @group Request Options
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
func (b OptionBuilder) Body(value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		setBody(r, value)
	}))
}

// JSON sets the request body as JSON.
// @group Request Options
func (b OptionBuilder) JSON(value any) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetBodyJsonMarshal(value)
	}))
}

// Form sets form data for the request.
// @group Request Options
func (b OptionBuilder) Form(values map[string]string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetFormData(values)
	}))
}

// Timeout sets a per-request timeout using context cancellation.
// @group Request Options
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
