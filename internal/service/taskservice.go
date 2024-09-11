package service

import "github.com/SevvyP/tasks_v1/pkg/model"

type TaskService interface {
	GetTasks(userId string) (*[]model.Task, error)
	GetTasksByID(userId string, id string) (*[]model.Task, error)
	CreateTask(userId string) error
	UpdateTask(userId string) error
	DeleteTask(userId string) error
}
