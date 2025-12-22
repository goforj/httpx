package httpx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

// File attaches a file from disk as multipart form data.
// @group Upload Options
//
// Applies to individual requests only.
// Example: upload a file
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.File("file", "/tmp/report.txt"))
func File(paramName, filePath string) OptionBuilder {
	return OptionBuilder{}.File(paramName, filePath)
}

func (b OptionBuilder) File(paramName, filePath string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetFile(paramName, filePath)
	}))
}

// Files attaches multiple files from disk as multipart form data.
// @group Upload Options
//
// Applies to individual requests only.
// Example: upload multiple files
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.Files(map[string]string{
//		"fileA": "/tmp/a.txt",
//		"fileB": "/tmp/b.txt",
//	}))
func Files(files map[string]string) OptionBuilder {
	return OptionBuilder{}.Files(files)
}

func (b OptionBuilder) Files(files map[string]string) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetFiles(files)
	}))
}

// FileBytes attaches a file from bytes as multipart form data.
// @group Upload Options
//
// Applies to individual requests only.
// Example: upload bytes as a file
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileBytes("file", "report.txt", []byte("hello")))
func FileBytes(paramName, filename string, content []byte) OptionBuilder {
	return OptionBuilder{}.FileBytes(paramName, filename, content)
}

func (b OptionBuilder) FileBytes(paramName, filename string, content []byte) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		r.SetFileUpload(req.FileUpload{
			ParamName: paramName,
			FileName:  filename,
			FileSize:  int64(len(content)),
			GetFileContent: func() (io.ReadCloser, error) {
				return io.NopCloser(bytes.NewReader(content)), nil
			},
		})
	}))
}

// FileReader attaches a file from a reader as multipart form data.
// @group Upload Options
//
// Applies to individual requests only.
// Example: upload from reader
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil, httpx.FileReader("file", "report.txt", strings.NewReader("hello")))
func FileReader(paramName, filename string, reader io.Reader) OptionBuilder {
	return OptionBuilder{}.FileReader(paramName, filename, reader)
}

func (b OptionBuilder) FileReader(paramName, filename string, reader io.Reader) OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		fileSize := int64(0)
		switch v := reader.(type) {
		case interface{ Size() int64 }:
			fileSize = v.Size()
		case interface{ Len() int }:
			fileSize = int64(v.Len())
		case io.Seeker:
			cur, err := v.Seek(0, io.SeekCurrent)
			if err == nil {
				end, err := v.Seek(0, io.SeekEnd)
				if err == nil {
					fileSize = end
					_, _ = v.Seek(cur, io.SeekStart)
				} else {
					_, _ = v.Seek(cur, io.SeekStart)
				}
			}
		}

		r.SetFileUpload(req.FileUpload{
			ParamName: paramName,
			FileName:  filename,
			FileSize:  fileSize,
			GetFileContent: func() (io.ReadCloser, error) {
				if rc, ok := reader.(io.ReadCloser); ok {
					return rc, nil
				}
				return io.NopCloser(reader), nil
			},
		})
	}))
}

// UploadCallback registers a callback for upload progress.
// @group Upload Options
//
// Applies to individual requests only.
// Example: track upload progress
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
//		httpx.File("file", "/tmp/report.bin"),
//		httpx.UploadCallback(func(info req.UploadInfo) {
//			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
//			fmt.Printf("\rprogress: %.1f%%", percent)
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}),
//	)
func UploadCallback(callback req.UploadCallback) OptionBuilder {
	return OptionBuilder{}.UploadCallback(callback)
}

func (b OptionBuilder) UploadCallback(callback req.UploadCallback) OptionBuilder {
	if callback == nil {
		return b
	}
	return b.add(requestOnly(func(r *req.Request) {
		var mu sync.Mutex
		var last req.UploadInfo
		var seen bool
		var completed bool
		r.SetUploadCallback(func(info req.UploadInfo) {
			mu.Lock()
			last = info
			seen = true
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				completed = true
			}
			mu.Unlock()
			callback(info)
		})
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			mu.Lock()
			info := last
			seenLocal := seen
			completedLocal := completed
			mu.Unlock()
			if !seenLocal {
				return nil
			}
			if !completedLocal {
				if info.FileSize == 0 {
					info.FileSize = info.UploadedSize
				}
				if info.FileSize > 0 {
					info.UploadedSize = info.FileSize
				}
				callback(info)
			}
			return nil
		})
	}))
}

