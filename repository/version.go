package repository

import (
	"asset-management/domain"
	"gorm.io/gorm"
)

type versionRepository struct {
	db *gorm.DB
}

func NewVersionRepository(db *gorm.DB) domain.VersionRepository {
	return &versionRepository{
		db: db,
	}
}

func (r *versionRepository) CreateVersion(req domain.Version) (domain.Version, error) {
	// TODO: implement this

	return domain.Version{}, nil
}

func (r *versionRepository) GetVersions() ([]domain.Version, error) {
	// TODO: implement this

	return []domain.Version{}, nil
}

func (r *versionRepository) GetVersion(id int) (domain.Version, error) {
	// TODO: implement this

	return domain.Version{}, nil
}

func (r *versionRepository) UpdateVersion(req domain.Version) error {
	// TODO: implement this

	return nil
}

func (r *versionRepository) DeleteVersion(id int) error {
	// TODO: implement this

	return nil
}
