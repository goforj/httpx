package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

// Header sets a header on a request.
// @group Request Options
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
	return func(r *req.Request) {
		r.SetHeaders(values)
	}
}

// Query adds query parameters as key/value pairs.
// @group Request Options
//
// Example: add query params
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go", "ok", "1"))
func Query(kv ...string) Option {
	return func(r *req.Request) {
		if len(kv)%2 != 0 {
			panic("httpx: Query expects even number of key/value arguments")
		}
		for i := 0; i < len(kv); i += 2 {
			r.AddQueryParam(kv[i], kv[i+1])
		}
	}
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
	return func(r *req.Request) {
		r.SetQueryParams(values)
	}
}

// Auth sets the Authorization header using a scheme and token.
// @group Auth
//
// Example: custom auth scheme
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.Auth("Token", "abc123"))
func Auth(scheme, token string) Option {
	return func(r *req.Request) {
		r.SetHeader("Authorization", scheme+" "+token)
	}
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
	return func(r *req.Request) {
		r.SetPathParam(key, fmt.Sprint(value))
	}
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
	return func(r *req.Request) {
		params := make(map[string]string, len(values))
		for key, value := range values {
			params[key] = fmt.Sprint(value)
		}
		r.SetPathParams(params)
	}
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
	return func(r *req.Request) {
		setBody(r, value)
	}
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
	return func(r *req.Request) {
		r.SetBodyJsonMarshal(value)
	}
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
	return func(r *req.Request) {
		r.SetFormData(values)
	}
}

// File attaches a file from disk as multipart form data.
// @group Upload Options
//
// Example: upload a file
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.File("file", "/tmp/report.txt"))
func File(paramName, filePath string) Option {
	return func(r *req.Request) {
		r.SetFile(paramName, filePath)
	}
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
	return func(r *req.Request) {
		r.SetFiles(files)
	}
}

// FileBytes attaches a file from bytes as multipart form data.
// @group Upload Options
//
// Example: upload bytes as a file
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileBytes("file", "report.txt", []byte("hello")))
func FileBytes(paramName, filename string, content []byte) Option {
	return func(r *req.Request) {
		r.SetFileBytes(paramName, filename, content)
	}
}

// FileReader attaches a file from a reader as multipart form data.
// @group Upload Options
//
// Example: upload from reader
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileReader("file", "report.txt", strings.NewReader("hello")))
func FileReader(paramName, filename string, reader io.Reader) Option {
	return func(r *req.Request) {
		r.SetFileReader(paramName, filename, reader)
	}
}

