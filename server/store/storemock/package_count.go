package storemock

import "private-conda-repo/store/models"

func (m MockStore) GetPackageCounts(channel, name string) ([]*models.PackageCount, error) {
	panic("implement me")
}

func (m MockStore) CreateInitialPackageCount(pkg *models.PackageCount) (*models.PackageCount, error) {
	panic("implement me")
}

func (m MockStore) IncreasePackageCount(channel, name, platform string) (*models.PackageCount, error) {
	panic("implement me")
}
