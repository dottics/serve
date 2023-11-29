package servemux

import (
	"context"
	"net/http"
	"strings"
)

func CORS(headers http.Header, payload []byte) ContextMiddleware {
	if headers == nil {
		headers = make(http.Header)
		headers.Set("access-control-allow-origin", "*")
		headers.Set("access-control-allow-methods", "*")
		headers.Set("access-control-allow-headers", "*")
		headers.Set("content-type", "application/json")
	}
	if payload == nil {
		payload = []byte(`{}`)
	}
	return func(c *context.Context, w http.ResponseWriter, r *http.Request) (bool, http.HandlerFunc) {
		if r.Method == "OPTIONS" {
			for k, v := range headers {
				w.Header().Set(k, strings.Join(v, "; "))
			}
			return true, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				_, _ = w.Write(payload)
			}
		}
		return false, nil
	}
}
