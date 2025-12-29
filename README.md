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
    <img src="https://img.shields.io/badge/tests-191-brightgreen" alt="Tests">
<!-- test-count:embed:end -->
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

## Browser Profiles

Browser profiles provide a simple way to match common client behavior without exposing low-level details.
Internally, profiles may apply transport and protocol settings as needed, but those details are intentionally abstracted.

```go
c := httpx.New(httpx.AsChrome())
_ = c
```

## Use Any req Feature

**httpx** is built on top of the incredible [req](https://github.com/imroc/req) library, and you can always drop down to it when you need something beyond httpx’s helpers. That means every example in req’s docs is available to you with `c.Req()` or `c.Raw()`.

While httpx provides ergonomic helpers for the most common use cases, req is a powerful and flexible HTTP client library with tons of features.

```go
c := httpx.New()

// Grab the underlying req client.
rc := c.Req()

// Now you can use any req feature from their docs.
// Example: enable trace, custom transports, cookie jars, proxies, etc.
rc.EnableTraceAll()
```

See the full req documentation here: https://req.cool/docs/prologue/quickstart/

Most users will only need the high-level APIs (browser profiles, request composition, retries, uploads).
When you need deep control over headers, transports, or protocol behavior, `req` is always there.

## Options in Practice

```go
c := httpx.New(httpx.BaseURL("https://api.example.com"))

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
- `httpx.DumpEachRequest()` enables per-request dumps on a client.

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
|------:|:-----------|
| **Auth** | [Auth](#auth) [Basic](#basic) [Bearer](#bearer) |
| **Browser Profiles** | [AsChrome](#aschrome) [AsFirefox](#asfirefox) [AsMobile](#asmobile) [AsSafari](#assafari) |
| **Client** | [Default](#default) [New](#new) [Raw](#raw) [Req](#req) |
| **Client Options** | [BaseURL](#baseurl) [CookieJar](#cookiejar) [ErrorMapper](#errormapper) [Middleware](#middleware) [Proxy](#proxy) [ProxyFunc](#proxyfunc) [Redirect](#redirect) [Transport](#transport) |
| **Debugging** | [Dump](#dump) [DumpAll](#dumpall) [DumpEachRequest](#dumpeachrequest) [DumpEachRequestTo](#dumpeachrequestto) [DumpTo](#dumpto) [DumpToFile](#dumptofile) [Trace](#trace) [TraceAll](#traceall) |
| **Download Options** | [OutputFile](#outputfile) |
| **Errors** | [Error](#error) |
| **Request Composition** | [Body](#body) [Form](#form) [Header](#header) [Headers](#headers) [JSON](#json) [Path](#path) [Paths](#paths) [Queries](#queries) [Query](#query) [UserAgent](#useragent) |
| **Request Control** | [Before](#before) [Timeout](#timeout) |
| **Requests** | [Delete](#delete) [Get](#get) [Head](#head) [Options](#options) [Patch](#patch) [Post](#post) [Put](#put) |
| **Requests (Context)** | [DeleteCtx](#deletectx) [GetCtx](#getctx) [HeadCtx](#headctx) [OptionsCtx](#optionsctx) [PatchCtx](#patchctx) [PostCtx](#postctx) [PutCtx](#putctx) |
| **Retry** | [RetryBackoff](#retrybackoff) [RetryCondition](#retrycondition) [RetryCount](#retrycount) [RetryFixedInterval](#retryfixedinterval) [RetryHook](#retryhook) [RetryInterval](#retryinterval) |
| **Retry (Client)** | [Retry](#retry) |
| **Upload Options** | [File](#file) [FileBytes](#filebytes) [FileReader](#filereader) [Files](#files) [UploadCallback](#uploadcallback) [UploadCallbackWithInterval](#uploadcallbackwithinterval) [UploadProgress](#uploadprogress) |
| **Advanced** | [TLSFingerprint](#tlsfingerprint) [TLSFingerprintAndroid](#tlsfingerprintandroid) [TLSFingerprintChrome](#tlsfingerprintchrome) [TLSFingerprintEdge](#tlsfingerprintedge) [TLSFingerprintFirefox](#tlsfingerprintfirefox) [TLSFingerprintIOS](#tlsfingerprintios) [TLSFingerprintRandomized](#tlsfingerprintrandomized) [TLSFingerprintSafari](#tlsfingerprintsafari) |


## Auth

### <a id="auth"></a>Auth

Auth sets the Authorization header using a scheme and token.

```go
// Apply to all requests
c := httpx.New(httpx.Auth("Token", "abc123"))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.Auth("Token", "abc123"))
```

### <a id="basic"></a>Basic

Basic sets HTTP basic authentication headers.

```go
// Apply to all requests
c := httpx.New(httpx.Basic("user", "pass"))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.Basic("user", "pass"))
```

### <a id="bearer"></a>Bearer

Bearer sets the Authorization header with a bearer token.

```go
// Apply to all requests
c := httpx.New(httpx.Bearer("token"))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.Bearer("token"))
```

## Browser Profiles

### <a id="aschrome"></a>AsChrome

AsChrome applies the Chrome browser profile (headers, TLS, and HTTP/2 behavior).

```go
c := httpx.New(httpx.AsChrome())
_ = c
```

### <a id="asfirefox"></a>AsFirefox

AsFirefox applies the Firefox browser profile (headers, TLS, and HTTP/2 behavior).

```go
c := httpx.New(httpx.AsFirefox())
_ = c
```

### <a id="asmobile"></a>AsMobile

AsMobile applies a mobile Chrome-like profile (headers, TLS, and HTTP/2 behavior).

```go
c := httpx.New(httpx.AsMobile())
_ = c
```

### <a id="assafari"></a>AsSafari

AsSafari applies the Safari browser profile (headers, TLS, and HTTP/2 behavior).

```go
c := httpx.New(httpx.AsSafari())
_ = c
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
var buf bytes.Buffer
c := httpx.New(httpx.
	BaseURL("https://api.example.com").
	Timeout(5*time.Second).
	Header("X-Trace", "1").
	Headers(map[string]string{
		"Accept": "application/json",
	}).
	Transport(http.RoundTripper(http.DefaultTransport)).
	Middleware(func(_ *req.Client, r *req.Request) error {
		r.SetHeader("X-Middleware", "1")
		return nil
	}).
	ErrorMapper(func(resp *req.Response) error {
		return fmt.Errorf("status %d", resp.StatusCode)
	}).
	DumpAll().
	DumpEachRequest().
	DumpEachRequestTo(&buf).
	Retry(func(rc *req.Client) {
		rc.SetCommonRetryCount(2)
	}).
	RetryCount(2).
	RetryFixedInterval(200 * time.Millisecond).
	RetryBackoff(100*time.Millisecond, 2*time.Second).
	RetryInterval(func(_ *req.Response, attempt int) time.Duration {
		return time.Duration(attempt) * 100 * time.Millisecond
	}).
	RetryCondition(func(resp *req.Response, _ error) bool {
		return resp != nil && resp.StatusCode == 503
	}).
	RetryHook(func(_ *req.Response, _ error) {}),
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

### <a id="baseurl"></a>BaseURL

BaseURL sets a base URL on the client.

```go
c := httpx.New(httpx.BaseURL("https://api.example.com"))
_ = c
```

### <a id="cookiejar"></a>CookieJar

CookieJar sets the cookie jar for the client.

```go
jar, _ := cookiejar.New(nil)
u, _ := url.Parse("https://example.com")
jar.SetCookies(u, []*http.Cookie{
	{Name: "session", Value: "abc123"},
})
c := httpx.New(httpx.CookieJar(jar))
_ = c
```

### <a id="errormapper"></a>ErrorMapper

ErrorMapper sets a custom error mapper for non-2xx responses.

```go
c := httpx.New(httpx.ErrorMapper(func(resp *req.Response) error {
	return fmt.Errorf("status %d", resp.StatusCode)
}))
_ = c
```

### <a id="middleware"></a>Middleware

Middleware adds request middleware to the client.

```go
c := httpx.New(httpx.Middleware(func(_ *req.Client, r *req.Request) error {
	r.SetHeader("X-Trace", "1")
	return nil
}))
_ = c
```

### <a id="proxy"></a>Proxy

Proxy sets a proxy URL for the client.

```go
c := httpx.New(httpx.Proxy("http://localhost:8080"))
_ = c
```

### <a id="proxyfunc"></a>ProxyFunc

ProxyFunc sets a proxy function for the client.

```go
c := httpx.New(httpx.ProxyFunc(http.ProxyFromEnvironment))
_ = c
```

### <a id="redirect"></a>Redirect

Redirect sets the redirect policy for the client.

```go
c := httpx.New(httpx.Redirect(req.NoRedirectPolicy()))
_ = c
```

### <a id="transport"></a>Transport

Transport wraps the underlying transport with a custom RoundTripper.

```go
c := httpx.New(httpx.Transport(http.RoundTripper(http.DefaultTransport)))
_ = c
```

## Debugging

### <a id="dump"></a>Dump

Dump enables req's request-level dump output.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Dump())
```

### <a id="dumpall"></a>DumpAll

DumpAll enables req's client-level dump output for all requests.

```go
c := httpx.New(httpx.DumpAll())
_ = c
```

### <a id="dumpeachrequest"></a>DumpEachRequest

DumpEachRequest enables request-level dumps for each request on the client.

```go
c := httpx.New(httpx.DumpEachRequest())
_ = c
```

### <a id="dumpeachrequestto"></a>DumpEachRequestTo

DumpEachRequestTo enables request-level dumps for each request and writes them to the provided output.

```go
var buf bytes.Buffer
c := httpx.New(httpx.DumpEachRequestTo(&buf))
_ = httpx.Get[string](c, "https://example.com")
_ = buf.String()
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

### <a id="trace"></a>Trace

Trace enables req's request-level trace output.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Trace())
```

### <a id="traceall"></a>TraceAll

TraceAll enables req's client-level trace output for all requests.

```go
c := httpx.New(httpx.TraceAll())
_ = c
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

## Request Composition

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

Header sets a header on a request or client.

```go
// Apply to all requests
c := httpx.New(httpx.Header("X-Trace", "1"))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.Header("X-Trace", "1"))
```

### <a id="headers"></a>Headers

Headers sets multiple headers on a request or client.

```go
// Apply to all requests
c := httpx.New(httpx.Headers(map[string]string{
	"X-Trace": "1",
	"Accept":  "application/json",
}))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.Headers(map[string]string{
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

### <a id="useragent"></a>UserAgent

UserAgent sets the User-Agent header on a request or client.

```go
// Apply to all requests
c := httpx.New(httpx.UserAgent("my-app/1.0"))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.UserAgent("my-app/1.0"))
```

## Request Control

### <a id="before"></a>Before

Before runs a hook before the request is sent.

```go
c := httpx.New()
_ = httpx.Get[string](c, "https://example.com", httpx.Before(func(r *req.Request) {
	r.EnableDump()
}))
```

### <a id="timeout"></a>Timeout

Timeout sets a per-request timeout using context cancellation.

```go
// Apply to all requests
c := httpx.New(httpx.Timeout(2 * time.Second))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.Timeout(2*time.Second))
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

