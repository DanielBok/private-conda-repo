package api_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"private-conda-repo/api/dto"
	"private-conda-repo/api/interfaces"
	"private-conda-repo/domain/condatypes"
	"private-conda-repo/domain/entity"
	"private-conda-repo/infrastructure/decompressor"
)

func NewTestRequest(method, target string, body io.Reader, routeParams map[string]string) *http.Request {
	r := httptest.NewRequest(method, target, body)

	if len(routeParams) > 0 {
		routeCtx := chi.NewRouteContext()
		for key, value := range routeParams {
			routeCtx.URLParams.Add(key, value)
		}

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, routeCtx))
	}

	return r
}

func NewMockDb() *MockDb {
	return &MockDb{
		channels:      make(map[string]*entity.Channel),
		packageCounts: make(map[string]*entity.PackageCount),
	}
}

type MockDb struct {
	channels      map[string]*entity.Channel
	packageCounts map[string]*entity.PackageCount
}

func (m *MockDb) Migrate() error {
	return nil
}

func (m *MockDb) CreateChannel(channel, password, email string) (*entity.Channel, error) {
	c := entity.NewChannel(channel, password, email)

	c.Id = len(m.channels) + 1
	m.channels[c.Channel] = c
	return c, nil
}

func (m *MockDb) GetChannel(channel string) (*entity.Channel, error) {
	if c, ok := m.channels[channel]; !ok {
		return nil, gorm.ErrRecordNotFound
	} else {
		return c, nil
	}
}

func (m *MockDb) RemoveChannel(id int) error {
	for key, c := range m.channels {
		if c.Id == id {
			delete(m.channels, key)
			return nil
		}
	}

	return errors.Errorf("no channel with id %d", id)
}

func (m *MockDb) GetAllChannels() ([]*entity.Channel, error) {
	var channels []*entity.Channel
	for _, c := range m.channels {
		channels = append(channels, c)
	}

	return channels, nil
}

func (m *MockDb) GetPackageCounts(channelId int, name string) ([]*entity.PackageCount, error) {
	var counts []*entity.PackageCount
	for _, p := range m.packageCounts {
		if p.ChannelId == channelId && p.Package == name {
			counts = append(counts, p)
		}
	}
	return counts, nil
}

func (m *MockDb) CreatePackageCount(pkg *entity.PackageCount) (*entity.PackageCount, error) {
	key := m.packageCountKey(pkg)
	if _, exists := m.packageCounts[key]; !exists {
		m.packageCounts[key] = pkg
		return pkg, nil
	} else {
		return nil, errors.New("package already exists")
	}
}

func (m *MockDb) IncreasePackageCount(pkg *entity.PackageCount) (*entity.PackageCount, error) {
	if p, exists := m.packageCounts[m.packageCountKey(pkg)]; !exists {
		return nil, errors.New("package does not exist")
	} else {
		p.Count += 1
		return p, nil
	}
}

func (m *MockDb) RemovePackageCount(pkg *entity.PackageCount) error {
	key := m.packageCountKey(pkg)
	if _, exists := m.packageCounts[key]; !exists {
		return errors.New("package does not exist")
	} else {
		delete(m.packageCounts, key)
		return nil
	}
}

func (m *MockDb) packageCountKey(pkg *entity.PackageCount) string {
	return strings.Join([]string{
		strconv.Itoa(pkg.ChannelId),
		pkg.Package,
		pkg.BuildString,
		strconv.Itoa(pkg.BuildNumber),
		pkg.Version,
		pkg.Platform,
	}, "::")
}

func NewMockDecompressor() *MockDecompressor {
	return &MockDecompressor{Filepath: "Replace this with file value when testing upload"}
}

type MockDecompressor struct {
	Filepath string
}

func (m MockDecompressor) RetrieveMetadata(io.ReadCloser) (*decompressor.MetaData, error) {
	return &decompressor.MetaData{
		Package: &dto.PackageDto{
			Name:        pkgName,
			Version:     "0.0.1",
			BuildString: "py",
			BuildNumber: 0,
			Platform:    "noarch",
		},
		Filepath: m.Filepath,
	}, nil
}

func NewMockFileSys() *MockFileSys {
	return &MockFileSys{
		channels: map[string]interfaces.Channel{},
	}
}

type MockFileSys struct {
	channels map[string]interfaces.Channel
}

func (m *MockFileSys) CreateChannel(channel string) (interfaces.Channel, error) {
	if _, exist := m.channels[channel]; exist {
		return nil, errors.New("channel exists")
	}

	chn := NewMockChannel(channel)
	m.channels[channel] = chn
	return chn, nil
}

func (m *MockFileSys) RenameChannel(oldName, newName string) (interfaces.Channel, error) {
	if _, exist := m.channels[oldName]; !exist {
		return nil, errors.New("channel does not exists")
	}

	m.channels[newName] = m.channels[oldName]
	delete(m.channels, oldName)
	return m.channels[newName], nil
}

func (m *MockFileSys) GetChannel(name string) (interfaces.Channel, error) {
	c, exist := m.channels[name]
	if !exist {
		return nil, errors.New("channel does not exists")
	}
	return c, nil
}

func (m *MockFileSys) ListAllChannels() ([]interfaces.Channel, error) {
	var channels []interfaces.Channel
	for _, c := range m.channels {
		channels = append(channels, c)
	}
	return channels, nil
}

func (m *MockFileSys) RemoveChannel(name string) error {
	if _, exist := m.channels[name]; !exist {
		return errors.New("channel does not exists")
	}

	delete(m.channels, name)
	return nil
}

func NewMockChannel(channel string) *MockChannel {
	return &MockChannel{channel: channel}
}

type MockChannel struct {
	channel string
}

func (m *MockChannel) AddPackage(_ io.Reader, pkg *dto.PackageDto, _ []string) (*dto.PackageDto, error) {
	return pkg, nil
}

func (m *MockChannel) Directory() string {
	return m.channel
}

func (m *MockChannel) GetChannelData() (*condatypes.ChannelData, error) {
	packages := make(map[string]condatypes.PackageData)

	for i := 0; i < nVerPerPkg; i++ {
		packages[pkgName] = condatypes.PackageData{}
	}

	return &condatypes.ChannelData{
		ChannelDataVersion: 0,
		Packages:           packages,
		Subdirs:            []string{"win-64", "noarch"},
	}, nil
}

func (m *MockChannel) Index([]string) error {
	return nil
}

func (m *MockChannel) Name() string {
	return m.channel
}

func (m *MockChannel) RemoveSinglePackage(*dto.PackageDto) error {
	return nil
}

func (m *MockChannel) RemovePackageAllVersions(_ string) (int, error) {
	return 1, nil
}
