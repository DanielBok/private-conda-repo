package storemock

import (
	"github.com/pkg/errors"

	"private-conda-repo/store/models"
)

var packageCounts = make(map[string]map[string]map[string]*models.PackageCount)

func (m MockStore) GetPackageCounts(channel, name string) ([]*models.PackageCount, error) {
	var counts []*models.PackageCount
	for _, p := range packageCounts[channel][name] {
		counts = append(counts, p)
	}
	return counts, nil
}

func (m MockStore) CreatePackageCount(pkg *models.PackageCount) (*models.PackageCount, error) {
	c := pkg.Channel
	n := pkg.Package
	p := pkg.Platform

	if _, exists := packageCounts[c]; !exists {
		packageCounts[c] = make(map[string]map[string]*models.PackageCount)
	}

	if _, exists := packageCounts[c][n]; !exists {
		packageCounts[c][n] = make(map[string]*models.PackageCount)
	}

	if _, exists := packageCounts[c][n][p]; !exists {
		packageCounts[c][n][p] = pkg
		return pkg, nil
	} else {
		return nil, errors.New("package already exists")
	}
}

func (m MockStore) IncreasePackageCount(channel, name, platform string) (*models.PackageCount, error) {
	if p, exists := packageCounts[channel][name][platform]; !exists {
		return nil, errors.Errorf("package '%s/%s/%s' does not exist", channel, name, platform)
	} else {
		p.Count += 1
		return p, nil
	}
}
