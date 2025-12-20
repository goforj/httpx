<p align="center">
  <img src="docs/images/logo.png" width="600" alt="httpx Logo">
</p>

A generics-first HTTP client wrapper for Go, built on top of the amazing `github.com/imroc/req/v3` library.
It keeps req's power and escape hatches, while making the 90% use case feel effortless.

<p align="center">
    <a href="https://pkg.go.dev/github.com/goforj/httpx"><img src="https://pkg.go.dev/badge/github.com/goforj/httpx.svg" alt="Go Reference"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT"></a>
    <a href="https://github.com/goforj/httpx/actions"><img src="https://github.com/goforj/httpx/actions/workflows/test.yml/badge.svg" alt="Go Test"></a>
    <a href="https://golang.org"><img src="https://img.shields.io/badge/go-1.18+-blue?logo=go" alt="Go version"></a>
    <img src="https://img.shields.io/github/v/tag/goforj/httpx?label=version&sort=semver" alt="Latest tag">
    <a href="https://goreportcard.com/report/github.com/goforj/httpx"><img src="https://goreportcard.com/badge/github.com/goforj/httpx" alt="Go Report Card"></a>
    <a href="https://codecov.io/gh/goforj/httpx" ><img src="https://codecov.io/gh/goforj/httpx/graph/badge.svg?token=R5O7LYAD4B"/></a>
<!-- test-count:embed:start -->
    <img src="https://img.shields.io/badge/tests-74-brightgreen" alt="Tests">
<!-- test-count:embed:end -->
    <a href="https://github.com/avelino/awesome-go?tab=readme-ov-file#parsersencodersdecoders"><img src="https://awesome.re/mentioned-badge-flat.svg" alt="Mentioned in Awesome Go"></a>
</p>

## Why httpx

- Typed, zero-ceremony requests with generics.
- Opinionated defaults (timeouts, result handling, safe error mapping).
- Built on req, with full escape hatches via `Client.Req()` and `Client.Raw()`.
- Ergonomic options for headers, query params, auth, retries, dumps, and uploads.

## Install

```bash
go get github.com/goforj/httpx
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"

	"github.com/goforj/httpx"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	c := httpx.New()

	// Simple typed GET.
	res := httpx.Get[User](c, "https://api.example.com/users/1")
	if res.Err != nil {
		panic(res.Err)
	}
	fmt.Println(res.Body.Name)

	// Context-aware GET.
	ctx := context.Background()
	res = httpx.GetCtx[User](c, ctx, "https://api.example.com/users/2")
	if res.Err != nil {
		panic(res.Err)
	}

	// Access the underlying response when you need it.
	_ = res.Response
}
```

## Escape Hatches (req is always available)

```go
c := httpx.New()

// Advanced req config.
c.Req().EnableDumpEachRequest()

// Drop down to raw req calls.
resp, err := c.Raw().R().Get("https://httpbin.org/uuid")
_, _ = resp, err
```

## Options in Practice

```go
c := httpx.New(httpx.WithBaseURL("https://api.example.com"))

res := httpx.Get[User](
	c,
	"/users/{id}",
	httpx.Path("id", "42"),
	httpx.Query("include", "teams", "active", "1"),
	httpx.Header("Accept", "application/json"),
)
_ = res
```

## Debugging and Tracing

- `HTTP_TRACE=1` enables request/response dumps for all requests.
- `httpx.Dump()` enables dump for a single request.
- `httpx.WithDumpEachRequest()` enables per-request dumps on a client.

## Examples

All runnable examples are generated from doc comments and live in `./examples`.
They are compiled by `example_compile_test.go` to keep docs and code in sync.

## Contributing

- Run `go run ./docs/examplegen` after updating doc examples.
- Run `go run ./docs/readme/main.go` to refresh the API index and test count.
- Run `go test ./...`.

<!-- api:embed:start -->

## API Index

