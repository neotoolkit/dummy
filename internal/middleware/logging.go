package middleware

import (
	"net/http"
	"time"

	"github.com/neotoolkit/dummy/internal/pkg/logger"
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
func Logging(next http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infow("request",
			"path", r.URL.Path,
			"method", r.Method,
			"header", r.Header,
			"body", r.Body,
		)

		wrapped := wrapResponseWriter(w)

		start := time.Now()

		next.ServeHTTP(wrapped, r)

		logger.Infow("response",
			"status-code", wrapped.Status(),
			"header", wrapped.Header(),
			"duration", time.Since(start),
		)
	})
}
