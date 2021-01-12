package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

const (
	errMsgErrorOnGetTaskByID = "error on GetTaskByID"

	idCondition = "id = ?"
)

func (i *impl) GetTaskByID(ctx context.Context, taskID int64) (task model.Task, err error) {
	if res := i.db.WithContext(ctx).First(&task, idCondition, taskID); res.Error != nil {
		return task, errors.Wrap(err, errMsgErrorOnGetTaskByID)
	}

	return task, nil
}
