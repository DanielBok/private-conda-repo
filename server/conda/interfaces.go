package conda

import (
	"io"

	"private-conda-repo/conda/condatypes"
)

type Conda interface {
	// Creates a new Conda channel. Essentially, a conda channel is a folder containing all the necessary platform folders
	CreateChannel(channel string) (Channel, error)

	// Retrieves the specified channel
	GetChannel(channel string) (Channel, error)

	// Removes the specified Conda channel
	RemoveChannel(channel string) error

	// Renames the conda channel
	ChangeChannelName(oldChannel, newChannel string) (Channel, error)
}

type Channel interface {
	// Returns the file location of the channel's filesys
	Dir() string

	// Indexes the channel. This should be done whenever a package is uploaded or removed from the channel.
	// During this process, the metadata in the channeldata.json file will be updated. This json file is used
	// by conda to know how to install the package requested by the client
	Index() error

	// Gets the channel's meta information
	GetMetaInfo() (*condatypes.ChannelMetaInfo, error)

	// Adds a package into the channel
	AddPackage(file io.Reader, pkg *condatypes.Package) (*condatypes.Package, error)

	// Removes a package from the channel
	RemoveSinglePackage(pkg *condatypes.Package) error

	// Removes all versions of package specified by name. Returns the number of packages removed
	RemovePackageAllVersions(name string) (int, error)
}
