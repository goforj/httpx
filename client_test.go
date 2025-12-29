package httpx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/imroc/req/v3"
)

type user struct {
	Name string `json:"name"`
}

type createUser struct {
	Name string `json:"name"`
}

func TestDefaultSharedClient(t *testing.T) {
	c1 := Default()
	c2 := Default()
	if c1 == nil || c2 == nil {
		t.Fatal("expected default client")
	}
	if c1 != c2 {
		t.Fatal("expected shared default client")
	}
}

func TestClientReqAndRaw(t *testing.T) {
	c := New()
	if c.Req() != c.req {
		t.Fatal("expected Req to return underlying client")
	}
	if c.Raw() != c.req {
		t.Fatal("expected Raw to return underlying client")
	}
}

func TestCloneNilAndNonNil(t *testing.T) {
	var nilClient *Client
	cloned := nilClient.clone()
	if cloned == nil || cloned.req == nil {
		t.Fatal("expected clone to return a new client")
	}

	base := New()
	base.errorMapper = func(*req.Response) error { return errors.New("mapped") }
	cloned2 := base.clone()
	if cloned2 == base || cloned2.req == base.req {
		t.Fatal("expected clone to duplicate client")
	}
	if cloned2.errorMapper == nil {
		t.Fatal("expected error mapper to be copied")
	}
}

func TestNilClientUsesDefault(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	t.Cleanup(server.Close)

	res, err := Get[string](nil, server.URL, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != "ok" {
		t.Fatalf("unexpected response: %s", res)
	}
}

func TestDoErrorWithSuccessEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	t.Cleanup(server.Close)

	client := New()
	res, err := Get[map[string]any](client, server.URL, Before(func(r *req.Request) {
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			return errors.New("after response")
		})
	}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected empty map")
	}
}

func TestGetTypedSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(w).Encode(user{Name: "roc"})
	}))
	t.Cleanup(server.Close)

	client := New()
	res, err := Get[user](client, server.URL+"/user")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Name != "roc" {
		t.Fatalf("unexpected response: %s", res.Name)
	}
}

func TestPostTypedSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload createUser
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(user{Name: payload.Name})
	}))
	t.Cleanup(server.Close)

	client := New()
	res, err := Post[createUser, user](client, server.URL, createUser{Name: "chris"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Name != "chris" {
		t.Fatalf("unexpected response: %s", res.Name)
	}
}

func TestRawString(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("plain"))
	}))
	t.Cleanup(server.Close)

	client := New()
	res, err := Get[string](client, server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != "plain" {
		t.Fatalf("unexpected response: %s", res)
	}
}

func TestHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad"))
	}))
	t.Cleanup(server.Close)

	client := New()
	_, err := Get[user](client, server.URL)
	if err == nil {
		t.Fatal("expected error")
	}

	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected HTTPError, got %T", err)
	}
	if httpErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("unexpected status: %d", httpErr.StatusCode)
	}
	if string(httpErr.Body) != "bad" {
		t.Fatalf("unexpected body: %s", string(httpErr.Body))
	}
}

func TestErrorMapper(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	t.Cleanup(server.Close)

	sentinel := errors.New("mapped")
	client := New()
	_, err := Get[user](client, server.URL, ErrorMapper(func(resp *req.Response) error {
		return sentinel
	}))
	if !errors.Is(err, sentinel) {
		t.Fatalf("expected mapped error, got %v", err)
	}
}

func TestHeadRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Fatalf("method = %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := New()
	_, err := Head[string](client, server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOptionsRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodOptions {
			t.Fatalf("method = %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := New()
	_, err := Options[string](client, server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHeadCtxRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Fatalf("method = %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	ctx := context.Background()
	client := New()
	_, err := HeadCtx[string](client, ctx, server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOptionsCtxRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodOptions {
			t.Fatalf("method = %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	ctx := context.Background()
	client := New()
	_, err := OptionsCtx[string](client, ctx, server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
func TestOptionsApplied(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/123" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.URL.Query().Get("q") != "ok" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Header.Get("X-Test") != "1" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(user{Name: "ok"})
	}))
	t.Cleanup(server.Close)

	client := New()
	res, err := Get[user](
		client,
		server.URL+"/users/{id}",
		Path("id", 123).Query("q", "ok", "ok2", "1").Header("X-Test", "1"),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Name != "ok" {
		t.Fatalf("unexpected response: %s", res.Name)
	}
}

func TestResultResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "1")
		_ = json.NewEncoder(w).Encode(user{Name: "ok"})
	}))
	t.Cleanup(server.Close)

	req := req.C().R().SetURL(server.URL)
	req.Method = http.MethodGet
	body, rawResp, err := Do[user](req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rawResp == nil || rawResp.Response == nil {
		t.Fatal("expected response")
	}
	if got := rawResp.Header.Get("X-Test"); got != "1" {
		t.Fatalf("unexpected header: %s", got)
	}
	if body.Name != "ok" {
		t.Fatalf("unexpected body: %s", body.Name)
	}
}

func TestDoNilRequest(t *testing.T) {
	_, _, err := Do[string](nil)
	if err == nil {
		t.Fatal("expected error for nil request")
	}
}

func TestDoHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	t.Cleanup(server.Close)

	req := req.C().R().SetURL(server.URL)
	req.Method = http.MethodGet
	_, _, err := Do[string](req)
	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("expected HTTPError, got %T", err)
	}
}

func TestDoRequestError(t *testing.T) {
	req := req.C().R().SetURL("http://127.0.0.1:0")
	req.Method = http.MethodGet
	_, _, err := Do[string](req)
	if err == nil {
		t.Fatal("expected request error")
	}
}

func TestDoRawStringSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	t.Cleanup(server.Close)

	req := req.C().R().SetURL(server.URL)
	req.Method = http.MethodGet
	res, _, err := Do[string](req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != "ok" {
		t.Fatalf("unexpected response: %s", res)
	}
}

func TestDoEmptyBodyEnsuresNonNil(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	t.Cleanup(server.Close)

	req := req.C().R().SetURL(server.URL)
	req.Method = http.MethodGet
	res, _, err := Do[map[string]any](req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected empty map")
	}
}

func TestEmptySliceOnEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := New()
	res, err := Get[[]user](client, server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("expected empty slice, got nil")
	}
	if len(res) != 0 {
		t.Fatalf("expected empty slice, got %d", len(res))
	}
}

func TestQueryOddPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()

	req := req.C().R()
	Query("q").applyRequest(req)
}

func TestRequestMethods(t *testing.T) {
	var gotMethod string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		_, _ = w.Write([]byte(`{"name":"ok"}`))
	}))
	t.Cleanup(server.Close)

	client := New()
	_, err := Put[createUser, user](client, server.URL, createUser{Name: "ok"})
	if err != nil {
		t.Fatalf("put error: %v", err)
	}
	if gotMethod != http.MethodPut {
		t.Fatalf("method = %q", gotMethod)
	}

	_, err = Patch[createUser, user](client, server.URL, createUser{Name: "ok"})
	if err != nil {
		t.Fatalf("patch error: %v", err)
	}
	if gotMethod != http.MethodPatch {
		t.Fatalf("method = %q", gotMethod)
	}

	_, err = Delete[user](client, server.URL)
	if err != nil {
		t.Fatalf("delete error: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Fatalf("method = %q", gotMethod)
	}
}

func TestContextMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"name":"ctx"}`))
	}))
	t.Cleanup(server.Close)

	ctx := context.Background()

	client := New()
	if _, err := GetCtx[user](client, ctx, server.URL); err != nil {
		t.Fatalf("getctx error: %v", err)
	}
	if _, err := PostCtx[createUser, user](client, ctx, server.URL, createUser{Name: "ctx"}); err != nil {
		t.Fatalf("postctx error: %v", err)
	}
	if _, err := PutCtx[createUser, user](client, ctx, server.URL, createUser{Name: "ctx"}); err != nil {
		t.Fatalf("putctx error: %v", err)
	}
	if _, err := PatchCtx[createUser, user](client, ctx, server.URL, createUser{Name: "ctx"}); err != nil {
		t.Fatalf("patchctx error: %v", err)
	}
	if _, err := DeleteCtx[user](client, ctx, server.URL); err != nil {
		t.Fatalf("deletectx error: %v", err)
	}
}

func TestSendUnknownMethod(t *testing.T) {
	_, err := send(req.C().R(), "TRACE-UNKNOWN", "http://example.com")
	if err == nil {
		t.Fatalf("expected error for unsupported method")
	}
}

func TestDumpExample(t *testing.T) {
	Dump("ok")
}

func TestDoIgnoresAfterResponseErrorOnEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := New()
	_, err := Get[user](client, server.URL, Before(func(r *req.Request) {
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			return errors.New("after response error")
		})
	}))
	if err != nil {
		t.Fatalf("expected error to be ignored, got %v", err)
	}
}

func TestNewWithHTTPTrace(t *testing.T) {
	t.Setenv("HTTP_TRACE", "1")
	c := New()
	clientVal := reflect.ValueOf(c.req).Elem()
	dumpField := clientVal.FieldByName("dumpOptions")
	if !dumpField.IsValid() {
		t.Fatalf("expected dumpOptions field")
	}
	if dumpField.IsNil() {
		t.Fatalf("expected dump options to be configured")
	}
}

func TestGetSkipsNilOption(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"name":"ok"}`))
	}))
	t.Cleanup(server.Close)

	client := New()
	_, err := Get[user](client, server.URL, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewSkipsNilOption(t *testing.T) {
	if c := New(nil); c == nil {
		t.Fatal("expected client")
	}
}
