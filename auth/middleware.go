package auth

import (
	"context"
	"fmt"
	"net/http"
)

func Middleware(c context.Context, w http.ResponseWriter, r *http.Request) (bool, http.HandlerFunc) {
	fmt.Println(c.Value("whitelisted"))
	whitelisted := c.Value("whitelisted").(bool)
	if whitelisted {
		return false, nil
	}
	err := ExtractAuthToken(r)
	if err != nil {
		return true, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("content-type", "application/json")
			_, _ = w.Write([]byte(`{"detail":"invalid auth token"}`))
		}
	}
	return false, nil
}
