package gzip

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGzipHandler(t *testing.T) {
	w := httptest.NewRecorder()
	handler := GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello\n")
	}))

	req := &http.Request{
		Method: "GET",
		Header: http.Header{
			"Accept-Encoding": []string{"gzip"},
		},
	}

	handler.ServeHTTP(w, req)

	if w.HeaderMap.Get("Content-Encoding") != "gzip" {
		t.Errorf("wrong content encoding %s", w.HeaderMap.Get("Content-Encoding"))
	}
}
