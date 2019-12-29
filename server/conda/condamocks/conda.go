package condamocks

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"private-conda-repo/conda"
)

var channels = make(map[string]conda.Channel)

func init() {
	conda.Register("mock", &MockConda{})
}

type MockConda struct {
	mock.Mock
}

func (m MockConda) CreateChannel(channel string) (conda.Channel, error) {
	channels[channel] = &MockChannel{name: channel}
	return channels[channel], nil
}

func (m MockConda) GetChannel(channel string) (conda.Channel, error) {
	if chn, ok := channels[channel]; !ok {
		return nil, errors.Errorf("channel '%s' does not exist", channel)
	} else {
		return chn, nil
	}
}

func (m MockConda) RemoveChannel(channel string) error {
	if _, ok := channels[channel]; !ok {
		return errors.Errorf("channel '%s' does not exist", channel)
	} else {
		delete(channels, channel)
	}

	return nil
}

func (m MockConda) ChangeChannelName(_, _ string) (conda.Channel, error) {
	return &MockChannel{}, nil
}

func (m MockConda) ListAllChannels() ([]conda.Channel, error) {
	var output []conda.Channel

	for _, c := range channels {
		output = append(output, c)
	}

	return output, nil
}
