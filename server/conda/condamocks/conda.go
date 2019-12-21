package condamocks

import (
	"github.com/stretchr/testify/mock"

	"private-conda-repo/conda"
)

func init() {
	conda.Register("mock", &MockConda{})
}

func New() (conda.Conda, error) {
	return &MockConda{}, nil
}

type MockConda struct {
	mock.Mock
}

func (m MockConda) CreateChannel(_ string) (conda.Channel, error) {
	return &MockChannel{}, nil
}

func (m MockConda) GetChannel(_ string) (conda.Channel, error) {
	return &MockChannel{}, nil
}

func (m MockConda) RemoveChannel(_ string) error {
	return nil
}

func (m MockConda) ChangeChannelName(_, _ string) (conda.Channel, error) {
	return &MockChannel{}, nil
}
