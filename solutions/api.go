package solutions

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

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
	values := r.URL.Query()
	isAdult := values.Get("isAdult")
	if len(isAdult) > 0 {
		var friends []Friend
		var isAdultBool bool
		if isAdult == "true" {
			isAdultBool = true
		} else if isAdult == "false" {
			isAdultBool = false
		} else {
			JSON(w, r, http.StatusBadRequest, "invalid isAdult parameter")
			return
		}
		for _, friend := range a.Friends {
			if isAdultBool && friend.Age > 18 {
				friends = append(friends, friend)
			} else if !isAdultBool && friend.Age < 18 {
				friends = append(friends, friend)
			}
		}
		JSON(w, r, http.StatusOK, friends)
		return
	}
	JSON(w, r, http.StatusOK, a.Friends)
}

func (a *App) getById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	JSON(w, r, http.StatusOK, a.Friends[id-1])
}

func (a *App) add(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	var friend Friend
	if err := json.Unmarshal(bodyBytes, &friend); err != nil {
		JSON(w, r, http.StatusBadRequest, "Invalid JSON")
		return
	}
	friend.Id = len(a.Friends) + 1
	a.Friends = append(a.Friends, friend)
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
