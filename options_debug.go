package httpx

import (
	"io"

	"github.com/imroc/req/v3"
)

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
