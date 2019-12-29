package postgres

import (
	"time"

	"github.com/pkg/errors"

	"private-conda-repo/store/models"
)

func (s *Store) GetPackageCounts(channel, name string) ([]*models.PackageCount, error) {
	var counts []*models.PackageCount
	if errs := s.db.
		Where("channel = ? AND package = ?", channel, name).
		Find(&counts).
		GetErrors(); len(errs) > 0 {
		return nil, errors.Wrapf(joinErrors(errs), "could not get count data from '%s' for package '%s'", channel, name)
	}
	return counts, nil
}

func (s *Store) CreatePackageCount(pkg *models.PackageCount) (*models.PackageCount, error) {
	if errs := s.db.
		Where(models.PackageCount{
			Channel:     pkg.Channel,
			Package:     pkg.Package,
			BuildString: pkg.BuildString,
			BuildNumber: pkg.BuildNumber,
			Version:     pkg.Version,
			Platform:    pkg.Platform,
		}).Assign(models.PackageCount{
		Count:      0,
		UploadDate: time.Now(),
	}).FirstOrCreate(pkg).
		GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}
	return pkg, nil
}

func (s *Store) IncreasePackageCount(channel, name, platform string) (*models.PackageCount, error) {
	var count models.PackageCount
	if errs := s.db.
		Where("channel = ? AND package = ? AND platform = ?", channel, name, platform).
		First(&count).
		GetErrors(); len(errs) > 0 {
		return nil, errors.Wrap(joinErrors(errs), "could not update count")
	}

	s.db.Model(&count).Update("count", count.Count+1)
	return &count, nil
}
