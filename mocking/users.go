package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const url = "https://jsonplaceholder.typicode.com/users"

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Company  Company `json:"company"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
}

type UserFetcher struct{}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (u *UserFetcher) GetUsers(client HTTPClient) ([]User, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get users request error: %v", err)
	}
	defer resp.Body.Close()
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("cannot decode response: %v", err)
	}
	return users, nil
}

type SimpleApp struct {
	UserFetcher *UserFetcher
}

func (a *SimpleApp) FilterByKeyword(keywords []string, client HTTPClient) ([]User, error) {
	users, err := a.UserFetcher.GetUsers(client)
	if err != nil {
		return nil, fmt.Errorf("cannot get users: %v", err)
	}
	var filteredUsers []User
	for _, user := range users {
		for _, keyword := range keywords {
			if strings.Contains(user.Company.CatchPhrase, keyword) {
				filteredUsers = append(filteredUsers, user)
				break
			}
		}
	}
	return filteredUsers, nil
}
