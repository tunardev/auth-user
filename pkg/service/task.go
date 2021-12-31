package service

import (
	"errors"
	"net/http"

	"github.com/tunardev/auth-user/pkg/entity"
	"github.com/tunardev/auth-user/pkg/repository"
)

type TaskService interface {
	All(user entity.User) (interface{}, error, int)
	Create(task entity.Task, user entity.User) (interface{}, error, int)
	Get(id string, user entity.User) (interface{}, error, int)
	Delete(id string, user entity.User) (interface{}, error, int)
	Update(id string, task entity.Task, user entity.User) (interface{}, error, int)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return taskService{repo}
}

func (service taskService) All(user entity.User) (interface{}, error, int) {
	tasks, err := service.repo.All(user.ID)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return tasks, nil, http.StatusOK
}

func (service taskService) Get(id string, user entity.User) (interface{}, error, int) {
	task, err := service.repo.Get(id, user.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return entity.Task{}, errors.New("Task not found"), http.StatusBadRequest
		}

		return entity.Task{}, err, http.StatusInternalServerError
	}

	return task, nil, http.StatusOK
}

func (service taskService) Create(task entity.Task, user entity.User) (interface{}, error, int) {
	err := task.IsValidCreate()
	if err != nil {
		return entity.Task{}, err, http.StatusBadRequest
	}
	task.UserId = user.ID

	id, err := service.repo.Create(task)
	if err != nil {
		return entity.Task{}, err, http.StatusInternalServerError
	}
	task.ID = id

	return task, nil, http.StatusOK
}

func (service taskService) Delete(id string, user entity.User) (interface{}, error, int) {
	err := service.repo.Delete(id, user.ID)
	if err != nil {
		return entity.Task{}, err, http.StatusInternalServerError
	}

	return map[string]bool{"success": true}, nil, http.StatusOK
}

func (service taskService) Update(id string, task entity.Task, user entity.User) (interface{}, error, int) {
	err := task.IsValidUpdate()
	if err != nil {
		return entity.Task{}, err, http.StatusBadRequest
	}
	task.UserId = user.ID

	err = service.repo.Update(id, task)
	if err != nil {
		return entity.Task{}, err, http.StatusInternalServerError
	}

	return map[string]bool{"success": true}, nil, http.StatusOK
}