| Group | Functions |
|------:|-----------|
| **Auth** | [Auth](#auth) [Basic](#basic) [Bearer](#bearer) |
| **Client** | [Default](#default) [New](#new) [Raw](#raw) [Req](#req) |
| **Client Options** | [WithBaseURL](#withbaseurl) [WithErrorMapper](#witherrormapper) [WithHeader](#withheader) [WithHeaders](#withheaders) [WithMiddleware](#withmiddleware) [WithTimeout](#withtimeout) [WithTransport](#withtransport) |
| **Debugging** | [Dump](#dump) [DumpTo](#dumpto) [DumpToFile](#dumptofile) [WithDumpAll](#withdumpall) [WithDumpEachRequest](#withdumpeachrequest) [WithDumpEachRequestTo](#withdumpeachrequestto) |
| **Download Options** | [OutputFile](#outputfile) |
| **Errors** | [Error](#error) |
| **Request Options** | [Before](#before) [Body](#body) [Form](#form) [Header](#header) [Headers](#headers) [JSON](#json) [Path](#path) [Paths](#paths) [Queries](#queries) [Query](#query) [Timeout](#timeout) |
| **Requests** | [Delete](#delete) [Get](#get) [Patch](#patch) [Post](#post) [Put](#put) |
| **Requests (Context)** | [DeleteCtx](#deletectx) [GetCtx](#getctx) [PatchCtx](#patchctx) [PostCtx](#postctx) [PutCtx](#putctx) |
| **Retry** | [RetryBackoff](#retrybackoff) [RetryCondition](#retrycondition) [RetryCount](#retrycount) [RetryFixedInterval](#retryfixedinterval) [RetryHook](#retryhook) [RetryInterval](#retryinterval) |
| **Retry (Client)** | [WithRetry](#withretry) [WithRetryBackoff](#withretrybackoff) [WithRetryCondition](#withretrycondition) [WithRetryCount](#withretrycount) [WithRetryFixedInterval](#withretryfixedinterval) [WithRetryHook](#withretryhook) [WithRetryInterval](#withretryinterval) |
| **Upload Options** | [File](#file) [FileBytes](#filebytes) [FileReader](#filereader) [Files](#files) [UploadCallback](#uploadcallback) [UploadCallbackWithInterval](#uploadcallbackwithinterval) [UploadProgress](#uploadprogress) |


## Auth

### <a id="auth"></a>Auth

Auth sets the Authorization header using a scheme and token.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Auth("Token", "abc123"))
```

### <a id="basic"></a>Basic

Basic sets HTTP basic authentication headers.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Basic("user", "pass"))
```

### <a id="bearer"></a>Bearer

Bearer sets the Authorization header with a bearer token.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Bearer("token"))
```

## Client

### <a id="default"></a>Default

Default returns the shared default client.

```go
c := httpx.Default()
_ = c
```

### <a id="new"></a>New

New creates a client with opinionated defaults and optional overrides.

```go
c := httpx.New(
	httpx.WithBaseURL("https://api.example.com"),
	httpx.WithTimeout(5*time.Second),
)
_ = c
```

### <a id="raw"></a>Raw

Raw returns the underlying req client for chaining raw requests.

```go
c := httpx.New()
resp, err := c.Raw().R().Get("https://httpbin.org/uuid")
_, _ = resp, err
```

### <a id="req"></a>Req

Req returns the underlying req client for advanced usage.

```go
c := httpx.New()
c.Req().EnableDumpEachRequest()
```

## Client Options

### <a id="withbaseurl"></a>WithBaseURL

WithBaseURL sets a base URL on the client.

```go
c := httpx.New(httpx.WithBaseURL("https://api.example.com"))
_ = c
```

### <a id="witherrormapper"></a>WithErrorMapper

WithErrorMapper sets a custom error mapper for non-2xx responses.

```go
c := httpx.New(httpx.WithErrorMapper(func(resp *req.Response) error {
	return fmt.Errorf("status %d", resp.StatusCode)
}))
_ = c
```

### <a id="withheader"></a>WithHeader

WithHeader sets a default header for all requests.

```go
c := httpx.New(httpx.WithHeader("X-Trace", "1"))
_ = c
```

### <a id="withheaders"></a>WithHeaders

WithHeaders sets default headers for all requests.

```go
c := httpx.New(httpx.WithHeaders(map[string]string{
	"X-Trace": "1",
	"Accept":  "application/json",
}))
_ = c
```

### <a id="withmiddleware"></a>WithMiddleware

WithMiddleware adds request middleware to the client.

```go
c := httpx.New(httpx.WithMiddleware(func(_ *req.Client, r *req.Request) error {
	r.SetHeader("X-Trace", "1")
	return nil
}))
_ = c
```

### <a id="withtimeout"></a>WithTimeout

WithTimeout sets the default timeout for the client.

```go
c := httpx.New(httpx.WithTimeout(3 * time.Second))
_ = c
```

### <a id="withtransport"></a>WithTransport

WithTransport wraps the underlying transport with a custom RoundTripper.

```go
c := httpx.New(httpx.WithTransport(http.RoundTripper(http.DefaultTransport)))
_ = c
```

## Debugging

### <a id="dump"></a>Dump

Dump enables req's request-level dump output.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Dump())
```

### <a id="dumpto"></a>DumpTo

DumpTo enables req's request-level dump output to a writer.

```go
var buf bytes.Buffer
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.DumpTo(&buf))
```

### <a id="dumptofile"></a>DumpToFile

DumpToFile enables req's request-level dump output to a file path.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.DumpToFile("httpx.dump"))
```

### <a id="withdumpall"></a>WithDumpAll

WithDumpAll enables req's client-level dump output for all requests.

```go
c := httpx.New(httpx.WithDumpAll())
_ = c
```

### <a id="withdumpeachrequest"></a>WithDumpEachRequest

WithDumpEachRequest enables request-level dumps for each request on the client.

```go
c := httpx.New(httpx.WithDumpEachRequest())
_ = c
```

### <a id="withdumpeachrequestto"></a>WithDumpEachRequestTo

WithDumpEachRequestTo enables request-level dumps for each request and writes

```go
var buf bytes.Buffer
c := httpx.New(httpx.WithDumpEachRequestTo(&buf))
_ = httpx.Get[string](c, "https://example.com")
_ = buf.String()
```

## Download Options

### <a id="outputfile"></a>OutputFile

OutputFile streams the response body to a file path.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com/file", httpx.OutputFile("/tmp/file.bin"))
```

## Errors

### <a id="error"></a>Error

Error returns a short, human-friendly summary of the HTTP error.

```go
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
res := httpx.Get[User](c, "https://example.com/users/1")
var httpErr *httpx.HTTPError
if errors.As(res.Err, &httpErr) {
	_ = httpErr.StatusCode
}
```

## Request Options

### <a id="before"></a>Before

Before runs a hook before the request is sent.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Before(func(r *req.Request) {
	r.EnableDump()
}))
```

### <a id="body"></a>Body

Body sets the request body and infers JSON for structs and maps.

```go
type Payload struct {
	Name string `json:"name"`
}

c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Body(Payload{Name: "Ana"}))
```

### <a id="form"></a>Form

Form sets form data for the request.

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.Form(map[string]string{
	"name": "Ana",
}))
```

### <a id="header"></a>Header

Header sets a header on a request.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Header("X-Trace", "1"))
```

### <a id="headers"></a>Headers

Headers sets multiple headers on a request.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Headers(map[string]string{
	"X-Trace": "1",
	"Accept":  "application/json",
}))
```

### <a id="json"></a>JSON

JSON sets the request body as JSON.

```go
type Payload struct {
	Name string `json:"name"`
}

c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com", nil, httpx.JSON(Payload{Name: "Ana"}))
```

### <a id="path"></a>Path

Path sets a path parameter by name.

```go
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
_ = httpx.Get[User](c, "https://example.com/users/{id}", httpx.Path("id", 42))
```

### <a id="paths"></a>Paths

Paths sets multiple path parameters.

```go
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
_ = httpx.Get[User](c, "https://example.com/orgs/{org}/users/{id}", httpx.Paths(map[string]any{
	"org": "goforj",
	"id":  42,
}))
```

### <a id="queries"></a>Queries

Queries adds multiple query parameters.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com/search", httpx.Queries(map[string]string{
	"q":  "go",
	"ok": "1",
}))
```

### <a id="query"></a>Query

Query adds query parameters as key/value pairs.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com/search", httpx.Query("q", "go", "ok", "1"))
```

### <a id="timeout"></a>Timeout

Timeout sets a per-request timeout using context cancellation.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Timeout(2*time.Second))
```

## Requests

### <a id="delete"></a>Delete

Delete issues a DELETE request using the provided client.

```go
type DeleteResponse struct {
	OK bool `json:"ok"`
}

c := httpx.New()
res := httpx.Delete[DeleteResponse](c, "https://api.example.com/users/1")
_, _ = res.Body, res.Err
```

### <a id="get"></a>Get

Get issues a GET request using the provided client.

```go
type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

c := httpx.New(httpx.WithHeader("Accept", "application/vnd.github+json"))
res := httpx.Get[[]PullRequest](c, "https://api.github.com/repos/goforj/httpx/pulls")
if res.Err != nil {
	return
}
godump.Dump(res.Body)
```

### <a id="patch"></a>Patch

Patch issues a PATCH request using the provided client.

```go
type UpdateUser struct {
	Name string `json:"name"`
}
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
res := httpx.Patch[UpdateUser, User](c, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
_, _ = res.Body, res.Err
```

### <a id="post"></a>Post

Post issues a POST request using the provided client.

```go
type CreateUser struct {
	Name string `json:"name"`
}
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
res := httpx.Post[CreateUser, User](c, "https://api.example.com/users", CreateUser{Name: "Ana"})
_, _ = res.Body, res.Err
```

### <a id="put"></a>Put

Put issues a PUT request using the provided client.

```go
type UpdateUser struct {
	Name string `json:"name"`
}
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
res := httpx.Put[UpdateUser, User](c, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
_, _ = res.Body, res.Err
```

## Requests (Context)

### <a id="deletectx"></a>DeleteCtx

DeleteCtx issues a DELETE request using the provided client and context.

```go
type DeleteResponse struct {
	OK bool `json:"ok"`
}

c := httpx.New()
ctx := context.Background()
res := httpx.DeleteCtx[DeleteResponse](c, ctx, "https://api.example.com/users/1")
_, _ = res.Body, res.Err
```

### <a id="getctx"></a>GetCtx

GetCtx issues a GET request using the provided client and context.

```go
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
ctx := context.Background()
res := httpx.GetCtx[User](c, ctx, "https://api.example.com/users/1")
_, _ = res.Body, res.Err
```

### <a id="patchctx"></a>PatchCtx

PatchCtx issues a PATCH request using the provided client and context.

```go
type UpdateUser struct {
	Name string `json:"name"`
}
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
ctx := context.Background()
res := httpx.PatchCtx[UpdateUser, User](c, ctx, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
_, _ = res.Body, res.Err
```

### <a id="postctx"></a>PostCtx

PostCtx issues a POST request using the provided client and context.

```go
type CreateUser struct {
	Name string `json:"name"`
}
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
ctx := context.Background()
res := httpx.PostCtx[CreateUser, User](c, ctx, "https://api.example.com/users", CreateUser{Name: "Ana"})
_, _ = res.Body, res.Err
```

### <a id="putctx"></a>PutCtx

PutCtx issues a PUT request using the provided client and context.

```go
type UpdateUser struct {
	Name string `json:"name"`
}
type User struct {
	Name string `json:"name"`
}

c := httpx.New()
ctx := context.Background()
res := httpx.PutCtx[UpdateUser, User](c, ctx, "https://api.example.com/users/1", UpdateUser{Name: "Ana"})
_, _ = res.Body, res.Err
```

## Retry

### <a id="retrybackoff"></a>RetryBackoff

RetryBackoff sets a capped exponential backoff retry interval for a request.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
```

### <a id="retrycondition"></a>RetryCondition

RetryCondition sets the retry condition for a request.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
	return resp != nil && resp.StatusCode == 503
}))
```

### <a id="retrycount"></a>RetryCount

RetryCount enables retry for a request and sets the maximum retry count.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.RetryCount(2))
```

### <a id="retryfixedinterval"></a>RetryFixedInterval

RetryFixedInterval sets a fixed retry interval for a request.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
```

### <a id="retryhook"></a>RetryHook

RetryHook registers a retry hook for a request.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
```

### <a id="retryinterval"></a>RetryInterval

RetryInterval sets a custom retry interval function for a request.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
	return time.Duration(attempt) * 100 * time.Millisecond
}))
```

## Retry (Client)

### <a id="withretry"></a>WithRetry

WithRetry applies a retry configuration to the client.

```go
c := httpx.New(httpx.WithRetry(func(rc *req.Client) {
	rc.SetCommonRetryCount(2)
}))
_ = c
```

### <a id="withretrybackoff"></a>WithRetryBackoff

WithRetryBackoff sets a capped exponential backoff retry interval for the client.

```go
c := httpx.New(httpx.WithRetryBackoff(100*time.Millisecond, 2*time.Second))
_ = c
```

### <a id="withretrycondition"></a>WithRetryCondition

WithRetryCondition sets the retry condition for the client.

```go
c := httpx.New(httpx.WithRetryCondition(func(resp *req.Response, _ error) bool {
	return resp != nil && resp.StatusCode == 503
}))
_ = c
```

### <a id="withretrycount"></a>WithRetryCount

WithRetryCount enables retry for the client and sets the maximum retry count.

```go
c := httpx.New(httpx.WithRetryCount(2))
_ = c
```

### <a id="withretryfixedinterval"></a>WithRetryFixedInterval

WithRetryFixedInterval sets a fixed retry interval for the client.

```go
c := httpx.New(httpx.WithRetryFixedInterval(200 * time.Millisecond))
_ = c
```

### <a id="withretryhook"></a>WithRetryHook

WithRetryHook registers a retry hook for the client.

```go
c := httpx.New(httpx.WithRetryHook(func(_ *req.Response, _ error) {}))
_ = c
```

### <a id="withretryinterval"></a>WithRetryInterval

WithRetryInterval sets a custom retry interval function for the client.

```go
c := httpx.New(httpx.WithRetryInterval(func(_ *req.Response, attempt int) time.Duration {
	return time.Duration(attempt) * 100 * time.Millisecond
}))
_ = c
```

## Upload Options

### <a id="file"></a>File

File attaches a file from disk as multipart form data.

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.File("file", "/tmp/report.txt"))
```

