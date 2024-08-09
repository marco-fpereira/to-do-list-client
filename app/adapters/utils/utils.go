package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marco-fpereira/to-do-list-client/app/config/grpc"
	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
)

func GetRequestId(ctx *gin.Context) string {
	requestId := ctx.Request.Header.Get("requestId")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	return requestId
}

func TaskResponseToTaskDomain(taskResponse *grpc.TaskDomain) model.TaskDomain {
	createdAt, err := time.Parse("2006-01-02", taskResponse.CreatedAt)
	if err != nil {
		createdAt = time.Now()
	}
	return model.TaskDomain{
		TaskId:          taskResponse.TaskId,
		TaskMessage:     taskResponse.TaskMessage,
		CreatedAt:       createdAt,
		IsTaskCompleted: taskResponse.IsTaskCompleted,
		UserId:          taskResponse.UserId,
	}
}
