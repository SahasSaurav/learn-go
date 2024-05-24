package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		ctx := context.WithValue(r.Context(), "hello", "world")
		req := r.WithContext(ctx)

		next.ServeHTTP(wrapped, req)

		slog.Info("",
			slog.Int("statusCode", wrapped.statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("duration", time.Since(start).String()),
		)
	})
}
