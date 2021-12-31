package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tunardev/auth-user/pkg/routes"
)

type App struct {
	DB *sql.DB
	Router *mux.Router
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	app := App{}

	app.Initialize()	
	app.Routes()

	app.Run(":8080")
}

func (app *App) Initialize() {
	var err error
	app.DB, err = sql.Open("postgres", "postgres://postgres:postgretunar2000@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

    app.Router = mux.NewRouter()
}

func (app *App) Routes() {
	routes.UserSetup(app.Router, app.DB)
	routes.TaskSetup(app.Router, app.DB)

	app.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "page not found", "status": 404})
	})
}

func (app *App) Run(port string) {
	go func() {
		if err := http.ListenAndServe(port, app.Router); err != nil {
			panic(err)
		}
	}()

	<- make(chan struct{})
}