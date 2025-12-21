package httpx

import (
	"io"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

// Header sets a header on a request or client.
// @group Request Options
//
// Example: apply a header
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Header("X-Trace", "1"))
func Header(key, value string) Option {
	return Opts().Header(key, value)
}

// Headers sets multiple headers on a request or client.
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
	return Opts().Headers(values)
}

// Query adds query parameters as key/value pairs.
// @group Request Options
//
// Example: add query params
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go", "ok", "1"))
func Query(kv ...string) Option {
	return Opts().Query(kv...)
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
	return Opts().Queries(values)
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
	return Opts().Path(key, value)
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
	return Opts().Paths(values)
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
	return Opts().Body(value)
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
	return Opts().JSON(value)
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
	return Opts().Form(values)
}

// Timeout sets a per-request timeout using context cancellation.
// @group Request Options
//
// Example: per-request timeout
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Timeout(2*time.Second))
func Timeout(d time.Duration) Option {
	return Opts().Timeout(d)
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
	return Opts().Before(fn)
}

// Auth sets the Authorization header using a scheme and token.
// @group Auth
//
// Example: custom auth scheme
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Auth("Token", "abc123"))
func Auth(scheme, token string) Option {
	return Opts().Auth(scheme, token)
}

// Bearer sets the Authorization header with a bearer token.
// @group Auth
//
// Example: bearer auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Bearer("token"))
func Bearer(token string) Option {
	return Opts().Bearer(token)
}

// Basic sets HTTP basic authentication headers.
// @group Auth
//
// Example: basic auth
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Basic("user", "pass"))
func Basic(user, pass string) Option {
	return Opts().Basic(user, pass)
}

// BaseURL sets a base URL on the client.
// @group Client Options
//
// Example: client base URL
//
//	c := httpx.New(httpx.BaseURL("https://api.example.com"))
//	_ = c
func BaseURL(url string) Option {
	return Opts().BaseURL(url)
}

// Transport wraps the underlying transport with a custom RoundTripper.
// @group Client Options
//
// Example: wrap transport
//
//	c := httpx.New(httpx.Transport(http.RoundTripper(http.DefaultTransport)))
//	_ = c
func Transport(rt http.RoundTripper) Option {
	return Opts().Transport(rt)
}

// Middleware adds request middleware to the client.
// @group Client Options
//
// Example: add request middleware
//
//	c := httpx.New(httpx.Middleware(func(_ *req.Client, r *req.Request) error {
//		r.SetHeader("X-Trace", "1")
//		return nil
//	}))
//	_ = c
func Middleware(mw ...req.RequestMiddleware) Option {
	return Opts().Middleware(mw...)
}

// ErrorMapper sets a custom error mapper for non-2xx responses.
// @group Client Options
//
// Example: map error responses
//
//	c := httpx.New(httpx.ErrorMapper(func(resp *req.Response) error {
//		return fmt.Errorf("status %d", resp.StatusCode)
//	}))
//	_ = c
func ErrorMapper(fn ErrorMapperFunc) Option {
	return Opts().ErrorMapper(fn)
}

// Dump enables req's request-level dump output.
// @group Debugging
//
// Example: dump a single request
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Dump())
func Dump() Option {
	return Opts().Dump()
}

// DumpTo enables req's request-level dump output to a writer.
// @group Debugging
//
// Example: dump to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.DumpTo(&buf))
func DumpTo(output io.Writer) Option {
	return Opts().DumpTo(output)
}

// DumpToFile enables req's request-level dump output to a file path.
// @group Debugging
//
// Example: dump to a file
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.DumpToFile("httpx.dump"))
func DumpToFile(filename string) Option {
	return Opts().DumpToFile(filename)
}

// DumpAll enables req's client-level dump output for all requests.
// @group Debugging
//
// Example: dump every request and response
//
//	c := httpx.New(httpx.DumpAll())
//	_ = c
func DumpAll() Option {
	return Opts().DumpAll()
}

// DumpEachRequest enables request-level dumps for each request on the client.
// @group Debugging
//
// Example: dump each request as it is sent
//
//	c := httpx.New(httpx.DumpEachRequest())
//	_ = c
func DumpEachRequest() Option {
	return Opts().DumpEachRequest()
}

// DumpEachRequestTo enables request-level dumps for each request and writes
// @group Debugging
// them to the provided output.
//
// Example: dump each request to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New(httpx.DumpEachRequestTo(&buf))
//	_ = httpx.Get[string](c, "https://example.com")
//	_ = buf.String()
func DumpEachRequestTo(output io.Writer) Option {
	return Opts().DumpEachRequestTo(output)
}