### <a id="filebytes"></a>FileBytes

FileBytes attaches a file from bytes as multipart form data.

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileBytes("file", "report.txt", []byte("hello")))
```

### <a id="filereader"></a>FileReader

FileReader attaches a file from a reader as multipart form data.

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileReader("file", "report.txt", strings.NewReader("hello")))
```

### <a id="files"></a>Files

Files attaches multiple files from disk as multipart form data.

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.Files(map[string]string{
	"fileA": "/tmp/a.txt",
	"fileB": "/tmp/b.txt",
}))
```

### <a id="uploadcallback"></a>UploadCallback

UploadCallback registers a callback for upload progress.

_Example: track upload progress_

```go
srv1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := make([]byte, 32*1024)
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			time.Sleep(10 * time.Millisecond)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}))
defer srv1.Close()

payload1 := bytes.Repeat([]byte("x"), 4*1024*1024)
pr1, pw1 := io.Pipe()
go func() {
	defer pw1.Close()
	chunk := 64 * 1024
	for i := 0; i < len(payload1); i += chunk {
		end := i + chunk
		if end > len(payload1) {
			end = len(payload1)
		}
		_, _ = pw1.Write(payload1[i:end])
		time.Sleep(50 * time.Millisecond)
	}
}()
c1 := httpx.New()
total1 := float64(len(payload1))
barWidth1 := 20
spin1 := []string{"|", "/", "-", "\\"}
spinIndex1 := 0
_ = httpx.Post[any, string](c1, srv1.URL+"/upload", nil,
	httpx.FileReader("file", "payload.bin", pr1),
	httpx.UploadCallback(func(info req.UploadInfo) {
		spinIndex1 = (spinIndex1 + 1) % len(spin1)
		percent := float64(info.UploadedSize) / total1 * 100
		filled := int(percent / 100 * float64(barWidth1))
		bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth1-filled)
		fmt.Printf("\r%s [%s] %.1f%%", spin1[spinIndex1], bar, percent)
		if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
			fmt.Print("\n")
		}
	}),
)
```

_Example: emit progress percent_

```go
srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := make([]byte, 32*1024)
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			time.Sleep(10 * time.Millisecond)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}))
defer srv2.Close()

