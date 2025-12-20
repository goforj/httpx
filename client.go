package httpx

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

const defaultTimeout = 10 * time.Second
const defaultUserAgent = "httpx (https://github.com/goforj/httpx)"

var (
	defaultClient *Client
	defaultOnce   sync.Once
)

// Client wraps a req client with typed helpers and defaults.
// @group Client
//
// Example: create a client
//
//	c := httpx.New()
//	_ = c
type Client struct {
	req         *req.Client
	errorMapper ErrorMapper
}

// New creates a client with opinionated defaults and optional overrides.
// @group Client
//
// Example: custom base URL and timeout
//
//	c := httpx.New(
//		httpx.WithBaseURL("https://api.example.com"),
//		httpx.WithTimeout(5*time.Second),
//	)
//	_ = c
func New(opts ...ClientOption) *Client {
	c := &Client{
		req: req.C().SetTimeout(defaultTimeout).SetUserAgent(defaultUserAgent),
	}
	for _, opt := range opts {
		opt(c)
	}
	if _, ok := os.LookupEnv("HTTP_TRACE"); ok {
		c.req.EnableDumpAll()
	}
	return c
}

// Default returns the shared default client.
// @group Client
//
// Example: use default client for quick calls
//
//	c := httpx.Default()
//	_ = c
func Default() *Client {
	defaultOnce.Do(func() {
		defaultClient = New()
	})
	return defaultClient
}

// Req returns the underlying req client for advanced usage.
// @group Client
//
// Example: enable req debugging
//
//	c := httpx.New()
//	c.Req().EnableDumpEachRequest()
func (c *Client) Req() *req.Client {
	return c.req
}

// Raw returns the underlying req client for chaining raw requests.
// @group Client
//
// Example: drop down to req
//
//	c := httpx.New()
//	resp, err := c.Raw().R().Get("https://httpbin.org/uuid")
//	_, _ = resp, err
func (c *Client) Raw() *req.Client {
	return c.req
}

// Get issues a GET request using the provided client.
// @group Requests
//
// Example: fetch GitHub pull requests
//
//	type PullRequest struct {
//		Number int    `json:"number"`
//		Title  string `json:"title"`
//	}
//
//	c := httpx.New(httpx.WithHeader("Accept", "application/vnd.github+json"))
//	res := httpx.Get[[]PullRequest](c, "https://api.github.com/repos/goforj/httpx/pulls")
//	if res.Err != nil {
//		return
//	}
//	godump.Dump(res.Body)
func Get[T any](client *Client, url string, opts ...Option) Result[T] {
	return do[T](client, nil, methodGet, url, nil, opts)
}

// Post issues a POST request using the provided client.
// @group Requests
//
// Example: typed POST
//
//	type CreateUser struct {
//		Name string `json:"name"`
//	}
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res := httpx.Post[CreateUser, User](c, "https://api.example.com/users", CreateUser{Name: "Ana"})
//	_, _ = res.Body, res.Err
func Post[In any, Out any](client *Client, url string, body In, opts ...Option) Result[Out] {
	return do[Out](client, nil, methodPost, url, body, opts)
}

