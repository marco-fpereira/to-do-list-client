package output

import (
	"context"

	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
)

type ToDoServerPort interface {
	Signup(
		ctx context.Context,
		userCredentials *model.UserCredentialsDomain,
		requestId string,
		token string,
	) error

	Login(
		ctx context.Context,
		userCredentials *model.UserCredentialsDomain,
		requestId string,
		token string,
	) (*string, error)

	GetAllTasks(
		ctx context.Context,
		userId string,
		requestId string,
		token string,
	) (*[]model.TaskDomain, error)

	CreateTask(
		ctx context.Context,
		userId string,
		taskMessage string,
		requestId string,
		token string,
	) (*model.TaskDomain, error)

	UpdateTaskMessage(
		ctx context.Context,
		userId string,
		taskId string,
		taskMessage string,
		requestId string,
		token string,
	) (*model.TaskDomain, error)

	UpdateTaskCompleteness(
		ctx context.Context,
		userId string,
		taskId string,
		requestId string,
		token string,
	) error

	DeleteTask(
		ctx context.Context,
		userId string,
		taskId string,
		requestId string,
		token string,
	) error
}