payload2 := bytes.Repeat([]byte("x"), 4*1024*1024)
pr2, pw2 := io.Pipe()
go func() {
	defer pw2.Close()
	chunk := 64 * 1024
	for i := 0; i < len(payload2); i += chunk {
		end := i + chunk
		if end > len(payload2) {
			end = len(payload2)
		}
		_, _ = pw2.Write(payload2[i:end])
		time.Sleep(50 * time.Millisecond)
	}
}()
c2 := httpx.New()
total2 := float64(len(payload2))
barWidth2 := 20
spin2 := []string{"|", "/", "-", "\\"}
spinIndex2 := 0
_ = httpx.Post[any, string](c2, srv2.URL+"/upload", nil,
	httpx.FileReader("file", "payload.bin", pr2),
	httpx.UploadCallback(func(info req.UploadInfo) {
		spinIndex2 = (spinIndex2 + 1) % len(spin2)
		percent := float64(info.UploadedSize) / total2 * 100
		filled := int(percent / 100 * float64(barWidth2))
		bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth2-filled)
		fmt.Printf("\r%s [%s] %.1f%%", spin2[spinIndex2], bar, percent)
		if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
			fmt.Print("\n")
		}
	}),
)
```

### <a id="uploadcallbackwithinterval"></a>UploadCallbackWithInterval

UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.

_Example: throttle upload progress updates_

```go
srv1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := make([]byte, 32*1024)
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			time.Sleep(20 * time.Millisecond)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}))
defer srv1.Close()

