package httpx

import (
	"io"

	"github.com/imroc/req/v3"
)

// EnableDump enables req's request-level dump output.
// @group Debugging
//
// Applies to individual requests only.
// Example: dump a single request
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.EnableDump())
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func EnableDump() OptionBuilder {
	return OptionBuilder{}.EnableDump()
}

func (b OptionBuilder) EnableDump() OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableDump()
	}))
}

// DumpTo enables req's request-level dump output to a writer.
// @group Debugging
//
// Applies to individual requests only.
// Example: dump to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.DumpTo(&buf))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
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
// Applies to individual requests only.
// Example: dump to a file
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.DumpToFile("httpx.dump"))
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
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
// Applies to the client configuration only.
// Example: dump every request and response
//
//	c := httpx.New(httpx.DumpAll())
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
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
// Applies to the client configuration only.
// Example: dump each request as it is sent
//
//	c := httpx.New(httpx.DumpEachRequest())
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
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
// Applies to the client configuration only.
// Example: dump each request to a buffer
//
//	var buf bytes.Buffer
//	c := httpx.New(httpx.DumpEachRequestTo(&buf))
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
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

// Trace enables req's request-level trace output.
// @group Debugging
//
// Applies to individual requests only.
// Example: trace a single request
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid", httpx.Trace())
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func Trace() OptionBuilder {
	return OptionBuilder{}.Trace()
}

func (b OptionBuilder) Trace() OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableTrace()
	}))
}

// TraceAll enables req's client-level trace output for all requests.
// @group Debugging
//
// Applies to the client configuration only.
// Example: trace all requests
//
//	c := httpx.New(httpx.TraceAll())
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/uuid")
//	_ = err
//	httpx.Dump(res) // dumps map[string]any
//	// #map[string]interface {} {
//	//   uuid => "<uuid>" #string
//	// }
func TraceAll() OptionBuilder {
	return OptionBuilder{}.TraceAll()
}

func (b OptionBuilder) TraceAll() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.EnableTraceAll()
	}))
}
