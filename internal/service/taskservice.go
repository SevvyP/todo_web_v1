package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SevvyP/tasks_v1/pkg/model"
)

type TaskConfig struct {
	BaseURL string `json:"baseurl"`
}

type TaskService interface {
	GetTasks(userId string) (*[]model.Task, error)
	CreateTask(userId string) error
	UpdateTask(userId string) error
	DeleteTask(userId string) error
}

type taskService struct {
	config       TaskConfig
	tokenService TokenService
}

func NewTaskService(config TaskConfig, tokenSevice TokenService) TaskService {
	return &taskService{
		config:       config,
		tokenService: tokenSevice,
	}
}

func (t *taskService) GetTasks(userId string) (*[]model.Task, error) {
	token, err := t.tokenService.GenerateToken()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", t.config.BaseURL+"/tasks?user_id="+userId, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Failed to get tasks from server status code: " + fmt.Sprint(resp.StatusCode))
	}
	tasks := []model.Task{}
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (t *taskService) CreateTask(userId string) error {
	return nil
}

func (t *taskService) UpdateTask(userId string) error {
	return nil
}

func (t *taskService) DeleteTask(userId string) error {
	return nil
}
