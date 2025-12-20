package httpx

import "github.com/imroc/req/v3"

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
