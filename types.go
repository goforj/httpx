package httpx

import "github.com/imroc/req/v3"

// Option applies configuration to a client or request.
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
// Example: chain options
//
//	opt := httpx.Header("X-Trace", "1").Query("q", "go")
//	_ = opt
type OptionBuilder struct {
	ops []Option
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
type ErrorMapperFunc func(*req.Response) error
