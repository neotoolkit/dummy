package middleware

import (
	"net/http"
	"time"

	"github.com/go-dummy/dummy/internal/logger"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler, logger *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		logger.Info().
			Str("path", req.URL.Path).
			Str("method", req.Method).
			Interface("header", req.Header).
			Interface("body", req.Body).
			Msg("request")

		wrapped := wrapResponseWriter(w)

		next.ServeHTTP(wrapped, req)

		logger.Info().
			Int("status-code", wrapped.Status()).
			Interface("header", wrapped.Header()).
			Interface("duration", time.Since(start)).
			Msg("response")
	})
}
