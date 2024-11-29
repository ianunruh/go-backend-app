package work

import (
	"context"

	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const TracerName = "github.com/ianunruh/go-backend-app/internal/work"

func asynqTracingMiddleware(tracerProvider *sdktrace.TracerProvider) asynq.MiddlewareFunc {
	tracer := tracerProvider.Tracer(TracerName)

	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			taskID, _ := asynq.GetTaskID(ctx)

			ctx, span := tracer.Start(ctx, "asynq.ProcessTask",
				trace.WithSpanKind(trace.SpanKindConsumer),
				trace.WithAttributes(
					attribute.String("task.id", taskID),
					attribute.String("task.type", task.Type()),
					attribute.String("task.payload", string(task.Payload())),
				))
			defer span.End()

			if err := next.ProcessTask(ctx, task); err != nil {
				span.SetStatus(codes.Error, err.Error())
				return err
			}

			return nil
		})
	}
}
