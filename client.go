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
//		BaseURL("https://httpbin.org").
//		Timeout(5*time.Second).
//		Header("X-Trace", "1").
//		Header("Accept", "application/json").
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
//		RetryFixedInterval(200*time.Millisecond).
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
func Default() *Client {
	defaultOnce.Do(func() {
		defaultClient = New()
	})
	return defaultClient
}

// Req returns the underlying req client for advanced usage.
// @group Client
func (c *Client) Req() *req.Client {
	return c.req
}

// Raw returns the underlying req client for chaining raw requests.
// @group Client
func (c *Client) Raw() *req.Client {
	return c.req
}

func (c *Client) clone() *Client {
	if c == nil {
		return New()
	}
	return &Client{
		req:         c.req.Clone(),
		errorMapper: c.errorMapper,
	}
}

// Get issues a GET request using the provided client.
// @group Requests
//
// Example: basic GET
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/get")
//	if err != nil {
//		return
//	}
//	httpx.Dump(res)
//	// #map[string]interface {} {
//	//   url => "https://httpbin.org/get" #string
//	// }
//
// Example: bind to a string body
//
//	resText, err := httpx.Get[string](c, "https://httpbin.org/get")
//	if err != nil {
//		return
//	}
//	_ = resText // resText is string
func Get[T any](client *Client, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, nil, methodGet, url, nil, opts)
	return body, err
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
//	res, err := httpx.Post[CreateUser, User](c, "https://httpbin.org/post", CreateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps User
//	// #User {
//	//   Name => "Ana" #string
//	// }
func Post[In any, Out any](client *Client, url string, body In, opts ...Option) (Out, error) {
	out, _, err := do[Out](client, nil, methodPost, url, body, opts)
	return out, err
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
//	res, err := httpx.Put[UpdateUser, User](c, "https://httpbin.org/put", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps User
//	// #User {
//	//   Name => "Ana" #string
//	// }
func Put[In any, Out any](client *Client, url string, body In, opts ...Option) (Out, error) {
	out, _, err := do[Out](client, nil, methodPut, url, body, opts)
	return out, err
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
//	res, err := httpx.Patch[UpdateUser, User](c, "https://httpbin.org/patch", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps User
//	// #User {
//	//   Name => "Ana" #string
//	// }
func Patch[In any, Out any](client *Client, url string, body In, opts ...Option) (Out, error) {
	out, _, err := do[Out](client, nil, methodPatch, url, body, opts)
	return out, err
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
//	res, err := httpx.Delete[DeleteResponse](c, "https://httpbin.org/delete")
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps DeleteResponse
//	// #DeleteResponse {
//	//   OK => true #bool
//	// }
func Delete[T any](client *Client, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, nil, methodDelete, url, nil, opts)
	return body, err
}

// Head issues a HEAD request using the provided client.
// @group Requests
//
// Example: HEAD request
//
//	c := httpx.New()
//	_, err := httpx.Head[string](c, "https://httpbin.org/get")
//	if err != nil {
//		return
//	}
func Head[T any](client *Client, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, nil, methodHead, url, nil, opts)
	return body, err
}

// Options issues an OPTIONS request using the provided client.
// @group Requests
//
// Example: OPTIONS request
//
//	c := httpx.New()
//	_, err := httpx.Options[string](c, "https://httpbin.org/get")
//	if err != nil {
//		return
//	}
func Options[T any](client *Client, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, nil, methodOptions, url, nil, opts)
	return body, err
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
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := httpx.GetCtx[User](c, ctx, "https://httpbin.org/get")
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps User
//	// #User {
//	//   Name => "Ana" #string
//	// }
func GetCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, ctx, methodGet, url, nil, opts)
	return body, err
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
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := httpx.PostCtx[CreateUser, User](c, ctx, "https://httpbin.org/post", CreateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps User
//	// #User {
//	//   Name => "Ana" #string
//	// }
func PostCtx[In any, Out any](client *Client, ctx context.Context, url string, body In, opts ...Option) (Out, error) {
	out, _, err := do[Out](client, ctx, methodPost, url, body, opts)
	return out, err
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
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := httpx.PutCtx[UpdateUser, User](c, ctx, "https://httpbin.org/put", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps User
//	// #User {
//	//   Name => "Ana" #string
//	// }
func PutCtx[In any, Out any](client *Client, ctx context.Context, url string, body In, opts ...Option) (Out, error) {
	out, _, err := do[Out](client, ctx, methodPut, url, body, opts)
	return out, err
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
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := httpx.PatchCtx[UpdateUser, User](c, ctx, "https://httpbin.org/patch", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps User
//	// #User {
//	//   Name => "Ana" #string
//	// }
func PatchCtx[In any, Out any](client *Client, ctx context.Context, url string, body In, opts ...Option) (Out, error) {
	out, _, err := do[Out](client, ctx, methodPatch, url, body, opts)
	return out, err
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
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := httpx.DeleteCtx[DeleteResponse](c, ctx, "https://httpbin.org/delete")
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps DeleteResponse
//	// #DeleteResponse {
//	//   OK => true #bool
//	// }
func DeleteCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, ctx, methodDelete, url, nil, opts)
	return body, err
}

// HeadCtx issues a HEAD request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware HEAD
//
//	ctx := context.Background()
//	c := httpx.New()
//	_, err := httpx.HeadCtx[string](c, ctx, "https://httpbin.org/get")
//	if err != nil {
//		return
//	}
func HeadCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, ctx, methodHead, url, nil, opts)
	return body, err
}

// OptionsCtx issues an OPTIONS request using the provided client and context.
// @group Requests (Context)
//
// Example: context-aware OPTIONS
//
//	ctx := context.Background()
//	c := httpx.New()
//	_, err := httpx.OptionsCtx[string](c, ctx, "https://httpbin.org/get")
//	if err != nil {
//		return
//	}
func OptionsCtx[T any](client *Client, ctx context.Context, url string, opts ...Option) (T, error) {
	body, _, err := do[T](client, ctx, methodOptions, url, nil, opts)
	return body, err
}

// Do executes a pre-configured req request and returns the decoded body and response.
// @group Requests
//
// Example: advanced request with response access
//
//	r := req.C().R().SetHeader("X-Trace", "1")
//	r.SetURL("https://httpbin.org/get")
//	r.Method = http.MethodGet
//
//	res, rawResp, err := httpx.Do[map[string]any](r)
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   url => "https://httpbin.org/get" #string
//	// }
//	_ = rawResp
//	_ = err
func Do[T any](r *req.Request) (T, *req.Response, error) {
	var out T

	if r == nil {
		return out, nil, fmt.Errorf("httpx: nil request")
	}
	rawKind := rawKindOf[T]()
	if rawKind == rawNone {
		r.SetSuccessResult(&out)
	}

	resp := r.Do()
	if resp.Err != nil {
		return out, resp, resp.Err
	}
	if resp.IsSuccessState() {
		if rawKind != rawNone {
			out = decodeRaw[T](resp)
		}
		if rawKind == rawNone && isEmptyBody(resp) {
			ensureNonNil(&out)
		}
		return out, resp, nil
	}

	return out, resp, newHTTPError(resp)
}

func do[T any](client *Client, ctx context.Context, method, url string, body any, opts []Option) (T, *req.Response, error) {
	var out T

	if client == nil {
		client = Default()
	}
	if len(opts) != 0 {
		client = client.clone()
		for _, opt := range opts {
			if opt == nil {
				continue
			}
			opt.applyClient(client)
		}
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
		req.SetSuccessResult(&out)
	}

	resp, err := send(req, method, url)
	if err != nil {
		if resp != nil && resp.IsSuccessState() && rawKind == rawNone && isEmptyBody(resp) {
			ensureNonNil(&out)
			return out, resp, nil
		}
		return out, resp, err
	}
	if resp.IsSuccessState() {
		if rawKind != rawNone {
			out = decodeRaw[T](resp)
		}
		if rawKind == rawNone {
			ensureNonNil(&out)
		}
		return out, resp, nil
	}

	return out, resp, client.mapError(resp)
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
