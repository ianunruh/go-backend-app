package work

import (
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func NewAsynqScheduler(redisOpt asynq.RedisClientOpt, log *zap.Logger) *asynq.Scheduler {
	return asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{
		Logger: newAsynqLogger(log),
		PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
			if err != nil {
				log.Error("Error enqueuing scheduled task", zap.Error(err))
			}
		},
	})
}

type SchedulerFunc func(sch *asynq.Scheduler) error

func RegisterScheduledTasks(sch *asynq.Scheduler, tasks ...SchedulerFunc) error {
	for _, task := range tasks {
		if err := task(sch); err != nil {
			return err
		}
	}
	return nil
}
