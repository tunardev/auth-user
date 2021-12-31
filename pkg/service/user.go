package service

import (
	"errors"
	"net/http"

	"github.com/tunardev/auth-user/pkg/entity"
	"github.com/tunardev/auth-user/pkg/repository"
	"github.com/tunardev/auth-user/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user entity.User) (interface{}, error, int, string)
	Login(user entity.User) (interface{}, error, int, string)
	Me(token string) (interface{}, error, int)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{repo}
}

func (service userService) Register(user entity.User) (interface{}, error, int, string) {
	err := user.IsValidRegister()
	if err != nil {
		return entity.User{}, err, http.StatusBadRequest, ""
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError, ""
	}
	user.Password = string(hashedPassword)

	id, err := service.repo.Insert(user)
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError, ""
	}
	user.ID = id

	token, err := utils.SignJWT(id)
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError, ""
	}

	return user.ToDTO(), nil, http.StatusOK, token
}

func (service userService) Login(user entity.User) (interface{}, error, int, string) {
	err := user.IsValidLogin()
	if err != nil {
		return entity.User{}, err, http.StatusBadRequest, ""
	}

	userData, err := service.repo.GetByEmail(user.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return entity.User{}, errors.New("User not found"), http.StatusBadRequest, ""
		}
		return entity.User{}, err, http.StatusInternalServerError, ""
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return entity.User{}, errors.New("Wrong password"), http.StatusBadRequest, ""
	}

	token, err := utils.SignJWT(userData.ID)
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError, ""
	}

	return userData.ToDTO(), nil, http.StatusOK, token
}

func (service userService) Me(token string) (interface{}, error, int) {
	id, err := utils.VerifyJWT(token)
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError
	}

	userData, err := service.repo.GetById(id)
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError
	}

	return userData.ToDTO(), nil, http.StatusOK
}

