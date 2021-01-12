package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

func (i *impl) TaskCreate(ctx context.Context, userID int64, task model.Task) (model.Task, error) {

	if res := i.db.WithContext(ctx).Create(&task); res.Error != nil {
		return task, errors.Wrap(res.Error, errMsgErrorOnTaskRemove)
	}

	return task, nil
}
