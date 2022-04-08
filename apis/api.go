package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type (
	Friend struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	App struct {
		// for protection against concurrent reading/writing.
		mtx     sync.Mutex
		Router  *mux.Router
		Friends []Friend
	}
)

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (a *App) getAll(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["isAdult"]
	var isAdult string
	if !(!ok || len(params[0]) < 1) {
		isAdult = params[0]
	}
	var response []Friend
	a.mtx.Lock()
	switch isAdult {
	case "true":
		for _, f := range a.Friends {
			if f.Age >= 18 {
				response = append(response, f)
			}
		}
	case "false":
		for _, f := range a.Friends {
			if f.Age < 18 {
				response = append(response, f)
			}
		}
	case "":
		response = a.Friends
	default:
		JSON(w, r, http.StatusBadRequest, "unrecognized input for isAdult")
		return
	}
	a.mtx.Unlock()
	if len(response) == 0 {
		JSON(w, r, http.StatusOK, "no friends match")
	}
	JSON(w, r, http.StatusOK, response)
}

func (a *App) getById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		JSON(w, r, http.StatusBadRequest, "missing id parameter")
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert id to int", err)
		JSON(w, r, http.StatusBadRequest, "wrong data type for id")
		return
	}
	a.mtx.Lock()
	var resp Friend
	for _, f := range a.Friends {
		if f.Id == intId {
			resp = f
		}
	}
	a.mtx.Unlock()
	JSON(w, r, http.StatusOK, resp)
}

func (a *App) add(w http.ResponseWriter, r *http.Request) {
	var f Friend
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		JSON(w, r, http.StatusBadRequest, err.Error())
	}
	if f.Name == "" {
		JSON(w, r, http.StatusBadRequest, "friend needs a name!")
	}
	f.Id = len(a.Friends) + 1
	a.mtx.Lock()
	a.Friends = append(a.Friends, f)
	a.mtx.Unlock()
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
