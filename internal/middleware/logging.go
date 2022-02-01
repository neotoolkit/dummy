package middleware

import (
	"net/http"
	"time"

	"github.com/neotoolkit/dummy/internal/logger"
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

// Logging -.
func Logging(next http.Handler, logger *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info().
			Str("path", r.URL.Path).
			Str("method", r.Method).
			Interface("header", r.Header).
			Interface("body", r.Body).
			Msg("request")

		wrapped := wrapResponseWriter(w)

		start := time.Now()

		next.ServeHTTP(wrapped, r)

		logger.Info().
			Int("status-code", wrapped.Status()).
			Interface("header", wrapped.Header()).
			Interface("duration", time.Since(start)).
			Msg("response")
	})
}
