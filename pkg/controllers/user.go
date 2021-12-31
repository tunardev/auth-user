package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/tunardev/auth-user/pkg/entity"
	"github.com/tunardev/auth-user/pkg/service"
	"github.com/tunardev/auth-user/pkg/utils"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request) 
	Login(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return userController{service}
}

func (controller userController) Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err, nil)
		return
	}

	var user entity.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, err, nil)
		return
	}

	data, err, status, token := controller.service.Register(user)
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}

	utils.Response(w, status, nil, map[string]interface{}{"user": data, "token": token})
}

func (controller userController) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err, nil)
		return
	}

	var user entity.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, err, nil)
		return
	}

	data, err, status, token := controller.service.Login(user)
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}

	utils.Response(w, status, nil, map[string]interface{}{"user": data, "token": token})
}

func (controller userController) Me(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "" {
		utils.Response(w, http.StatusUnauthorized, errors.New("Wrong Authorization"), nil)
		return
	}

	if strings.Split(r.Header.Get("Authorization"), " ")[1] == "" {
		utils.Response(w, http.StatusUnauthorized, errors.New("Wrong Authorization"), nil)
		return
	}
	
	data, err, status := controller.service.Me(strings.Split(r.Header.Get("Authorization"), " ")[1])
	if err != nil {
		utils.Response(w, status, err, nil)
		return
	}

	utils.Response(w, status, nil, data)
}
