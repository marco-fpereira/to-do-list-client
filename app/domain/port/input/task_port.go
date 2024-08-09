package input

import (
	"context"

	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
)

type TaskPort interface {
	GetAllTasks(
		ctx context.Context,
		userId string,
		requestId string,
	) (*[]model.TaskDomain, error)

	CreateTask(
		ctx context.Context,
		userId string,
		taskMessage string,
		requestId string,
	) (*model.TaskDomain, error)

	UpdateTaskMessage(
		ctx context.Context,
		userId string,
		taskId string,
		taskMessage string,
		requestId string,
	) (*model.TaskDomain, error)

	UpdateTaskCompleteness(
		ctx context.Context,
		userId string,
		taskId string,
		requestId string,
	) error

	DeleteTask(
		ctx context.Context,
		userId string,
		taskId string,
		requestId string,
	) error
}