c := httpx.New(httpx.Header("Accept", "application/vnd.github+json"))
res := httpx.Get[[]PullRequest](c, "https://api.github.com/repos/goforj/httpx/pulls")
if res.Err != nil {
	return
}
godump.Dump(res.Body)
```

### <a id="head"></a>Head

Head issues a HEAD request using the provided client.

```go
c := httpx.New()
res := httpx.Head[string](c, "https://example.com")
_ = res
```

### <a id="options"></a>Options

Options issues an OPTIONS request using the provided client.

```go
c := httpx.New()
res := httpx.Options[string](c, "https://example.com")
_ = res
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

### <a id="headctx"></a>HeadCtx

HeadCtx issues a HEAD request using the provided client and context.

```go
c := httpx.New()
ctx := context.Background()
res := httpx.HeadCtx[string](c, ctx, "https://example.com")
_ = res
```

### <a id="optionsctx"></a>OptionsCtx

OptionsCtx issues an OPTIONS request using the provided client and context.

```go
c := httpx.New()
ctx := context.Background()
res := httpx.OptionsCtx[string](c, ctx, "https://example.com")
_ = res
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
// Apply to all requests
c := httpx.New(httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryBackoff(100*time.Millisecond, 2*time.Second))
```

### <a id="retrycondition"></a>RetryCondition

