package logging

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// statusResponseWriter is a wrapper around http.ResponseWriter that keeps track of the status code
type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader will track the status code for future logging
func (w *statusResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func InterceptorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := GetZap(r.Context())
		start := time.Now()

		// Wrap existing response writer with our custom one
		sw := &statusResponseWriter{ResponseWriter: w}
		next.ServeHTTP(sw, r)
		lFn := log.Debug
		if sw.statusCode >= http.StatusBadRequest {
			lFn = log.Warn
		}

		lFn("Request handled",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status", sw.statusCode),
			zap.Duration("duration", time.Since(start)))
	})
}