// UploadCallback registers a callback for upload progress.
// @group Upload Options
//
// Example: track upload progress
//
//	srv1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		buf := make([]byte, 32*1024)
//		for {
//			n, err := r.Body.Read(buf)
//			if n > 0 {
//				time.Sleep(10 * time.Millisecond)
//			}
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				return
//			}
//		}
//		w.WriteHeader(http.StatusOK)
//	}))
//	defer srv1.Close()
//
//	payload1 := bytes.Repeat([]byte("x"), 4*1024*1024)
//	pr1, pw1 := io.Pipe()
//	go func() {
//		defer pw1.Close()
//		chunk := 64 * 1024
//		for i := 0; i < len(payload1); i += chunk {
//			end := i + chunk
//			if end > len(payload1) {
//				end = len(payload1)
//			}
//			_, _ = pw1.Write(payload1[i:end])
//			time.Sleep(50 * time.Millisecond)
//		}
//	}()
//	c1 := httpx.New()
//	total1 := float64(len(payload1))
//	barWidth1 := 20
//	spin1 := []string{"|", "/", "-", "\\"}
//	spinIndex1 := 0
//	_ = httpx.Post[any, string](c1, srv1.URL+"/upload", nil,
//		httpx.FileReader("file", "payload.bin", pr1),
//		httpx.UploadCallback(func(info req.UploadInfo) {
//			spinIndex1 = (spinIndex1 + 1) % len(spin1)
//			percent := float64(info.UploadedSize) / total1 * 100
//			filled := int(percent / 100 * float64(barWidth1))
//			bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth1-filled)
//			fmt.Printf("\r%s [%s] %.1f%%", spin1[spinIndex1], bar, percent)
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}),
//	)
//
// Example: emit progress percent
//
//	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		buf := make([]byte, 32*1024)
//		for {
//			n, err := r.Body.Read(buf)
//			if n > 0 {
//				time.Sleep(10 * time.Millisecond)
//			}
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				return
//			}
//		}
//		w.WriteHeader(http.StatusOK)
//	}))
//	defer srv2.Close()
//
//	payload2 := bytes.Repeat([]byte("x"), 4*1024*1024)
//	pr2, pw2 := io.Pipe()
//	go func() {
//		defer pw2.Close()
//		chunk := 64 * 1024
//		for i := 0; i < len(payload2); i += chunk {
//			end := i + chunk
//			if end > len(payload2) {
//				end = len(payload2)
//			}
//			_, _ = pw2.Write(payload2[i:end])
//			time.Sleep(50 * time.Millisecond)
//		}
//	}()
//	c2 := httpx.New()
//	total2 := float64(len(payload2))
//	barWidth2 := 20
//	spin2 := []string{"|", "/", "-", "\\"}
//	spinIndex2 := 0
//	_ = httpx.Post[any, string](c2, srv2.URL+"/upload", nil,
//		httpx.FileReader("file", "payload.bin", pr2),
//		httpx.UploadCallback(func(info req.UploadInfo) {
//			spinIndex2 = (spinIndex2 + 1) % len(spin2)
//			percent := float64(info.UploadedSize) / total2 * 100
//			filled := int(percent / 100 * float64(barWidth2))
//			bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth2-filled)
//			fmt.Printf("\r%s [%s] %.1f%%", spin2[spinIndex2], bar, percent)
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}),
//	)
func UploadCallback(callback req.UploadCallback) Option {
	return func(r *req.Request) {
		if callback == nil {
			return
		}
		var mu sync.Mutex
		var last req.UploadInfo
		var seen bool
		var completed bool
		r.SetUploadCallback(func(info req.UploadInfo) {
			mu.Lock()
			last = info
			seen = true
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				completed = true
			}
			mu.Unlock()
			callback(info)
		})
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			mu.Lock()
			info := last
			seenLocal := seen
			completedLocal := completed
			mu.Unlock()
			if !seenLocal {
				return nil
			}
			if !completedLocal && info.FileSize > 0 {
				info.UploadedSize = info.FileSize
				callback(info)
			}
			return nil
		})
	}
}

// UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.
// @group Upload Options
//
// Example: throttle upload progress updates
//
//	srv1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		buf := make([]byte, 32*1024)
//		for {
//			n, err := r.Body.Read(buf)
//			if n > 0 {
//				time.Sleep(20 * time.Millisecond)
//			}
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				return
//			}
//		}
//		w.WriteHeader(http.StatusOK)
//	}))
//	defer srv1.Close()
//
//	payload1 := bytes.Repeat([]byte("x"), 4*1024*1024)
//	pr1, pw1 := io.Pipe()
//	go func() {
//		defer pw1.Close()
//		chunk := 64 * 1024
//		for i := 0; i < len(payload1); i += chunk {
//			end := i + chunk
//			if end > len(payload1) {
//				end = len(payload1)
//			}
//			_, _ = pw1.Write(payload1[i:end])
//			time.Sleep(50 * time.Millisecond)
//		}
//	}()
//	c1 := httpx.New()
//	total1 := float64(len(payload1))
//	_ = httpx.Post[any, string](c1, srv1.URL+"/upload", nil,
//		httpx.FileReader("file", "payload.bin", pr1),
//		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
//			percent := float64(info.UploadedSize) / total1 * 100
//			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}, 200*time.Millisecond),
//	)
//
// Example: report filename and bytes
//
//	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		buf := make([]byte, 32*1024)
//		for {
//			n, err := r.Body.Read(buf)
//			if n > 0 {
//				time.Sleep(20 * time.Millisecond)
//			}
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				return
//			}
//		}
//		w.WriteHeader(http.StatusOK)
//	}))
//	defer srv2.Close()
//
//	payload2 := bytes.Repeat([]byte("x"), 4*1024*1024)
//	pr2, pw2 := io.Pipe()
//	go func() {
//		defer pw2.Close()
//		chunk := 64 * 1024
//		for i := 0; i < len(payload2); i += chunk {
//			end := i + chunk
//			if end > len(payload2) {
//				end = len(payload2)
//			}
//			_, _ = pw2.Write(payload2[i:end])
//			time.Sleep(50 * time.Millisecond)
//		}
//	}()
//	c2 := httpx.New()
//	total2 := float64(len(payload2))
//	_ = httpx.Post[any, string](c2, srv2.URL+"/upload", nil,
//		httpx.FileReader("file", "payload.bin", pr2),
//		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
//			percent := float64(info.UploadedSize) / total2 * 100
//			fmt.Printf("\r%s: %.1f%% (%d bytes)", info.FileName, percent, info.UploadedSize)
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}, 200*time.Millisecond),
//	)
func UploadCallbackWithInterval(callback req.UploadCallback, minInterval time.Duration) Option {
	return func(r *req.Request) {
		if callback == nil {
			return
		}
		var mu sync.Mutex
		var last req.UploadInfo
		var seen bool
		var completed bool
		r.SetUploadCallbackWithInterval(func(info req.UploadInfo) {
			mu.Lock()
			last = info
			seen = true
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				completed = true
			}
			mu.Unlock()
			callback(info)
		}, minInterval)
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			mu.Lock()
			info := last
			seenLocal := seen
			completedLocal := completed
			mu.Unlock()
			if !seenLocal {
				return nil
			}
			if !completedLocal && info.FileSize > 0 {
				info.UploadedSize = info.FileSize
				callback(info)
			}
			return nil
		})
	}
}

