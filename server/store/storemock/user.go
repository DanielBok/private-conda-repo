package storemock

import (
	"github.com/pkg/errors"

	"private-conda-repo/store/models"
)

var users = make(map[string]*models.User)

func (m MockStore) AddUser(name, password string) (*models.User, error) {
	m.Called(name, password)
	u, err := models.NewUser(name, password)
	if err != nil {
		return nil, err
	}

	users[u.Name] = u
	return users[u.Name], nil
}

func (m *MockStore) GetAllUsers() ([]*models.User, error) {
	var userList []*models.User
	for _, u := range users {
		userList = append(userList, u)
	}

	return userList, nil
}

func (m MockStore) GetUser(name string) (*models.User, error) {
	if u, ok := users[name]; !ok {
		return nil, errors.Errorf("user '%s' does not exist", name)
	} else {
		return u, nil
	}
}

func (m MockStore) RemoveUser(name, password string) error {
	m.Called(name, password)
	u, ok := users[name]
	if !ok {
		return errors.Errorf("user '%s' does not exist", name)
	}
	if u.HasValidPassword(password) {
		delete(users, name)
		return nil
	} else {
		return errors.New("password does not match")
	}
}
