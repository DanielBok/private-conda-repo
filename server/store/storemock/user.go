package storemock

import "private-conda-repo/store/models"

func (m MockStore) AddUser(name, password string) (*models.User, error) {
	m.Called(name, password)
	return models.NewUser(name, password)
}

func (m *MockStore) GetAllUsers() ([]*models.User, error) {
	return []*models.User{
		{
			Id:   1,
			Name: "Daniel",
		}, {
			Id:   2,
			Name: "Pikachu",
		},
	}, nil
}

func (m MockStore) GetUser(name string) (*models.User, error) {
	panic("implement me")
}

func (m MockStore) RemoveUser(name, password string) error {
	m.Called(name, password)
	return nil
}
