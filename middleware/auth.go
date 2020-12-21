package middleware

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

const (
	apiKeyHeader = "X-Api-Key"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(apiKeyHeader)

		if authHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			_, _ = fmt.Fprintln(w, "No X-Api-Key provided")
			return
		}

		apiKeys := strings.Split(os.Getenv("API_KEYS"), ",")
		sort.Strings(apiKeys)
		if sort.SearchStrings(apiKeys, authHeader) > len(apiKeys) {
			w.WriteHeader(http.StatusForbidden)
			_, _ = fmt.Fprintln(w, "Forbidden")
			return
		}

		next.ServeHTTP(w, r)
	})
}
