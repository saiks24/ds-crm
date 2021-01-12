package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

const (
	errMsgErrorOnChangeTaskStatus = "error on ChangeTaskStatus"
)

func (i *impl) ChangeTaskStatus(ctx context.Context, taskID, userID int64, newStatus string) (model.Task, error) {

	if task, err := i.GetTaskByID(ctx, taskID); err != nil {
		return model.Task{}, errors.Wrap(err, errMsgErrorOnChangeTaskStatus)
	} else {
		// TODO не строка а stateID
		if res := i.db.Model(&task).WithContext(ctx).Update("CurrentStateID", newStatus); res.Error != nil {
			return model.Task{}, errors.Wrap(err, errMsgErrorOnChangeTaskStatus)
		} else {
			return task, nil
		}
	}
}
