package work

import (
	"encoding/json"
	"errors"

	"github.com/hibiken/asynq"
)

func IsDuplicateTaskErr(err error) bool {
	return errors.Is(err, asynq.ErrTaskIDConflict)
}

func NewTask(typename string, payload any, opts ...asynq.Option) (*asynq.Task, error) {
	encoded, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(typename, encoded, opts...)
	return task, nil
}

func DecodeTask(task *asynq.Task, out any) error {
	encoded := task.Payload()
	return json.Unmarshal(encoded, out)
}
