package telemetry

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type TracingConfig struct {
	EndpointURL string `yaml:"endpointURL" env:"ENDPOINT_URL" envDefault:"http://localhost:4318"`
}

func NewOTLPTraceExporter(ctx context.Context, cfg TracingConfig) (sdktrace.SpanExporter, error) {
	return otlptracehttp.New(ctx,
		otlptracehttp.WithEndpointURL(cfg.EndpointURL))
}

func NewOTELTracerProvider(ctx context.Context, exporter sdktrace.SpanExporter) (*sdktrace.TracerProvider, error) {
	res, err := resource.New(
		ctx,
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithTelemetrySDK(),
	)
	if err != nil {
		return nil, fmt.Errorf("creating resource: %w", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	return provider, nil
}
