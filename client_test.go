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
	res := Get[user](client, server.URL+"/user")
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body.Name != "roc" {
		t.Fatalf("unexpected response: %s", res.Body.Name)
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
	res := Post[createUser, user](client, server.URL, createUser{Name: "chris"})
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body.Name != "chris" {
		t.Fatalf("unexpected response: %s", res.Body.Name)
	}
}

func TestRawString(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("plain"))
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[string](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body != "plain" {
		t.Fatalf("unexpected response: %s", res.Body)
	}
}

func TestHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad"))
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[user](client, server.URL)
	if res.Err == nil {
		t.Fatal("expected error")
	}

	var httpErr *HTTPError
	if !errors.As(res.Err, &httpErr) {
		t.Fatalf("expected HTTPError, got %T", res.Err)
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
	client := New(ErrorMapper(func(resp *req.Response) error {
		return sentinel
	}))

	res := Get[user](client, server.URL)
	if !errors.Is(res.Err, sentinel) {
		t.Fatalf("expected mapped error, got %v", res.Err)
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
	res := Head[string](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
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
	res := Options[string](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
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
	res := Get[user](
		client,
		server.URL+"/users/{id}",
		Path("id", 123).Query("q", "ok", "ok2", "1").Header("X-Test", "1"),
	)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body.Name != "ok" {
		t.Fatalf("unexpected response: %s", res.Body.Name)
	}
}

func TestResultResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "1")
		_ = json.NewEncoder(w).Encode(user{Name: "ok"})
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[user](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Response == nil || res.Response.Response == nil {
		t.Fatal("expected response")
	}
	if got := res.Response.Header.Get("X-Test"); got != "1" {
		t.Fatalf("unexpected header: %s", got)
	}
	if res.Body.Name != "ok" {
		t.Fatalf("unexpected body: %s", res.Body.Name)
	}
}

func TestEmptySliceOnEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[[]user](client, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body == nil {
		t.Fatal("expected empty slice, got nil")
	}
	if len(res.Body) != 0 {
		t.Fatalf("expected empty slice, got %d", len(res.Body))
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

func TestDefaultReqRaw(t *testing.T) {
	c1 := Default()
	c2 := Default()
	if c1 == nil || c2 == nil || c1 != c2 {
		t.Fatalf("expected default client singleton")
	}
	if c1.Req() == nil {
		t.Fatalf("expected req client")
	}
	if c1.Raw() == nil {
		t.Fatalf("expected raw client")
	}
}

func TestRequestMethods(t *testing.T) {
	var gotMethod string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		_, _ = w.Write([]byte(`{"name":"ok"}`))
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Put[createUser, user](client, server.URL, createUser{Name: "ok"})
	if res.Err != nil {
		t.Fatalf("put error: %v", res.Err)
	}
	if gotMethod != http.MethodPut {
		t.Fatalf("method = %q", gotMethod)
	}

	res = Patch[createUser, user](client, server.URL, createUser{Name: "ok"})
	if res.Err != nil {
		t.Fatalf("patch error: %v", res.Err)
	}
	if gotMethod != http.MethodPatch {
		t.Fatalf("method = %q", gotMethod)
	}

	resDel := Delete[user](client, server.URL)
	if resDel.Err != nil {
		t.Fatalf("delete error: %v", resDel.Err)
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

	client := New()
	ctx := context.Background()

	if res := GetCtx[user](client, ctx, server.URL); res.Err != nil {
		t.Fatalf("getctx error: %v", res.Err)
	}
	if res := PostCtx[createUser, user](client, ctx, server.URL, createUser{Name: "ctx"}); res.Err != nil {
		t.Fatalf("postctx error: %v", res.Err)
	}
	if res := PutCtx[createUser, user](client, ctx, server.URL, createUser{Name: "ctx"}); res.Err != nil {
		t.Fatalf("putctx error: %v", res.Err)
	}
	if res := PatchCtx[createUser, user](client, ctx, server.URL, createUser{Name: "ctx"}); res.Err != nil {
		t.Fatalf("patchctx error: %v", res.Err)
	}
	if res := DeleteCtx[user](client, ctx, server.URL); res.Err != nil {
		t.Fatalf("deletectx error: %v", res.Err)
	}
}

func TestSendUnknownMethod(t *testing.T) {
	_, err := send(req.C().R(), "TRACE-UNKNOWN", "http://example.com")
	if err == nil {
		t.Fatalf("expected error for unsupported method")
	}
}

func TestDumpExample(t *testing.T) {
	dumpExample("ok")
}

func TestDoIgnoresAfterResponseErrorOnEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := New()
	res := Get[user](client, server.URL, Before(func(r *req.Request) {
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			return errors.New("after response error")
		})
	}))
	if res.Err != nil {
		t.Fatalf("expected error to be ignored, got %v", res.Err)
	}
	if res.Response == nil {
		t.Fatalf("expected response")
	}
}

func TestNilClientDefaults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(user{Name: "default"})
	}))
	t.Cleanup(server.Close)

	res := Get[user](nil, server.URL)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Body.Name != "default" {
		t.Fatalf("unexpected response: %s", res.Body.Name)
	}
}

func TestNewWithHTTPTrace(t *testing.T) {
	t.Setenv("HTTP_TRACE", "1")
	c := New()
	clientVal := reflect.ValueOf(c.Req()).Elem()
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
	res := Get[user](client, server.URL, nil)
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
}

func TestNewSkipsNilOption(t *testing.T) {
	if c := New(nil); c == nil {
		t.Fatal("expected client")
	}
}