payload1 := bytes.Repeat([]byte("x"), 4*1024*1024)
pr1, pw1 := io.Pipe()
go func() {
	defer pw1.Close()
	chunk := 64 * 1024
	for i := 0; i < len(payload1); i += chunk {
		end := i + chunk
		if end > len(payload1) {
			end = len(payload1)
		}
		_, _ = pw1.Write(payload1[i:end])
		time.Sleep(50 * time.Millisecond)
	}
}()
c1 := httpx.New()
total1 := float64(len(payload1))
_ = httpx.Post[any, string](c1, srv1.URL+"/upload", nil,
	httpx.FileReader("file", "payload.bin", pr1),
	httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
		percent := float64(info.UploadedSize) / total1 * 100
		fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
		if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
			fmt.Print("\n")
		}
	}, 200*time.Millisecond),
)
```

_Example: report filename and bytes_

```go
srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := make([]byte, 32*1024)
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			time.Sleep(20 * time.Millisecond)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}))
defer srv2.Close()

payload2 := bytes.Repeat([]byte("x"), 4*1024*1024)
pr2, pw2 := io.Pipe()
go func() {
	defer pw2.Close()
	chunk := 64 * 1024
	for i := 0; i < len(payload2); i += chunk {
		end := i + chunk
		if end > len(payload2) {
			end = len(payload2)
		}
		_, _ = pw2.Write(payload2[i:end])
		time.Sleep(50 * time.Millisecond)
	}
}()
c2 := httpx.New()
total2 := float64(len(payload2))
_ = httpx.Post[any, string](c2, srv2.URL+"/upload", nil,
	httpx.FileReader("file", "payload.bin", pr2),
	httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
		percent := float64(info.UploadedSize) / total2 * 100
		fmt.Printf("\r%s: %.1f%% (%d bytes)", info.FileName, percent, info.UploadedSize)
		if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
			fmt.Print("\n")
		}
	}, 200*time.Millisecond),
)
```

### <a id="uploadprogress"></a>UploadProgress

UploadProgress enables a default progress spinner and bar for uploads.

```go
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := make([]byte, 32*1024)
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			time.Sleep(10 * time.Millisecond)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}))
defer srv.Close()

payload := bytes.Repeat([]byte("x"), 4*1024*1024)
pr, pw := io.Pipe()
go func() {
	defer pw.Close()
	chunk := 64 * 1024
	for i := 0; i < len(payload); i += chunk {
		end := i + chunk
		if end > len(payload) {
			end = len(payload)
		}
		_, _ = pw.Write(payload[i:end])
		time.Sleep(50 * time.Millisecond)
	}
}()

c := httpx.New()
_ = httpx.Post[any, string](c, srv.URL+"/upload", nil,
	httpx.FileReader("file", "payload.bin", pr),
	httpx.UploadProgress(),
)
```
<!-- api:embed:end -->
