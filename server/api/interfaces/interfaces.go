package interfaces

import (
	"io"

	"private-conda-repo/api/dto"
	"private-conda-repo/domain/condatypes"
	"private-conda-repo/domain/entity"
	"private-conda-repo/infrastructure/decompressor"
)

type DataAccessLayer interface {
	Migrate() error

	CreateChannel(channel, password, email string) (*entity.Channel, error)
	GetChannel(channel string) (*entity.Channel, error)
	RemoveChannel(id int) error
	GetAllChannels() ([]*entity.Channel, error)

	GetPackageCounts(channelId int, name string) ([]*entity.PackageCount, error)
	CreatePackageCount(pkg *entity.PackageCount) (*entity.PackageCount, error)
	IncreasePackageCount(pkg *entity.PackageCount) (*entity.PackageCount, error)
	RemovePackageCount(pkg *entity.PackageCount) error
}

type Decompressor interface {
	// Retrieves MetaData from the .tar.bz2 file
	RetrieveMetadata(file io.ReadCloser) (*decompressor.MetaData, error)
}

type Indexer interface {
	// Indexes the directory via `conda index`. This should be run whenever a package is added,
	// removed or updated. It will update the current_repodata.json and repodata.json files in
	// the repository so that when we `conda install`, the dependency solver will know how to
	// look for files
	Index(dir string) error

	// Applies a series of fixes to repodata.json and current_repodata.json such as removal of
	// `python_abi` from the dependency lists. A list of instructions (fixes) must be provided
	// to determine which fixes to apply. Presently the supported values are
	//
	// - no-abi : Removes "python_abi *" dependencies from the uploaded package
	FixRepoData(dir string, fixes []string) error

	// Updates the indexer. This should not need to be called most times
	Update() error
}

type Channel interface {
	// Adds a package to the channel
	AddPackage(file io.Reader, pkg *dto.PackageDto, fixes []string) (*dto.PackageDto, error)

	// Returns the absolute path of channel
	Directory() string

	// Returns channel's channeldata.json file. This is useful for debugging the current state of the
	// channel
	GetChannelData() (*condatypes.ChannelData, error)

	// Reindex the channel folder. This should be called whenever there are changes to the packages
	// in the channel.
	Index(fixes []string) error

	// Returns the name of the channel
	Name() string

	// Removes a single package from the channel
	RemoveSinglePackage(pkg *dto.PackageDto) error

	// Removes all packages of the same name from the channel That is if you have a package called
	// numpy with different versions, this method will remove all versions of 'numpy'. Other packages
	// like 'scipy' will remain intact. Returns the number of packages removed
	RemovePackageAllVersions(name string) (int, error)
}

type FileSys interface {
	CreateChannel(channel string) (Channel, error)
	RenameChannel(oldName, newName string) (Channel, error)
	GetChannel(name string) (Channel, error)
	ListAllChannels() ([]Channel, error)
	RemoveChannel(name string) error
}
