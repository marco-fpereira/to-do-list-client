package usecase

import (
	"context"

	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
	inputPort "github.com/marco-fpereira/to-do-list-client/app/domain/port/input"
	outputPort "github.com/marco-fpereira/to-do-list-client/app/domain/port/output"
)

type TaskUseCase struct {
	output        outputPort.ToDoServerPort
	authenticator outputPort.AuthenticationPort
}

func NewTaskUseCase(
	output outputPort.ToDoServerPort,
	authenticator outputPort.AuthenticationPort,
) inputPort.TaskPort {
	return &TaskUseCase{
		output:        output,
		authenticator: authenticator,
	}
}

func (t *TaskUseCase) GetAllTasks(
	ctx context.Context,
	userId string,
	requestId string,
) (*[]model.TaskDomain, error) {
	token, err := t.authenticator.GenerateTokenWithClaim(
		outputPort.Claim{
			ClaimName:  "userId",
			ClaimValue: userId,
		},
	)
	if err != nil {
		return nil, err
	}
	taskDomainList, err := t.output.GetAllTasks(ctx, userId, requestId, *token)
	if err != nil {
		return nil, err
	}

	return taskDomainList, nil
}

func (t *TaskUseCase) CreateTask(
	ctx context.Context,
	userId string,
	taskMessage string,
	requestId string,
) (*model.TaskDomain, error) {
	token, err := t.authenticator.GenerateTokenWithClaim(
		outputPort.Claim{
			ClaimName:  "userId",
			ClaimValue: userId,
		},
	)
	if err != nil {
		return nil, err
	}
	taskDomain, err := t.output.CreateTask(ctx, userId, taskMessage, requestId, *token)
	if err != nil {
		return nil, err
	}

	return taskDomain, nil
}

func (t *TaskUseCase) UpdateTaskMessage(
	ctx context.Context,
	userId string,
	taskId string,
	taskMessage string,
	requestId string,
) (*model.TaskDomain, error) {
	token, err := t.authenticator.GenerateTokenWithClaim(
		outputPort.Claim{
			ClaimName:  "userId",
			ClaimValue: userId,
		},
	)
	if err != nil {
		return nil, err
	}
	taskDomain, err := t.output.UpdateTaskMessage(ctx, userId, taskId, taskMessage, requestId, *token)
	if err != nil {
		return nil, err
	}

	return taskDomain, nil
}

func (t *TaskUseCase) UpdateTaskCompleteness(
	ctx context.Context,
	userId string,
	taskId string,
	requestId string,
) error {
	token, err := t.authenticator.GenerateTokenWithClaim(
		outputPort.Claim{
			ClaimName:  "userId",
			ClaimValue: userId,
		},
	)
	if err != nil {
		return err
	}
	err = t.output.UpdateTaskCompleteness(ctx, userId, taskId, requestId, *token)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskUseCase) DeleteTask(
	ctx context.Context,
	userId string,
	taskId string,
	requestId string,
) error {
	token, err := t.authenticator.GenerateTokenWithClaim(
		outputPort.Claim{
			ClaimName:  "userId",
			ClaimValue: userId,
		},
	)
	if err != nil {
		return err
	}
	err = t.output.DeleteTask(ctx, userId, taskId, requestId, *token)
	if err != nil {
		return err
	}

	return nil
}
