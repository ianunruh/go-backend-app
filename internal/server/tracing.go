package server

import (
	"net/http"

	"go.opentelemetry.io/otel/propagation"
)

func tracePropagationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		propagator := propagation.TraceContext{}
		ctx := propagator.Extract(req.Context(), propagation.HeaderCarrier(req.Header))
		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
