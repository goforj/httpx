package httpx

import "github.com/imroc/req/v3"

// OutputFile streams the response body to a file path.
// @group Download Options
//
// Applies to individual requests only.
// Example: download to file
//
//	c := httpx.New()
//	_ = httpx.Get[string](c, "https://example.com/file", httpx.OutputFile("/tmp/file.bin"))
func OutputFile(path string) OptionBuilder {
	return OptionBuilder{}.OutputFile(path)
}

func (b OptionBuilder) OutputFile(path string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetOutputFile(path)
	}))
}
