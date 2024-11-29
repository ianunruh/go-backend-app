package work

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

const MeterName = "github.com/ianunruh/go-backend-app/internal/work"

func asynqMetricsMiddleware(meterProvider *sdkmetric.MeterProvider) asynq.MiddlewareFunc {
	meter := meterProvider.Meter(MeterName)

	taskDuration, _ := meter.Float64Histogram(
		"worker.task.duration",
		metric.WithDescription("Duration of task processing"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5, 10),
	)

	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			start := time.Now()

			err := next.ProcessTask(ctx, task)

			elapsed := time.Since(start).Seconds()

			taskDuration.Record(ctx, elapsed, metric.WithAttributes(
				attribute.String("task.type", task.Type()),
				attribute.String("task.status", taskStatusAttribute(err)),
			))

			return err
		})
	}
}

func taskStatusAttribute(err error) string {
	if err != nil {
		return "failure"
	}
	return "success"
}
