package httpx

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	"github.com/imroc/req/v3"
)

func uploadServer(t *testing.T, handler func(*http.Request)) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if handler != nil {
			handler(r)
		}
		w.WriteHeader(http.StatusOK)
	}))
}

func TestFileAndFiles(t *testing.T) {
	fileA := writeTempFile(t, "fileA.txt", "alpha")
	fileB := writeTempFile(t, "fileB.txt", "bravo")
	defer os.Remove(fileA)
	defer os.Remove(fileB)

	var got map[string]string
	srv := uploadServer(t, func(r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			t.Fatalf("parse multipart: %v", err)
		}
		got = map[string]string{}
		for name, files := range r.MultipartForm.File {
			if len(files) == 0 {
				continue
			}
			f, err := files[0].Open()
			if err != nil {
				t.Fatalf("open file: %v", err)
			}
			data, _ := io.ReadAll(f)
			_ = f.Close()
			got[name] = string(data)
		}
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil,
		File("file", fileA).Files(map[string]string{"fileB": fileB}),
	)
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if got["file"] != "alpha" {
		t.Fatalf("file content = %q", got["file"])
	}
	if got["fileB"] != "bravo" {
		t.Fatalf("fileB content = %q", got["fileB"])
	}
}

func TestFilesFunction(t *testing.T) {
	file := writeTempFile(t, "payload.txt", "payload")
	defer os.Remove(file)

	var got string
	srv := uploadServer(t, func(r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			t.Fatalf("parse multipart: %v", err)
		}
		f, _, err := r.FormFile("file")
		if err != nil {
			t.Fatalf("form file: %v", err)
		}
		data, _ := io.ReadAll(f)
		_ = f.Close()
		got = string(data)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil, Files(map[string]string{"file": file}))
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if got != "payload" {
		t.Fatalf("file content = %q", got)
	}
}

func TestFileBytes(t *testing.T) {
	var content string
	srv := uploadServer(t, func(r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			t.Fatalf("parse multipart: %v", err)
		}
		f, _, err := r.FormFile("file")
		if err != nil {
			t.Fatalf("form file: %v", err)
		}
		data, _ := io.ReadAll(f)
		_ = f.Close()
		content = string(data)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil, FileBytes("file", "report.txt", []byte("hello")))
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if content != "hello" {
		t.Fatalf("content = %q", content)
	}
}

type sizeReader struct {
	r    io.Reader
	size int64
}

func (s *sizeReader) Read(p []byte) (int, error) { return s.r.Read(p) }
func (s *sizeReader) Size() int64                { return s.size }

type seekerReader struct {
	r io.ReadSeeker
}

func (s *seekerReader) Read(p []byte) (int, error) { return s.r.Read(p) }
func (s *seekerReader) Seek(o int64, w int) (int64, error) {
	return s.r.Seek(o, w)
}

type lenOnlyReader struct {
	buf *bytes.Buffer
}

func (l *lenOnlyReader) Read(p []byte) (int, error) { return l.buf.Read(p) }
func (l *lenOnlyReader) Len() int                   { return l.buf.Len() }

func TestFileReaderSizes(t *testing.T) {
	cases := []struct {
		name   string
		reader io.Reader
	}{
		{
			name:   "size",
			reader: &sizeReader{r: bytes.NewReader([]byte("abc")), size: 3},
		},
		{
			name:   "len",
			reader: &lenOnlyReader{buf: bytes.NewBufferString("abcd")},
		},
		{
			name:   "seeker",
			reader: &seekerReader{r: bytes.NewReader([]byte("abcde"))},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var gotSize int64
			srv := uploadServer(t, func(r *http.Request) {
				_, _ = io.Copy(io.Discard, r.Body)
			})
			defer srv.Close()

			c := New()
			res := Post[any, string](c, srv.URL, nil,
				FileReader("file", "payload.bin", tc.reader).UploadCallback(func(info req.UploadInfo) {
					if info.FileSize > 0 {
						gotSize = info.FileSize
					}
				}),
			)
			if res.Err != nil {
				t.Fatalf("upload failed: %v", res.Err)
			}
			if gotSize == 0 {
				t.Fatalf("file size not detected")
			}
		})
	}
}

