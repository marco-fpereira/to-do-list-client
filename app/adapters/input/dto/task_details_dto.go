package dto

import "time"

type TaskDetailsDTO struct {
	TaskId          string    `json:"taskId" copier:"TaskId"`
	TaskMessage     string    `json:"taskMessage" copier:"TaskMessage"`
	CreatedAt       time.Time `json:"createdAt" copier:"CreatedAt"`
	IsTaskCompleted bool      `json:"isTaskCompleted" copier:"IsTaskCompleted" default:"false"`
	UserId          string    `json:"userId" copier:"UserId"`
}
