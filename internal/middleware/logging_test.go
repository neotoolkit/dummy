package middleware_test

import (
	"github.com/neotoolkit/dummy/internal/logger"
	"github.com/neotoolkit/dummy/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogging(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	l := logger.NewLogger("DEBUG")

	handler := middleware.Logging(next, l)

	req := httptest.NewRequest("GET", "https://", nil)

	handler.ServeHTTP(httptest.NewRecorder(), req)
}
