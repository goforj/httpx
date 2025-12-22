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

// Option applies configuration to a client or a request.
type Option interface {
	applyClient(*Client)
	applyRequest(*req.Request)
}

type option struct {
	clientFn  func(*Client)
	requestFn func(*req.Request)
}

func (o option) applyClient(c *Client) {
	if o.clientFn == nil {
		return
	}
	o.clientFn(c)
}

func (o option) applyRequest(r *req.Request) {
	if o.requestFn == nil {
		return
	}
	o.requestFn(r)
}

func clientOnly(fn func(*Client)) Option {
	return option{clientFn: fn}
}

func requestOnly(fn func(*req.Request)) Option {
	return option{requestFn: fn}
}

func bothOption(clientFn func(*Client), requestFn func(*req.Request)) Option {
	return option{
		clientFn:  clientFn,
		requestFn: requestFn,
	}
}

// OptionBuilder chains request and client options.
// @group Options
//
// Example: build options
//
//	opt := httpx.Opts().Header("X-Trace", "1").Query("q", "go")
//	_ = opt
type OptionBuilder struct {
	ops []Option
}

// Opts creates a chainable option builder.
// @group Options
//
// For single options you can skip `Opts()` and call helpers like `httpx.Header` directly;
// they simply forward to the builder internally.
//
// Example: chain options
//
//	opt := httpx.Opts().Header("X-Trace", "1").Query("q", "go")
//	_ = opt
func Opts() OptionBuilder {
	return OptionBuilder{}
}

func (b OptionBuilder) add(opt Option) OptionBuilder {
	b.ops = append(b.ops, opt)
	return b
}

func (b OptionBuilder) applyClient(c *Client) {
	for _, opt := range b.ops {
		if opt == nil {
			continue
		}
		opt.applyClient(c)
	}
}

func (b OptionBuilder) applyRequest(r *req.Request) {
	for _, opt := range b.ops {
		if opt == nil {
			continue
		}
		opt.applyRequest(r)
	}
}

// ErrorMapperFunc customizes error creation for non-2xx responses.
//
// Example: map error responses
//
//	c := httpx.New(httpx.Opts().ErrorMapper(func(resp *req.Response) error {
//		return fmt.Errorf("status %d", resp.StatusCode)
//	}))
//	_ = c
type ErrorMapperFunc func(*req.Response) error
