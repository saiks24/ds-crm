package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

const (
	statusCondition = "CurrentStateID = ?"

	errMsgErrorOnGetTasksOnStatus = "error on GetTasksOnStatus"
)

func (i *impl) GetTasksOnStatus(ctx context.Context, status string) (tasks []model.Task, err error) {
	if res := i.db.WithContext(ctx).Find(&tasks, statusCondition, status); res.Error != nil {
		return tasks, errors.Wrap(err, errMsgErrorOnGetTasksOnStatus)
	}

	return tasks, nil
}
