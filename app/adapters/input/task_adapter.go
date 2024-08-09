package input

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/marco-fpereira/to-do-list-client/app/adapters/input/dto"
	"github.com/marco-fpereira/to-do-list-client/app/adapters/utils"
	"github.com/marco-fpereira/to-do-list-client/app/domain/port/input"
)

type TaskAdapterInterface interface {
	GetAllTasks(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTaskMessage(ctx *gin.Context)
	UpdateTaskCompleteness(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

type TaskAdapter struct {
	TaskPort input.TaskPort
}

func NewTaskAdapter(
	taskPort input.TaskPort,
) TaskAdapterInterface {
	return &TaskAdapter{
		TaskPort: taskPort,
	}
}

func (t *TaskAdapter) GetAllTasks(ctx *gin.Context) {
	userId := ctx.Param("userId")
	tasksDomain, err := t.TaskPort.GetAllTasks(
		ctx,
		userId,
		utils.GetRequestId(ctx),
	)
	if err != nil {
		ctx.Error(err)
		return
	}
	tasksAdapter := []dto.TaskDetailsDTO{}
	copier.Copy(&tasksAdapter, &tasksDomain)
	ctx.JSON(http.StatusOK, tasksAdapter)
}

func (t *TaskAdapter) CreateTask(ctx *gin.Context) {
	userId := ctx.Param("userId")
	var taskDetailsDTO dto.TaskDetailsDTO
	if err := ctx.BindJSON(&taskDetailsDTO); err != nil {
		ctx.Error(err)
		return
	}

	taskDomain, err := t.TaskPort.CreateTask(
		ctx,
		userId,
		taskDetailsDTO.TaskMessage,
		utils.GetRequestId(ctx),
	)
	if err != nil {
		ctx.Error(err)
		return
	}
	var taskAdapter dto.TaskDetailsDTO
	copier.Copy(&taskAdapter, &taskDomain)
	ctx.JSON(http.StatusCreated, taskAdapter)
}

func (t *TaskAdapter) UpdateTaskMessage(ctx *gin.Context) {
	userId := ctx.Param("userId")
	taskId := ctx.Param("taskId")
	var taskDetailsDTO dto.TaskDetailsDTO
	if err := ctx.BindJSON(&taskDetailsDTO); err != nil {
		ctx.Error(err)
		return
	}

	taskDomain, err := t.TaskPort.UpdateTaskMessage(
		ctx,
		userId,
		taskId,
		taskDetailsDTO.TaskMessage,
		utils.GetRequestId(ctx),
	)
	if err != nil {
		ctx.Error(err)
		return
	}
	var taskAdapter dto.TaskDetailsDTO
	copier.Copy(&taskAdapter, &taskDomain)
	ctx.JSON(http.StatusCreated, taskAdapter)
}

func (t *TaskAdapter) UpdateTaskCompleteness(ctx *gin.Context) {
	userId := ctx.Param("userId")
	taskId := ctx.Param("taskId")
	err := t.TaskPort.UpdateTaskCompleteness(ctx, userId, taskId, utils.GetRequestId(ctx))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (t *TaskAdapter) DeleteTask(ctx *gin.Context) {
	userId := ctx.Param("userId")
	taskId := ctx.Param("taskId")
	err := t.TaskPort.DeleteTask(ctx, userId, taskId, utils.GetRequestId(ctx))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}
