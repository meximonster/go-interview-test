package task

import (
	"testing"
)

type FakeUserFetcher struct{}

func (f *FakeUserFetcher) GetUsers() ([]User, error) {
	return []User{{
		Id:       1,
		Name:     "John",
		Username: "Stankovic",
		Email:    "js@example.com",
		Company: Company{
			Name:        "Marketing Company",
			CatchPhrase: "best multimedia ever!",
		},
	}, {
		Id:       2,
		Name:     "Mike",
		Username: "Malone",
		Email:    "mm@example.com",
		Company: Company{
			Name:        "Spinalonga Records",
			CatchPhrase: "tasks within a minute!",
		},
	}, {
		Id:       3,
		Name:     "Ritchie",
		Username: "Finestra",
		Email:    "rf@example.com",
		Company: Company{
			Name:        "Vinyl",
			CatchPhrase: "random",
		},
	}, {
		Id:       4,
		Name:     "Nicola",
		Username: "Jokic",
		Email:    "nj@example.com",
		Company: Company{
			Name:        "Basketball",
			CatchPhrase: "only but net",
		},
	}}, nil
}

func TestMyApp_FilterByKeyword(t *testing.T) {
	userFetcher := FakeUserFetcher{}
	app := SimpleApp{&userFetcher}
	keywords := []string{"server", "net", "multimedia", "task"}
	users, err := app.FilterByKeyword(keywords)
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(users) != 3 {
		t.Errorf("Expected 3 users, got %d", len(users))
	}
}
