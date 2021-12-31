package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/tunardev/auth-user/pkg/config"
	"github.com/tunardev/auth-user/pkg/middleware"
	"github.com/tunardev/auth-user/pkg/repository"
)

func TaskSetup(router *mux.Router, db *sql.DB) {
	taskRoutes := config.TaskInit(db)
	repo := repository.NewUserRepository(db)
	
	router.HandleFunc("/tasks", middleware.IsAuth(taskRoutes.All, repo)).Methods("GET")
	router.HandleFunc("/tasks/{id}", middleware.IsAuth(taskRoutes.Get, repo)).Methods("GET")
	router.HandleFunc("/tasks", middleware.IsAuth(taskRoutes.Create, repo)).Methods("POST")
	router.HandleFunc("/tasks/{id}", middleware.IsAuth(taskRoutes.Update, repo)).Methods("PUT")
	router.HandleFunc("/tasks/{id}", middleware.IsAuth(taskRoutes.Delete, repo)).Methods("DELETE")
}