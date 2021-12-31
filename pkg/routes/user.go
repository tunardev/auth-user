package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/tunardev/auth-user/pkg/config"
)

func UserSetup(router *mux.Router, db *sql.DB) {
	userRoutes := config.UserInit(db)

	router.HandleFunc("/users/register", userRoutes.Register).Methods("POST")
	router.HandleFunc("/users/login", userRoutes.Login).Methods("POST")
	router.HandleFunc("/users/me", userRoutes.Me).Methods("GET")
}