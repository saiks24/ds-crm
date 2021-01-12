package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

const (
	errMsgErrorOnGetUserTasks = "error on GetUserTasks"

	executorIDCondition = "ExecutorID = ?"
)

func (i *impl) GetUserTasks(ctx context.Context, userID int64) (tasks []model.Task, err error) {
	if res := i.db.WithContext(ctx).Find(&tasks, executorIDCondition, userID); res.Error != nil {
		return tasks, errors.Wrap(err, errMsgErrorOnGetUserTasks)
	}

	return tasks, nil
}
