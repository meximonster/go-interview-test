package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	Friend struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	Friends = []Friend{{
		Id:   1,
		Name: "Michael",
	}, {
		Id:   2,
		Name: "Angela",
	}, {
		Id:   3,
		Name: "Maria",
	}}
)

func health(w http.ResponseWriter, r *http.Request) {
	JSON(w, r, http.StatusOK, "hello.")
}

func handleFriends(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		JSON(w, r, http.StatusOK, Friends)
	case http.MethodPost:
		// fill your code here for a friend addition!
		JSON(w, r, http.StatusCreated, nil)
	}
}

func getById(w http.ResponseWriter, r *http.Request) {
	// write your code here to return a friend by his/her id!
	JSON(w, r, http.StatusNotFound, nil)
}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Println("failed to marshal ErrorResp:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleRequests() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/friends", handleFriends)
	http.HandleFunc("/friends/", getById)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// Lets the API run in the background. Then some actions for checking API functionality are taken.
	go handleRequests()

	// check if health request returns 200
	if err := checkHealth(); err != nil {
		log.Fatalf("health request error: %s", err)
	}

	// check if GET /friends request returns a Friend slice of length 3.
	friendList, err := doGetAll()
	if err != nil {
		log.Fatalf("handleFriends request error: %s", err)
	}
	if len(friendList) != 3 {
		log.Fatalf("friend list length should be 3!")
	}

	// checks if GET /friends/1 request returns friend with id 1.
	friend, err := doGetOne(1)
	if err != nil {
		log.Fatalf("get one error: %s", err)
	}
	if friend.Id != 1 {
		log.Fatalf("getting friend with id 1 failed. API returned: %v", friend)
	}

	// checks if POST /friends actually adds a new friend and the Friends slice length is incremented by 1.
	if err := doAddOne(); err != nil {
		log.Fatalf("adding a new friend failed: %s", err)
	}
	if len(Friends) != 4 {
		log.Fatalf("adding a new friend failed. Friends slice length is different than 4. Current length: %d",
			len(Friends))
	}
	friend, err = doGetOne(4)
	if err != nil {
		log.Fatalf("get one error: %s", err)
	}
	if friend.Id != 4 {
		log.Fatalf("getting friend with id 4 failed. API returned: %v", friend)
	}

	// Every check succeeds. You're the best!
	fmt.Printf("SUCCESS!")
}

func checkHealth() error {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/health", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot perform request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	return nil
}

func doGetAll() ([]Friend, error) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/friends", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot perform request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %s", err)
	}
	defer resp.Body.Close()
	var friendsList []Friend
	if err := json.Unmarshal(bodyBytes, &friendsList); err != nil {
		return nil, fmt.Errorf("cannot unmarshal body: %s", err)
	}
	return friendsList, nil
}

func doGetOne(id int) (Friend, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/friends/%d", id), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Friend{}, fmt.Errorf("cannot perform request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return Friend{}, fmt.Errorf("friend with id %d not found", id)
		}
		return Friend{}, fmt.Errorf("bad status code: %d. Expected: %d", resp.StatusCode, http.StatusCreated)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Friend{}, fmt.Errorf("cannot read body: %s", err)
	}
	defer resp.Body.Close()
	var friend Friend
	if err := json.Unmarshal(bodyBytes, &friend); err != nil {
		return Friend{}, fmt.Errorf("cannot unmarshal body: %s", err)
	}
	return friend, nil
}

func doAddOne() error {
	friend := Friend{
		Id:   4,
		Name: "Pablo",
	}
	body, _ := json.Marshal(friend)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/friends", bytes.NewBuffer(body))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot perform request: %s", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("bad status code: %d. Expected: %d", resp.StatusCode, http.StatusCreated)
	}
	return nil
}
