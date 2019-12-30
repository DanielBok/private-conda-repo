package storemock

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"private-conda-repo/store/models"
)

var packageCounts = make(map[string]*models.PackageCount)

func (m MockStore) GetPackageCounts(channel, name string) ([]*models.PackageCount, error) {
	var counts []*models.PackageCount
	for _, p := range packageCounts {
		if p.Channel == channel && p.Package == name {
			counts = append(counts, p)
		}
	}
	return counts, nil
}

func (m MockStore) CreatePackageCount(pkg *models.PackageCount) (*models.PackageCount, error) {
	key := formKey(pkg)
	if _, exists := packageCounts[key]; !exists {
		packageCounts[key] = pkg
		return pkg, nil
	} else {
		return nil, errors.New("package already exists")
	}
}

func (m MockStore) IncreasePackageCount(pkg *models.PackageCount) (*models.PackageCount, error) {
	if p, exists := packageCounts[formKey(pkg)]; !exists {
		return nil, errors.New("package does not exist")
	} else {
		p.Count += 1
		return p, nil
	}
}

func (m MockStore) RemovePackageCount(pkg *models.PackageCount) error {
	key := formKey(pkg)
	if _, exists := packageCounts[key]; !exists {
		return errors.New("package does not exist")
	} else {
		delete(packageCounts, key)
		return nil
	}
}

func formKey(pkg *models.PackageCount) string {
	return strings.Join([]string{
		pkg.Channel,
		pkg.Package,
		pkg.BuildString,
		strconv.Itoa(pkg.BuildNumber),
		pkg.Version,
		pkg.Platform,
	}, ":::")
}
