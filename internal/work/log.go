package work

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func newAsynqLogger(log *zap.Logger) asynq.Logger {
	return asynqLogger{
		log: log.WithOptions(zap.AddCallerSkip(3)).Sugar(),
	}
}

type asynqLogger struct {
	log *zap.SugaredLogger
}

func (l asynqLogger) Debug(args ...any) {
	l.log.Debug(args...)
}

func (l asynqLogger) Info(args ...any) {
	l.log.Info(args...)
}

func (l asynqLogger) Warn(args ...any) {
	l.log.Warn(args...)
}

func (l asynqLogger) Error(args ...any) {
	l.log.Error(args...)
}

func (l asynqLogger) Fatal(args ...any) {
	l.log.Fatal(args...)
}

func asynqLogMiddleware(log *zap.Logger) asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			taskID, _ := asynq.GetTaskID(ctx)
			taskType := task.Type()

			log.Debug("Started task",
				zap.String("id", taskID),
				zap.String("type", taskType))

			start := time.Now()
			err := next.ProcessTask(ctx, task)
			elapsed := time.Since(start)

			if err != nil {
				log.Error("Error during task",
					zap.String("id", taskID),
					zap.String("type", taskType),
					zap.Duration("elapsed", elapsed),
					zap.Error(err))
				return err
			}

			log.Info("Finished task",
				zap.String("id", taskID),
				zap.String("type", taskType),
				zap.Duration("elapsed", elapsed))

			return nil
		})
	}
}
