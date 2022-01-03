package middleware

import (
	"net/http"

	"github.com/go-dummy/dummy/internal/logger"
)

func Logging(next http.Handler, logger *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger.Info().
			Str("path", req.URL.Path).
			Str("method", req.Method).
			Interface("header", req.Header).
			Interface("body", req.Body).
			Msg("request")

		next.ServeHTTP(w, req)
	})
}