// UploadCallbackWithInterval registers a callback for upload progress with a minimum interval.
// @group Upload Options
//
// Applies to individual requests only.
// Example: throttle upload progress updates
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
//		httpx.File("file", "/tmp/report.bin"),
//		httpx.UploadCallbackWithInterval(func(info req.UploadInfo) {
//			percent := float64(info.UploadedSize) / float64(info.FileSize) * 100
//			fmt.Printf("\rprogress: %.1f%% (%.0f bytes)", percent, float64(info.UploadedSize))
//			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
//				fmt.Print("\n")
//			}
//		}, 200*time.Millisecond),
//	)
func UploadCallbackWithInterval(callback req.UploadCallback, minInterval time.Duration) OptionBuilder {
	return OptionBuilder{}.UploadCallbackWithInterval(callback, minInterval)
}

func (b OptionBuilder) UploadCallbackWithInterval(callback req.UploadCallback, minInterval time.Duration) OptionBuilder {
	if callback == nil {
		return b
	}
	return b.add(requestOnly(func(r *req.Request) {
		var mu sync.Mutex
		var last req.UploadInfo
		var seen bool
		var completed bool
		r.SetUploadCallbackWithInterval(func(info req.UploadInfo) {
			mu.Lock()
			last = info
			seen = true
			if info.FileSize > 0 && info.UploadedSize >= info.FileSize {
				completed = true
			}
			mu.Unlock()
			callback(info)
		}, minInterval)
		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			mu.Lock()
			info := last
			seenLocal := seen
			completedLocal := completed
			mu.Unlock()
			if !seenLocal {
				return nil
			}
			if !completedLocal {
				if info.FileSize == 0 {
					info.FileSize = info.UploadedSize
				}
				if info.FileSize > 0 {
					info.UploadedSize = info.FileSize
				}
				callback(info)
			}
			return nil
		})
	}))
}

// UploadProgress enables a default progress spinner and bar for uploads.
// @group Upload Options
// Applies to individual requests only.
//
// Example: upload with automatic progress
//
//	c := httpx.New()
//	_ = httpx.Post[any, string](c, "https://example.com/upload", nil,
//		httpx.File("file", "/tmp/report.bin"),
//		httpx.UploadProgress(),
//	)
func UploadProgress() OptionBuilder {
	return OptionBuilder{}.UploadProgress()
}

func (b OptionBuilder) UploadProgress() OptionBuilder {
	return b.add(requestOnly(func(r *req.Request) {
		var mu sync.Mutex
		var last string
		var total int64
		spin := []string{"|", "/", "-", "\\"}
		spinIndex := 0
		barWidth := 20

		r.SetUploadCallback(func(info req.UploadInfo) {
			mu.Lock()
			defer mu.Unlock()

			if info.FileSize > 0 {
				total = info.FileSize
			}
			spinIndex = (spinIndex + 1) % len(spin)

			if total > 0 {
				percent := float64(info.UploadedSize) / float64(total) * 100
				filled := int(percent / 100 * float64(barWidth))
				bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth-filled)
				last = fmt.Sprintf(
					"\r%s upload [%s] %.1f%% (%s/%s)",
					spin[spinIndex],
					bar,
					percent,
					formatBytes(info.UploadedSize),
					formatBytes(total),
				)
			} else {
				last = fmt.Sprintf("\r%s upload %s", spin[spinIndex], formatBytes(info.UploadedSize))
			}
			fmt.Print(last)
		})

		r.OnAfterResponse(func(_ *req.Client, _ *req.Response) error {
			mu.Lock()
			defer mu.Unlock()
			if last == "" {
				return nil
			}
			if total > 0 {
				fmt.Printf(
					"\r%s upload [%s] 100.0%% (%s/%s)\n",
					spin[spinIndex],
					strings.Repeat("=", barWidth),
					formatBytes(total),
					formatBytes(total),
				)
			} else {
				fmt.Print("\n")
			}
			return nil
		})
	}))
}

func formatBytes(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	units := []string{"KiB", "MiB", "GiB", "TiB", "PiB"}
	value := float64(size)
	unit := "B"
	for _, u := range units {
		value /= 1024
		unit = u
		if value < 1024 {
			break
		}
	}
	return fmt.Sprintf("%.1f %s", value, unit)
}
