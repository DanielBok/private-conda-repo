package condamocks

import (
	"private-conda-repo/conda/types"
)

func New() (types.Conda, error) {
	return &MockConda{}, nil
}

type MockConda struct {
}

func (m MockConda) CreateChannel(_ string) (types.Channel, error) {
	return &MockChannel{}, nil
}

func (m MockConda) GetChannel(_ string) (types.Channel, error) {
	return &MockChannel{}, nil
}

func (m MockConda) RemoveChannel(_ string) error {
	return nil
}

func (m MockConda) ChangeChannelName(_, _ string) (types.Channel, error) {
	return &MockChannel{}, nil
}
