package config

import (
	"database/sql"

	"github.com/tunardev/auth-user/pkg/controllers"
	"github.com/tunardev/auth-user/pkg/repository"
	"github.com/tunardev/auth-user/pkg/service"
)

func UserInit(db *sql.DB) controllers.UserController {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	controller := controllers.NewUserController(service)

	return controller
}

func TaskInit(db *sql.DB) controllers.TaskController {
	repo := repository.NewTaskRepository(db)
	service := service.NewTaskService(repo)
	controller := controllers.NewTaskController(service)

	return controller
}