RetryCondition sets the retry condition for a request.

```go
// Apply to all requests
c := httpx.New(httpx.RetryCondition(func(resp *req.Response, _ error) bool {
	return resp != nil && resp.StatusCode == 503
}))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryCondition(func(resp *req.Response, _ error) bool {
	return resp != nil && resp.StatusCode == 503
}))
```

### <a id="retrycount"></a>RetryCount

RetryCount enables retry for a request and sets the maximum retry count.

```go
// Apply to all requests
c := httpx.New(httpx.RetryCount(2))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryCount(2))
```

### <a id="retryfixedinterval"></a>RetryFixedInterval

RetryFixedInterval sets a fixed retry interval for a request.

```go
// Apply to all requests
c := httpx.New(httpx.RetryFixedInterval(200 * time.Millisecond))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryFixedInterval(200*time.Millisecond))
```

### <a id="retryhook"></a>RetryHook

RetryHook registers a retry hook for a request.

```go
// Apply to all requests
c := httpx.New(httpx.RetryHook(func(_ *req.Response, _ error) {}))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryHook(func(_ *req.Response, _ error) {}))
```

### <a id="retryinterval"></a>RetryInterval

RetryInterval sets a custom retry interval function for a request.

```go
// Apply to all requests
c := httpx.New(httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
	return time.Duration(attempt) * 100 * time.Millisecond
}))
httpx.Get[string](c, "https://example.com")

// Apply to a single request
httpx.Get[string](httpx.Default(), "https://example.com", httpx.RetryInterval(func(_ *req.Response, attempt int) time.Duration {
	return time.Duration(attempt) * 100 * time.Millisecond
}))
```

