package work

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var (
	CriticalQueue = asynq.Queue("critical")
	HighQueue     = asynq.Queue("high")
	LowQueue      = asynq.Queue("low")
)

type Queue interface {
	Enqueue(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
	Close() error
}

func NewQueue(r asynq.RedisConnOpt, meterProvider *sdkmetric.MeterProvider, tracerProvider *sdktrace.TracerProvider, log *zap.Logger) Queue {
	meter := meterProvider.Meter(MeterName)

	enqueueCount, _ := meter.Int64Counter(
		"client.task.enqueue.count",
		metric.WithDescription("Number of tasks enqueued"),
	)

	enqueueDuration, _ := meter.Float64Histogram(
		"client.task.enqueue.duration",
		metric.WithDescription("Duration of task enqueuing"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5, 10),
	)

	return &asynqQueue{
		client: asynq.NewClient(r),
		log:    log,
		tracer: tracerProvider.Tracer(TracerName),

		enqueueCount:    enqueueCount,
		enqueueDuration: enqueueDuration,
	}
}

type asynqQueue struct {
	client *asynq.Client
	log    *zap.Logger
	tracer trace.Tracer

	enqueueCount    metric.Int64Counter
	enqueueDuration metric.Float64Histogram
}

func (q *asynqQueue) Enqueue(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	ctx, span := q.tracer.Start(ctx, "asynq.Enqueue",
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("task.type", task.Type()),
			attribute.String("task.payload", string(task.Payload())),
		))
	defer span.End()

	start := time.Now()
	info, err := q.client.EnqueueContext(ctx, task, opts...)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	span.SetAttributes(
		attribute.String("task.id", info.ID),
		attribute.String("task.queue", info.Queue),
	)

	elapsed := time.Since(start)

	q.enqueueCount.Add(ctx, 1, metric.WithAttributes(
		attribute.String("task.type", info.Type),
		attribute.String("task.queue", info.Queue),
	))
	q.enqueueDuration.Record(ctx, elapsed.Seconds())

	payload := json.RawMessage(task.Payload())

	q.log.Info("Enqueued task",
		zap.String("id", info.ID),
		zap.String("queue", info.Queue),
		zap.String("type", task.Type()),
		zap.Any("payload", payload),
		zap.Duration("elapsed", elapsed))

	return info, nil
}

func (q *asynqQueue) Close() error {
	return q.client.Close()
}
