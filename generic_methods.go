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
//	res, err := c.Get[GetResponse]("https://httpbin.org/get")
//	fmt.Println(err == nil, res.URL)
//	// true https://httpbin.org/get
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
//	fmt.Println(res.JSON.Name)
//	// Ana
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
//	fmt.Println(res.JSON.Name)
//	// Ana
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
//	fmt.Println(res.JSON.Name)
//	// Ana
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
//	fmt.Println(res.URL)
//	// https://httpbin.org/delete
func (c *Client) Delete[Out any](url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, nil, methodDelete, url, nil, opts)
	return body, err
}

// Head issues a HEAD request and decodes its response using the established response contract.
// @group Requests
//
// Example: client HEAD request
//
//	c := httpx.New()
//	_, err := c.Head[[]byte]("https://httpbin.org/get")
//	fmt.Println(err == nil)
//	// true
func (c *Client) Head[Out any](url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, nil, methodHead, url, nil, opts)
	return body, err
}

// Options issues an OPTIONS request and decodes its response.
// @group Requests
//
// Example: client OPTIONS request
//
//	c := httpx.New()
//	_, err := c.Options[[]byte]("https://httpbin.org/get")
//	fmt.Println(err == nil)
//	// true
func (c *Client) Options[Out any](url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, nil, methodOptions, url, nil, opts)
	return body, err
}

// GetCtx issues a context-aware GET request and decodes its response.
// @group Requests (Context)
//
// Example: client context-aware GET
//
//	type GetResponse struct {
//		URL string `json:"url"`
//	}
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := c.GetCtx[GetResponse](ctx, "https://httpbin.org/get")
//	fmt.Println(err == nil, res.URL)
//	// true https://httpbin.org/get
func (c *Client) GetCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodGet, url, nil, opts)
	return body, err
}

// PostCtx issues a context-aware POST request and decodes its response.
// @group Requests (Context)
//
// Example: client context-aware POST
//
//	type CreateUser struct { Name string `json:"name"` }
//	type CreateUserResponse struct { JSON CreateUser `json:"json"` }
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := c.PostCtx[CreateUserResponse](ctx, "https://httpbin.org/post", CreateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	fmt.Println(res.JSON.Name)
//	// Ana
func (c *Client) PostCtx[Out any](ctx context.Context, url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, ctx, methodPost, url, body, opts)
	return out, err
}

// PutCtx issues a context-aware PUT request and decodes its response.
// @group Requests (Context)
//
// Example: client context-aware PUT
//
//	type UpdateUser struct { Name string `json:"name"` }
//	type UpdateUserResponse struct { JSON UpdateUser `json:"json"` }
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := c.PutCtx[UpdateUserResponse](ctx, "https://httpbin.org/put", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	fmt.Println(res.JSON.Name)
//	// Ana
func (c *Client) PutCtx[Out any](ctx context.Context, url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, ctx, methodPut, url, body, opts)
	return out, err
}

// PatchCtx issues a context-aware PATCH request and decodes its response.
// @group Requests (Context)
//
// Example: client context-aware PATCH
//
//	type UpdateUser struct { Name string `json:"name"` }
//	type UpdateUserResponse struct { JSON UpdateUser `json:"json"` }
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := c.PatchCtx[UpdateUserResponse](ctx, "https://httpbin.org/patch", UpdateUser{Name: "Ana"})
//	if err != nil {
//		return
//	}
//	fmt.Println(res.JSON.Name)
//	// Ana
func (c *Client) PatchCtx[Out any](ctx context.Context, url string, body any, opts ...Option) (Out, error) {
	out, _, err := do[Out](c, ctx, methodPatch, url, body, opts)
	return out, err
}

// DeleteCtx issues a context-aware DELETE request and decodes its response.
// @group Requests (Context)
//
// Example: client context-aware DELETE
//
//	type DeleteResponse struct {
//		URL string `json:"url"`
//	}
//	ctx := context.Background()
//	c := httpx.New()
//	res, err := c.DeleteCtx[DeleteResponse](ctx, "https://httpbin.org/delete")
//	if err != nil {
//		return
//	}
//	fmt.Println(res.URL)
//	// https://httpbin.org/delete
func (c *Client) DeleteCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodDelete, url, nil, opts)
	return body, err
}

// HeadCtx issues a context-aware HEAD request and decodes its response using the established response contract.
// @group Requests (Context)
//
// Example: client context-aware HEAD
//
//	ctx := context.Background()
//	c := httpx.New()
//	_, err := c.HeadCtx[[]byte](ctx, "https://httpbin.org/get")
//	fmt.Println(err == nil)
//	// true
func (c *Client) HeadCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodHead, url, nil, opts)
	return body, err
}

// OptionsCtx issues a context-aware OPTIONS request and decodes its response.
// @group Requests (Context)
//
// Example: client context-aware OPTIONS
//
//	ctx := context.Background()
//	c := httpx.New()
//	_, err := c.OptionsCtx[[]byte](ctx, "https://httpbin.org/get")
//	fmt.Println(err == nil)
//	// true
func (c *Client) OptionsCtx[Out any](ctx context.Context, url string, opts ...Option) (Out, error) {
	body, _, err := do[Out](c, ctx, methodOptions, url, nil, opts)
	return body, err
}
