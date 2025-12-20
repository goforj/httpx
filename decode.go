package httpx

import (
	"bytes"
	"reflect"

	"github.com/imroc/req/v3"
)

type rawKind int

const (
	rawNone rawKind = iota
	rawString
	rawBytes
)

func rawKindOf[T any]() rawKind {
	t := reflect.TypeOf((*T)(nil)).Elem()
	switch t.Kind() {
	case reflect.String:
		return rawString
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return rawBytes
		}
	}
	return rawNone
}

func decodeRaw[T any](resp *req.Response) T {
	var out T
	switch rawKindOf[T]() {
	case rawString:
		out = any(resp.String()).(T)
	case rawBytes:
		out = any(resp.Bytes()).(T)
	}
	return out
}

func ensureNonNil[T any](out *T) {
	if out == nil {
		return
	}
	t := reflect.TypeOf((*T)(nil)).Elem()
	v := reflect.ValueOf(out).Elem()
	switch t.Kind() {
	case reflect.Slice:
		if v.IsNil() {
			v.Set(reflect.MakeSlice(t, 0, 0))
		}
	case reflect.Map:
		if v.IsNil() {
			v.Set(reflect.MakeMapWithSize(t, 0))
		}
	}
}

func isEmptyBody(resp *req.Response) bool {
	if resp == nil {
		return false
	}
	return len(bytes.TrimSpace(resp.Bytes())) == 0
}
