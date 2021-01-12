package db

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

const errMsgErrorOnTaskRemove = "error on TaskRemove"

func (i *impl) TaskRemove(ctx context.Context, userID, taskID int64) error {
	task := model.Task{
		Model: gorm.Model{ID: uint(taskID)},
	}
	if res := i.db.WithContext(ctx).Delete(&task, taskID); res.Error != nil {
		return errors.Wrap(res.Error, errMsgErrorOnTaskRemove)
	}

	return nil
}