// OutputFile streams the response body to a file path.
// @group Download Options
//
// Example: download to file
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/file", httpx.OutputFile("/tmp/file.bin"))
func OutputFile(path string) Option {
	return func(r *req.Request) {
		r.SetOutputFile(path)
	}
}

// UploadProgress enables a default progress spinner and bar for uploads.
// @group Upload Options
//
// Example: upload with automatic progress
//
//	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		buf := make([]byte, 32*1024)
//		for {
//			n, err := r.Body.Read(buf)
//			if n > 0 {
//				time.Sleep(10 * time.Millisecond)
//			}
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				return
//			}
//		}
//		w.WriteHeader(http.StatusOK)
//	}))
//	defer srv.Close()
//
//	payload := bytes.Repeat([]byte("x"), 4*1024*1024)
//	pr, pw := io.Pipe()
//	go func() {
//		defer pw.Close()
//		chunk := 64 * 1024
//		for i := 0; i < len(payload); i += chunk {
//			end := i + chunk
//			if end > len(payload) {
//				end = len(payload)
//			}
//			_, _ = pw.Write(payload[i:end])
//			time.Sleep(50 * time.Millisecond)
//		}
//	}()
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, srv.URL+"/upload", nil,
//		httpx.FileReader("file", "payload.bin", pr),
//		httpx.UploadProgress(),
//	)
func UploadProgress() Option {
	return func(r *req.Request) {
		var mu sync.Mutex
		var last string
		var total int64
		spin := []string{"|", "/", "-", "\\"}
		spinIndex := 0
		barWidth := 20

		r.SetUploadCallback(func(info req.UploadInfo) {
			mu.Lock()
			defer mu.Unlock()

			if info.FileSize > 0 {
				total = info.FileSize
			}
			spinIndex = (spinIndex + 1) % len(spin)

			if total > 0 {
				percent := float64(info.UploadedSize) / float64(total) * 100
				filled := int(percent / 100 * float64(barWidth))
				bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth-filled)
				last = fmt.Sprintf(
					"\r%s upload [%s] %.1f%% (%s/%s)",
					spin[spinIndex],
					bar,
					percent,
					formatBytes(info.UploadedSize),
					formatBytes(total),
				)
			} else {
				last = fmt.Sprintf("\r%s upload %s", spin[spinIndex], formatBytes(info.UploadedSize))
			}
			fmt.Print(last)
		})

		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			mu.Lock()
			defer mu.Unlock()
			if last == "" {
				return nil
			}
			if total > 0 {
				fmt.Printf(
					"\r%s upload [%s] 100.0%% (%s/%s)\n",
					spin[spinIndex],
					strings.Repeat("=", barWidth),
					formatBytes(total),
					formatBytes(total),
				)
			} else {
				fmt.Print("\n")
			}
			return nil
		})
	}
}

