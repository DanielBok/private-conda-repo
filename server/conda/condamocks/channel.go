package condamocks

import (
	"io"

	"github.com/stretchr/testify/mock"

	"private-conda-repo/conda/condatypes"
)

type MockChannel struct {
	mock.Mock
}

func (m MockChannel) AddPackage(file io.Reader, platform string, name string) (*condatypes.Package, error) {
	panic("implement me")
}

func (m MockChannel) RemoveSinglePackage(pkg *condatypes.Package) error {
	panic("implement me")
}

func (m MockChannel) RemovePackageAllVersions(name string) error {
	panic("implement me")
}

func (m MockChannel) Dir() string {
	return ""
}

func (m MockChannel) Index() error {
	return nil
}

func (m MockChannel) GetMetaInfo() (*condatypes.ChannelMetaInfo, error) {
	return &condatypes.ChannelMetaInfo{
		ChannelDataVersion: 0,
		Packages: map[string]condatypes.ChannelMetaPackageInfo{"perfana": {
			Subdirs:      []string{"noarch"},
			Version:      nil,
			ActivateD:    false,
			BinaryPrefix: false,
			DeactivateD:  false,
			Description:  nil,
			DevUrl:       nil,
			DocSourceUrl: nil,
			DocUrl:       nil,
			Home:         nil,
			IconHash:     nil,
			IconUrl:      nil,
			Identifiers:  nil,
			Keywords:     nil,
			License:      nil,
			PostLink:     false,
			PreLink:      false,
			PreUnlink:    false,
			RecipeOrigin: nil,
			SourceGitUrl: nil,
			SourceUrl:    nil,
			Summary:      nil,
			Tags:         nil,
			TextPrefix:   nil,
			Timestamp:    0,
		}},
		Subdirs: []string{"dir1", "dir2"},
	}, nil
}
