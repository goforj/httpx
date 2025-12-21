package httpx

import "github.com/imroc/req/v3"

// OutputFile streams the response body to a file path.
// @group Download Options
func (b OptionBuilder) OutputFile(path string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetOutputFile(path)
	}))
}
