package work

import (
	"github.com/hibiken/asynq"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

func NewServeMux(
	meterProvider *sdkmetric.MeterProvider,
	tracerProvider *sdktrace.TracerProvider,
	log *zap.Logger,
) *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.Use(asynqLogMiddleware(log))
	mux.Use(asynqMetricsMiddleware(meterProvider))
	mux.Use(asynqTracingMiddleware(tracerProvider))
	return mux
}
