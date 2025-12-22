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
func Dump() OptionBuilder {
	return OptionBuilder{}.Dump()
}

func (b OptionBuilder) Dump() OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableDump()
	}))
}

// DumpTo enables req's request-level dump output to a writer.
// @group Debugging
//
// Example: dump to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.DumpTo(&buf))
func DumpTo(output io.Writer) OptionBuilder {
	return OptionBuilder{}.DumpTo(output)
}

func (b OptionBuilder) DumpTo(output io.Writer) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableDumpTo(output)
	}))
}

// DumpToFile enables req's request-level dump output to a file path.
// @group Debugging
//
// Example: dump to a file
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com", httpx.DumpToFile("httpx.dump"))
func DumpToFile(filename string) OptionBuilder {
	return OptionBuilder{}.DumpToFile(filename)
}

func (b OptionBuilder) DumpToFile(filename string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableDumpToFile(filename)
	}))
}

// DumpAll enables req's client-level dump output for all requests.
// @group Debugging
//
// Example: dump every request and response
//
//	c := httpx.New(httpx.DumpAll())
//	_ = c
func DumpAll() OptionBuilder {
	return OptionBuilder{}.DumpAll()
}

func (b OptionBuilder) DumpAll() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.EnableDumpAll()
	}))
}

// DumpEachRequest enables request-level dumps for each request on the client.
// @group Debugging
//
// Example: dump each request as it is sent
//
//	c := httpx.New(httpx.DumpEachRequest())
//	_ = c
func DumpEachRequest() OptionBuilder {
	return OptionBuilder{}.DumpEachRequest()
}

func (b OptionBuilder) DumpEachRequest() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.EnableDumpEachRequest()
	}))
}

// DumpEachRequestTo enables request-level dumps for each request and writes them to the provided output.
// @group Debugging
//
// Example: dump each request to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New(httpx.DumpEachRequestTo(&buf))
//	_ = httpx.Get[string](c, "https://example.com")
//	_ = buf.String()
func DumpEachRequestTo(output io.Writer) OptionBuilder {
	return OptionBuilder{}.DumpEachRequestTo(output)
}

func (b OptionBuilder) DumpEachRequestTo(output io.Writer) OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
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
	}))
}
