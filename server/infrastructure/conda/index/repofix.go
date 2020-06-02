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
func fixRepoDataFiles(channelPath string, fixes []string) error {
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
			fp := filepath.Join(folder, file)
			if !libs.PathExists(fp) {
				continue
			}

			data, hasChanges, err := FixRepoData(fp, fixes)
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

func FixRepoData(fp string, fixes []string) (repoData *condatypes.RepoData, hasChanges bool, err error) {
	fixSet := make(map[string]bool)
	for _, name := range fixes {
		fixSet[strings.ToLower(strings.TrimSpace(name))] = true
	}

	if len(fixSet) == 0 {
		return
	}

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

			if fixSet["no-abi"] && strings.HasPrefix(dependency, "python_abi") {
				// remove pesky python_abi which is caused when conda-build puts in a specific abi dependency
				// when it is not needed thus rendering the package impossible to download.
				hasChanges = true // do not add python_abi dependency
				continue
			}

			depends = append(depends, dependency)

		}

		repoData.Packages[k].Depends = depends
	}

	return
}

func overwriteRepoDataFile(data *condatypes.RepoData, fp string) error {
	dir, file := filepath.Split(fp)
	path := filepath.Join(filepath.Base(dir), file)

	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "could not overwrite %s", path)
	}
	defer libs.IOCloser(f)

	enc := json.NewEncoder(f)
	enc.SetIndent(" ", "  ")
	err = enc.Encode(&data)
	if err != nil {
		return errors.Wrapf(err, "could not write data to %s", path)
	}

	return nil
}

func (d *DockerIndex) FixRepoData(dir string, fixes []string) error {
	return fixRepoDataFiles(dir, fixes)
}

func (s *ShellIndex) FixRepoData(dir string, fixes []string) error {
	return fixRepoDataFiles(dir, fixes)
}
