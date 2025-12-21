package httpx

import (
	"io"

	"github.com/imroc/req/v3"
)

// Dump enables req's request-level dump output.
// @group Debugging
func (b OptionBuilder) Dump() OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableDump()
	}))
}

// DumpTo enables req's request-level dump output to a writer.
// @group Debugging
func (b OptionBuilder) DumpTo(output io.Writer) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableDumpTo(output)
	}))
}

// DumpToFile enables req's request-level dump output to a file path.
// @group Debugging
func (b OptionBuilder) DumpToFile(filename string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.EnableDumpToFile(filename)
	}))
}

// DumpAll enables req's client-level dump output for all requests.
// @group Debugging
func (b OptionBuilder) DumpAll() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.EnableDumpAll()
	}))
}

// DumpEachRequest enables request-level dumps for each request on the client.
// @group Debugging
func (b OptionBuilder) DumpEachRequest() OptionBuilder {
	return b.add(clientOnly(func(c *Client) {
		c.req.EnableDumpEachRequest()
	}))
}

// DumpEachRequestTo enables request-level dumps for each request and writes
// @group Debugging
// them to the provided output.
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
