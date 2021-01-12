package app

import (
	"context"

	"github.com/saiks24/ds-crm/internal/pkg/model"
)

type ServiceProvider interface {
	GetUserTasks(ctx context.Context, userID int64) ([]model.Task, error)
	GetTasksOnStatus(ctx context.Context, status string) ([]model.Task, error)
	GetTaskByID(ctx context.Context, taskID int64) (model.Task, error)

	TaskCreate(ctx context.Context, userID int64, task model.Task) (model.Task, error)
	TaskRemove(ctx context.Context, userID, taskID int64) error

	DoCommentTask(ctx context.Context, taskID, userID int64, comment string) (model.Task, error)
	ChangeTaskStatus(ctx context.Context, taskID, userID int64, newStatus string) (model.Task, error)

	CreateFlow(ctx context.Context, userID int64, template model.FlowPartTemplate) (model.FlowPartTemplate, error)
	UpdateFlow(ctx context.Context, userID int64, template model.FlowPartTemplate) (model.FlowPartTemplate, error)
	RemoveFlow(ctx context.Context, userID, flowID int64) error
}
