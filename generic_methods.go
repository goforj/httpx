package httpx

import "context"

// Get issues a GET request and decodes its response.
// @group Requests
//
// Example: bind to a struct
//
//	type GetResponse struct {
//		URL string `json:"url"`
//	}
//
//	c := httpx.New()
//	res, _ := c.Get[GetResponse]("https://httpbin.org/get")
//	httpx.Dump(res) // dumps GetResponse
func (c *Client) Get[Out any](url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, nil, methodGet, url, nil, opts)
	return body, err
}

// Post issues a POST request and decodes its response.
// @group Requests
//
// Example: typed POST
//
//	type CreateUser struct { Name string `json:"name"` }
//	type CreateUserResponse struct { JSON CreateUser `json:"json"` }
//
//	c := httpx.New()
//	res, err := c.Post[CreateUserResponse]("https://httpbin.org/post", CreateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps CreateUserResponse
func (c *Client) Post[Out any](url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, nil, methodPost, url, body, opts)
	return out, err
}

// Put issues a PUT request and decodes its response.
// @group Requests
//
// Example: typed PUT
//
//	type UpdateUser struct { Name string `json:"name"` }
//	type UpdateUserResponse struct { JSON UpdateUser `json:"json"` }
//
//	c := httpx.New()
//	res, err := c.Put[UpdateUserResponse]("https://httpbin.org/put", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps UpdateUserResponse
func (c *Client) Put[Out any](url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, nil, methodPut, url, body, opts)
	return out, err
}

// Patch issues a PATCH request and decodes its response.
// @group Requests
//
// Example: typed PATCH
//
//	type UpdateUser struct { Name string `json:"name"` }
//	type UpdateUserResponse struct { JSON UpdateUser `json:"json"` }
//
//	c := httpx.New()
//	res, err := c.Patch[UpdateUserResponse]("https://httpbin.org/patch", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps UpdateUserResponse
func (c *Client) Patch[Out any](url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, nil, methodPatch, url, body, opts)
	return out, err
}

// Delete issues a DELETE request and decodes its response.
// @group Requests
//
// Example: typed DELETE
//
//	type DeleteResponse struct { URL string `json:"url"` }
//
//	c := httpx.New()
//	res, err := c.Delete[DeleteResponse]("https://httpbin.org/delete")
//	if err != nil {
//		return
//	}
//	httpx.Dump(res) // dumps DeleteResponse
func (c *Client) Delete[Out any](url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, nil, methodDelete, url, nil, opts)
	return body, err
}

// Head issues a HEAD request and decodes its response using the established response contract.
// @group Requests
func (c *Client) Head[Out any](url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, nil, methodHead, url, nil, opts)
	return body, err
}

// Options issues an OPTIONS request and decodes its response.
// @group Requests
func (c *Client) Options[Out any](url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, nil, methodOptions, url, nil, opts)
	return body, err
}

// GetCtx issues a context-aware GET request and decodes its response.
// @group Requests (Context)
func (c *Client) GetCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodGet, url, nil, opts)
	return body, err
}

// PostCtx issues a context-aware POST request and decodes its response.
// @group Requests (Context)
func (c *Client) PostCtx[Out any](ctx context.Context, url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, ctx, methodPost, url, body, opts)
	return out, err
}

// PutCtx issues a context-aware PUT request and decodes its response.
// @group Requests (Context)
func (c *Client) PutCtx[Out any](ctx context.Context, url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, ctx, methodPut, url, body, opts)
	return out, err
}

// PatchCtx issues a context-aware PATCH request and decodes its response.
// @group Requests (Context)
func (c *Client) PatchCtx[Out any](ctx context.Context, url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, ctx, methodPatch, url, body, opts)
	return out, err
}

// DeleteCtx issues a context-aware DELETE request and decodes its response.
// @group Requests (Context)
func (c *Client) DeleteCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodDelete, url, nil, opts)
	return body, err
}

// HeadCtx issues a context-aware HEAD request and decodes its response using the established response contract.
// @group Requests (Context)
func (c *Client) HeadCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodHead, url, nil, opts)
	return body, err
}

// OptionsCtx issues a context-aware OPTIONS request and decodes its response.
// @group Requests (Context)
func (c *Client) OptionsCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodOptions, url, nil, opts)
	return body, err
}