type errorSeeker struct{}

func (e *errorSeeker) Read(p []byte) (int, error) { return 0, io.EOF }
func (e *errorSeeker) Seek(int64, int) (int64, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestFileReaderSeekError(t *testing.T) {
	var gotSize int64
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil,
		FileReader("file", "payload.bin", &errorSeeker{}).UploadCallback(func(info req.UploadInfo) {
			gotSize = info.FileSize
		}),
	)
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if gotSize != 0 {
		t.Fatalf("expected size 0, got %d", gotSize)
	}
}

type partialSeeker struct {
	seekCount int
}

func (p *partialSeeker) Read([]byte) (int, error) { return 0, io.EOF }

func (p *partialSeeker) Seek(int64, int) (int64, error) {
	p.seekCount++
	if p.seekCount == 1 {
		return 0, nil
	}
	return 0, io.ErrUnexpectedEOF
}

func TestFileReaderSeekEndError(t *testing.T) {
	var gotSize int64
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil,
		FileReader("file", "payload.bin", &partialSeeker{}).UploadCallback(func(info req.UploadInfo) {
			gotSize = info.FileSize
		}),
	)
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if gotSize != 0 {
		t.Fatalf("expected size 0, got %d", gotSize)
	}
}

func TestFileReaderReadCloser(t *testing.T) {
	var gotSize int64
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	reader := io.NopCloser(bytes.NewReader([]byte("abc")))
	c := New()
	res := Post[any, string](c, srv.URL, nil,
		FileReader("file", "payload.bin", reader).UploadCallback(func(info req.UploadInfo) {
			gotSize = info.FileSize
		}),
	)
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if gotSize != 0 {
		t.Fatalf("expected size 0, got %d", gotSize)
	}
}

func TestFileReaderGetFileContentBranches(t *testing.T) {
	r := req.C().R()
	FileReader("file", "payload.bin", bytes.NewReader([]byte("abc"))).applyRequest(r)

	reqVal := reflect.ValueOf(r).Elem()
	uploadsField := reqVal.FieldByName("uploadFiles")
	uploadsField = reflect.NewAt(uploadsField.Type(), unsafe.Pointer(uploadsField.UnsafeAddr())).Elem()
	if uploadsField.Len() == 0 {
		t.Fatalf("expected upload file")
	}
	upload := uploadsField.Index(0).Elem()
	getFile := upload.FieldByName("GetFileContent")
	getFile = reflect.NewAt(getFile.Type(), unsafe.Pointer(getFile.UnsafeAddr())).Elem()
	fn := getFile.Interface().(req.GetContentFunc)
	rc, err := fn()
	if err != nil || rc == nil {
		t.Fatalf("expected reader")
	}
	_ = rc.Close()

	r2 := req.C().R()
	FileReader("file", "payload.bin", io.NopCloser(bytes.NewReader([]byte("abc")))).applyRequest(r2)
	reqVal2 := reflect.ValueOf(r2).Elem()
	uploadsField2 := reqVal2.FieldByName("uploadFiles")
	uploadsField2 = reflect.NewAt(uploadsField2.Type(), unsafe.Pointer(uploadsField2.UnsafeAddr())).Elem()
	upload2 := uploadsField2.Index(0).Elem()
	getFile2 := upload2.FieldByName("GetFileContent")
	getFile2 = reflect.NewAt(getFile2.Type(), unsafe.Pointer(getFile2.UnsafeAddr())).Elem()
	fn2 := getFile2.Interface().(req.GetContentFunc)
	rc2, err := fn2()
	if err != nil || rc2 == nil {
		t.Fatalf("expected readcloser")
	}
	_ = rc2.Close()
}

