package postgres

import (
	"time"

	"github.com/pkg/errors"

	"private-conda-repo/domain/entity"
)

func (p *Postgres) GetPackageCounts(channelId int, name string) ([]*entity.PackageCount, error) {
	var counts []*entity.PackageCount
	if errs := p.db.
		Where("channel_id = ? AND package = ?", channelId, name).
		Find(&counts).
		GetErrors(); len(errs) > 0 {
		return nil, errors.Wrapf(joinErrors(errs), "could not get count data for channel '%d' for package '%s'", channelId, name)
	}
	return counts, nil
}

func (p *Postgres) CreatePackageCount(pkg *entity.PackageCount) (*entity.PackageCount, error) {
	if errs := p.db.
		Where(entity.PackageCount{
			ChannelId:   pkg.ChannelId,
			Package:     pkg.Package,
			BuildString: pkg.BuildString,
			BuildNumber: pkg.BuildNumber,
			Version:     pkg.Version,
			Platform:    pkg.Platform,
		}).Assign(entity.PackageCount{
		Count:      0,
		UploadDate: time.Now().UTC(),
	}).FirstOrCreate(pkg).
		GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}
	return pkg, nil
}

func (p *Postgres) IncreasePackageCount(pkg *entity.PackageCount) (*entity.PackageCount, error) {
	var count entity.PackageCount
	if errs := p.db.
		Where(entity.PackageCount{
			ChannelId:   pkg.ChannelId,
			Package:     pkg.Package,
			BuildString: pkg.BuildString,
			BuildNumber: pkg.BuildNumber,
			Version:     pkg.Version,
			Platform:    pkg.Platform,
		}).
		First(&count).
		GetErrors(); len(errs) > 0 {
		return nil, errors.Wrap(joinErrors(errs), "could not update count")
	}

	p.db.Model(&count).Update("count", count.Count+1)
	return &count, nil
}

func (p *Postgres) RemovePackageCount(pkg *entity.PackageCount) error {
	var record entity.PackageCount
	if errs := p.db.Where(entity.PackageCount{
		ChannelId:   pkg.ChannelId,
		Package:     pkg.Package,
		BuildString: pkg.BuildString,
		BuildNumber: pkg.BuildNumber,
		Version:     pkg.Version,
		Platform:    pkg.Platform,
	}).First(&record).GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	if errs := p.db.Delete(record).GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	return nil
}
