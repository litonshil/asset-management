package repository

import (
	"asset-management/domain"
	"gorm.io/gorm"
)

type dependencyRepository struct {
	db *gorm.DB
}

func NewDependencyRepository(db *gorm.DB) domain.DependencyRepository {
	return &dependencyRepository{
		db: db,
	}
}

func (r *dependencyRepository) CreateDependency(req domain.Dependency) (domain.Dependency, error) {
	// TODO: implement this

	return domain.Dependency{}, nil
}

func (r *dependencyRepository) GetDependencies(assetID int) ([]domain.Dependency, error) {
	// TODO: implement this

	return []domain.Dependency{}, nil
}

func (r *dependencyRepository) DeleteDependency(assetID int, dependencyID int) error {
	// TODO: implement this

	return nil
}
