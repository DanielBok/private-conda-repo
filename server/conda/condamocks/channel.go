package condamocks

import (
	"io"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"private-conda-repo/conda/condatypes"
)

type MockChannel struct {
	mock.Mock
}

var meta = &condatypes.ChannelMetaInfo{
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
		TextPrefix:   false,
		Timestamp:    0,
	}},
	Subdirs: []string{"dir1", "dir2"},
}

func (m MockChannel) AddPackage(file io.Reader, platform string, name string) (*condatypes.Package, error) {
	panic("implement me")
}

func (m MockChannel) RemoveSinglePackage(pkg *condatypes.Package) error {
	panic("implement me")
}

func (m MockChannel) RemovePackageAllVersions(name string) (int, error) {
	m.Called(name)
	if _, exists := meta.Packages[name]; exists {
		return 1, nil
	}

	return 0, errors.New("mock does not have package")
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
			TextPrefix:   false,
			Timestamp:    0,
		}},
		Subdirs: []string{"dir1", "dir2"},
	}, nil
}

func (m *MockChannel) GetPackageDetails(name string) ([]*condatypes.Package, error) {
	if _, exists := meta.Packages[name]; !exists {
		return nil, errors.Errorf("Package '%s' does not exist", name)
	}

	return []*condatypes.Package{
		{
			Name:        name,
			Version:     "0.1.0",
			BuildString: "py_a187v872",
			BuildNumber: 0,
			Platform:    "noarch",
		},
	}, nil
}
