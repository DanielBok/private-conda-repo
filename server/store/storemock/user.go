package storemock

import "private-conda-repo/store/models"

func (m MockStore) AddUser(name, password string) (*models.User, error) {
	args := m.Called(name, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m MockStore) GetUser(name string) (*models.User, error) {
	panic("implement me")
}

func (m MockStore) RemoveUser(name, password string) error {
	panic("implement me")
}
