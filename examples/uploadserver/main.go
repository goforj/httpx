//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()
		_, _ = io.Copy(io.Discard, r.Body)
		fmt.Fprintln(w, "ok")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