## Retry (Client)

### <a id="retry"></a>Retry

Retry applies a custom retry configuration to the client.

```go
c := httpx.New(httpx.Retry(func(rc *req.Client) {
	rc.SetCommonRetryCount(2)
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

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
	httpx.File("file", "/tmp/report.bin"),
	httpx.UploadCallback(func(info req.UploadInfo) {
		percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
		fmt.Printf("\rprogress: %.1f%%", percent)
		if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
			fmt.Print("\n")
		}
	}),
)
```

### <a id="uploadcallbackwithinterval"></a>UploadCallbackWithInterval

UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
	httpx.File("file", "/tmp/report.bin"),
	httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
		percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
		fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
		if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
			fmt.Print("\n")
		}
	}, 200*time.Millisecond),
)
```

### <a id="uploadprogress"></a>UploadProgress

UploadProgress enables a default progress spinner and bar for uploads.

```go
c := httpx.New()
_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
	httpx.File("file", "/tmp/report.bin"),
	httpx.UploadProgress(),
)
```

## Advanced

### <a id="tlsfingerprint"></a>TLSFingerprint

TLSFingerprint applies a TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprint(httpx.TLSFingerprintChromeKind))
_ = c
```

### <a id="tlsfingerprintandroid"></a>TLSFingerprintAndroid

TLSFingerprintAndroid applies the Android TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprintAndroid())
_ = c
```

### <a id="tlsfingerprintchrome"></a>TLSFingerprintChrome

TLSFingerprintChrome applies the Chrome TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprintChrome())
_ = c
```

### <a id="tlsfingerprintedge"></a>TLSFingerprintEdge

TLSFingerprintEdge applies the Edge TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprintEdge())
_ = c
```

### <a id="tlsfingerprintfirefox"></a>TLSFingerprintFirefox

TLSFingerprintFirefox applies the Firefox TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprintFirefox())
_ = c
```

### <a id="tlsfingerprintios"></a>TLSFingerprintIOS

TLSFingerprintIOS applies the iOS TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprintIOS())
_ = c
```

### <a id="tlsfingerprintrandomized"></a>TLSFingerprintRandomized

TLSFingerprintRandomized applies a randomized TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprintRandomized())
_ = c
```

### <a id="tlsfingerprintsafari"></a>TLSFingerprintSafari

TLSFingerprintSafari applies the Safari TLS fingerprint preset.

```go
c := httpx.New(httpx.TLSFingerprintSafari())
_ = c
```
<!-- api:embed:end -->
