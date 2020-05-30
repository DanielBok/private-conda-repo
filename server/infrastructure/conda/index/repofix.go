package index

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"private-conda-repo/domain/condatypes"
	"private-conda-repo/domain/enum"
	"private-conda-repo/libs"
)

// This function does additional formatting on current_repodata.json and repodata.json.
// These additional transformations are required to fix some issues that conda index
// introduces. For example, conda index will unnecessarily create a dependency on
// python_abi which is usually not what we want. This function maybe deprecated in
// a future version after checking that conda index doesn't do anything weird.
func fixRepoDataFiles(channelPath string) error {
	if !libs.PathExists(channelPath) {
		return errors.Errorf("channel path: '%s' does not exist", channelPath)
	}

	for _, platform := range []enum.Platform{
		enum.LINUX32,
		enum.LINUX64,
		enum.WIN32,
		enum.WIN64,
		enum.OSX64,
		enum.NOARCH,
	} {
		folder := filepath.Join(channelPath, string(platform))
		if !libs.PathExists(folder) {
			continue
		}

		for _, file := range []string{"current_repodata.json", "repodata.json"} {
			if !libs.PathExists(file) {
				continue
			}

			fp := filepath.Join(folder, file)
			data, hasChanges, err := FixRepoData(fp)
			if err != nil {
				return err
			}

			if hasChanges {
				err = overwriteRepoDataFile(data, fp)
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

func overwriteRepoDataFile(data *condatypes.RepoData, fp string) error {
	dir, file := filepath.Split(fp)
	path := filepath.Join(filepath.Base(dir), file)

	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "could not overwrite %s", path)
	}
	defer func() { _ = f.Close() }()

	err = json.NewEncoder(f).Encode(&data)
	if err != nil {
		return errors.Wrapf(err, "could not write data to %s", path)
	}

	return nil
}

func FixRepoData(fp string) (repoData *condatypes.RepoData, hasChanges bool, err error) {
	dir, file := filepath.Split(fp)
	path := filepath.Join(filepath.Base(dir), file) // return "informative: filepath. i.e. noarch/repodata.json

	f, err := os.Open(fp)
	if err != nil {
		return nil, false, errors.Wrapf(err, "could not read %s", path)
	}

	err = json.NewDecoder(f).Decode(&repoData)
	if err != nil {
		return nil, false, errors.Wrapf(err, "could not decode %s", path)
	}

	for k, v := range repoData.Packages {
		var depends []string
		for _, dependency := range v.Depends {
			dependency = strings.ToLower(dependency)

			// remove pesky python_abi which is caused when conda-build puts in a specific abi dependency
			// when it is not needed thus rendering the package impossible to download.
			if strings.HasPrefix(dependency, "python_abi") {
				hasChanges = true // do no add this dependency skip
			} else {
				depends = append(depends, dependency)
			}

		}

		repoData.Packages[k].Depends = depends
	}

	return
}

func (d *DockerIndex) FixRepoData(dir string) error {
	return fixRepoDataFiles(dir)
}

func (s *ShellIndex) FixRepoData(dir string) error {
	return fixRepoDataFiles(dir)
}