// Put issues a PUT request using the provided client.
// @group Requests
//
// Example: typed PUT
//
//	type UpdateUser struct {
//		Name string `json:"name"`
//	}
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res := httpx.Put[UpdateUser, User](c, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
//	_, _ = res.Body, res.Err
func Put[In any, Out any](client *Client, url string, body In, opts ...Option) Result[Out] {
	return do[Out](client, nil, methodPut, url, body, opts)
}

// Patch issues a PATCH request using the provided client.
// @group Requests
//
// Example: typed PATCH
//
//	type UpdateUser struct {
//		Name string `json:"name"`
//	}
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res := httpx.Patch[UpdateUser, User](c, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
//	_, _ = res.Body, res.Err
func Patch[In any, Out any](client *Client, url string, body In, opts ...Option) Result[Out] {
	return do[Out](client, nil, methodPatch, url, body, opts)
}

// Delete issues a DELETE request using the provided client.
// @group Requests
//
// Example: typed DELETE
//
//	type DeleteResponse struct {
//		OK bool `json:"ok"`
//	}
//
//	c := httpx.New()
//	res := httpx.Delete[DeleteResponse](c, "https://api.example.com/users/1")
//	_, _ = res.Body, res.Err
func Delete[T any](client *Client, url string, opts ...Option) Result[T] {
	return do[T](client, nil, methodDelete, url, nil, opts)
}

// GetCtx issues a GET request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware GET
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	ctx := context.Background()
//	res := httpx.GetCtx[User](c, ctx, "https://api.example.com/users/1")
//	_, _ = res.Body, res.Err
func GetCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) Result[T] {
	return do[T](client, ctx, methodGet, url, nil, opts)
}

// PostCtx issues a POST request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware POST
//
//	type CreateUser struct {
//		Name string `json:"name"`
//	}
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	ctx := context.Background()
//	res := httpx.PostCtx[CreateUser, User](c, ctx, "https://api.example.com/users", CreateUser{Name: "Ana"})
//	_, _ = res.Body, res.Err
func PostCtx[In any, Out any](client *Client, ctx context.Context, url string, body In, opts ...Option) Result[Out] {
	return do[Out](client, ctx, methodPost, url, body, opts)
}

// PutCtx issues a PUT request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware PUT
//
//	type UpdateUser struct {
//		Name string `json:"name"`
//	}
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	ctx := context.Background()
//	res := httpx.PutCtx[UpdateUser, User](c, ctx, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
//	_, _ = res.Body, res.Err
func PutCtx[In any, Out any](client *Client, ctx context.Context, url string, body In, opts ...Option) Result[Out] {
	return do[Out](client, ctx, methodPut, url, body, opts)
}

// PatchCtx issues a PATCH request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware PATCH
//
//	type UpdateUser struct {
//		Name string `json:"name"`
//	}
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	ctx := context.Background()
//	res := httpx.PatchCtx[UpdateUser, User](c, ctx, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
//	_, _ = res.Body, res.Err
func PatchCtx[In any, Out any](client *Client, ctx context.Context, url string, body In, opts ...Option) Result[Out] {
	return do[Out](client, ctx, methodPatch, url, body, opts)
}

// DeleteCtx issues a DELETE request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware DELETE
//
//	type DeleteResponse struct {
//		OK bool `json:"ok"`
//	}
//
//	c := httpx.New()
//	ctx := context.Background()
//	res := httpx.DeleteCtx[DeleteResponse](c, ctx, "https://api.example.com/users/1")
//	_, _ = res.Body, res.Err
func DeleteCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) Result[T] {
	return do[T](client, ctx, methodDelete, url, nil, opts)
}

func do[T any](client *Client, ctx context.Context, method, url string, body any, opts []Option) Result[T] {
	var res Result[T]

	if client == nil {
		client = Default()
	}
	req := client.req.R()
	if ctx != nil {
		req.SetContext(ctx)
	}
	if body != nil {
		setBody(req, body)
	}
	for _, opt := range opts {
		opt(req)
	}

	rawKind := rawKindOf[T]()
	if rawKind == rawNone {
		req.SetSuccessResult(&res.Body)
	}

	resp, err := send(req, method, url)
	if err != nil {
		if resp != nil && resp.IsSuccessState() && rawKind == rawNone && isEmptyBody(resp) {
			ensureNonNil(&res.Body)
			res.Response = resp
			return res
		}
		res.Err = err
		res.Response = resp
		return res
	}
	if resp.IsSuccessState() {
		if rawKind != rawNone {
			res.Body = decodeRaw[T](resp)
		}
		if rawKind == rawNone {
			ensureNonNil(&res.Body)
		}
		res.Response = resp
		return res
	}

	res.Response = resp
	res.Err = client.mapError(resp)
	return res
}

func (c *Client) mapError(resp *req.Response) error {
	if c.errorMapper != nil {
		return c.errorMapper(resp)
	}
	return newHTTPError(resp)
}

const (
	methodGet    = "GET"
	methodPost   = "POST"
	methodPut    = "PUT"
	methodPatch  = "PATCH"
	methodDelete = "DELETE"
)

func send(r *req.Request, method, url string) (*req.Response, error) {
	switch method {
	case methodGet:
		return r.Get(url)
	case methodPost:
		return r.Post(url)
	case methodPut:
		return r.Put(url)
	case methodPatch:
		return r.Patch(url)
	case methodDelete:
		return r.Delete(url)
	default:
		return nil, fmt.Errorf("httpx: unsupported method %s", method)
	}
}
