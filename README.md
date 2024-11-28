# go-backend-app

Template for a Go backend app with OpenAPI and OpenTelemetry metrics/tracing.

## Development

### Docker Compose

App dependencies can be run locally using Docker Compose.

### OpenAPI codegen

The `api` directory contains the OpenAPI spec and tools for generating the API client and server. Once changes have been made to the OpenAPI spec, you can generate the complete API spec and the API server code by running `make generate`.

## Observability

### Debug/profiling

This template includes a separate debug HTTP server that exposes [pprof](https://pkg.go.dev/net/http/pprof) for profiling app CPU usage, goroutines, heap, mutexes, and more.

### Metrics

This template uses OpenTelemetry metrics for collecting metrics. It comes with a built-in Prometheus metrics server that exposes the metrics for scraping.

### Logging

This template uses Zap for logging. In development mode, the logs will be written to the console in a human-friendly format. In production, logs are written to the console in JSON format for easy parsing.

The log level can be changed at runtime by calling the `/debug/log/level` endpoint on the debug server using Zap's [AtomicLevel](https://pkg.go.dev/go.uber.org/zap#AtomicLevel.ServeHTTP) API.

### Tracing

This template uses OpenTelemetry tracing for collecting traces. It comes with a built-in exporter that can be configured to send traces to any OTLP-compatible collector.

For demonstration purposes, it comes with a Jaeger all-in-one server that can be run locally. Traces can be viewed in the Jaeger UI at http://localhost:16686.
