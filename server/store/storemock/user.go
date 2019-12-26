package storemock

import (
	"github.com/pkg/errors"

	"private-conda-repo/store/models"
)

var users = make(map[string]*models.User)

func (m MockStore) AddUser(channel, password string) (*models.User, error) {
	m.Called(channel, password)
	u, err := models.NewUser(channel, password)
	if err != nil {
		return nil, err
	}

	users[u.Channel] = u
	return users[u.Channel], nil
}

func (m *MockStore) GetAllUsers() ([]*models.User, error) {
	var userList []*models.User
	for _, u := range users {
		userList = append(userList, u)
	}

	return userList, nil
}

func (m MockStore) GetUser(channel string) (*models.User, error) {
	if u, ok := users[channel]; !ok {
		return nil, errors.Errorf("user '%s' does not exist", channel)
	} else {
		return u, nil
	}
}

func (m MockStore) RemoveUser(channel, password string) error {
	m.Called(channel, password)
	u, ok := users[channel]
	if !ok {
		return errors.Errorf("user '%s' does not exist", channel)
	}
	if u.HasValidPassword(password) {
		delete(users, channel)
		return nil
	} else {
		return errors.New("password does not match")
	}
}