func formatBytes(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	units := []string{"KiB", "MiB", "GiB", "TiB", "PiB"}
	value := float64(size)
	unit := "B"
	for _, u := range units {
		value /= 1024
		unit = u
		if value < 1024 {
			break
		}
	}
	return fmt.Sprintf("%.1f %s", value, unit)
}

// Dump enables req's request-level dump output.
// @group Debugging
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
// @group Debugging
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
// @group Debugging
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
// @group Request Options
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
// @group Auth
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
// @group Auth
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
// @group Request Options
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
// @group Client Options
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
// @group Client Options
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
// @group Client Options
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
// @group Client Options
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

// RetryCount enables retry for a request and sets the maximum retry count.
// @group Retry
//
// Example: request retry count
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryCount(2))
func RetryCount(count int) Option {
	return func(r *req.Request) {
		r.SetRetryCount(count)
	}
}

// RetryFixedInterval sets a fixed retry interval for a request.
// @group Retry
//
// Example: request retry interval
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
func RetryFixedInterval(interval time.Duration) Option {
	return func(r *req.Request) {
		r.SetRetryFixedInterval(interval)
	}
}

// RetryBackoff sets a capped exponential backoff retry interval for a request.
// @group Retry
//
// Example: request retry backoff
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
func RetryBackoff(min, max time.Duration) Option {
	return func(r *req.Request) {
		r.SetRetryBackoffInterval(min, max)
	}
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
	return func(r *req.Request) {
		r.SetRetryInterval(fn)
	}
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
	return func(r *req.Request) {
		r.SetRetryCondition(condition)
	}
}

// RetryHook registers a retry hook for a request.
// @group Retry
//
// Example: hook on retry
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
func RetryHook(hook req.RetryHookFunc) Option {
	return func(r *req.Request) {
		r.SetRetryHook(hook)
	}
}

// WithRetry applies a retry configuration to the client.
// @group Retry (Client)
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

// WithRetryCount enables retry for the client and sets the maximum retry count.
// @group Retry (Client)
//
// Example: client retry count
//
//	c := httpx.New(httpx.WithRetryCount(2))
//	_ = c
func WithRetryCount(count int) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryCount(count)
	}
}

// WithRetryFixedInterval sets a fixed retry interval for the client.
// @group Retry (Client)
//
// Example: client retry interval
//
//	c := httpx.New(httpx.WithRetryFixedInterval(200 * time.Millisecond))
//	_ = c
func WithRetryFixedInterval(interval time.Duration) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryFixedInterval(interval)
	}
}

// WithRetryBackoff sets a capped exponential backoff retry interval for the client.
// @group Retry (Client)
//
// Example: client retry backoff
//
//	c := httpx.New(httpx.WithRetryBackoff(100*time.Millisecond, 2*time.Second))
//	_ = c
func WithRetryBackoff(min, max time.Duration) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryBackoffInterval(min, max)
	}
}

// WithRetryInterval sets a custom retry interval function for the client.
// @group Retry (Client)
//
// Example: client retry interval
//
//	c := httpx.New(httpx.WithRetryInterval(func(_ *req.Response, attempt int) time.Duration {
//		return time.Duration(attempt) * 100 * time.Millisecond
//	}))
//	_ = c
func WithRetryInterval(fn req.GetRetryIntervalFunc) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryInterval(fn)
	}
}

// WithRetryCondition sets the retry condition for the client.
// @group Retry (Client)
//
// Example: retry on 503
//
//	c := httpx.New(httpx.WithRetryCondition(func(resp *req.Response, _ error) bool {
//		return resp != nil && resp.StatusCode == 503
//	}))
//	_ = c
func WithRetryCondition(condition req.RetryConditionFunc) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryCondition(condition)
	}
}

// WithRetryHook registers a retry hook for the client.
// @group Retry (Client)
//
// Example: hook on retry
//
//	c := httpx.New(httpx.WithRetryHook(func(_ *req.Response, _ error) {}))
//	_ = c
func WithRetryHook(hook req.RetryHookFunc) ClientOption {
	return func(c *Client) {
		c.req.SetCommonRetryHook(hook)
	}
}

// WithTransport wraps the underlying transport with a custom RoundTripper.
// @group Client Options
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
// @group Debugging
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
// @group Debugging
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
// @group Debugging
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
// @group Client Options
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
// @group Client Options
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
