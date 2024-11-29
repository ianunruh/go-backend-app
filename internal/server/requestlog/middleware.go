package requestlog

import (
	"bytes"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Config struct {
	ErrorBody bool `yaml:"errorBody" env:"ERROR_BODY"`
}

func Middleware(next http.Handler, cfg Config, log *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		wrapped := &logResponseWriter{
			Wrapped: w,
		}

		start := time.Now()
		next.ServeHTTP(wrapped, req)
		elapsed := time.Since(start)

		fields := []zap.Field{
			zap.String("method", req.Method),
			zap.Stringer("url", req.URL),
			zap.Int("status", wrapped.StatusCode),
			zap.Int("written", wrapped.Written),
			zap.String("client", req.RemoteAddr),
			zap.String("userAgent", req.UserAgent()),
			zap.Duration("elapsed", elapsed),
		}

		if cfg.ErrorBody && wrapped.StatusCode > 399 {
			fields = append(fields, zap.ByteString("body", wrapped.Body.Bytes()))
		}

		log.Info("Handled HTTP request", fields...)
	})
}

type logResponseWriter struct {
	Wrapped    http.ResponseWriter
	StatusCode int
	Written    int

	Body bytes.Buffer
}

func (w *logResponseWriter) Header() http.Header {
	return w.Wrapped.Header()
}

func (w *logResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	written, err := w.Wrapped.Write(b)
	w.Written += written
	return written, err
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.Wrapped.WriteHeader(statusCode)
}
