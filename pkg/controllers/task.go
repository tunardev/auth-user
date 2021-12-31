package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tunardev/auth-user/pkg/entity"
	"github.com/tunardev/auth-user/pkg/service"
	"github.com/tunardev/auth-user/pkg/utils"
)

type TaskController interface {
	All(w http.ResponseWriter, r *http.Request, user entity.User)
	Get(w http.ResponseWriter, r *http.Request, user entity.User)
	Create(w http.ResponseWriter, r *http.Request, user entity.User)
	Update(w http.ResponseWriter, r *http.Request, user entity.User)
	Delete(w http.ResponseWriter, r *http.Request, user entity.User)
}

type taskController struct {
	service service.TaskService
}

func NewTaskController(service service.TaskService) TaskController {
	return taskController{service}
}

func (controller taskController) All(w http.ResponseWriter, r *http.Request, user entity.User) {
	data, err, status := controller.service.All(user)
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}

	utils.Response(w, status, nil, data)
}

func (controller taskController) Get(w http.ResponseWriter, r *http.Request, user entity.User) {
	params := mux.Vars(r)
	id := params["id"]
	
	data, err, status := controller.service.Get(id, user)
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}

	utils.Response(w, status, nil, data)
}

func (controller taskController) Create(w http.ResponseWriter, r *http.Request, user entity.User) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err, nil)
		return
	}

	var task entity.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, err, nil)
		return
	}

	data, err, status := controller.service.Create(task, user)
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}

	utils.Response(w, status, nil, data)
}

func (controller taskController) Update(w http.ResponseWriter, r *http.Request, user entity.User) {
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err, nil)
		return
	}

	var task entity.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, err, nil)
		return
	}

	data, err, status := controller.service.Update(id, task, user)
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}

	utils.Response(w, status, nil, data)
}

func (controller taskController) Delete(w http.ResponseWriter, r *http.Request, user entity.User) {
	params := mux.Vars(r)
	id := params["id"]

	data, err, status := controller.service.Delete(id, user)
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}
	
	utils.Response(w, status, nil, data)
}