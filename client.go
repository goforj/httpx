package httpx

import (
	"context"
	"fmt"
	"github.com/goforj/godump"
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
	errorMapper ErrorMapperFunc
}

// New creates a client with opinionated defaults and optional overrides.
// @group Client
//
// Example: configure all client options
//
//	var buf bytes.Buffer
//	c := httpx.New(httpx.
//		BaseURL("https://api.example.com").
//		Timeout(5*time.Second).
//		Header("X-Trace", "1").
//		Headers(map[string]string{
//			"Accept": "application/json",
//		}).
//		Transport(http.RoundTripper(http.DefaultTransport)).
//		Middleware(func(_ *req.Client, r *req.Request) error {
//			r.SetHeader("X-Middleware", "1")
//			return nil
//		}).
//		ErrorMapper(func(resp *req.Response) error {
//			return fmt.Errorf("status %d", resp.StatusCode)
//		}).
//		DumpAll().
//		DumpEachRequest().
//		DumpEachRequestTo(&buf).
//		Retry(func(rc *req.Client) {
//			rc.SetCommonRetryCount(2)
//		}).
//		RetryCount(2).
//		RetryFixedInterval(200 * time.Millisecond).
//		RetryBackoff(100*time.Millisecond, 2*time.Second).
//		RetryInterval(func(_ *req.Response, attempt int) time.Duration {
//			return time.Duration(attempt) * 100 * time.Millisecond
//		}).
//		RetryCondition(func(resp *req.Response, _ error) bool {
//			return resp != nil && resp.StatusCode == 503
//		}).
//		RetryHook(func(_ *req.Response, _ error) {}),
//	)
//	_ = c
func New(opts ...Option) *Client {
	c := &Client{
		req: req.C().SetTimeout(defaultTimeout).SetUserAgent(defaultUserAgent),
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt.applyClient(c)
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
// Example: fetch GitHub pull requests (typed)
//
//	type PullRequest struct {
//		Number int    `json:"number"`
//		Title  string `json:"title"`
//	}
//
//	c := httpx.New(httpx.Header("Accept", "application/vnd.github+json"))
//	res := httpx.Get[[]PullRequest](c, "https://api.github.com/repos/goforj/httpx/pulls")
//	if res.Err != nil {
//		return
//	}
//	godump.Dump(res.Body)
//
// Example: bind to a string body
//
//	c2 := httpx.New()
//	res2 := httpx.Get[string](c2, "https://httpbin.org/uuid")
//	_, _ = res2.Body, res2.Err // Body is string
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
//	_, _ = res.Body, res.Err // Body is User
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
//	_, _ = res.Body, res.Err // Body is User
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
//	_, _ = res.Body, res.Err // Body is User
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
//	_, _ = res.Body, res.Err // Body is DeleteResponse
func Delete[T any](client *Client, url string, opts ...Option) Result[T] {
	return do[T](client, nil, methodDelete, url, nil, opts)
}

// Head issues a HEAD request using the provided client.
// @group Requests
//
// Example: HEAD request
//
//	c := httpx.New()
//	res := httpx.Head[string](c, "https://example.com")
//	_ = res
func Head[T any](client *Client, url string, opts ...Option) Result[T] {
	return do[T](client, nil, methodHead, url, nil, opts)
}

// Options issues an OPTIONS request using the provided client.
// @group Requests
//
// Example: OPTIONS request
//
//	c := httpx.New()
//	res := httpx.Options[string](c, "https://example.com")
//	_ = res
func Options[T any](client *Client, url string, opts ...Option) Result[T] {
	return do[T](client, nil, methodOptions, url, nil, opts)
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
//	_, _ = res.Body, res.Err // Body is User
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
//	_, _ = res.Body, res.Err // Body is User
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
//	_, _ = res.Body, res.Err // Body is User
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
//	_, _ = res.Body, res.Err // Body is User
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
//	_, _ = res.Body, res.Err // Body is DeleteResponse
func DeleteCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) Result[T] {
	return do[T](client, ctx, methodDelete, url, nil, opts)
}

// HeadCtx issues a HEAD request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware HEAD
//
//	c := httpx.New()
//	ctx := context.Background()
//	res := httpx.HeadCtx[string](c, ctx, "https://example.com")
//	_ = res
func HeadCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) Result[T] {
	return do[T](client, ctx, methodHead, url, nil, opts)
}

// OptionsCtx issues an OPTIONS request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware OPTIONS
//
//	c := httpx.New()
//	ctx := context.Background()
//	res := httpx.OptionsCtx[string](c, ctx, "https://example.com")
//	_ = res
func OptionsCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) Result[T] {
	return do[T](client, ctx, methodOptions, url, nil, opts)
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
		if opt == nil {
			continue
		}
		opt.applyRequest(req)
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
	methodGet     = "GET"
	methodPost    = "POST"
	methodPut     = "PUT"
	methodPatch   = "PATCH"
	methodDelete  = "DELETE"
	methodHead    = "HEAD"
	methodOptions = "OPTIONS"
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
	case methodHead:
		return r.Head(url)
	case methodOptions:
		return r.Options(url)
	default:
		return nil, fmt.Errorf("httpx: unsupported method %s", method)
	}
}

// dumpExample is a no-op wrapper to keep the godump import live for doc examples.
func dumpExample(values ...interface{}) {
	godump.Dump(values...)
}
