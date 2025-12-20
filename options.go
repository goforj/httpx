package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

// Header sets a header on a request.
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

// Query adds a single query parameter.
//
// Example: add query param
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go"))
func Query(key, value string) Option {
	return func(r *req.Request) {
		r.SetQueryParam(key, value)
	}
}

// Queries adds multiple query parameters.
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

// Dump enables req's request-level dump output.
//
// Example: dump a single request
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Dump())
func Dump() Option {
	return func(r *req.Request) {
		r.EnableDump()
	}
}

// DumpTo enables req's request-level dump output to a writer.
//
// Example: dump to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.DumpTo(&buf))
func DumpTo(output io.Writer) Option {
	return func(r *req.Request) {
		r.EnableDumpTo(output)
	}
}

// DumpToFile enables req's request-level dump output to a file path.
//
// Example: dump to a file
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.DumpToFile("httpx.dump"))
func DumpToFile(filename string) Option {
	return func(r *req.Request) {
		r.EnableDumpToFile(filename)
	}
}

// Timeout sets a per-request timeout using context cancellation.
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

// Bearer sets the Authorization header with a bearer token.
//
// Example: bearer auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Bearer("token"))
func Bearer(token string) Option {
	return func(r *req.Request) {
		r.SetBearerAuthToken(token)
	}
}

// Basic sets HTTP basic authentication headers.
//
// Example: basic auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Basic("user", "pass"))
func Basic(user, pass string) Option {
	return func(r *req.Request) {
		r.SetBasicAuth(user, pass)
	}
}

// Before runs a hook before the request is sent.
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

// WithBaseURL sets a base URL on the client.
//
// Example: client base URL
//
//	c := httpx.New(httpx.WithBaseURL("https://api.example.com"))
//	_ = c
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.req.SetBaseURL(url)
	}
}

// WithTimeout sets the default timeout for the client.
//
// Example: client timeout
//
//	c := httpx.New(httpx.WithTimeout(3 * time.Second))
//	_ = c
func WithTimeout(d time.Duration) ClientOption {
	return func(c *Client) {
		c.req.SetTimeout(d)
	}
}

// WithHeader sets a default header for all requests.
//
// Example: client header
//
//	c := httpx.New(httpx.WithHeader("X-Trace", "1"))
//	_ = c
func WithHeader(key, value string) ClientOption {
	return func(c *Client) {
		c.req.SetCommonHeader(key, value)
	}
}

// WithHeaders sets default headers for all requests.
//
// Example: client headers
//
//	c := httpx.New(httpx.WithHeaders(map[string]string{
//		"X-Trace": "1",
//		"Accept":  "application/json",
//	}))
//	_ = c
func WithHeaders(values map[string]string) ClientOption {
	return func(c *Client) {
		c.req.SetCommonHeaders(values)
	}
}

// RetryOption configures retry behavior on the underlying req client.
type RetryOption func(*req.Client)

// WithRetry applies a retry configuration to the client.
//
// Example: set retry count
//
//	c := httpx.New(httpx.WithRetry(func(rc *req.Client) {
//		rc.SetCommonRetryCount(2)
//	}))
//	_ = c
func WithRetry(opt RetryOption) ClientOption {
	return func(c *Client) {
		if opt != nil {
			opt(c.req)
		}
	}
}

// WithTransport wraps the underlying transport with a custom RoundTripper.
//
// Example: wrap transport
//
//	c := httpx.New(httpx.WithTransport(http.RoundTripper(http.DefaultTransport)))
//	_ = c
func WithTransport(rt http.RoundTripper) ClientOption {
	return func(c *Client) {
		if rt == nil {
			return
		}
		c.req.Transport.WrapRoundTrip(func(http.RoundTripper) http.RoundTripper {
			return rt
		})
	}
}

// WithDumpAll enables req's client-level dump output for all requests.
//
// Example: dump every request and response
//
//	c := httpx.New(httpx.WithDumpAll())
//	_ = c
func WithDumpAll() ClientOption {
	return func(c *Client) {
		c.req.EnableDumpAll()
	}
}

// WithDumpEachRequest enables request-level dumps for each request on the client.
//
// Example: dump each request as it is sent
//
//	c := httpx.New(httpx.WithDumpEachRequest())
//	_ = c
func WithDumpEachRequest() ClientOption {
	return func(c *Client) {
		c.req.EnableDumpEachRequest()
	}
}

// WithDumpEachRequestTo enables request-level dumps for each request and writes
// them to the provided output.
//
// Example: dump each request to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New(httpx.WithDumpEachRequestTo(&buf))
//	_ = httpx.Get[string](c, "https://example.com")
//	_ = buf.String()
func WithDumpEachRequestTo(output io.Writer) ClientOption {
	return func(c *Client) {
		if output == nil {
			return
		}
		c.req.EnableDumpEachRequest()
		c.req.OnAfterResponse(func(_ *req.Client, resp *req.Response) error {
			if resp == nil {
				return nil
			}
			_, _ = output.Write([]byte(resp.Dump()))
			return nil
		})
	}
}

// WithMiddleware adds request middleware to the client.
//
// Example: add request middleware
//
//	c := httpx.New(httpx.WithMiddleware(func(_ *req.Client, r *req.Request) error {
//		r.SetHeader("X-Trace", "1")
//		return nil
//	}))
//	_ = c
func WithMiddleware(mw ...req.RequestMiddleware) ClientOption {
	return func(c *Client) {
		for _, m := range mw {
			c.req.OnBeforeRequest(m)
		}
	}
}

// WithErrorMapper sets a custom error mapper for non-2xx responses.
//
// Example: map error responses
//
//	c := httpx.New(httpx.WithErrorMapper(func(resp *req.Response) error {
//		return fmt.Errorf("status %d", resp.StatusCode)
//	}))
//	_ = c
func WithErrorMapper(fn ErrorMapper) ClientOption {
	return func(c *Client) {
		c.errorMapper = fn
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