// RetryCount enables retry for a request and sets the maximum retry count.
// @group Retry
//
// Example: request retry count
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCount(2))
func RetryCount(count int) Option {
	return Opts().RetryCount(count)
}

// RetryFixedInterval sets a fixed retry interval for a request.
// @group Retry
//
// Example: request retry interval
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
func RetryFixedInterval(interval time.Duration) Option {
	return Opts().RetryFixedInterval(interval)
}

// RetryBackoff sets a capped exponential backoff retry interval for a request.
// @group Retry
//
// Example: request retry backoff
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
func RetryBackoff(min, max time.Duration) Option {
	return Opts().RetryBackoff(min, max)
}

// RetryInterval sets a custom retry interval function for a request.
// @group Retry
//
// Example: custom retry interval
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
//		return time.Duration(attempt) * 100 * time.Millisecond
//	}))
func RetryInterval(fn req.GetRetryIntervalFunc) Option {
	return Opts().RetryInterval(fn)
}

// RetryCondition sets the retry condition for a request.
// @group Retry
//
// Example: retry on 503
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
//		return resp != nil && resp.StatusCode == 503
//	}))
func RetryCondition(condition req.RetryConditionFunc) Option {
	return Opts().RetryCondition(condition)
}

// RetryHook registers a retry hook for a request.
// @group Retry
//
// Example: hook on retry
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
func RetryHook(hook req.RetryHookFunc) Option {
	return Opts().RetryHook(hook)
}

// Retry applies a custom retry configuration to the client.
// @group Retry (Client)
//
// Example: configure client retry
//
//	c := httpx.New(httpx.Retry(func(rc *req.Client) {
//		rc.SetCommonRetryCount(2)
//	}))
//	_ = c
func Retry(fn func(*req.Client)) Option {
	return Opts().Retry(fn)
}

// File attaches a file from disk as multipart form data.
// @group Upload Options
//
// Example: upload a file
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.File("file", "/tmp/report.txt"))
func File(paramName, filePath string) Option {
	return Opts().File(paramName, filePath)
}

// Files attaches multiple files from disk as multipart form data.
// @group Upload Options
//
// Example: upload multiple files
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.Files(map[string]string{
//		"fileA": "/tmp/a.txt",
//		"fileB": "/tmp/b.txt",
//	}))
func Files(files map[string]string) Option {
	return Opts().Files(files)
}

// FileBytes attaches a file from bytes as multipart form data.
// @group Upload Options
//
// Example: upload bytes as a file
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileBytes("file", "report.txt", []byte("hello")))
func FileBytes(paramName, filename string, content []byte) Option {
	return Opts().FileBytes(paramName, filename, content)
}

// FileReader attaches a file from a reader as multipart form data.
// @group Upload Options
//
// Example: upload from reader
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileReader("file", "report.txt", strings.NewReader("hello")))
func FileReader(paramName, filename string, reader io.Reader) Option {
	return Opts().FileReader(paramName, filename, reader)
}

// UploadCallback registers a callback for upload progress.
// @group Upload Options
//
// Example: track upload progress
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
//		httpx.File("file", "/tmp/report.bin"),
//		httpx.UploadCallback(func(info req.UploadInfo) {
//			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
//			fmt.Printf("\rprogress: %.1f%%", percent)
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}),
//	)
func UploadCallback(callback req.UploadCallback) Option {
	return Opts().UploadCallback(callback)
}

// UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.
// @group Upload Options
//
// Example: throttle upload progress updates
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
//		httpx.File("file", "/tmp/report.bin"),
//		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
//			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
//			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}, 200*time.Millisecond),
//	)
func UploadCallbackWithInterval(callback req.UploadCallback, minInterval time.Duration) Option {
	return Opts().UploadCallbackWithInterval(callback, minInterval)
}

// UploadProgress enables a default progress spinner and bar for uploads.
// @group Upload Options
//
// Example: upload with automatic progress
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
//		httpx.File("file", "/tmp/report.bin"),
//		httpx.UploadProgress(),
//	)
func UploadProgress() Option {
	return Opts().UploadProgress()
}

// OutputFile streams the response body to a file path.
// @group Download Options
//
// Example: download to file
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/file", httpx.OutputFile("/tmp/file.bin"))
func OutputFile(path string) Option {
	return Opts().OutputFile(path)
}
