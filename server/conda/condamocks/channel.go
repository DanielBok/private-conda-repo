package condamocks

import (
	"io"

	"private-conda-repo/conda/condatypes"
)

type MockChannel struct {
}

func (m MockChannel) Dir() string {
	return ""
}

func (m MockChannel) Index() error {
	return nil
}

func (m MockChannel) GetMetaInfo() (*condatypes.ChannelMetaInfo, error) {
	return &condatypes.ChannelMetaInfo{
		ChannelVersion: 0,
		Packages: map[string]struct {
			Subdirs []string `json:"subdirs"`
			Version string   `json:"version"`
		}{"package1": {
			Subdirs: []string{"dir1", "dir2"},
			Version: "0.1.0",
		}},
		Subdirs: []string{"dir1", "dir2"},
	}, nil
}

func (m MockChannel) AddPackage(_ io.Reader, _, _ string) error {
	return nil
}

func (m MockChannel) RemovePackage(_, _ string) error {
	return nil
}
