package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Friend struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	App struct {
		Router  *mux.Router
		Friends []Friend
	}
)

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (a *App) getAll(w http.ResponseWriter, r *http.Request) {
	// place your code and add an ?isAdult=true or ?isAdult=false filter
	JSON(w, r, http.StatusOK, a.Friends)
}

func (a *App) getById(w http.ResponseWriter, r *http.Request) {
	// place your code here to get a friend by id
	JSON(w, r, http.StatusOK, nil)
}

func (a *App) add(w http.ResponseWriter, r *http.Request) {
	// place your code here to add a new friend
	JSON(w, r, http.StatusCreated, nil)
}

func NewApp() *App {
	app := &App{Friends: []Friend{
		{Id: 1, Name: "Gina", Age: 20},
		{Id: 2, Name: "John", Age: 13},
		{Id: 3, Name: "Angela", Age: 42},
	}}

	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/health", app.healthCheck).Methods("GET")
	app.Router.HandleFunc("/friends", app.getAll).Methods("GET")
	app.Router.HandleFunc("/friends/{id}", app.getById).Methods("GET")
	app.Router.HandleFunc("/friends", app.add).Methods("POST")

	return app
}

func (a *App) Run() {
	log.Fatalln(http.ListenAndServe(":8080", a.Router))
}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Println("failed to marshal ErrorResp:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
