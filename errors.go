package httpx

import (
	"fmt"
	"net/http"

	"github.com/imroc/req/v3"
)

type HTTPError struct {
	StatusCode int
	Status     string
	Body       []byte
	Header     http.Header
}

// Error returns a short, human-friendly summary of the HTTP error.
// @group Errors
//
// Example: check for HTTP errors
//
//	type User struct {
//		Name string `json:"name"`
//	}
//
//	c := httpx.New()
//	res, err := httpx.Get[map[string]any](c, "https://httpbin.org/status/404")
//	httpx.Dump(res) // dumps map[string]any
//	// map[string]interface {}(nil)
//	var httpErr *httpx.HTTPError
//	if errors.As(err, &httpErr) {
//		_ = httpErr.StatusCode
//	}
func (e *HTTPError) Error() string {
	return fmt.Sprintf("httpx: http %d %s", e.StatusCode, e.Status)
}

func newHTTPError(resp *req.Response) *HTTPError {
	if resp == nil {
		return &HTTPError{StatusCode: 0, Status: "missing response"}
	}
	return &HTTPError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       resp.Bytes(),
		Header:     resp.Header,
	}
}
