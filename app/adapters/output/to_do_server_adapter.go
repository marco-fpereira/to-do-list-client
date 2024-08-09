package output

import (
	"context"

	"github.com/marco-fpereira/to-do-list-client/app/adapters/utils"
	"github.com/marco-fpereira/to-do-list-client/app/config/grpc"
	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
	outputPort "github.com/marco-fpereira/to-do-list-client/app/domain/port/output"
)

type ToDoServerAdapter struct {
	accountClient grpc.AccountClient
	taskClient    grpc.TaskClient
}

func NewToDoServerAdapter(
	accountClient grpc.AccountClient,
	taskClient grpc.TaskClient,
) outputPort.ToDoServerPort {
	return &ToDoServerAdapter{
		accountClient: accountClient,
		taskClient:    taskClient,
	}
}

func (t *ToDoServerAdapter) Signup(
	ctx context.Context,
	userCredentials *model.UserCredentialsDomain,
	requestId string,
	token string,
) error {
	userCredentialsRequest := grpc.UserCredentialsRequest{
		Username:  userCredentials.Username,
		Password:  userCredentials.Password,
		RequestId: requestId,
		Token:     token,
	}
	_, err := t.accountClient.Signup(ctx, &userCredentialsRequest)
	if err != nil {
		return err
	}
	return nil
}

func (t *ToDoServerAdapter) Login(
	ctx context.Context,
	userCredentials *model.UserCredentialsDomain,
	requestId string,
	token string,
) (*string, error) {
	userCredentialsRequest := grpc.UserCredentialsRequest{
		Username:  userCredentials.Username,
		Password:  userCredentials.Password,
		RequestId: requestId,
		Token:     token,
	}
	userId, err := t.accountClient.Login(ctx, &userCredentialsRequest)
	if err != nil {
		return nil, err
	}
	return &userId.UserId, nil
}

func (t *ToDoServerAdapter) CreateTask(
	ctx context.Context,
	userId string,
	taskMessage string,
	requestId string,
	token string,
) (*model.TaskDomain, error) {
	request := grpc.CreateTaskRequest{
		UserId:      userId,
		TaskMessage: taskMessage,
		RequestId:   requestId,
		Token:       token,
	}
	taskResponse, err := t.taskClient.CreateTask(ctx, &request)
	if err != nil {
		return nil, err
	}
	taskDomain := utils.TaskResponseToTaskDomain(taskResponse)
	return &taskDomain, nil
}

func (t *ToDoServerAdapter) GetAllTasks(
	ctx context.Context,
	userId string,
	requestId string,
	token string,
) (*[]model.TaskDomain, error) {
	request := grpc.GetAllTasksRequest{
		UserId:    userId,
		RequestId: requestId,
		Token:     token,
	}
	tasksResponseList, err := t.taskClient.GetAllTasks(ctx, &request)
	if err != nil {
		return nil, err
	}
	var taskDomainList []model.TaskDomain
	for _, taskResponse := range tasksResponseList.TaskDomain {
		taskDomainList = append(
			taskDomainList,
			utils.TaskResponseToTaskDomain(taskResponse),
		)
	}
	return &taskDomainList, err
}

func (t *ToDoServerAdapter) UpdateTaskCompleteness(
	ctx context.Context,
	userId string,
	taskId string,
	requestId string,
	token string,
) error {
	request := grpc.UpdateTaskCompletenessRequest{
		TaskId:    taskId,
		RequestId: requestId,
		Token:     token,
	}
	_, err := t.taskClient.UpdateTaskCompleteness(ctx, &request)
	return err
}

func (t *ToDoServerAdapter) UpdateTaskMessage(
	ctx context.Context,
	userId string,
	taskId string,
	taskMessage string,
	requestId string,
	token string,
) (*model.TaskDomain, error) {
	request := grpc.UpdateTaskMessageRequest{
		TaskId:      taskId,
		TaskMessage: taskMessage,
		RequestId:   requestId,
		Token:       token,
	}
	taskResponse, err := t.taskClient.UpdateTaskMessage(ctx, &request)
	if err != nil {
		return nil, err
	}
	taskDomain := utils.TaskResponseToTaskDomain(taskResponse)
	return &taskDomain, nil
}

func (t *ToDoServerAdapter) DeleteTask(
	ctx context.Context,
	userId string,
	taskId string,
	requestId string,
	token string,
) error {
	request := grpc.DeleteMessageRequest{
		TaskId:    taskId,
		RequestId: requestId,
		Token:     token,
	}
	_, err := t.taskClient.DeleteMessage(ctx, &request)
	return err
}
