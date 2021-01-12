package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

const (
	errMsgErrorOnDoCommentTask = "error on DoCommentTask"
)

func (i *impl) DoCommentTask(ctx context.Context, taskID, userID int64, comment string) (task model.Task, err error) {
	if task, err = i.GetTaskByID(ctx, taskID); err != nil {
		return model.Task{}, errors.Wrap(err, errMsgErrorOnDoCommentTask)
	}

	cmt := model.Comment{
		UserID: userID,
		TaskID: taskID,
		Text:   comment,
	}

	if err := i.db.WithContext(ctx).Create(&cmt); err.Error != nil {
		return model.Task{}, errors.Wrap(err.Error, errMsgErrorOnDoCommentTask)
	}

	return task, nil
}
