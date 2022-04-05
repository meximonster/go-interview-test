package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var MyApp = NewApp()

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	MyApp.Router.ServeHTTP(rr, req)

	return rr
}

func TestHealth(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	resp := executeRequest(req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.Code)
	}
}

func TestGetFriends(t *testing.T) {
	req, _ := http.NewRequest("GET", "/friends", nil)
	resp := executeRequest(req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.Code)
	}
	var friends []Friend
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading body: %v", err)
	}
	if err := json.Unmarshal(bodyBytes, &friends); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	if len(friends) != 3 {
		t.Errorf("Expected 3 friends. Got %d", len(friends))
	}
}

func TestGetAdultFriends(t *testing.T) {
	req, _ := http.NewRequest("GET", "/friends?isAdult=true", nil)
	resp := executeRequest(req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.Code)
	}
	var friends []Friend
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading body: %v", err)
	}
	if err := json.Unmarshal(bodyBytes, &friends); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	if len(friends) != 2 {
		t.Errorf("Expected 2 adult friends. Got %d", len(friends))
	}
	for _, friend := range friends {
		if friend.Age < 18 {
			t.Errorf("Expected adult friend. Got %v", friend)
		}
	}
}

func TestGetNonAdultFriends(t *testing.T) {
	req, _ := http.NewRequest("GET", "/friends?isAdult=false", nil)
	resp := executeRequest(req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.Code)
	}
	var friends []Friend
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading body: %v", err)
	}
	if err := json.Unmarshal(bodyBytes, &friends); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	if len(friends) != 1 {
		t.Errorf("Expected 1 non-adult friends. Got %d", len(friends))
	}
	for _, friend := range friends {
		if friend.Age >= 18 {
			t.Errorf("Expected non-adult friend. Got %v", friend)
		}
	}
}

func TestGetFriendById(t *testing.T) {
	req, _ := http.NewRequest("GET", "/friends/1", nil)
	resp := executeRequest(req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.Code)
	}
	var friend Friend
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading body: %v", err)
	}
	if err := json.Unmarshal(bodyBytes, &friend); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	if friend.Age != 20 || friend.Name != "Gina" {
		t.Errorf("Expected friend with name 'Gina' and age 20. Got %v", friend)
	}
}

func TestAddFriend(t *testing.T) {
	friend := Friend{Name: "Lebron", Age: 37}
	jsonStr, _ := json.Marshal(friend)
	req, _ := http.NewRequest("POST", "/friends", bytes.NewBuffer(jsonStr))
	resp := executeRequest(req)
	if resp.Code != http.StatusCreated {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, resp.Code)
	}
	if len(MyApp.Friends) != 4 {
		t.Errorf("Expected 4 friends. Got %d", len(MyApp.Friends))
	}
	if len(MyApp.Friends) == 4 {
		if MyApp.Friends[len(MyApp.Friends)-1].Name != "Lebron" || MyApp.Friends[len(MyApp.Friends)-1].Age != 37 ||
			MyApp.Friends[len(MyApp.Friends)-1].Id != 4 {
			t.Errorf("Expected friend with id 4, name 'Lebron' and age 37. Got %v", MyApp.Friends[4])
		}
	}
}