func TestUploadCallbackFinal(t *testing.T) {
	var last req.UploadInfo
	var calls int32
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil,
		FileBytes("file", "report.bin", []byte("payload")).UploadCallback(func(info req.UploadInfo) {
			atomic.AddInt32(&calls, 1)
			last = info
		}),
	)
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if atomic.LoadInt32(&calls) == 0 {
		t.Fatalf("callback not invoked")
	}
	if last.FileSize == 0 || last.UploadedSize != last.FileSize {
		t.Fatalf("final callback size = %d/%d", last.UploadedSize, last.FileSize)
	}
}

func TestUploadCallbackForcesFinal(t *testing.T) {
	calledFinal := false
	request := req.C().R()
	UploadCallback(func(info req.UploadInfo) {
		if info.FileSize == 10 && info.UploadedSize == 10 {
			calledFinal = true
		}
	}).applyRequest(request)

	reqVal := reflect.ValueOf(request).Elem()
	cbField := reqVal.FieldByName("uploadCallback")
	cbField = reflect.NewAt(cbField.Type(), unsafe.Pointer(cbField.UnsafeAddr())).Elem()
	cb := cbField.Interface().(req.UploadCallback)
	cb(req.UploadInfo{FileSize: 10, UploadedSize: 5})

	afterField := reqVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	_ = mw(nil, nil)

	if !calledFinal {
		t.Fatalf("expected forced final callback")
	}
}

func TestUploadCallbackForcesFinalZeroSize(t *testing.T) {
	calledFinal := false
	request := req.C().R()
	UploadCallback(func(info req.UploadInfo) {
		if info.FileSize == 5 && info.UploadedSize == 5 {
			calledFinal = true
		}
	}).applyRequest(request)

	reqVal := reflect.ValueOf(request).Elem()
	cbField := reqVal.FieldByName("uploadCallback")
	cbField = reflect.NewAt(cbField.Type(), unsafe.Pointer(cbField.UnsafeAddr())).Elem()
	cb := cbField.Interface().(req.UploadCallback)
	cb(req.UploadInfo{UploadedSize: 5})

	afterField := reqVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	_ = mw(nil, nil)

	if !calledFinal {
		t.Fatalf("expected forced final callback")
	}
}

func TestUploadCallbackNoFile(t *testing.T) {
	called := false
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, "ok", UploadCallback(func(info req.UploadInfo) {
		called = true
	}))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if called {
		t.Fatalf("expected callback not to run")
	}
}

func TestUploadCallbackNil(t *testing.T) {
	r := req.C().R()
	UploadCallback(nil).applyRequest(r)
}

func TestUploadCallbackWithIntervalFinal(t *testing.T) {
	var last req.UploadInfo
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, nil,
		FileBytes("file", "report.bin", []byte("payload")).UploadCallbackWithInterval(func(info req.UploadInfo) {
			last = info
		}, 5*time.Millisecond),
	)
	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if last.FileSize == 0 || last.UploadedSize != last.FileSize {
		t.Fatalf("final callback size = %d/%d", last.UploadedSize, last.FileSize)
	}
}

func TestUploadCallbackWithIntervalForcesFinal(t *testing.T) {
	calledFinal := false
	request := req.C().R()
	UploadCallbackWithInterval(func(info req.UploadInfo) {
		if info.FileSize == 10 && info.UploadedSize == 10 {
			calledFinal = true
		}
	}, 10*time.Millisecond).applyRequest(request)

	reqVal := reflect.ValueOf(request).Elem()
	cbField := reqVal.FieldByName("uploadCallback")
	cbField = reflect.NewAt(cbField.Type(), unsafe.Pointer(cbField.UnsafeAddr())).Elem()
	cb := cbField.Interface().(req.UploadCallback)
	cb(req.UploadInfo{FileSize: 10, UploadedSize: 5})

	afterField := reqVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	_ = mw(nil, nil)

	if !calledFinal {
		t.Fatalf("expected forced final callback")
	}
}

func TestUploadCallbackWithIntervalNoFile(t *testing.T) {
	called := false
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	c := New()
	res := Post[any, string](c, srv.URL, "ok", UploadCallbackWithInterval(func(info req.UploadInfo) {
		called = true
	}, 10*time.Millisecond))
	if res.Err != nil {
		t.Fatalf("request failed: %v", res.Err)
	}
	if called {
		t.Fatalf("expected callback not to run")
	}
}

func TestUploadCallbackWithIntervalForcesFinalZeroSize(t *testing.T) {
	calledFinal := false
	request := req.C().R()
	UploadCallbackWithInterval(func(info req.UploadInfo) {
		if info.FileSize == 5 && info.UploadedSize == 5 {
			calledFinal = true
		}
	}, 10*time.Millisecond).applyRequest(request)

	reqVal := reflect.ValueOf(request).Elem()
	cbField := reqVal.FieldByName("uploadCallback")
	cbField = reflect.NewAt(cbField.Type(), unsafe.Pointer(cbField.UnsafeAddr())).Elem()
	cb := cbField.Interface().(req.UploadCallback)
	cb(req.UploadInfo{UploadedSize: 5})

	afterField := reqVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	_ = mw(nil, nil)

	if !calledFinal {
		t.Fatalf("expected forced final callback")
	}
}
func TestUploadCallbackWithIntervalNil(t *testing.T) {
	r := req.C().R()
	UploadCallbackWithInterval(nil, 10*time.Millisecond).applyRequest(r)
}

func TestUploadProgress(t *testing.T) {
	srv := uploadServer(t, func(r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
	})
	defer srv.Close()

	buf := &bytes.Buffer{}
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w

	c := New()
	res := Post[any, string](c, srv.URL, nil,
		FileBytes("file", "report.bin", bytes.Repeat([]byte("x"), 1024)).UploadProgress(),
	)
	_ = w.Close()
	os.Stdout = stdout
	_, _ = io.Copy(buf, r)
	_ = r.Close()

	if res.Err != nil {
		t.Fatalf("upload failed: %v", res.Err)
	}
	if !strings.Contains(buf.String(), "upload") {
		t.Fatalf("expected progress output, got %q", buf.String())
	}
}

func TestUploadProgressNoCallback(t *testing.T) {
	request := req.C().R()
	UploadProgress().applyRequest(request)

	reqVal := reflect.ValueOf(request).Elem()
	afterField := reqVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	if err := mw(nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUploadProgressUnknownSize(t *testing.T) {
	buf := &bytes.Buffer{}
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w
	request := req.C().R()
	UploadProgress().applyRequest(request)

	reqVal := reflect.ValueOf(request).Elem()
	cbField := reqVal.FieldByName("uploadCallback")
	cbField = reflect.NewAt(cbField.Type(), unsafe.Pointer(cbField.UnsafeAddr())).Elem()
	cb := cbField.Interface().(req.UploadCallback)
	cb(req.UploadInfo{UploadedSize: 5})

	afterField := reqVal.FieldByName("afterResponse")
	afterField = reflect.NewAt(afterField.Type(), unsafe.Pointer(afterField.UnsafeAddr())).Elem()
	if afterField.Len() == 0 {
		t.Fatalf("expected afterResponse middleware")
	}
	mw := afterField.Index(afterField.Len() - 1).Interface().(req.ResponseMiddleware)
	_ = mw(nil, nil)

	_ = w.Close()
	os.Stdout = stdout
	_, _ = io.Copy(buf, r)
	_ = r.Close()

	if !strings.Contains(buf.String(), "upload") {
		t.Fatalf("expected progress output, got %q", buf.String())
	}
}
func TestFormatBytes(t *testing.T) {
	if got := formatBytes(1); got != "1 B" {
		t.Fatalf("formatBytes(1) = %q", got)
	}
	if got := formatBytes(1024); !strings.Contains(got, "KiB") {
		t.Fatalf("formatBytes(1024) = %q", got)
	}
}

func writeTempFile(t *testing.T, name, contents string) string {
	file, err := os.CreateTemp("", name)
	if err != nil {
		t.Fatalf("temp file: %v", err)
	}
	if _, err := file.WriteString(contents); err != nil {
		_ = file.Close()
		t.Fatalf("write file: %v", err)
	}
	_ = file.Close()
	return file.Name()
}
